package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)
	var c int
	for _, b := range s {
		if b == '1' {
			c++
		}
	}
	fmt.Println(c)
}
