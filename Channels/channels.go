package main

import (
	"fmt"
	"time"
)

func main() {
	//Um canal pode enviar e receber dados
	//Usamos esse envio e recebimento pra sincronizar as Goroutines
	fmt.Println("Canais")
	canal := make(chan string)
	go escrever("Olá Mundo", canal)

	fmt.Println("Depois da função escrever ser executada")

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do programa")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
