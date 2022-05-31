package 二分查找

/**
 * @Description

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
请必须使用时间复杂度为 O(log n) 的算法。

 * @Author ZzzWw
 * @Date 2022-05-25 13:52
 **/

//nums = [1,3,5,6], target = 2
//输出: 1

func searchInsert(nums []int, target int) int{
    //手动模拟一遍，就知道插入位置在哪了
    //具体看题目的要求
    left,right:=0,len(nums)-1
    for left<=right{
        mid:=left+(right-left)/2
        if nums[mid]==target{
            return mid
        }else if nums[mid]>target{
            right=mid-1
        }else{
            left=mid+1
        }
    }
    //如果不存在，返回的位置可以手动模拟一遍
    return left
    //return right+1
}


func main() {

}
