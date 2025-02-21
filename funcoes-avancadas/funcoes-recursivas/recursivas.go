package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções recursivas")

	posicao := uint(10)
	fmt.Println(fibonacci(posicao))

	// 1 + 1 + 2 + 3 + 5 + 8 + 13 + 21  --> fibonacci
}

//função recursiva serve para chamar ela mesmo.
