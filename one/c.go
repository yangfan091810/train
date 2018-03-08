package main

import (
	"fmt"
	"time"
)
const K = 10000000
func main(){
	var max int = 0
	var n int = 10
	var a = [10]int{ 2, 8, -3, 5, 6, -9, -4, 3, -9, 1}
	t := time.Now().Unix()
	for i := 0; i < K; i++ {
		max = f1(a, n)
	}
	t1 := time.Now().Unix()
	fmt.Println(t1-t)
	fmt.Println(max) 
}
//求给定数组的最大子列和，如下函数的时间复杂度为O(N3)
func f(A [10]int, N int) int {
	var ThisSum, MaxSum int = 0, 0
	var i, j, k int
	for i = 0; i < N; i++ {
		for j = i; j < N; j++ {
			ThisSum = 0
			for k = i; k < j; k++{
				ThisSum += A[k]
				if ThisSum > MaxSum {
					MaxSum = ThisSum
				}
			}
		}
	}
	return MaxSum
}
//时间复杂度为O(N2)
func f1(A [10]int, N int) int {
	var ThisSum, MaxSum int = 0, 0
	var i, j int
	for i = 0; i < N; i++ {
		ThisSum = 0
		for j = i; j < N; j++ {
			ThisSum += A[j]
			if ThisSum > MaxSum {
				MaxSum = ThisSum
			}
		}
	}
	return MaxSum
}







