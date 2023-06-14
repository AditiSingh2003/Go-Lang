package main

import "fmt"

const Login string = "Aditi4" //constant variable

func main ()  {
	var username string = "Aditi"
	fmt.Println("Hello", username)
	fmt.Printf("The type of var is: %T \n", username)

	var smalluint uint16 = 255
	fmt.Println(smalluint)
	fmt.Printf("The type of var is: %T \n", smalluint)

	var smallfloat float32 = 374.904
	fmt.Println(smallfloat)
	fmt.Printf("The type of var is: %T \n", smallfloat)

	var istrue bool = false
	fmt.Println(istrue)
	fmt.Printf("The type of var is: %T \n", istrue)

	//no garbage value

	var small int 
	fmt.Println(small)

	var str string 
	fmt.Println(str)

	//no type conversion
	var small2 = 12
	fmt.Println(small2)

	var str2 = "Aditi2"
	fmt.Println(str2)

	//no var
	small3 := 13  //:= is a valoruous operator only used inside the function main
	fmt.Println(small3)

	str3 := "Aditi3"
	fmt.Println(str3)

	fmt.Println(Login)

}