package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Ola Mundo", canal)

	fmt.Println("Depois da função escrever")

	//for {
	//	mensagem, aberto := <-canal
	//	if !aberto {
	//		break
	//	}
	//	fmt.Println(mensagem)
	//}

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")
}

func escrever(text string, canal chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- text
		time.Sleep(time.Second)
	}

	close(canal)
}
