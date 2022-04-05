package main

import "fmt"

func alunoEstaAprovado(n1, n2 float64) bool {
	defer fmt.Println("media calculada")
	media := n1 + n2/2

	if media > 6 {
		return true
	} else {
		return false
	}
}

func main() {
	alunoEstaAprovado(5, 7)
}
