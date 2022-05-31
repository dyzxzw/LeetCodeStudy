package 滑动窗口

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
 * @Description
输入一个正整数 target ，
输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列

输入：target = 9
输出：[[2,3,4],[4,5]]

 * @Author ZzzWw
 * @Date 2022-05-31 15:36
 **/

//双指针+滑动窗口

func findContinuousSequence(target int) [][]int {
	left,right:=1,1 //滑动窗口的左 右 边界、
	sum:=0
	var res [][]int
	/*
	  为什么边界判断是left<=target/2
	  因为left最大为target/2,且要求至少含有两个数字
	  当 left=target/2时 ;下一个数字最小为：left+1=target/2+1
	  此时和sum=left+left+1=target+1>target，超过了target，不满足
	*/

	for left<=target/2{
		//1.如果sum<target,说明要移动窗口的右边界，增加sum
		if sum<target{
			sum+=right
			right++
		}else if sum>target{
			//2.如果sum>target,说明值大了，要减小，缩小窗口，移动左边界
			sum-=left
			left++
		}else{
			//3.sum==target,把找到的答案加入结果集合中
			tmp:=[]int{}
			for i:=left;i<right;i++{
				tmp=append(tmp,i)
			}
			res=append(res,tmp)
			sum-=left //加完不要忘记再次移动一次左边界
			left++
		}
	}
	return res
}

func findContinuousSequenceTest() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan() //读取target
	target, _ := strconv.Atoi(input.Text())
	res := findContinuousSequence(target)
	fmt.Println(res)
}
