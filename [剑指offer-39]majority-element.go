
"""
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。

 

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

 

示例 1:

输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
输出: 2
 

限制：

1 <= 数组长度 <= 50000

来源：剑指offer-39
链接：https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof

"""

func majorityElement(nums []int) int {
    var major, count int
    for _, v := range nums{
        if count == 0{
            major = v
        }
        if v == major{
            count++
        }
        if v != major{
            count--
        }
    }
    return major

}