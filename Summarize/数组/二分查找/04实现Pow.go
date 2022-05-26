package 二分查找

/**
 * @Description
实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn ）

关键思想：
1.n的二进制位数 ：[log2n]+1
所以x^n=x^1 * x^2 * x^4.... x^2log2n ===>时间复杂度 O(log2n)

根据以上推导，可把计算 x^n转化为解决以下两个问题：
1.计算 x^1 * x^2 * x^4.... x^2log2n：循环赋值操作x=x^2
2.观察得知，当n & 1==1时，结果要累乘x

比如3^13=3^8 *3^4 *3^1

为什么是3^8 *3^4 *3^1？
因为 n& 1==1


//记住 快速幂模板
LL binPow(LL x,LL k)
{
 LL ans=1;
 while(k){
   if(k&1==1) ans*=x;
   x*=x;
   k>>=1;
}
return ans;
}

 * @Author 赵稳
 * @Date 2022-05-25 14:04
 **/

func myPow(x float64, n int) float64 {
	flag:=false
	if n<0{
		n=-n
		flag=true
	}
	var res float64
	res=1
	for n>0{
		if n&1==1{
			res*=x
		}
		x*=x
		n>>=1
	}
	if flag==true{
		return 1/res
	}
	return res
}