package 不重复的数

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-06-13 10:01
 **/

//用哈希表记录次数吧

func NumberOfN(nums[]int)(res []int){
	if nums==nil{
		return nil
	}
	hs:=map[int]int{}
	for i:=0;i<len(nums);i++{
		hs[nums[i]]++
	}
	//遍历哈希表，找出次数为1的
	for key,value:=range hs{
		if value==1{
			res=append(res,key)
		}
	}
	return res
}