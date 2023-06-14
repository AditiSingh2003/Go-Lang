package main

import (
	"fmt"
)

func main(){
	fmt.Println("Welcome to Functions in golang")
	greeter()

	result := adder(3,5)

	fmt.Println("Result is : ",result)

	proRes, myMess := proAdder(2,5,6,8,5,4)
	fmt.Println("Pro Result is: ", proRes)
	fmt.Println("Pro mess is" , myMess)
}

func greeter(){
	fmt.Println("Namstey from golang")
}

func adder (a int, b int) int{
	return a+b
}

func proAdder(values ...int) (int,string){
	total := 0
	for _,val := range values{
		total += val
	}
	return total, "hi pro res"
}