package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u *usuario) setNewIdade(newIdade int) {
	u.idade = uint8(newIdade)
}

func main() {
	usuario := usuario{"Vitor", 20}

	fmt.Println(usuario)
	usuario.setNewIdade(25)

	fmt.Println(usuario)
}
