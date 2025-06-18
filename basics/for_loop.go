package basics

import "fmt"

func forLoop() {

	//Simple for loop
	// for i := 1; i<=5; i++ {
	// 	fmt.Println("Iteration:", i)
	// }

	// Iterating over a slice
	// numbers := [] int {1, 2, 3, 4, 5, 6}
	// for index, value := range numbers {
	// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// }

	// for i:=1; i<=10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("Odd number:", i)
	// 	if i== 5 {
	// 		break
	// 	}
	// }

	// Asterisk pattern printing
	// rows := 5

	// //Outer loop
	// for i := 1; i<=rows; i++ {
	// 	//Inner loop for spaces before stars
	// 	for j:=1; j<= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	//Inner loop for stars
	// 	for k:=1; k<=2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println() // New line after each row
	// }

	for i := range 10 {
		i++
		fmt.Println(10 - i)
	}
	fmt.Println("We have a lift off!")
}
