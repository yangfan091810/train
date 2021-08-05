package linkedlist

import "testing"

func TestLinkedList_InsertToHead(t *testing.T) {
	l := NewLinkedList()
	for i := 0 ;i < 10; i++ {
		l.InsertToHead(i+1)
	}
	l.Print()
}

func TestLinkedList_InsertToTail(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
}

func TestLinkedList_FindByIndex(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
	t.Log(l.FindByIndex(5).GetValue())
	t.Log(l.FindByIndex(1).GetValue())
	t.Log(l.FindByIndex(8).GetValue())
}

func TestLinkedList_DeleteNode(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
	cur := l.head.next
	for i := 0; i < 3; i++ {
		cur = cur.next
	}
	t.Log(l.DeleteNode(cur))
	l.Print()
}