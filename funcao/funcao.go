package main

import "fmt"


func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculoMatematico(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(20, 30)
	fmt.Println(soma)

	var f = func ()  {
		fmt.Println("Função f")
	}

	f()

	resultadoSoma, resultadoSubtracao := calculoMatematico(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}
