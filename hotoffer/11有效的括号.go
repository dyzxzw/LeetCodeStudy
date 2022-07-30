package hotoffer

/**
 * @Description
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

 * @Author ZzzWw
 * @Date 2022-07-30 10:52
 **/

func isValid(s string) bool {
      hmap:=map[byte]byte{
		  ')':'(',
		  ']':'[',
		  '}':'{',
	  }
	  st:=make([]byte,0)
	  if len(s)<=1{
		  return false
	  }
	  if s==""{
		  return false
	  }
	  for i:=0;i<len(s);i++{
		  if s[i]=='(' || s[i]=='[' || s[i]=='{'{
			  st=append(st,s[i])
		  }else{
			  if len(st)>0 && hmap[s[i]]==st[len(st)-1]{
				  st=st[:len(st)-1]
			  }else{
				  return false
			  }
		  }
	  }
	  return len(st)==0
}