package 二分查找

/**
 * @Description
给你一个非负整数 x ，计算并返回x的 算术平方根 。

由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

注意：不允许使用任何内置指数函数和算符，
例如 pow(x, 0.5) 或者 x ** 0.5 。
 * @Author 赵稳
 * @Date 2022-05-25 14:00
 **/
func mySqrt(x int) int {
	//手动模拟一遍，和搜索插入位置类似，最后找不到应该返回的位置在哪
	//最后找不到应该返回的是right
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}
func main() {

}
