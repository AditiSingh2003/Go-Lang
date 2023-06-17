package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main(){
	fmt.Println("Welcome to my module")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func greeter(){
	fmt.Println("Hello from my greeter module")
}

func serveHome(w http.ResponseWriter, r *http.Request){
	fmt.Println("Serving home")
	w.Write([]byte("<h1>Welcome to golang module</h1>"))

}