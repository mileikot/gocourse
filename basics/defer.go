package basics

// import "fmt"

func deferSection() {
	process(1)
}

// func process(i int) {
// 	defer fmt.Println("defer i value", i)
// 	defer fmt.Println("First deferred execution of this function")
// 	defer fmt.Println("Secong deferred execution of this function")
// 	defer fmt.Println("Third deferred execution of this function")
// 	i++
// 	fmt.Println("Normal execusion of process function")
// 	fmt.Println("Final value of i:", i)
// }