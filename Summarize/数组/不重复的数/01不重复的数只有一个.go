package 不重复的数

/**
 * @Description
查找数组中一个非成对的数字
假设找出数组arr[]={1,1,2,2,3,3,4}中不成对出现的数字，既4.
异或的性质：a^a=0;b^0=b.
将数组中的所以值安位异或起来就可以找到唯一一个不成对出现的数字，既1^1^2^2^3^3^4=4
 * @Author ZzzWw
 * @Date 2022-06-13 9:46
 **/

func NumberOf1(nums[]int)int{
	if nums==nil{
		return -1
	}
	if len(nums)<=2{
		return -1
	}
	res:=0
	for i:=0;i<len(nums);i++{
		res^=nums[i]
	}
	return res
}