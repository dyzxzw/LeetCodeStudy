package 二分查找

/**
 * @Description
给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，
则返回 true ，否则返回 false 。
 * @Author 赵稳
 * @Date 2022-05-25 14:01
 **/
func isPerfectSquare(num int) bool {
	//和02 题 类似
	left, right := 1, num
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == num {
			return true
		} else if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

