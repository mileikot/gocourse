package basics

import (
	"errors"
	"fmt"
)

func multipleReturnValues() {
	// func functionName(parameter1 type1, parameter2 type2, ...) (returnType1, returnType2, ...) {
	// 	code block
	// 	return value1, value2
	// }

	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)

	result, err := compare(3, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Comparison Result:", result)
	}

}

func divide(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

func compare(a, b int) (string, error) {
	if ( a > b) {
		return "a is greater than b", nil
	} else if (a < b) {
		return "a is less than b", nil
	} else {
		return "", errors.New("unable to compare, a is equal to b")
	}
}