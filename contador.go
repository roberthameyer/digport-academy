package main

import "fmt"

func main() {
	var contador int

	for contador = 10; contador != 0; contador-- {
		fmt.Println("Contagem regressiva:", contador)
	}
	fmt.Println("A contagem terminou!")
}
