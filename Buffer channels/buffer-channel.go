package main

import "fmt"

//Buffer vai especificar uma capacidade para seu Canal
func main() {
	fmt.Println("Canais com Buffer")

	canal := make(chan string, 2)
	canal <- "Olá Amigo!"
	canal <- "Hello  my Friend!"
	//Se for enviado mais uma mensagem vai dar deadlock
	//vai estourar a capacidade
	//canal <- "Deadlock!"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
