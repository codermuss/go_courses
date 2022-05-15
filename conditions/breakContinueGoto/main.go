package main

import "fmt"

func main() {
	/*i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println("i value: ", i)
		i++

	}*/
	for i := 0; i < 7; i++ {
		if i == 3 {
			continue
		} else if i == 5 {
			break
		}
		fmt.Println("i value: ", i)

	}
}
