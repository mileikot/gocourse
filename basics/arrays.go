package basics

import "fmt"

func arrays() {

	// var arrayName [size]elementType
	var numbers [5]int
	fmt.Println("Initial array:", numbers)

	numbers[4] = 20
	fmt.Println("Updated array:", numbers)

	numbers[0] = 9
	fmt.Println("After setting first element:", numbers)

	fruits := [4]string{"Apple", "Banana", "Cherry", "Date"}
	fmt.Println("Fruits array:", fruits)

	fmt.Println("Third element in fruits:", fruits[2])

	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray // This creates a copy of the array

	copiedArray[0] = 100
	fmt.Println("Original Array:", originalArray)
	fmt.Println("Copied Array:", copiedArray)

	for i:=0; i < len(numbers); i++ {
		fmt.Printf("Element at index %d: %d\n", i, numbers[i])
	}

	// underscore is blank identifier, used to ignore a value
	for _, value := range numbers {
		fmt.Printf("Element: %d\n", value)
	}

	a, b := someFunction()
	fmt.Printf("Values returned from someFunction: a = %d, b = %d\n", a, b)

	c, _ := someFunction()
	fmt.Printf("Values returned from someFunction: c = %d", c)

	fmt.Println("The length of numbers array is:", len(numbers))

	//compare Arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{0, 2, 3}

	fmt.Println("Are array1 and array2 equal?", array1 == array2)

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matrix:", matrix)

	var copiedArray2 *[3]int
	copiedArray2 = &originalArray // This is how you can assign a pointer to an array
	copiedArray2[0] = 2000 // Modifying the original array through the pointer
	fmt.Println("Original array:", originalArray)
	fmt.Println("Pointer to original array:", copiedArray2)
}
func someFunction() (int, int) {
	return 1, 2
}