package basics

import (
	"fmt"
	"math"
)

func main() {
	//Variable declaration
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Remainder:", result)

	const p float64 = 22/7.0
	fmt.Println("Constant value of p:", p)

	//Overflow with signed integers
	var maxInt int64 = 9223372036854775807
	fmt.Println("Maximum value of int64:", maxInt)

	maxInt = maxInt + 1
	fmt.Println("Overflowed value of int64:", maxInt)

	//Overflow with unsigned integers
	var uMaxInt uint64 = 18446744073709551615
	fmt.Println("Maximum value of uint64:", uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println("Overflowed value of uint64:", uMaxInt)

	//Floating point arithmetic
	var smallFloat float64 = 1.0e-323
	fmt.Println("Small float value before division:", smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println("Small float value after division:", smallFloat)
}