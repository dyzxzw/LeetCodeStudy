package 不重复的数

/**
 * @Description
假设要找出数组arr[]={2,2,3,3,5,6}中不成对出现的两个数，既5和6

如果可以把5和6分别到两组里，然后将这两组的数异或在一起，就可以得到这两个数了

首先将数组的所有值异或在一起的到5和6异或的结果，既2^2^3^3^5^6=3

然后找到3二进制位任意一个为1的位置，因为异或的性质是相异为1，那么如果按照这个二进制位对数组进行分组，那么5和6必然被分开了

3的二进制中第一位就是1，那么我们按这位进行分组。

第一位二进制位位0:{2,2,6};第一位二进制位位1:{3,3,5}

最后将这两组数分别异或并输出就可以得到结果了

 * @Author ZzzWw
 * @Date 2022-06-13 9:50
 **/

func NumberOf2(nums[]int)[]int{
	if nums==nil{
		return nil
	}
	if len(nums)<=2{
		return nil
	}
	res:=0
	x,y:=0,0 //记录结果
	for i:=0;i<len(nums);i++{
		res^=nums[i]
	}
	index:=findFirstOne(res)
	for _,n:=range nums{
		if isOne(n,index){
			x^=n
		}else{
			y^=n
		}
	}
	return []int{x,y}
}

//找到最低位为1的位置
func findFirstOne(n int) int{
	i:=0
	for n&1==0{
		i++
		n>>=1
	}
	return i
}
//判断数字的那一个位置是否为1
func isOne(num int,index int)bool{
	if (num>>index) & 1==1{
		return true
	}
	return false
}