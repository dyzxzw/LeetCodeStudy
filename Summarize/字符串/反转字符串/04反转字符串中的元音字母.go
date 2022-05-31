package 反转字符串

/**
 * @Description
给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。

元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现。

 * @Author ZzzWw
 * @Date 2022-05-23 17:18
 **/

func reverseVowels(s string) string {
    b:=[]byte(s)
	left,right:=0,len(s)-1;
  for left<right{
	  for left<right && !IsVowels(b[left]){
		  left++
	  }
	  for left<right && !IsVowels(b[right]){
		 right--
	  }
	  b[left],b[right]=b[right],b[left]
	  left++
	  right--
  }
  return string(b)
}
func IsVowels(s byte)bool{
	if s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' ||
		s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U' {
		return true
	}
	return false
}
func main() {

}
