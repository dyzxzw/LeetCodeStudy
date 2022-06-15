package 链表相交与成环问题

/**
 * @Description

编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」定义为：

对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果这个过程 结果为1，那么这个数就是快乐数。
如果 n 是 快乐数 就返回 true ；不是，则返回 false

 * @Author ZzzWw
 * @Date 2022-06-15 15:04
 **/

//和寻找重复数类似，快慢指针方法YY DS

func isHappy(n int) bool {
	slow:=n
	fast:=getSum(n)
	for fast!=1 && fast!=slow{
		slow=getSum(slow)
		fast=getSum(getSum(fast))
	}
	return fast == 1
}

func getSum(n int) int {
	res := 0
	for n > 0 {
		tmp := n % 10
		res += tmp * tmp
		n = n / 10
	}

	return res
}