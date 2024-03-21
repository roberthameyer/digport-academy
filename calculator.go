package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operacao string

	// Solicita os dois números ao usuário
	fmt.Println("Digite o primeiro número: ")
	fmt.Scanln(&num1)
	fmt.Println("Digite o segundo número: ")
	fmt.Scanln(&num2)

	// Solicita a operação ao usuário
	fmt.Println("Escolha a operação + para soma, - para subtração, * para multiplicação, / para divisão): ")
	fmt.Scanln(&operacao)

	// Realiza a operação escolhida e exibe o resultado
	switch operacao {
	case "+":
		fmt.Printf("Resultado: %v\n", num1+num2)
	case "-":
		fmt.Printf("Resultado: %v\n", num1-num2)
	case "*":
		fmt.Printf("Resultado: %v\n", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Resultado: %v\n", num1/num2)
		} else {
			fmt.Println("Erro: divisão por zero.")
		}
	default:
		fmt.Println("Operação inválida.")
	}
}
