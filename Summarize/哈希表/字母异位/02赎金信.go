package 字母异位

/**
 * @Description

给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。

如果可以，返回 true ；否则返回 false 。

magazine 中的每个字符只能在 ransomNote 中使用一次。


 * @Author ZzzWw
 * @Date 2022-06-29 15:51
 **/

func canConstruct(ransomNote string, magazine string) bool {
	record := [26]int{}
	for i := 0; i < len(magazine); i++ {
		record[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		record[ransomNote[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if record[i] < 0 {
			return false
		}
	}
	return true
}
