package main

import "fmt"

func main() {
	var variavel1 string = "Váriavel 1"
	variavel2 := "Variável 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variavel3"
		variavel4 string = "Variavel4"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "Variavel5", "Variavel6"
	fmt.Println(variavel5, variavel6)

	const contante1 string = "Contante 1"
	fmt.Println(contante1)

	variavel5, variavel6 = variavel6, variavel5
}
