package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Select")

	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		/* Usando o Select o canal 1 não precisa esperar os 2 segundos
		do canal 2 para ser executado */
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)

		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}

	}
}
