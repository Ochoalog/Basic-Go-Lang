package main

import "fmt"

func main() {
	fmt.Println("Aula ponteiros")

	//PONTEIRO É UMA REFERENCIA DE MEMORIA
	var variavel1 int
	var ponteiro *int

	variavel1 = 100
	ponteiro = &variavel1

	fmt.Println(variavel1, *ponteiro)
}
