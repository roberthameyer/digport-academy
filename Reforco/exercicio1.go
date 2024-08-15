package main

import "fmt"

// Função que imprime a tabuada do número fornecido
func imprimirTabuada(num int) {
	for i := 1; i <= 10; i++ {
		resultado := i * num
		fmt.Printf("%d x %d = %d\n", i, num, resultado)
	}
}

func main() {
	// Exemplo de uso da função imprimirTabuada
	numero := 4
	imprimirTabuada(numero)
	fmt.Printf("Essa é a tabuada do %d\n", numero)
}
