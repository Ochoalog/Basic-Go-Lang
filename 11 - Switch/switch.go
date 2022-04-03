package main

import "fmt"

func selecionaDiaDaSemana(numero int) string {
	var diaDaSemana string
	switch numero {
	case 1:
		diaDaSemana = "Domingo"
	case 2:
		diaDaSemana = "Segunda"
	case 3:
		diaDaSemana = "Terça feira"
	case 4:
		diaDaSemana = "Quarta Feira"
	default:
		diaDaSemana = "Numero invalido"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := selecionaDiaDaSemana(2)

	fmt.Println(dia)
}
