package main

import (
	"fmt"
	"time"
)

func main(){
	t := time.Now().Unix()
	printn2(1000000)
	fmt.Println()
	t1 := time.Now().Unix()
	fmt.Println(t1 - t)
}

func printn( n int){
	for i:= 0; i < n; i++ {
		fmt.Print(i)
	}
}

func printn2( n int) {
	if n > 0 {
		printn2( n - 1)
		fmt.Print(n)
	}
}
