package main

import "fmt"

func somar(a int8, b int8) int8 {
	return a + b
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := somar(10, 20)

	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println("funcao f")
		return txt
	}

	resultado := f("teste")

	fmt.Println(resultado)

	resultadoSoma, _ := calculosMatematicos(10, 20)

	fmt.Println(resultadoSoma)
}
