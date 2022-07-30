package hotoffer

/**
 * @Description
给你两个整数 a 和 b ，不使用 运算符 + 和 - ，计算并返回两整数之和。
 * @Author ZzzWw
 * @Date 2022-07-29 21:54
 **/

func getSum(a int, b int) int {
     for b!=0{
		 temp:=a ^b //计算无进位的结果
		 b = (a&b)<<1 //计算有进位的结果
		 a = temp
	 }
	 return a
}