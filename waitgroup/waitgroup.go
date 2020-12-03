package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Wait Group")

	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	//goroutine Wait Group
	go func() {
		escrever("olá Mundo...")
		waitGroup.Done() //quando faz o done -1 no add
	}()
	go func() {
		escrever("Programando em Golang!")
		waitGroup.Done() //quando faz o done -1 no add
	}()
	go func() {
		escrever("Goroutine 3")
		waitGroup.Done() //quando faz o done -1 no add
	}()
	go func() {
		escrever("Goroutine 4")
		waitGroup.Done() //quando faz o done -1 no add
	}()

	waitGroup.Wait() //Espera a contagem das goroutines zerar

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
