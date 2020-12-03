package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Goroutines")
	// Concorrência != Paralelismo

	go escrever("olá Mundo...") //goroutine
	escrever("Programando em Golang!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
