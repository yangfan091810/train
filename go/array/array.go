package array

import "errors"

type Array struct {
	data []int
	length int
}

func NewArray(capacity int) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data: make([]int, capacity, capacity),
		length: 0,
	}
}

func (this *Array)Len() int {
	return this.length
}

func (this *Array) isIndexOutOfRange(index int) bool {
	if index >= int(cap(this.data)) {
		return true
	}
	return false
}

func (this *Array) Find(index int) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return this.data[index], nil
}

func (this *Array) Insert(index int, val int) error {
	if this.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}
	if this.length == int(cap(this.data)) {
		return errors.New("array full")
	}
	for i := this.length - 1; i >= index; i-- {
		this.data[i+1] = this.data[i]
	}
	this.data[index] = val
	this.length++
	return nil
}

func (this *Array) Delete(index int) error {
	if this.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}
	for i := index; i < this.length-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return nil
}
