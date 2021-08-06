package linkedlist

import "testing"

func TestPalindrome(t *testing.T)  {
	strs := []string{"heooeh", "hello", "heoeh", "a"}
	for _, str1 := range strs{
		l := NewLinkedList()
		for _,c := range str1 {
			l.InsertToTail(string(c))
		}
		l.Print()
		t.Log(isPalindrome(l))
		l.Print()
	}
}