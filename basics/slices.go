package basics

import (
	"fmt"
	"slices"
)

func slicesFunc() {

	// var sliceName []elementType
	// var numbers []int
	// var numbers1 = []int{1, 2, 3, 4, 5}

	// numbers2 := []int{6, 7, 8, 9, 10}

	// slice := make([]int, 5) // Creates a slice of length 5 with zero values

	a := [5]int{1, 2, 3, 4, 5}
	slice1 := a[1:4] // Creates a slice from index 1 to 3 (4 is exclusive)
	fmt.Println("Slice from array:", slice1)

	slice1 = append(slice1, 6, 7) // Appending elements to the slice
	fmt.Println("Slice after appending:", slice1)

	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1) // Copying the slice
	fmt.Println("Copied slice:", sliceCopy)

	// var nilSlice []int

	for i, v := range slice1 {
		fmt.Printf("Element at index %d: %d\n", i, v)
	}

	fmt.Println("Element at index 2 in slice1:", slice1[2])

	// slice1[3] = 50
	// fmt.Println("Slice after modifying index 3:", slice1)

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("slice1 and sliceCopy are equal")
	} else {
		fmt.Println("slice1 and sliceCopy are not equal")
	}

	twoD := make([][]int, 3) // Creating a 2D slice
	for i := 0; i<3; i++ {
		innerLength := i + 1
		twoD[i] = make([]int, innerLength) // Each inner slice has a different length
		for j := 0; j< innerLength; j++ {
			twoD[i][j] = i + j // Assigning values to the 2D slice
		}
	}

	fmt.Println("2D Slice:", twoD)

	// slice[low:high]
	slice2 := slice1[2:4]
	fmt.Println("Slice from index 2 to 4:", slice2)

	fmt.Println("The capacity of slice2 is:", cap(slice2))
}