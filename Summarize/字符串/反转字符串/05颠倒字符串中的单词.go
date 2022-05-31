package 反转字符串

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
 * @Description
给你一个字符串 s ，颠倒字符串中 单词 的顺序。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

输入：s = "a good  example"
输出："example good a"
解释：如果两个单词间有多余的空格，颠倒后的字符串需要将单词间的空格减少到仅有一个。


 * @Author ZzzWw
 * @Date 2022-05-23 17:26
 **/
func reverseWords2(s string) string {
         //1.移除多余空格
	   s=removeExtraSpace(s)
	   tmp:=[]byte(s)
	   //2.反转整个字符串
	   reverseAllStr(tmp,0,len(tmp)-1)
	   //3.反转单个单词
	   for i:=0;i<len(tmp);i++{
		   j:=i
		   for j<len(tmp)&&tmp[j]!=' '{
			   j++
		   }
		   reverseAllStr(tmp,i,j-1)
		   i=j
	   }
	   return string(tmp)
}
//移除多余空格
func removeExtraSpace(s string)string{
	b,slow,fast:=[]byte(s),0,0
	//1.移除首部空格
	for len(b)>0&&b[fast]==' '&&fast<len(b){
		fast++
	}
	//2.移除中间多余空格
	for ;fast<len(b);fast++{
		if fast>1 && b[fast]==b[fast-1] && b[fast]==' '{
			continue
		}else{
			b[slow]=b[fast]
			slow++
		}
	}
	//3.移除尾部空格
	if b[slow-1]==' '&&slow-1>0{
		b=b[:slow-1]
	}else{
		b=b[:slow]
	}
	return string(b)
}
//反转字符串
func reverseAllStr(s []byte,left,right int){
	for left<right{
		s[left],s[right]=s[right],s[left]
		left++
		right--
	}
}

func main() {
  input:=bufio.NewScanner(os.Stdin)
  input.Scan()//读取k
  k,_:=strconv.Atoi(input.Text())
  for k>0{
	  input.Scan()//读取str
	  str:=input.Text()
	  fmt.Println(reverseWords2(str))
	  k--
  }
}
