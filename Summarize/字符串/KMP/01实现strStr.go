package KMP

/**
 * @Description
实现strStr()函数。

给你两个字符串haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。
如果不存在，则返回 -1 。

核心思想：
一个KMP函数：通过Pat构造dp数组
一个search函数：借助dp数组匹配txt
1：
dp[j][c] = next
0 <= j < M，代表当前的状态
0 <= c < 256，代表遇到的字符（ASCII 码）
0 <= next <= M，代表下一个状态

dp[4]['A'] = 3 表示：
当前是状态 4，如果遇到字符 A，
pat 应该转移到状态 3

dp[1]['B'] = 2 表示：
当前是状态 1，如果遇到字符 B，
pat 应该转移到状态 2

2.构建状态转移图：
for 0 <= j < M: # 状态
    for 0 <= c < 256: # 字符
        dp[j][c] = next

int X # 影子状态
for 0 <= j < M:
    for 0 <= c < 256:
        if c == pat[j]:
            # 状态推进
            dp[j][c] = j + 1
        else:
            # 状态重启
            # 委托 X 计算重启位置
   #只能回退，而且 KMP 就是要尽可能少的回退，以免多余的计算。
#那么 j 就可以去问问和自己具有相同前缀的 X，如果 X 遇见 “A” 可以进行「状态推进」，那就转移过去，因为这样回退最少。
            dp[j][c] = dp[X][c]



 * @Author 赵稳
 * @Date 2022-05-23 17:47
 **/


type KMP struct {
	 dp [][]int
	 pat string
}
func (this*KMP)NewKmp(pat string){
	this.pat=pat
	M:=len(this.pat) //pat是要匹配的字符串
	//dp[当前状态][遇到的字符]=下个状态
	this.dp=make([][]int,M)
	for i:=0;i<M;i++{
		this.dp[i]=make([]int,256)
	}
	X:=0 //影子状态，初始化为0
	this.dp[0][this.pat[0]]=1 //初始状态:只有遇到 pat[0] 这个字符才能使状态从 0 转移到 1,遇到其它字符的话还是停留在状态 0
	//当前状态j=1开始
	for j:=1;j<M;j++{
		for c:=0;c<256;c++{
			this.dp[j][c]=this.dp[X][c]
		}
		this.dp[j][this.pat[j]]=j+1
		X=this.dp[X][this.pat[j]]
	}
}

func(this*KMP)Search(txt string)int{
	 M:=len(this.pat)
	 N:=len(txt)
	 j:=0
	for i:=0;i<N;i++{
		//当前状态是j,遇到字符txt[i]
		//pat应该转移到的状态
		j=this.dp[j][txt[i]]
		//如果到达终止态，返回匹配开头的索引
		if j==M {
			return i-M+1
		}
	}
	//没用到达终止态，匹配失败
	return -1
}
func strStr(haystack string, needle string) int {
	if needle==""{
		return 0
	}
	kmp:=new(KMP)
	kmp.NewKmp(needle)
	return kmp.Search(haystack)
}
func main() {

}
