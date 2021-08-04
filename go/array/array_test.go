package array

import "testing"

func TestInsert(t *testing.T) {
	capacity := 10
	arr := NewArray(capacity)
	for i :=10;i < capacity - 2; i++ {
		err := arr.Insert(i, i+1)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}