package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	maxCnt := math.MaxInt64
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		cnt := 0
		for a%2 == 0  {
			a /= 2
			cnt++
		}

		if cnt < maxCnt {
			maxCnt = cnt
		}
	}
	fmt.Println(maxCnt)
}