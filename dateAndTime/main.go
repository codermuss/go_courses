package main

import (
	"fmt"
	"time"
)

func main() {
	//Console : Date & Time operations

	// We can create our date data using Date() method

	t := time.Date(2022, time.November, 10, 20, 0, 0, 0, time.UTC)
	fmt.Printf("Go Launch at %s\n", t)
	fmt.Println("-----------")

	//We can get live time information using time.Now() method
	now := time.Now()

	//We are putting this result into screen
	fmt.Printf("The time now is %s\n", now)
	fmt.Println("-----------")

	//We are showing the day,month, and year information from t we created first

	fmt.Println("The day is", t.Day())
	fmt.Println("The weekday is", t.Weekday())
	fmt.Println("The month is", t.Month())
	fmt.Println("The year is", t.Year())
	fmt.Println("--------------------")

	//Add 1 day to date
	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v %v %v %v\n", tomorrow.Weekday(), tomorrow.Day(), tomorrow.Month(), tomorrow.Year())
	fmt.Println("--------------------")

	// Determine a Format
	longFormat := "Monday, January 2, 2006"
	fmt.Println("Tomorrow is", tomorrow.Format(longFormat))

	fmt.Println("--------------------")

	shortFormat := "1/2/06"
	fmt.Printf("Tomorrow is", tomorrow.Format(shortFormat))

}
