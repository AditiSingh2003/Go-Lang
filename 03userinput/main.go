package main

import (
	"bufio"
	"os"
)

func main() {

	welcome := "Welcome to my golang program"
	println(welcome)

	reader:= bufio.NewReader(os.Stdin)
	print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	println("Your name is: ", name)
}