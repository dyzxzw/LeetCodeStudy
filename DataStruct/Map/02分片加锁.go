package Map

import "sync"

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-06-14 14:25
 **/

// SharedMap 并发安全的小map,ShardCount 个这样的小map数组组成一个大map
type SharedMap struct {
	items        map[string]interface{}
	sync.RWMutex // 读写锁，保护items
}
// ShardCount 底层小shareMap数量
var ShardCount = 32

// ConcurrentHashMap 并发安全的大map，由 ShardCount 个小map数组组成，方便实现分段锁机制
type ConcurrentHashMap []*SharedMap

// fnv32 hash函数
func fnv32(key string) uint32 {
	// 著名的fnv哈希函数，由 Glenn Fowler、Landon Curt Noll和 Kiem-Phong Vo 创建
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	keyLength := len(key)
	for i := 0; i < keyLength; i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

// New 创建一个新的concurrent map.
func New() ConcurrentHashMap {
	m := make(ConcurrentHashMap, ShardCount)
	for i := 0; i < ShardCount; i++ {
		m[i] = &SharedMap{items: make(map[string]interface{})}
	}
	return m
}

// GetShardMap 返回给定key的sharedMap
func (m ConcurrentHashMap) GetShardMap(key string) *SharedMap {
	return m[uint(fnv32(key))%uint(ShardCount)]
}
// Set 添加 key-value
func (m ConcurrentHashMap) Set(key string, value interface{}) {
	// Get map shard.
	shard := m.GetShardMap(key)
	shard.Lock()
	shard.items[key] = value
	shard.Unlock()
}

// Get 返回指定key的value值
func (m ConcurrentHashMap) Get(key string) (interface{}, bool) {
	shard := m.GetShardMap(key)
	shard.RLock()
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}

// Remove 删除一个元素
func (m ConcurrentHashMap) Remove(key string) {
	// Try to get shard.
	shard := m.GetShardMap(key)
	shard.Lock()
	delete(shard.items, key)
	shard.Unlock()
}

// Has 判断元素是否存在
func (m ConcurrentHashMap) Has(key string) bool {
	// Get shard
	shard := m.GetShardMap(key)
	shard.RLock()
	// See if element is within shard.
	_, ok := shard.items[key]
	shard.RUnlock()
	return ok
}