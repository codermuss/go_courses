package main

import (
	"fmt"
	"os"
)

func main() {

	/*for _, env := range os.Environ() {
		fmt.Println(env)
	}*/

	uName := os.Getenv("USERNAME")

	fmt.Println(uName)
}

// isimlendirmede küçük harfle başlarsan private olur ve diğer dosyalardan erişilemez
// isimlendirmede büyük harfle başlarsan public olur ve diğer dosyalardan erişilebilir
var name string = "golang"
var Version string = "1.2.3"
