package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Aula sobre herança.")

	p1 := pessoa{
		"Vitor",
		"Ochoa",
		20,
		160,
	}

	aluno := estudante{
		p1,
		"Engenharia",
		"UERGS",
	}

	fmt.Println(aluno)
}
