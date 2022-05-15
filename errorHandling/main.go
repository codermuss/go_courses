package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	//if there is an error err variable can't be empty. So err!nil expression give us error data
	if err != nil {
		fmt.Println("err.Error()")
	} else {
		fmt.Println(file.Name())
	}

	myError := errors.New("This is an error message!")
	fmt.Println(myError)
	_, errr := os.Open("abc.rar")
	checkError(errr)
}
func checkError(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
