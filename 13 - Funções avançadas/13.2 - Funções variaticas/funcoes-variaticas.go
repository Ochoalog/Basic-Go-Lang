package main

import "fmt"

//soma n numeros
func somaNNumeros(numeros ...int) (soma int) {
	for _, numero := range numeros {
		soma += numero
	}

	return
}

func main() {
	fmt.Println(somaNNumeros(10, 20, 30, 40, 5050))
}
