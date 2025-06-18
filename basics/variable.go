package basics

import "fmt"

var middleName = "Cane"

func variable() {
	// var age int
	// var name string = "John"
	// var name1 = "Doe"

	// count := 10
	// lastName := "Smith"

	middleName := "Mayor"
	
	fmt.Println(middleName)

	// Default values
	// Numeric: 0
	// Boolean: false
	// String: ""
	// Pointer, slices, maps, functions and structs: nil

	// ---- SCOPE

	// fmt.Println(firstName)

}

func printname() {
	firstName := "Michael"
	fmt.Println(firstName)
}