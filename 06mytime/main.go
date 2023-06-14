package main

import (
	"fmt"
	"time"
)

func main (){

	fmt.Println("Welcome to My Time Zone")
	presentTime := time.Now()
	fmt.Println("The time now is", presentTime)

	// time package
	fmt.Println("The month is", presentTime.Month())

	// custom date
	fmt.Println("The date is", presentTime.Format("01-02-2006 15:04:05 Monday"))

	//create a date
	myBirthday := time.Date(2003, time.May, 1,15,30,0,0,time.UTC)
	fmt.Println("My birthday is", myBirthday)
	fmt.Println("My birthday is", myBirthday.Format("01-02-2006 Monday 15:04:05 PM "))

}