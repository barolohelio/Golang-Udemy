package main

import "fmt"

type usuarioType struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
	bairro     string
	cidade     string
}

// Coleção de dados - vários campos r
func main() {
	fmt.Println("Arquivo structs")

	var usuario01 usuarioType
	fmt.Println(usuario01)
	usuario01.nome = "Hélio"
	usuario01.idade = 23
	fmt.Println(usuario01)

	usuario02 := usuarioType{
		"Guilherme",
		28,
		endereco{
			logradouro: "Oudezijds Voorburgwal",
			bairro:     "1012 GG",
			numero:     90,
		},
	}
	fmt.Println(usuario02)

	usuario03 := usuarioType{nome: "David Guetta"}
	fmt.Println(usuario03)
}
