package _linkedlist

import (
	"fmt"
	"testing"
)

func TestInsertToHead(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToHead(i + 1)
	}
	l.Print()
}

func TestInsertToTail(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
}

func TestFindByIndex(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	t.Log(l.FindByIndex(0))
	t.Log(l.FindByIndex(9))
	t.Log(l.FindByIndex(5))
	t.Log(l.FindByIndex(11))
}

func TestDeleteNode(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 3; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()

	t.Log(l.DeleteNode(l.head.next))
	l.Print()

	t.Log(l.DeleteNode(l.head.next.next))
	l.Print()
}

func TestLinkedListReverse(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 5; i++ {
		l.InsertToTail(i)
	}
	l.Print()

	LinkedListReverse(l)
	fmt.Print("reverse:")
	l.Print()
}

func TestLinkedListHasCycle(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 5; i++ {
		l.InsertToTail(i)
	}
	l.Print()

	node := l.head.next
	for node.next != nil {
		node = node.next
	}
	node.next = l.head

	b := LinkedListHasCycle(l)
	fmt.Println(b)
}

func TestLinkListMerge(t *testing.T) {
	fmt.Println("merage:")
	l := NewLinkedList()
	for i := 0; i < 10; i += 2 {
		l.InsertToTail(i)
	}
	l.Print()

	j := NewLinkedList()
	for i := 1; i < 10; i += 2 {
		j.InsertToTail(i)
	}
	j.Print()

	k := LinkListMerge(l, j)
	k.Print()
}

func TestLinkedListDelNode(t *testing.T) {
	fmt.Println("delete")

	l := NewLinkedList()
	for i := 0; i < 5; i++ {
		l.InsertToTail(i)
	}
	l.Print()

	fmt.Println("5")
	LinkedListDelNode(l, 1)
	l.Print()
}

func TestLinkedListGetMidNode(t *testing.T) {
	fmt.Println("midNode")
	l := NewLinkedList()
	for i := 0; i < 5; i++ {
		l.InsertToTail(i)
	}
	l.Print()

	mnode := LinkedListGetMidNode(l)
	fmt.Println(mnode.GetValue())

	j := NewLinkedList()
	for i := 0; i < 6; i++ {
		j.InsertToTail(i)
	}
	j.Print()

	node := LinkedListGetMidNode(l)
	fmt.Println(node.GetValue())
}

func TestFindMiddleNode(t *testing.T) {
	fmt.Println("TestFindMiddleNode")
	l := NewLinkedList()
	for i := 0; i < 5; i++ {
		l.InsertToTail(i)
	}
	l.Print()

	mnode := l.FindMiddleNode()
	fmt.Println(mnode.GetValue())

	j := NewLinkedList()
	for i := 0; i < 6; i++ {
		j.InsertToTail(i)
	}
	j.Print()

	node := j.FindMiddleNode()
	fmt.Println(node.GetValue())
}
