package mergersort

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-06-10 14:31
 **/

func mergerSort(nums[]int,left,right int){
	if left>=right{
		return
	}
	mid:=left+(right-left)/2
	mergerSort(nums,left,mid)
	mergerSort(nums,mid+1,right)
	mergerArray(nums,left,mid,right)
}
func mergerArray(nums[]int,left,mid,right int){
	tmp:=make([]int,right-left+1)
	i,j,k:=left,mid+1,0
	for i<=mid&&j<=right{
		if nums[i]<=nums[j]{
			tmp[k]=nums[i]
			i++
			k++
		}else{
			tmp[k]=nums[j]
			j++
			k++
		}
	}
	for i <= mid {
		tmp[k] = nums[i]
		k++
		i++
	}
	for j <= mid {
		tmp[k] = nums[j]
		k++
		j++
	}
	for c := 0; c < k; c++ {
		nums[c+left] = tmp[c]
	}
}