package 字母异位

/**
 * @Description

给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。


 * @Author ZzzWw
 * @Date 2022-06-29 15:36
 **/

func isAnagram(s string, t string) bool {
	record := [26]int{}
	for i := 0; i < len(s); i++ {
		record[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		record[t[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if record[i] != 0 {
			return false
		}
	}
	return true
}