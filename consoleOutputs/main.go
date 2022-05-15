package main

import "fmt"

func main() {
	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog"

	aNumber := 42
	isTrue := true
	//Println ekrana çıktı verir ve çıktıda yazılan metnin kaç karakter olduğunu int olarak dönüyor
	// burada stringLength Println de birleşip ekrana yazdırılan metnin kaç karakter olduğunu int olarak tutuyor
	stringLength, _ := fmt.Println(str1, str2, str3)

	fmt.Println("String length", stringLength)
	// %v placeholder olarak yani yer tutucu olarak iş görür. Buradaki aNumber değeri %v yerine gelecek
	fmt.Printf("Value of aNumber: %v\n", aNumber)
	fmt.Printf("Value of isTrue: %v\n", isTrue)
	fmt.Printf("Value of aNumber as float: %.2f\n", float64(aNumber))

	// %T yine yer tutucudur. Gönderilen parametrenin tipini tutar
	fmt.Printf("Data types: %T %T %T %T %T\n", str1, str2, str3, aNumber, isTrue)
	myString := fmt.Sprintf("Data types as var: %T %T %T %T %T", str1, str2, str3, aNumber, isTrue)

	fmt.Printf(myString)
}
