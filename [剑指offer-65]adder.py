"""
写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。

 

示例:

输入: a = 1, b = 1
输出: 2
 

提示：

a, b 均可能是负数或 0
结果不会溢出 32 位整数

来源：剑指offer-65
链接：https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof
"""

func add(a int, b int) int {
    for b != 0{
        a, b = a ^ b, a & b << 1 // a^b为非进位和，a&b<<1 为进位
    }
    return a
}