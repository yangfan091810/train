package main

import (
	"fmt"
	"math"
	"time"
)
const N = 1000
func main(){
	var b [N]int
	for i := 0; i < N; i++ {
		b[i] = i+1
	}
	var x int = 2
	t := time.Now().Unix()
	r := f(N, b, x)
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
