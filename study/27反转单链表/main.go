package main

import "fmt"

//ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next //0.把下一节点进行缓存
		cur.Next = pre   //1.把当前指向前面一个
		pre = cur        //2.把当前值
		cur = next
	}
	return pre
}

func main() {
	//0.创建一个单链表
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	// var ret *
	ret := reverseList(head)
	for ret != nil {
		fmt.Print(ret.Val, "->")
		ret = ret.Next
	}
	fmt.Println("\ndada")
}
