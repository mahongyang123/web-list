package _linkedlist

import "fmt"

/*单链表基本操作*/

//ListNode 单链表节点
type ListNode struct {
	next  *ListNode
	value interface{}
}

//LinkedList 单链表头结点
type LinkedList struct {
	head   *ListNode
	length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

func (list *ListNode) GetNext() *ListNode {
	return list.next
}

func (list *ListNode) GetValue() interface{} {
	return list.value
}

// func NewLinkedList() *LinkedList {
// 	return &LinkedList{&ListNode{nil, 0}, 0}
// }

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

//InsertAfter 在某个节点后面插入节点
func (list *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}

	newNode := NewListNode(v)
	newNode.next = p.GetNext()
	p.next = newNode
	list.length++

	return true
}

//InsertBefore 在某个节点前面插入节点
func (list *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if p == nil || p == list.head {
		return false
	}

	perNode := list.head
	curNode := perNode.GetNext()

	for curNode != nil {
		if curNode == p {
			newNode := NewListNode(v)
			newNode.next = curNode
			perNode.next = newNode
			list.length++
			return true
		}
		perNode = curNode
		curNode = curNode.GetNext()
	}

	return false
}

//InsertToHead 在链表头部插入节点
func (list *LinkedList) InsertToHead(v interface{}) bool {
	return list.InsertAfter(list.head, v)
}

//InsertToTail 在链表尾部插入节点
func (list *LinkedList) InsertToTail(v interface{}) bool {
	cur := list.head
	for nil != cur.next {
		cur = cur.next
	}
	return list.InsertAfter(cur, v)
}

//FindByIndex 通过索引查找节点
func (list *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= list.length {
		return nil
	}
	cur := list.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

//DeleteNode 删除传入的节点
func (list *LinkedList) DeleteNode(p *ListNode) bool {
	if p == nil || list.head == nil || list.head.next == nil {
		return false
	}

	perNode := list.head
	curNode := perNode.GetNext()

	for curNode != nil {
		if curNode == p {
			perNode.next = p.next
			p = nil
			list.length--
			return true
		}
		perNode = curNode
		curNode = curNode.next
	}

	return false
}

//Print 打印链表
func (list *LinkedList) Print() {
	cur := list.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

// package main

// import "fmt"

// //ListNode 节点
// type ListNode struct {
// 	next  *ListNode
// 	value interface{}
// }

// //CreateLinkedList 单链表创建
// func CreateLinkedList() *ListNode {
// 	head := ListNode{}
// 	head.next = nil
// 	return &head
// }

// //LinkedListAddNode 为单链表增加node
// func LinkedListAddNode(head *ListNode, newNode *ListNode) {
// 	node := head
// 	for node.next != nil {
// 		node = node.next
// 	}

// 	node.next = newNode
// 	newNode.next = nil
// }

// //LinkedListDelNode 删除node
// func LinkedListDelNode(head *ListNode, value interface{}) {
// 	node := head
// 	for node.next != nil {
// 		if node.next.value == value {
// 			node.next = node.next.next
// 			break
// 		}
// 		node = node.next
// 	}
// }

// //LinkedListAddNodeBefor
// func (head *ListNode) LinkedListAddNodeBefor(node *ListNode, value interface{}) {
// 	per := head
// 	cur := head
// 	for cur != nil {
// 		if cur == node {
// 			newNode := ListNode{per.next, value}
// 			per.next = &newNode
// 			break
// 		}
// 		per = cur
// 		cur = cur.next
// 	}
// }

// //LinkedListAddNodeAfter
// func (head *ListNode) LinkedListAddNodeAfter(node *ListNode, value interface{}) {
// 	cur := head
// 	for cur != nil {
// 		if cur == node {
// 			newNode := ListNode{cur.next, value}
// 			cur.next = &newNode
// 			break
// 		}
// 		cur = cur.next
// 	}
// }

// //LinkedListFindNode find node

// //LinkedListPrintNode printf node
// func LinkedListPrintNode(head *ListNode) {
// 	node := head.next
// 	for node != nil {
// 		fmt.Println(node.value)
// 		node = node.next
// 	}
// }

// func (head *ListNode) LinkListGetTail() *ListNode {
// 	node := head
// 	for node.next != nil {
// 		node = node.next
// 	}
// 	return node
// }

// func main() {
// 	head := CreateLinkedList()

// 	for i := 0; i < 10; i++ {
// 		//node := ListNode{nil, i}
// 		//LinkedListAddNode(head, &node)
// 		head.LinkedListAddNodeBefor(head, i)
// 	}
// 	LinkedListPrintNode(head)

// 	for i := 0; i < 10; i++ {
// 		//LinkedListDelNode(head, i)
// 		tail := head.LinkListGetTail()
// 		head.LinkedListAddNodeAfter(tail, i)
// 		//fmt.Println("del:", i)
// 	}
// 	LinkedListPrintNode(head)
// }

//LinkedListReverse 1. 单链表反转
func LinkedListReverse(list *LinkedList) {
	if list == nil || list.head == nil {
		return
	}

	head := list.head
	node := head.next
	for node.next != nil {
		iNode := node.next
		node.next = iNode.next

		iNode.next = head.GetNext()
		head.next = iNode
	}

	return
}

//LinkedListHasCycle 2. 链表中环的检测
func LinkedListHasCycle(list *LinkedList) bool {
	iNode := list.head
	jNode := list.head

	for iNode != nil && jNode != nil {
		iNode = iNode.GetNext()
		jNode = jNode.GetNext()
		if jNode != nil {
			jNode = jNode.GetNext()
		} else {
			return false
		}

		if iNode == jNode {
			return true
		}
	}
	return false
}

//LinkListMerge 3. 两个有序的链表合并
func LinkListMerge(l *LinkedList, j *LinkedList) *LinkedList {
	if nil == l || nil == l.head || nil == l.head.next {
		return j
	}
	if nil == j || nil == j.head || nil == j.head.next {
		return l
	}

	lNode, jNode := l.head, j.head

	lperNode, jperNode := lNode, jNode
	lcurNode, jcurNode := lperNode.next, jperNode.next

	for lcurNode != nil && jcurNode != nil {
		if lcurNode.value.(int) < jcurNode.value.(int) {
			lperNode = lcurNode
			lcurNode = lcurNode.next
		} else {
			inode := jcurNode
			jperNode.next = jcurNode.next
			jcurNode = jperNode.next

			inode.next = lcurNode
			lperNode.next = inode

			l.length++
		}
	}

	if jcurNode != nil && lcurNode == nil {
		lperNode.next = jcurNode
	}

	return l
}

//LinkedListDelNode 4. 删除链表倒数第 n 个结点
func LinkedListDelNode(l *LinkedList, n uint) {
	if l == nil || l.length < n || n == 0 {
		return
	}

	jnode := l.head
	for i := 0; uint(i) < n; i++ {
		jnode = jnode.next
	}

	ipernode := l.head
	icurnode := l.head
	for jnode != nil {
		jnode = jnode.next
		ipernode = icurnode
		icurnode = icurnode.next
	}

	ipernode.next = icurnode.next
}

//LinkedListGetMidNode 5. 求链表的中间结点
func LinkedListGetMidNode(l *LinkedList) *ListNode {
	inode := l.head
	jnode := l.head

	for jnode != nil && inode != nil {
		inode = inode.next
		jnode = jnode.next
		if jnode != nil {
			jnode = jnode.next
		} else {
			break
		}
	}

	return inode
}

func (this *LinkedList) FindMiddleNode() *ListNode {
	if nil == this.head || nil == this.head.next {
		return nil
	}
	if nil == this.head.next.next {
		return this.head.next
	}

	slow, fast := this.head, this.head
	for nil != fast && nil != fast.next {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
