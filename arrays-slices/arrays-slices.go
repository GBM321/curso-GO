package main

import "fmt"


func main() {
	fmt.Println("Arrays e slices")

	var array1[5] string
	array1[0] = "posição 1"
	fmt.Println(array1)

	// outra forma de declarar array
	array2 := [5]string{"posição 1", "posição 2", "posição 3", "posição 4", "posição 5"}
	fmt.Println(array2)

	//slice
	slice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(slice)

	slice = append(slice, 16)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "posição alterada"
	fmt.Println(slice2)


	// Arrays Internos
	fmt.Println("---------")
	slice3 := make([] float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	slice4 := make([] float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 6) // adicionando mais um valor
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}

// Array --> lista de valores.
// Slaice --> é uma estrutura baseada no array, porém, com algumas flexibilidades, principalmente na questão do tamanho.
//função append --> adiciona um item no slice e retorna um slice novo com esse item novo adicionado. 