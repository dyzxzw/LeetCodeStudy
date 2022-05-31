package 移除元素相关

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * @Description
给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，
使每个元素 只出现一次 ，返回删除后数组的新长度。
元素的 相对顺序 应该保持 一致 。

 * @Author 赵稳
 * @Date 2022-05-31 14:57
 **/

//属于经典题目，要记住方法

func removeDuplicates(nums []int) int {
	n:=len(nums)
	if n<=1{
		return n
	}
	slow,fast:=1,1
	for fast<n{
		if nums[slow-1]!=nums[fast]{
			nums[slow]=nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func removeDuplicatesTest(){
	input:=bufio.NewScanner(os.Stdin)
	input.Scan()
	var nums []int
	str:=strings.Split(input.Text(),",")
	for _,s:=range str{
		v,_:=strconv.Atoi(s)
		nums=append(nums,v)
	}

	newLen:=removeDuplicates(nums)
	fmt.Println("新长度为:",newLen)
	fmt.Println(nums[:newLen])
}