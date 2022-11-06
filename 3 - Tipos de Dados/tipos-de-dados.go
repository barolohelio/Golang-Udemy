package main

import (
	"errors"
	"fmt"
)

// Somente com int, ele utiliza de acordo com arquitetura do computador
func main() {

	// NUMERO INTEIRO
	numero00 := -500000000

	var numero01 int = 1000000000000000
	fmt.Println(numero00)
	fmt.Println(numero01)

	// uint nao aceita negativo
	var numero02 uint32 = 1000
	fmt.Println(numero02)

	//alias
	// INT32 = RUNE
	var numero03 rune = 1234
	fmt.Println(numero03)

	//Byte = UINT8
	var numero04 byte = 123
	fmt.Println(numero04)
	// FIM NUMERO INTEIRO

	// NUMERO REAL
	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)

	var numeroReal01 float64 = 12345000000000.45
	fmt.Println(numeroReal01)

	numeroReal02 := 12345.67
	fmt.Println(numeroReal02)
	//FIM DE NUMERO REAL

	// STRING
	var str string = "Hello World!!"
	fmt.Println(str)

	str2 := "Γεια σου κόσμε!!"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)
	//FIM STRING

	var texto int16
	fmt.Println(texto)

	// Em Go todo tipo de dado tem o valor inical(Conhecido como valor 0)
	var booleano bool
	fmt.Println(booleano)

	var erro error
	fmt.Println(erro)

	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)
}
