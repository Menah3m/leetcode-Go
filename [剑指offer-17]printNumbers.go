"""
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

示例 1:

输入: n = 1
输出: [1,2,3,4,5,6,7,8,9]
 

说明：

用返回一个整数列表来代替打印
n 为正整数

Go



链接：https://leetcode-cn.com/leetbook/read/illustrate-lcof/xzvgc2/
来源：剑指offer-17


"""


func printNumbers(n int) []int {
    res := make([]int, 0)
    c := 1
    for i:=0;i<n;i++{
        c *= 10
    }
    for j:=1;j<c;j++{
        res = append(res, j)
    }
    return res
}