package 移除元素相关

/**
 * @Description
网易互娱一面变形题： 将指定的元素1、6、3移动到末尾，
并且需要保持所有元素的相对顺序。
 * @Author ZzzWw
 * @Date 2022-05-31 15:15
 **/
func moveZ(nums []int, k int) {
	slowIndex := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != k {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}
	for i := slowIndex; i < len(nums); i++ {
		nums[i] = k
	}
}
func moveOneSixThree(nums []int) {
	moveZ(nums, 1)
	moveZ(nums, 6)
	moveZ(nums, 3)
}