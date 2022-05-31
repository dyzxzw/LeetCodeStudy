package 螺旋矩阵

/**
 * @Description
给你一个 m 行 n 列的矩阵 matrix ，
请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
 * @Author ZzzWw
 * @Date 2022-05-27 14:22
 **/
func spiralOrder(matrix [][]int) []int {
	res:=[]int{}
	if matrix==nil{
		return nil
	}
	m,n:=len(matrix),len(matrix[0])
	top,left,right,down:=0,0,n-1,m-1
	for{
		//top,left---->top,right
		for i:=left;i<=right;i++{
			res=append(res,matrix[top][i])
		}
		top++
		if top>down{
			break
		}
		//top,right--->down,right
		for i:=top;i<=down;i++{
			res=append(res,matrix[i][right])
		}
		right--
		if right<left{
			break
		}
		//down,right--->down,left
		for i:=right;i>=left;i--{
			res=append(res,matrix[down][i])
		}
		down--
		if down<top{
			break
		}
		//down,left---->top,left
		for i:=down;i>=top;i--{
			res=append(res,matrix[i][left])
		}
		left++
		if left>right{
			break
		}
	}
	return res
}