package heapsort

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-06-10 14:25
 **/


func heapSort(nums[]int,len int){
	for i:=(len-1)/2;i>=0;i--{
		heapify(nums,i,len-1)
	}
	for i:=len-1;i>0;i--{
		nums[0],nums[i]=nums[i],nums[0]
		heapify(nums,0,i-1)
	}
}

func heapify(nums[]int,start,end int){
	dad:=start
	son:=2*dad+1
	for son<=end{
		if son+1<=end && nums[son]<nums[son+1]{
			son++
		}
		if nums[dad]>nums[son]{
			break
		}
		nums[son],nums[dad]=nums[dad],nums[son]
		dad=son
		son=2*dad+1
	}
}