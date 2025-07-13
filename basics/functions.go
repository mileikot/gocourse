package basics

import "fmt"

func functions() {
	// func <name>(parameters list) returnType {
	// 	code block
	// 	return value
	// }

	// sum := add(5, 10)
	// fmt.Println("Sum:", add(5, 10))

	// greet := func() {
	// 	fmt.Println("Anonymous function executed")
	// }

	// greet()

	// operation := add

	// result := operation(20, 30)
	// fmt.Println("Result of operation:", result)

	//Passing function as an argument
	result := applyOperation(5, 3, add)
	fmt.Println("Result of applyOperation(5 + 3):", result)

	// return and using function
	multiplyBy2 := createMiltiplier(2)
	fmt.Println("Result of multiplyBy2(6 * 2):", multiplyBy2(6))
}

func add(a, b int) int {
	return a + b
}

// Function that takes a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// function that returns a function
func createMiltiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}