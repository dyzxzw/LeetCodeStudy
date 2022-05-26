package 反转字符串

/**
 * @Description
给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。

如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。


 * @Author 赵稳
 * @Date 2022-05-23 15:40
 **/

func reverseStr(s string, k int) string {
//1.字符串转化为切片
	b:=[]byte(s)
	for i:=0;i<len(s);i+=2*k{
		if len(s)-i>=k{
			reverse(b[i:i+k])
		}else{
			reverse(b[:])
		}
	}
	return string(b)
}
func reverse(s []byte)  {
	left,right:=0,len(s)-1
	for left<right{
		s[left],s[right]=s[right],s[left]
		left++
		right--
	}
}
func main() {

}
