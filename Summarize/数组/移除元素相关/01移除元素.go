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
给你一个数组 nums和一个值 val，
你需要 原地 移除所有数值等于val的元素，并返回移除后数组的新长度。
不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

典型的快慢指针用法,动手模拟一次就知道了

 * @Author 赵稳
 * @Date 2022-05-31 14:50
 **/

func removeElement(nums []int, val int) int {
	slowIndex:=0
	for fastIndex:=0;fastIndex<len(nums);fastIndex++{
		if nums[fastIndex]!=val{
			nums[slowIndex]=nums[fastIndex]
			slowIndex++
		}
	}
	return slowIndex
}

func removeElementTest(){
	input:=bufio.NewScanner(os.Stdin)
	input.Scan()
	var nums []int
	str:=strings.Split(input.Text(),",")
	for _,s:=range str{
		v,_:=strconv.Atoi(s)
		nums=append(nums,v)
	}
	input.Scan()//读入val
	val,_:=strconv.Atoi(input.Text())
	newLen:=removeElement(nums,val)
	fmt.Println(nums[:newLen])
}