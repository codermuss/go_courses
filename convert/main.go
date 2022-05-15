package main

import (
	"fmt"
	"strconv"
)

func main() {
	var myString = "1"
	var x = 10
	var f float32 = 2.0
	fmt.Println(myString, x, f)

	//Stringten int'e dönüştürme

	number, _ := strconv.Atoi(myString)

	fmt.Println(number)

	result := number + 2
	fmt.Println(result)
	//Int to string
	str := strconv.Itoa(result)
	fmt.Println("Sonuç= " + str)

	//CASTING

	var i int = 55
	var fl float64 = float64(i)
	var u uint = uint(fl)

	fmt.Println(i, fl, u)

}
