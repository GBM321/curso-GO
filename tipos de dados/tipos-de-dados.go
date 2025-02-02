package main

import (
	"errors"
	"fmt"
)


func main() {
	// TIPOS DE DADOS:
	// int --> números inteiros
	// float --> números fracionários, reais

	var numero int = 200
	fmt.Println(numero)

	var numero1 float64 = 1.14
	fmt.Println(numero1)

	// strings:
	var filha string = "Mavie"
	fmt.Println(filha)
	
	vida := "Mavie"
	fmt.Println(vida)

	char := 'M'
	fmt.Println(char)

	// bool true ou false:
	var booleano1 bool
	fmt.Println(booleano1)

	// error:
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}