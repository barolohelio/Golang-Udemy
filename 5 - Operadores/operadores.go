package main

import "fmt"

func main() {
	//Aritméticos
	soma := 9713 + 3179
	subtracao := 33333 - 456
	divisao := 10 / 5
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 3

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	resultado := numero1 + numero2
	fmt.Println(resultado)

	// FIM DOS ARITMÉTICOS

	// ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2" //Usando Inferencia de tipos

	fmt.Println(variavel1, variavel2)
	//FIM DOS OPERADORES DE ATRIBUIÇÃO

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(2 >= 2)

	fmt.Println(1 == 2)

	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)

	fmt.Println(1 != 2)
	//FIM DOS OPERADORES RELACIONAIS

	//OPERADORES LÓGICOS
	fmt.Println("---------------//-------------------//------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && true)  // AND
	fmt.Println(falso && false)      // AND
	fmt.Println(falso && verdadeiro) // AND

	fmt.Println(verdadeiro || falso) // OR

	fmt.Println(!verdadeiro) //NEGAÇÃO
	fmt.Println(!falso)      //NEGAÇÃO
	//FIM DOS OPERADORES LÓGICOS

	fmt.Println("---------------//-------------------//------------------")
	//OPERADORES UNÁRIOS
	numero := 0
	numero++
	numero++
	fmt.Println(numero)

	numero += 8
	fmt.Println(numero)

	numero--
	numero--
	numero--
	fmt.Println(numero)

	numero -= 20
	fmt.Println(numero)

	fmt.Println("~~~~~~~~~~~~~~//~~~~~~~~~~~~~~~~~~~~~~~~~//~~~~~~~~~~~~")
	numero *= 3
	fmt.Println(numero)
	numero *= 10
	fmt.Println(numero)
	numero %= 3
	fmt.Println(numero)
	fmt.Println("~~~~~~~~~~~~~~//~~~~~~~~~~~~~~~~~~~~~~~~~//~~~~~~~~~~~~")
	//FIM DOS OPERADORES UNÁRIOS
}
