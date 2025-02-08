package main

import "fmt"

type usuario struct {
	nome string
	idade int
	endereco endereco
}

type endereco struct {	//uma struct dentro de uma struct
	rua string
	numero int
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Gean"
	u.idade = 25
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua madalenna", 21}

	usuario2 := usuario{"Mavie", 1, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Samila"} //pra usar somente 1 valor
	fmt.Println(usuario3)
}