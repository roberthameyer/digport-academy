package main

import "fmt"

func main() {
	var numero int

	fmt.Println("Digite um número inteiro:")
	fmt.Scanf("%d", &numero)

	if numero > 0 {
		fmt.Println(numero, "é Positivo!")
	} else if numero < 0 {
		fmt.Println(numero, "é Negativo!")
	} else {
		fmt.Println(numero, "é Zero!")
	}
}
