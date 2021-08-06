package linkedlist

func isPalindrome(l *LinkedList) bool{
	lLen := l.length
	if lLen == 0 {
		return false
	}
	if lLen == 1 {
		return true
	}
	isPalindrome := true
	step := lLen/2
	var pre *ListNode = nil  //前一个节点，用于反转链表用
	cur := l.head.next //
	//next := l.head.next.next
	for i := uint(1); i <= step; i++ {
		tmp := cur.GetNext() //获取下一个节点，临时存储
		cur.next = pre //反转节点指向，将当前节点的next指向pre
		pre = cur //前一个节点更新
		cur = tmp //当前节点更新
		//next = cur.GetNext() //一直往前走的指针
	}
	mid := cur //中间节点指针
	var left,right *ListNode = pre, nil //给左右中间节点赋值
	if lLen%2 != 0 { //基数
		right = mid.GetNext()
	}else {
		right = mid
	}
	//依次比较
	for left != nil && right != nil {
		if left.GetValue().(string) != right.GetValue().(string) {
			isPalindrome = false
			break
		}
		left = left.GetNext()
		right = right.GetNext()
	}

	//复原链表 此时 cur = mid ,重置cur 为pre ,pre 为mid ,即 cur 和 pre 互换位置
	cur = pre
	pre = mid
	for cur != nil {
		next := cur.GetNext()
		cur.next = pre
		pre = cur
		cur = next
	}
	l.head.next = pre
	return isPalindrome
}
