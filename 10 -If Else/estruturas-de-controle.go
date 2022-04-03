package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//não consigo acessar fora do if
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("outro numero é maior que zero")
	}
}
