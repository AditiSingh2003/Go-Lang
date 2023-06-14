package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("Welcome to slices in Golang")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Println("FruitList is", fruitList)
	fmt.Printf("Type of fruit List is %T \n", fruitList) 
	
	fruitList = append(fruitList, "mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highscores := make([]int, 4)
	highscores[0] = 234
	highscores[1] = 945
	highscores[2] = 465
	highscores[3] = 867
	// highscores[4] = 777
	highscores = append(highscores, 555, 666, 777)
	fmt.Println(highscores)

	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	// remove the element from slice
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby", "php"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}