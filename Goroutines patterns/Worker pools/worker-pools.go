package main

import "fmt"

func main() {
	fmt.Println("Worker Pools")

	//Criando um canal com Buffer
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	//Quatro processos executando
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

//recebe: (dados <-chan int)
//envia: (dados chan<- int)
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

//Função recursiva que pega um número em uma determinada posição na Sequência de Fibonacci
func fibonacci(posicao int) int {
	//condição de parada
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
