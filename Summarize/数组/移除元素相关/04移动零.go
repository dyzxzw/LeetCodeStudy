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
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，
同时保持非零元素的相对顺序。
 * @Author 赵稳
 * @Date 2022-05-31 15:07
 **/

//移除元素类似，移除val=0;最后补足0即可

func moveZeroes(nums []int){
	slowIndex:=0
	for fastIndex:=0;fastIndex<len(nums);fastIndex++{
		if nums[fastIndex]!=0{
			nums[slowIndex]=nums[fastIndex]
			slowIndex++
			}
	}
	for i:=slowIndex;i<len(nums);i++{
		nums[i]=0
	}
}
func moveZeroesTest(){
	input:=bufio.NewScanner(os.Stdin)
	input.Scan()
	var nums []int
	str:=strings.Split(input.Text(),",")
	for _,s:=range str{
		v,_:=strconv.Atoi(s)
		nums=append(nums,v)
	}

	fmt.Println("移除前的数组为：",nums)
	moveZeroes(nums)
	fmt.Println("移除后的数组为：",nums)
}