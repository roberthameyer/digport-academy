package main

import "fmt"

func main() {
	var num1 float64 = 12.67
	var num2 float64

	fmt.Println("Digite um número decimal:")
	fmt.Scanf("%v", &num2)

	fmt.Println("O resultado da soma é:", num1+num2)

}
