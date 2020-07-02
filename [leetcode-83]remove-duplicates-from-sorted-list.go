
"""
   Definition for singly-linked list.
   type ListNode struct {
       Val int
       Next *ListNode
"""
"""
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3

来源：LeetCode-83
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list

"""


 func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	//遍历链表
    for cur != nil{
		//碰到值相同时删除节点
        for cur.Next != nil && cur.Val == cur.Next.Val{
            cur.Next = cur.Next.Next
        }
        cur = cur.Next
    }
    return head
}