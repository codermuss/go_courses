package main

import "fmt"

func main() {
	/*
		a := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
		for i := range a {
			fmt.Printf("Array index: %v and value: %v\n", i, a[i])

		}

		for j := range a {
			fmt.Println("Array index: and value ", j, a[j])

		}*/

	//MAP
	/*capitalCity := map[string]string{"Turkey": "Istanbul", "France": "Paris", "Germany": "Berlin", "Italy": "Rome"}
	for key := range capitalCity {
		fmt.Println("Map item: Capital of", key, "is ", capitalCity[key])
	}*/

	capitalCity := map[string]string{"Turkey": "Istanbul", "France": "Paris", "Germany": "Berlin", "Italy": "Rome"}
	for key, value := range capitalCity {
		fmt.Println("Map item: Capital of", key, "is ", value)
	}
}
