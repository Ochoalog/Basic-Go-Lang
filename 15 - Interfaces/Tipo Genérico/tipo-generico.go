package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

type pessoa struct {
	nome  string
	idade uint8
}

func main() {
	generica("Teste")
	generica(1)
	generica(true)

}
