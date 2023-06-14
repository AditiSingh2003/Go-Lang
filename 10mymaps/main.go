package main

import "fmt"

func main ()  {
	fmt.Println("Welcome to maps in Golang")

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"
	 fmt.Println(languages)

	delete(languages, "RB")
	fmt.Println(languages)

	// Iterating over maps
	for key, value := range languages {
		fmt.Printf("For key %v the value is %v \n", key, value)
	}

}