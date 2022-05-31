package 反转字符串

/**
 * @Description
给定一个字符串 s ，你需要反转字符串中每个单词的字符顺序
同时仍保留空格和单词的初始顺序

输入：s = "Let's take LeetCode contest"
输出："s'teL ekat edoCteeL tsetnoc"

 * @Author ZzzWw
 * @Date 2022-05-23 17:09
 **/

func reverseWords(s string) string {
         n,i:=len(s),0
		 b:=[]byte(s)
		 for i<n{
			 start:=i
			 for i<n && b[i]!=' '{
				 i++
			 }
			 left,right:=start,i-1
			 for left<right{
				 b[left],b[right]=b[right],b[left]
				 left++
				 right--
			 }
			 for i<n && b[i]==' '{
				 i++
			 }
		 }
		 return string(b)
}

func main() {

}
