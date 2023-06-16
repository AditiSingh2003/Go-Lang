package main

import (
	"encoding/json"
	"fmt"
)

type course struct{
	Name string `json:"coursename"`
	Price int `json:"price"`
	Platform string `json:"platform"`
	password string `json:"-"`
	tags []string `json:"tags, omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson(){
	//to store and print the json data in a proper format
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "Youtube", "abc1", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "Youtube", "abc2", []string{"web-dev", "js", "react"}},
		{"Angular Bootcamp", 399, "Youtube", "abc3", nil},
	}
	finalJson , err := json.MarshalIndent(lcoCourses,"","\t")
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson(){
	//
	jsonData := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"price": 299,
		"platform": "Youtube",
		"tags": [
			"web-dev",
			"js"
		]
	}
	`)
	var courses course
	chackValid := json.Valid(jsonData)
	if chackValid{
		fmt.Println("JSON was valid")
	}else{
		fmt.Println("JSON was not valid")
	}
	check := json.Unmarshal(jsonData, &courses)
	if check!=nil{
		panic(check)
	}
	fmt.Printf("Course: %#v\n", courses)

	// some cases where we just want to add data to the key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData, &myOnlineData)
	fmt.Printf("Data: %#v\n", myOnlineData)

	for k,v := range myOnlineData{
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v,v)
	}
}