package basics

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	// Structs, interfaces, enums
	//Eg. CalculateArea, UserInfo, NewHTTPRequest

	// snake_case
	//Eg. user_id, first_name, calculate_area

	// UPPERCASE
	// Constants
	// Eg. MAX_VALUE, DEFAULT_TIMEOUT, API_KEY

	// misedCase
	// Eg. javaScript, htmlDocument, isValid

	const MAXRETRIES = 5

	var employeeID = 10001
	fmt.Println("Employee ID:", employeeID)
}