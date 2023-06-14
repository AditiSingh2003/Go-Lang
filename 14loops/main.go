package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang")

	// days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i:= range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	rougueValue := 1
	for rougueValue < 10 {
		fmt.Println("Value is", rougueValue)
		rougueValue++
		if rougueValue == 5 {
			rougueValue++
			goto lco
		} else if rougueValue == 7 {
			break
		}
	}

lco:
	fmt.Println("Jumping to LCO")
}