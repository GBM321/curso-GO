package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := -1

	if numero > 14 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

		//atribuindo outro falor
		if outroNumero := numero; outroNumero > 0 {
			fmt.Println("Número é maior que zero")
		} else {
			fmt.Println("Número é menor que zero")
		}
}