package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Gean",
		"sobrenome": "Barros",
	}

	fmt.Println(usuario["nome"]) //pra acessar só o nome

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Mavie",
			"ultimo":   "Moreira",
		},
		"mae": {
			"primeiro": "Smila",
			"ultimo":   "Souza",
		},
	}

	fmt.Println(usuario2)
}
