"""
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

注意:
不能使用代码库中的排序函数来解决这道题。

示例:

输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]
进阶：

一个直观的解决方案是使用计数排序的两趟扫描算法。
首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
你能想出一个仅使用常数空间的一趟扫描算法吗？

来源：LeetCode-75
链接：https://leetcode-cn.com/problems/sort-colors

"""


// 1次遍历（计数）+ 1次遍历（填值）
func sortColors(nums []int)  {
    count_0 := 0
    count_1 := 0
    count_2 := 0
    
    for _, v := range nums {
        if v == 0 {
            count_0 += 1
        }
        if v == 1 {
            count_1 += 1
        }
        if v == 2 {
            count_2 += 1
        }
    }
    for i:=0;i<count_0;i++{
        nums[i] = 0
    }
    
    for j:=count_0;j<count_0+count_1;j++{
        nums[j] = 1
    }
    for k:=count_0+count_1;k<count_0+count_1+count_2;k++{
        nums[k] = 2
    }
    
    return 
}

// 双指针（交换0 2的位置）
func sortColors(nums []int)  {
    l := 0
    r := len(nums)-1
    for i:=0; i<=r; i++ {
        if nums[i] == 0{
            nums[i], nums[l] = nums[l], nums[i]
            l += 1
        }
        if nums[i] == 2{
            nums[i], nums[r] = nums[r], nums[i]
            i -= 1
            r -= 1
        }
    }
    return 

}