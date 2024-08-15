package main

import "fmt"

// Função que retorna uma string com base no valor do número
func verificarNumero(num int) string {
	if num > 0 {
		return "Greater than zero"
	} else if num == 0 {
		return "Zero"
	} else {
		return "Less than zero"
	}
}

// Função principal renomeada para 'executar'
func executar() {
	// Exemplos de uso da função verificarNumero
	numeros := []int{10, 0, -5}

	for _, num := range numeros {
		resultado := verificarNumero(num)
		fmt.Printf("O número %d é %s\n", num, resultado)
	}
}

// Função main padrão para execução do programa
func main() {
	executar()
}
