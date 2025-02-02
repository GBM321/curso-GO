package main

import "fmt"


func main() {
	// duas formas para declarar variável
	// 1 forma:
	var variavel1 string = "variável 1"
	fmt.Println(variavel1)

	// 2 forma:
	variavel2 := "variável 2"
	fmt.Println(variavel2)

	// 3 forma:
	var (
		variavel3 string = "Samila"
		variavel4 string = "Mavie"
	)
	fmt.Println(variavel3, variavel4)
}