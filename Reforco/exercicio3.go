package main

import "fmt"

// Função que realiza operações matemáticas e imprime o resultado com base no valor do número
func processarNumero(num int) {
	if num > 0 {
		resultado := num + 10
		fmt.Printf("Número %d é maior que zero. Resultado: %d\n", num, resultado)
	} else if num == 0 {
		resultado := num + 2
		fmt.Printf("Número %d é zero. Resultado: %d\n", num, resultado)
	} else {
		resultado := num + 23
		fmt.Printf("Número %d é menor que zero. Resultado: %d\n", num, resultado)
	}
}

// Função para processar uma lista de números
func processarLista(numeros []int) {
	for _, num := range numeros {
		processarNumero(num)
	}
}

// Função principal
func main() {
	lista := []int{10, 0, -5, 7, -3}
	processarLista(lista)
}
