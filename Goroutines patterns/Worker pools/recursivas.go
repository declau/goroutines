package main

import "fmt"

//Função recursiva que pega um número em uma determinada posição na Sequência de Fibonacci
func fibonacci(posicao uint) uint {
	//condição de parada
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	posicao := uint(3)
	fmt.Println(fibonacci(posicao))

	fmt.Println("-------------------------------")

	//Printa a sequencia de fibonacci até a posição final
	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}

}
