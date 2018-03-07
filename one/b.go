package main

import (
	"fmt"
	"math"
	"time"
)
const N = 100
const K = 10000000
func main(){
	var b [N]int
	for i := 0; i < N; i++ {
		b[i] = i+1
	}
	var x int = 2
	var r float64
	t := time.Now().Unix()
	for j := 0; j < K; j++ {
		r = f(N, b, x)
	}
	t1 := time.Now().Unix()
	fmt.Println(t1 - t)
	fmt.Println(r)
}

func f (n int, b [N]int, x int) float64 {
	var res =float64(b[0])
	for i := 0; i < n; i++ {
		res += float64(b[i]) * math.Pow(float64(x), float64(i))
	}
	return res
}

func f1(n int, b [N]int, x int) float64 {
	var i int
	var p float64 = float64(b[n-1])
	for i = n; i > 0; i-- {
		p += float64(b[i-1]) + p * float64(x)
	}
	return p
}
