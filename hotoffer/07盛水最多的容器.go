package hotoffer

/**
 * @Description
给定一个长度为 n 的整数数组height。有n条垂线，第 i 条线的两个端点是(i, 0)和(i, height[i])。

找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

 * @Author ZzzWw
 * @Date 2022-07-30 10:27
 **/

func maxArea(height []int) int {
 left,right,ans:=0,len(height)-1,0
 for left<right{
  minVal:=min(height[left],height[right])
  area:=minVal*(right-left)
  ans=max(area,ans)
  if height[left]<=height[right]{
   left++
  }else{
   right--
  }
 }
 return ans
}
func min(a,b int)int{
 if a<b{
  return a
 }
 return b
}