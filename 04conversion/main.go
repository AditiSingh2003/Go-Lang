package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main ()  {
	welcome := "Welcome to Pizza shop"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the rating for our pizza: ")

	// comma ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)

	// type conversion
	newRating, err := strconv.ParseFloat(strings.TrimSpace(input),64)
	if(err != nil){
		fmt.Println(err)
	} else {
		fmt.Println("Added new rating", newRating+1)
	}
}