package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20

	numeroInvertido := inverterSinal(numero)

	fmt.Println(numeroInvertido)

	//o valor de numero continua igual
	fmt.Println(numero)

	novoNumero := 40

	inverterSinalPonteiro(&novoNumero)
	fmt.Println(novoNumero)

}
