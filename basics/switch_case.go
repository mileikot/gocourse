package basics

import "fmt"

func switch_case() {

	// Example of a switch statement in Go (switch case default) (fallthrough)
	// switch expression {
	// case value1:
	// 	// Code to execute if expression equals value1
	// case value2:
	// 	// Code to execute if expression equals value2
	// fallthrough
	// case value3:
	// 	// Code to execute if expression equals value3
	// default:
	// 	// Code to execute if expression does not match any case
	// }

	// Example of a switch statement in other programming languages (switch case break default)
	// switch expression {
	// case value1:
	// 	// Code to execute if expression equals value1
	// 	break // Optional break to exit the switch
	// case value2:
	// 	// Code to execute if expression equals value2
	// 	break // Optional break to exit the switch
	// case value3:
	// 	// Code to execute if expression equals value3
	// 	break // Optional break to exit the switch
	// default:
	// 	// Code to execute if expression does not match any case
	// 	break // Optional break to exit the switch
	// }

	// fruit := "pinapple"

	// switch fruit {
	// case "apple":
	// 	fmt.Println("This is an apple.")
	// case "banana":
	// 	fmt.Println("This is a banana.")
	// default:
	// 	fmt.Println("This is some other fruit.")
	// }

	// // Example of a switch statement with multiple cases
	// day := "Monday"
	// switch day {
	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("It's a weekday.")
	// case "Saturday", "Sunday":
	// 	fmt.Println("It's the weekend.")
	// default:
	// 	fmt.Println("Invalid day.")
	// }

	// number := 15

	// switch {
	// case number < 10:
	// 	fmt.Println("The number is less than 10.")
	// case number >= 10 && number < 20:
	// 	fmt.Println("The number is between 10 and 19.")
	// default:
	// 	fmt.Println("The number is 20 or greater.")
	// }

	// num := 2

	// switch {
	// case num > 1:
	// 	fmt.Println("The number is greater than 1.")
	// 	fallthrough
	// case num == 2:
	// 	fmt.Println("The number is equal to 2.")
	// default:
	// 	fmt.Println("Not Two")
	// }

	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("x is an int")
	case int32:
		fmt.Println("x is an int32")
	case float64:
		fmt.Println("x is a float64")
	case string:
		fmt.Println("x is a string")
	default:
		fmt.Println("x is of a different type")
	}
}