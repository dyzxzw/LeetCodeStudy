package hotoffer

/**
 * @Description
给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。
答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
0 <= digits.length <= 4

 * @Author ZzzWw
 * @Date 2022-07-30 10:40
 **/

func letterCombinations(digits string) []string {
    length:=len(digits)
    if length==0 || length>4{
        return nil
    }
    digitsMap:=[10]string{
        "", // 0
        "", // 1
        "abc", // 2
        "def", // 3
        "ghi", // 4
        "jkl", // 5
        "mno", // 6
        "pqrs", // 7
        "tuv", // 8
        "wxyz", // 9
    }

    res:=make([]string,0)
    recursion("",digits,0,digitsMap,&res)
    return res
}

func recursion(tempString,digits string,Index int,digitsMap [10]string,res*[]string){
    if len(tempString)==len(digits){
        //递归终止条件
        *res=append(*res,tempString)
        return
    }
    tmpK:=digits[Index]-'0'
    letter:=digitsMap[tmpK]
    for i:=0;i<len(letter);i++{
        tempString=tempString+string(letter[i])
        recursion(tempString,digits,Index+1,digitsMap,res)
        tempString=tempString[:len(tempString)-1]
    }
}