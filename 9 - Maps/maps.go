package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Vitor",
		"Sobrenome": "Ochoa",
	}

	fmt.Println(usuario)

	fmt.Println(usuario["nome"])

	user := map[string]map[string]string{
		"nomeMap": {
			"nome":      "Vitor",
			"sobrenome": "Ochoa",
		},
	}

	fmt.Println(user)

	delete(user, "nomeMap")
}
