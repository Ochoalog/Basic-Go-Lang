package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint8
}

func main() {
	fmt.Println("Aula de Structs")

	var user usuario
	user.nome = "Vitor"
	user.idade = 20

	fmt.Println(user)

	endereco1 := endereco{
		"Rua do golang",
		85,
	}

	user2 := usuario{
		"Vitor",
		20,
		endereco1,
	}

	fmt.Println(user2)

	user3 := usuario{
		nome: "Vitor",
	}

	fmt.Println(user3)
}
