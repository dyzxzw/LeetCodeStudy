package 移除元素相关

/**
 * @Description
给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，
如果两者相等，返回 true 。# 代表退格字符。

注意：如果对空文本输入退格字符，文本继续为空。

输入：s = "ab#c", t = "ad#c"
输出：true
解释：s 和 t 都会变成 "ac"。

 * @Author ZzzWw
 * @Date 2022-05-31 15:16
 **/
func getStr(s string) string {
	var tmp []rune
	for _, c := range s {
		if c != '#' {
			tmp = append(tmp, c)
		} else if len(tmp) > 0 {
			tmp = tmp[:len(tmp)-1]
		}
	}
	return string(tmp)
}
func backspaceCompare(s string, t string) bool {
	return getStr(s) == getStr(t)
}