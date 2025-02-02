package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)


func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	//validação de email
	erro := checkmail.ValidateFormat("geanbarros13@gmail.com")
	fmt.Println(erro)
}