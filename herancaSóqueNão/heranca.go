package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade int
	altura float64
}

type estudante struct {
	pessoa		//aqui eu trouxe a struct pessoa para a struct estudante. aqui seria a "herança"
	curso string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Gean", "Barros", 25, 1.75}
	fmt.Println(p1)

	e1 := estudante{p1, "ADS", "Facimp"}
	fmt.Println(e1)
}