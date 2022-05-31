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
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，
使每个元素 最多出现两次 ，返回删除后数组的新长度。
 * @Author 赵稳
 * @Date 2022-05-31 15:02
 **/
//典型题目，需要记忆

func removeDuplicates2(nums []int) int{
	n:=len(nums)
	if n<=2{
		return n
	}
	slow,fast:=2,2
	for fast<n{
		if nums[slow-2]!=nums[fast]{
			nums[slow]=nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func removeDuplicates2Test(){
	input:=bufio.NewScanner(os.Stdin)
	input.Scan()
	var nums []int
	str:=strings.Split(input.Text(),",")
	for _,s:=range str{
		v,_:=strconv.Atoi(s)
		nums=append(nums,v)
	}

	newLen:=removeDuplicates2(nums)
	fmt.Println("新长度为:",newLen)
	fmt.Println(nums[:newLen])
}

//同理，最多出现K次
func removeDuplicatesK(nums []int, k int) int{
	n:=len(nums)
	if n<=k{
		return n
	}
	slow,fast:=k,k
	for fast<n{
		if nums[slow-k]!=nums[fast]{
			nums[slow]=nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func removeDuplicatesKTest(){
	input:=bufio.NewScanner(os.Stdin)
	input.Scan()
	var nums []int
	str:=strings.Split(input.Text(),",")
	for _,s:=range str{
		v,_:=strconv.Atoi(s)
		nums=append(nums,v)
	}
    input.Scan()//读入k
	k,_:=strconv.Atoi(input.Text())
	newLen:=removeDuplicatesK(nums,k)
	fmt.Println("新长度为:",newLen)
	fmt.Println(nums[:newLen])
}