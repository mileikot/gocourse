package basics

import "fmt"

func rangeFunc() {
	message := "Hello, World!"
	for i, v := range message {
		println("Index:", i, "Unicode:", v)
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}
}