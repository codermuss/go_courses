package main

import "fmt"

func main() {
	/*myArray := [...]int{45, 23, 43}
	mySlice1 := myArray[:]
	fmt.Println(mySlice1)
	mySlice1[0] = 11
	fmt.Println(mySlice1)
	fmt.Println(myArray)*/

	/*
		//
		var colors = []string{"Red", "Green", "Blue"}
		fmt.Println(colors)
		colors = append(colors, "Purple")
		fmt.Println(colors)
		colors = append(colors[1:len(colors)])
		fmt.Println(colors)
		colors = append(colors[:len(colors)-1])
		fmt.Println(colors)
	*/
	numbers := make([]int, 5, 5)
	numbers[0] = 0
	numbers[1] = 1
	numbers[2] = 2
	numbers[3] = 3
	numbers[4] = 4
	fmt.Println(numbers)
	numbers = append(numbers, 342)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))
	fmt.Println(len(numbers))

}
