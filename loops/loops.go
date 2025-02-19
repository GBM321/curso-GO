package main

import (
	"fmt"
	//"time"
)

func main() {
	fmt.Println("Loops")

	//i := 0

	//for i < 10 {
	//	i++
	//	fmt.Println("Incrementando i", i)
	//	time.Sleep(time.Second)
	//}

	//fmt.Println(i)

	//for j := 0; j < 10; j++ {
	//	fmt.Println("Incrementando j", j)
	//	time.Sleep(time.Second)
	//}

	nomes := [3]string{"Gean", "Samila", "Mavie"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome": "Gean",
		"sobrenome": "Barros",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
