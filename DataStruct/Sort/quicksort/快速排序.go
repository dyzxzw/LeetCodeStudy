package quicksort

import (
	"fmt"
	"math/rand"
)

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-06-10 14:33
 **/

func QuickSort(nums[]int,left,right int){
	if left>=right{
		return
	}
	base:=division(nums,left,right)
	QuickSort(nums,left,base-1)
	QuickSort(nums,base+1,right)
}
func division(nums[]int,left,right int)int{
	//定义随机下标
	pos:=rand.Intn(right-left+1)+left
	nums[left],nums[pos]=nums[pos],nums[left]
	base:=nums[left]
	for left<right{
		for left<right&&nums[right]>=base{
			right--
		}
		nums[left]=nums[right]
		for left<right&&nums[left]<=base{
			left++
		}
		nums[right]=nums[left]
	}
	nums[left]=base
	return left
}

func test(){
	nums := []int{4, 2, 5, 6, 7, 3, 34, 5, 6}
	QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}