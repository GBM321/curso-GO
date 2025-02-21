package main

import "fmt"


func main() {
	
	//FUNÇÃO ANÔNIMA
	retorno := func(texto string) string  {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parâmetro") //o segredo para funcao anonima é o ()
	fmt.Println(retorno)
}