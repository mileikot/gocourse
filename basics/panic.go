package basics

import "fmt"

func panicLesson() {
	// panic(interface{})

	//Example of valid input
	process(10)

	//Example of invalid input
	process(-3)
}

func process(input int) {
	defer fmt.Println("Defer1")
	defer fmt.Println("Defer2")

	if input < 0 {
		fmt.Println("Before panic")
		panic("input must be non-negative")
		// fmt.Println("This line will not be executed due to panic")
		// defer fmt.Println("This defer will not execute due to panic")
	}
	fmt.Println("Processing input:", input)
}