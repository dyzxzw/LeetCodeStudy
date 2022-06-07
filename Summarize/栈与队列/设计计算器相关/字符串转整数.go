package 设计计算器相关

/**
 * @Description
s="123" 正整数 转化为整数
实现strconv.Atoi()
 * @Author ZzzWw
 * @Date 2022-05-31 20:14
 **/

func strToNum(s string)int{
	n:=0
	for  i := 0; i < len(s); i++ {
		c:=s[i]
		n = 10 * n + int(c - '0')
	}
  return n
}