package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("For loop")
	i := 0

	for i < 10 {
		i++
		time.Sleep(time.Second)
		fmt.Println(i)
	}

	for j := 0; j < 10; j++ {

	}

	nomes := [3]string{"vitor", "mateus", "golang"}

	for index, nome := range nomes {
		fmt.Println(index, nome)
	}

	for _, letra := range "Palavra" {
		fmt.Println(string(letra))
	}

	usuario := map[string]string{
		"Vitor":     "Vitor",
		"sobrenome": "Ochoa",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
