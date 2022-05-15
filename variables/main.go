package main

import "fmt"

// scope dışında değişken tanımlama
var sayi = 4

//Globalda böyle bir tanımlama yapılamaz. Bu tanımlama sadece fonksiyon içinde geçerli
//sayi:=4

//const değeri sonradan değiştirilemez bu yüzden başlangıçta değer ataması gereklidir.
const aciklama = "Go app"
const pi = 3.14

func main() {

	/*
		var message string
		message = "Hello, world!"

		var message string = "Hello, world!"
		Verinin tipini bilmiyorsan alttaki gibi tanımlayabilirsin.
		var message= "Hello, world!"


		var a, b, c int = 3, 4, 5


		var message = "Hello, world!"
		var a, b, c = 3, true, 4.5

		var c int
		var k,o string="abc","xyz"

		a := "Hello, world!"
		b := 'M'
		c := "Metin"

	*/

	var myFloat32 float32 = 44.
	myComplex := complex(3, 4)

	fmt.Println(myFloat32)
	println(myComplex)

}
