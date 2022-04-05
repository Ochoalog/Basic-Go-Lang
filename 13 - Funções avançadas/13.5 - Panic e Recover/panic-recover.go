package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Recuperando a execução")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()

	if media := (n1 + n2) / 2; media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A media é exatamente 6!")
}

func main() {
	alunoEstaAprovado(6, 6)

	fmt.Println("pos execução")
}
