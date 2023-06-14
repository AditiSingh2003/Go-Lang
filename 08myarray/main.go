package main

import "fmt"

func main ()  {
	fmt.Println("WElcome to array in GoLonag")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"
	// fruitList[2] = "Mango"

	fmt.Println("FruitList is", fruitList)
	fmt.Println("FriutList is", len(fruitList))

	var vegList = [5]string{"Potato", "Beans", "Onion"}
	fmt.Println("VegList is", vegList)
	fmt.Println("VegList is", len(vegList))

}