package main

import (
	"fmt"
)

func main() {
	theInts := []int{0: 10}
	for i := range theInts {
		if i%2 == 0 {
			fmt.Println(i, " is even")
		} else {
			fmt.Println(i, " is odd")
		}
	}
}
