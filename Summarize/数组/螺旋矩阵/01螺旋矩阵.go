package 螺旋矩阵

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
 * @Description
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，
且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
 * @Author 赵稳
 * @Date 2022-05-27 14:22
 **/
func generateMatrix(n int) [][]int{
	//构造n x n的矩阵
	res:=make([][]int,n)
	for i:=0;i<n;i++{
		res[i]=make([]int,n)
	}
	num:=1
	left,right,top,down:=0,n-1,0,n-1
	for num<=n*n{
		//top,left--->top,right
		for i:=left;i<=right;i++{
			res[top][i]=num
			num++
		}
		top++
		//top,right--->down,right
		for i:=top;i<=down;i++{
			res[i][right]=num
			num++
		}
		right--
		//down,right---->down,left
		for i:=right;i>=left;i--{
			res[down][i]=num
			num++
		}
		down--
		//down,left---->top,left
		for i:=down;i>=top;i--{
			res[i][left]=num
			num++
		}
		left++
	}
	return res
}
func generateMatrixTest(){
	input:=bufio.NewScanner(os.Stdin)
	input.Scan()//读取N
	n,_:=strconv.Atoi(input.Text())
	res:=generateMatrix(n)
	fmt.Println(res)
}