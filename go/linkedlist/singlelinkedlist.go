package linkedlist

import "fmt"

// ListNode 单链表节点
type ListNode struct {
	next *ListNode
	value interface{}
}

// LinkedList 链表
type LinkedList struct {
	head *ListNode
	length uint
}

// NewListNode 创建节点
func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

// GetNext 获取下一个节点
func (this *ListNode) GetNext() *ListNode {
	return this.next
}

// GetValue 获取节点值
func (this *ListNode) GetValue() interface{} {
	return this.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		NewListNode(0),
		0,
	}
}

// InsertAfter 在链表的某个节点后插入节点
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if p == nil {
		return false
	}
	newNode := NewListNode(v)
	oldNode := p.next
	newNode.next = oldNode
	p.next = newNode
	this.length++
	return true
}

// InsertBefore 在某个节点前面插入节点
func (this *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if p == nil || p == this.head {
		return false
	}
	cur := this.head.next
	pre := this.head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	newNode := NewListNode(v)
	newNode.next = pre.next
	pre.next = newNode

	this.length--
	return true
}

// InsertToHead 在链表头部插入节点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

// InsertToTail 在链表尾部插入节点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for cur.next != nil {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

// FindByIndex 通过索引查找节点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= this.length {
		return nil
	}
	cur := this.head.next //第一个节点
	var i uint = 1
	for ; i <= index; i++ {
		cur = cur.next
	}
	return cur
}

// DeleteNode 删除传入的节点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if p == nil {
		return false
	}
	cur := this.head.next
	for cur != nil {
		if cur.next == p {
			break
		}
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	cur.next = p.next
	this.length--
	return true
}

func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if cur != nil {
			format += " -> "
		}
	}
	fmt.Println(format)
}


