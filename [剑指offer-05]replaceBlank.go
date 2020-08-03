
"""
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

 

示例 1：

输入：s = "We are happy."
输出："We%20are%20happy."
 


链接：https://leetcode-cn.com/leetbook/read/illustrate-lcof/xz2cf5/
来源：剑指offer-05

"""


func replaceSpace(s string) string {
    var res string = ""
    for _, v := range s{
        if v == ' '{
            res = res + "%20"
        } else {
            res = res + string(v)  //v本质是rune类型，需要转换成string
        }
    }
    return res

}