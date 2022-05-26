package 合并数组问题_涉及时空复杂度问题

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * @Description

给你两个按 非递减顺序 排列的整数数组nums1 和 nums2，
另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：
nums1 的初始长度为 m + n
其中前 m 个元素表示应合并的元素
后 n 个元素为 0 ，应忽略



 * @Author 赵稳
 * @Date 2022-05-25 14:13
 **/

/*
首先第一种提法：
合并两个有序数组，要求时间复杂度 0(m+n) 空间复杂度O(m+n)
*/

func merge1(nums1 []int, m int, nums2 []int, n int){
	i,j:=0,0
	//构造临时数组
	tmp:=make([]int,0,m+n)//千万不要忘记了 len=0
	for {
		if nums1[i]<nums2[j]{
			tmp=append(tmp,nums1[i])
			i++
		}else {
			tmp=append(tmp,nums2[j])
			j++
		}
		if i==m{
			tmp=append(tmp,nums2[j:]...)
			break
		}
		if j==n{
			tmp=append(tmp,nums1[i:]...)
			break
		}
	}
	copy(nums1,tmp)

}

/*
第二种提法：
合并两个有序数组，要求时间复杂度 0(m+n) 空间复杂度O(1)
即：只能在原数组上修改
*/
func merge2(nums1 []int, m int, nums2 []int, n int){
	i,j,p:=m-1,n-1,m+n-1 //从后向前的指针
	for i>=0 && j>=0{
		if nums1[i]>nums2[j]{
			nums1[p]=nums1[i]
			i--
			p--
		}else{
			nums1[p]=nums2[j]
			j--
			p--
		}
	}
	//注意只需考虑j>=0的情况
	for j>=0{
		nums1[p]=nums2[j]
		j--
		p--
	}
}

/*
第三种提法：
合并两个有序数组，并去重，要求时间复杂度 0(m+n) 空间复杂度O(m+n)
这种方法有缺陷，不知道去重后的数组长度，会返回原始长度，包含多余数据
*/
func merge3(nums1 []int, m int, nums2 []int, n int){
	i,j:=0,0
	tmp := make([]int, 0, m+n) //构造临时数组
	for {
		if nums1[i]<nums2[j]{
			tmp=append(tmp,nums1[i])
			i++
		}else if nums1[i]>nums2[j]{
			tmp=append(tmp,nums2[j])
			j++
		}else{
			tmp=append(tmp,nums1[i]) //相等的时候，任意加入一个元素，指针同时移动
			i++
			j++
		}
		if i == m {
			tmp = append(tmp, nums2[j:]...)
			break
		}
		if j == n {
			tmp = append(tmp, nums1[i:]...)
			break
		}
	}
	copy(nums1, tmp)
}


/*
第四种提法：
合并两个有序数组，并去重，要求时间复杂度 0(m+n) 空间复杂度O(1)
这就要自己写个去重函数
后期扩展：可能还会考到去重函数相关，比如只保留一个重复元素，暴力两个，保留n个？
*/
func merge4(nums1 *[]int, m int, nums2 []int, n int){
	i,j,p:=m-1,n-1,m+n-1
	for i>=0 && j>=0{
		if (*nums1)[i]>nums2[j]{
			(*nums1)[p]=(*nums1)[i]
			i--
			p--
		}else{
			(*nums1)[p]=nums2[j]
			j--
			p--
		}
	}
	for j>=0{
		(*nums1)[p]=nums2[j]
		j--
		p--
	}
	newLen:=removeDuplicates(nums1)
	*nums1=(*nums1)[:newLen]
}
//去重函数,返回的是去重后的数组长度
func removeDuplicates(nums *[]int)int{
	if len(*nums)<=1{
		return len(*nums)
	}
	slow,fast:=1,1
	for fast<len(*nums){
		if (*nums)[slow-1]!=(*nums)[fast]{
			(*nums)[slow]=(*nums)[fast]
			slow++
		}
		fast++
	}
	return slow
}

func mergeTest(){
	input := bufio.NewScanner(os.Stdin)
	input.Scan() //读取一行m
	m, _ := strconv.Atoi(input.Text())
	input.Scan() //读取一行n
	n, _ := strconv.Atoi(input.Text())
	nums1 := make([]int, m+n)
	input.Scan() //读取nums1
	str := strings.Split(input.Text(), ",")
	//fmt.Println(str)
	for i, v := range str {
		value, _ := strconv.Atoi(v)
		nums1[i] = value
	}

	input.Scan() //读取nums2
	nums2 := make([]int, n)
	str2 := strings.Split(input.Text(), ",")
	for i, v := range str2 {
		value, _ := strconv.Atoi(v)
		nums2[i] = value
	}
	//merge1(nums1, m, nums2, n)
	//merge2(nums1,m,nums2,n)
	//merge3(nums1,m,nums2,n)
	merge4(&nums1,m,nums2,n)
	fmt.Println(nums1)
}