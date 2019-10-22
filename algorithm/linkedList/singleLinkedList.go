// package _6_linkedlist

// import "fmt"

// /*
// 单链表基本操作
// author:leo
// */

// type ListNode struct {
// 	next  *ListNode
// 	value interface{}
// }

// type LinkedList struct {
// 	head   *ListNode
// 	length uint
// }

// func NewListNode(v interface{}) *ListNode {
// 	return &ListNode{nil, v}
// }

// func (list *ListNode) GetNext() *ListNode {
// 	return list.next
// }

// func (list *ListNode) GetValue() interface{} {
// 	return list.value
// }

// func NewLinkedList() *LinkedList {
// 	return &LinkedList{NewListNode(0), 0}
// }

// //InsertAfter 在某个节点后面插入节点
// func (list *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
// 	if nil == p {
// 		return false
// 	}
// 	newNode := NewListNode(v)
// 	oldNext := p.next
// 	p.next = newNode
// 	newNode.next = oldNext
// 	list.length++
// 	return true
// }

// //InsertBefore 在某个节点前面插入节点
// func (list *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
// 	if nil == p || p == list.head {
// 		return false
// 	}
// 	cur := list.head.next
// 	pre := list.head
// 	for nil != cur {
// 		if cur == p {
// 			break
// 		}
// 		pre = cur
// 		cur = cur.next
// 	}
// 	if nil == cur {
// 		return false
// 	}
// 	newNode := NewListNode(v)
// 	pre.next = newNode
// 	newNode.next = cur
// 	list.length++
// 	return true
// }

// //InsertToHead 在链表头部插入节点
// func (list *LinkedList) InsertToHead(v interface{}) bool {
// 	return list.InsertAfter(list.head, v)
// }

// //InsertToTail 在链表尾部插入节点
// func (list *LinkedList) InsertToTail(v interface{}) bool {
// 	cur := list.head
// 	for nil != cur.next {
// 		cur = cur.next
// 	}
// 	return list.InsertAfter(cur, v)
// }

// //FindByIndex 通过索引查找节点
// func (list *LinkedList) FindByIndex(index uint) *ListNode {
// 	if index >= list.length {
// 		return nil
// 	}
// 	cur := list.head.next
// 	var i uint = 0
// 	for ; i < index; i++ {
// 		cur = cur.next
// 	}
// 	return cur
// }

// //DeleteNode 删除传入的节点
// func (list *LinkedList) DeleteNode(p *ListNode) bool {
// 	if nil == p {
// 		return false
// 	}
// 	cur := list.head.next
// 	pre := list.head
// 	for nil != cur {
// 		if cur == p {
// 			break
// 		}
// 		pre = cur
// 		cur = cur.next
// 	}
// 	if nil == cur {
// 		return false
// 	}
// 	pre.next = p.next
// 	p = nil
// 	list.length--
// 	return true
// }

// //Print 打印链表
// func (list *LinkedList) Print() {
// 	cur := list.head.next
// 	format := ""
// 	for nil != cur {
// 		format += fmt.Sprintf("%+v", cur.GetValue())
// 		cur = cur.next
// 		if nil != cur {
// 			format += "->"
// 		}
// 	}
// 	fmt.Println(format)
// }

package main

import "fmt"

//ListNode 节点
type ListNode struct {
	next  *ListNode
	value interface{}
}

//CreateLinkedList 单链表创建
func CreateLinkedList() *ListNode {
	head := ListNode{}
	head.next = nil
	return &head
}

//LinkedListAddNode 为单链表增加node
func LinkedListAddNode(head *ListNode, newNode *ListNode) {
	node := head
	for node.next != nil {
		node = node.next
	}

	node.next = newNode
	newNode.next = nil
}

//LinkedListDelNode 删除node
func LinkedListDelNode(head *ListNode, value interface{}) {
	node := head
	for node.next != nil {
		if node.next.value == value {
			node.next = node.next.next
			break
		}
		node = node.next
	}
}

//LinkedListAddNodeBefor
func (head *ListNode) LinkedListAddNodeBefor(node *ListNode, value interface{}) {
	per := head
	cur := head
	for cur != nil {
		if cur == node {
			newNode := ListNode{per.next, value}
			per.next = &newNode
			break
		}
		per = cur
		cur = cur.next
	}
}

//LinkedListAddNodeAfter
func (head *ListNode) LinkedListAddNodeAfter(node *ListNode, value interface{}) {
	cur := head
	for cur != nil {
		if cur == node {
			newNode := ListNode{cur.next, value}
			cur.next = &newNode
			break
		}
		cur = cur.next
	}
}

//LinkedListFindNode find node

//LinkedListPrintNode printf node
func LinkedListPrintNode(head *ListNode) {
	node := head.next
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}

func (head *ListNode) LinkListGetTail() *ListNode {
	node := head
	for node.next != nil {
		node = node.next
	}
	return node
}

func main() {
	head := CreateLinkedList()

	for i := 0; i < 10; i++ {
		//node := ListNode{nil, i}
		//LinkedListAddNode(head, &node)
		head.LinkedListAddNodeBefor(head, i)
	}
	LinkedListPrintNode(head)

	for i := 0; i < 10; i++ {
		//LinkedListDelNode(head, i)
		tail := head.LinkListGetTail()
		head.LinkedListAddNodeAfter(tail, i)
		//fmt.Println("del:", i)
	}
	LinkedListPrintNode(head)
}
