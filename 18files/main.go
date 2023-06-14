package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main()  {
	fmt.Println("Welcome to files in Golang")
	content := "This needs to go in a file - LearnCodeOnline.in"
	file , err := os.Create("./myFile.txt")
	if err != nil {
		fmt.Println("Error creating file")
		panic(err) // exit main function
	}
	// checkNilErr(err)
	length , err := io.WriteString(file, content)
	if err != nil {
		fmt.Println("Error writing to file")
		panic(err)
	}
	fmt.Println("length is: ",length)
	defer file.Close()
	readFile("./myFile.txt")

}

func readFile(filename string){
	hui, err :=ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("File contents:", string(hui))
}

func checkNilErr(err error){
	if(err != nil){
		panic(err)
	}
}