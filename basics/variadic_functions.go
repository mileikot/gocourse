package basics

import "fmt"

func variadicFunctions() {
	// ... Ellipsis
	// func functionName(param1 ...type1) returnType {
	// 	code block
	// 	return value
	// }

	// fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))

	sequence1, total1 := sum(1, 1, 2, 3, 4, 5, 6)
	fmt.Println(sequence1, total1)

	numbers := []int{1, 3, 4, 5, 9}

	sequensce2, total2 := sum(2, numbers...)
	fmt.Println(sequensce2, total2)
}

func sum(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}

	return sequence, total
}