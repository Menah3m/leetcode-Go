"""
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]  
   1
    \
     2
    /
   3 

输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

来源：LeetCode-144
链接：https://leetcode-cn.com/problems/binary-tree-preorder-traversal

"""


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归 实现
 func preorderTraversal(root *TreeNode) []int {
    res := helper(root)
    return res

}

func helper(root *TreeNode) []int{
    if root == nil{
        return []int{}
    }
    res := append([]int{root.Val}, helper(root.Left)...)
    res = append(res, helper(root.Right)...)
    return res
    
}