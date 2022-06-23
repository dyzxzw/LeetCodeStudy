package 单调栈

/**
 * @Description
arr = [1,2,3] 的下一个排列是 [1,3,2] 。

31
找下一个更大排列的基本思想：
1 从后向前找第一个升序对i,j (一对，两个数字）
2 后从向前找第一个K,k>i
3 交换i,k
4 逆序j,end


同理 如果要找上一个排列，
   只需要将所有的不等式符号全部反转,不要带等号

 * @Author ZzzWw
 * @Date 2022-06-23 16:26
 **/

func nextPermutation(nums []int){
   if nums==nil || len(nums)<=1{
	   return
   }
   i:=len(nums)-1
   for ;;{
	   j:=i
	   i--
	   if nums[i]<nums[j]{
		   k:=len(nums)//注意这个地方，不要写成len(nums)-1
		   for ;;{
			   k--
			   if nums[k]>nums[i]{
				   nums[i],nums[k]=nums[k],nums[i]
				   //Reverse
				   end := len(nums) - 1
				   for j < end {
					   nums[j], nums[end] = nums[end], nums[j]
					   j++
					   end--
				   }
				   return
			   }
		   }
	   }
	   if  i==0{
		   //从小到大排序
		   end := len(nums) - 1
		   for i < end {
			   nums[i], nums[end] = nums[end], nums[i]
			   i++
			   end--
		   }
		   return
	   }
   }
	return
}