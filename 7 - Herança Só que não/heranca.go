package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa    // Pegando todos os campos de pessoa e colocando em estudante
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	pessoa00 := pessoa{"Vinicius", "Junior", 22, 176}
	fmt.Println(pessoa00)

	estudante00 := estudante{pessoa00, "Arte do Futebol", "Real Madrid"}
	fmt.Println(estudante00)
}
