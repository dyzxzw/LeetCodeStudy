package 字母异位

import "sort"

/**
 * @Description

给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

 * @Author ZzzWw
 * @Date 2022-06-29 15:53
 **/

//这题还是有点意思 哈希表
func groupAnagrams(strs []string) (res [][]string){
	//哈希表 关键字：排序后的单词，值 :切片
	hashmap:=make(map[string][]string)
	for i:=0;i<len(strs);i++{
		word:=[]byte(strs[i])
		//将单词进行排序，用到切片的排序
		sort.Slice(word,func(i,j int)bool{
			return word[i]<word[j]
		})
		sortedWord:=string(word)
		hashmap[sortedWord]=append(hashmap[sortedWord],strs[i])
	}
	for _,v:=range hashmap{
		res=append(res,v)
	}
	return res
}