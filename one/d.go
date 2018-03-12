//分治法求最大子列和问题
package main

import (
	"fmt"
	"math"
)

func main(){
	var list = [8]int{4,-3,5,-2,-1,2,6,-2}
	max := fenzhi(list, 0, 7)
	fmt.Println(max)
}

func max3(a int, b int, c int) int {
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	} else {
		if b > c {
			return b
		} else {
			return c
		}
	}
}

func fenzhi( A [8]int, left int, right int) int {
	var maxLeftSum , maxRightSum int
	var maxLeftBorderSum , maxRightBorderSum int

	var LeftBorderSum , RightBorderSum int
	var center , i int

	if left == right {
		if A[left] > 0 {
			return A[left]
		} else {
			return 0
		}
	}

	center = int(math.Floor(float64((left+right)/2)))

	maxLeftSum = fenzhi(A, left, center)
	maxRightSum = fenzhi(A, center+1, right)

	//分界线往左求最大子数列
	LeftBorderSum = 0
	maxLeftBorderSum = 0
	for i=center; i >= left; i-- {
		LeftBorderSum += A[i]
		if LeftBorderSum > maxLeftBorderSum {
			maxLeftBorderSum = LeftBorderSum
		}
	}

	//以分界线往右求最大子列和
	RightBorderSum = 0
	maxRightBorderSum = 0
	for j:= center+1; j<=right; j++ {
		RightBorderSum += A[j]
		if RightBorderSum > maxRightBorderSum {
			maxRightBorderSum = RightBorderSum
		}
	}
	
	maxBorderSum := maxLeftBorderSum + maxRightBorderSum
	return max3(maxLeftSum, maxRightSum, maxBorderSum)
}














