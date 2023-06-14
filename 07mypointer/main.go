package main

import "fmt"

func main(){
	fmt.Println("Welcome to a class on pointers")
	var myInt int
	fmt.Println("myInt is", myInt)
	myInt = 42
	fmt.Println("myInt is", myInt)
	var myIntPointer *int
	fmt.Println("myIntPointer is", myIntPointer)
	myIntPointer = &myInt  // & is the address of operator
	fmt.Println("myIntPointer is", myIntPointer)

	fmt.Println("The value of myIntPointer is", *myIntPointer) // * is the dereference operator
	*myIntPointer = 21 // * is the dereference operator
	fmt.Println("The value of myIntPointer is", *myIntPointer) // the value of myIntPointer is changed
	fmt.Println("The value of myInt is", myInt) // the value of myInt is changed
	
}