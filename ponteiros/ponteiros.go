package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int
	var ponteiro *int

  // o valor zero de um ponteiro é <nil>

  variavel3 = 100
  ponteiro = &variavel3

  fmt.Println(variavel3, *ponteiro) // * desreferenciação

	
}