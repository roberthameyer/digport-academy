package main

import "fmt"

func name() {
	var nome string

	fmt.Println("Qual o seu nome?")
	fmt.Scanf("%s", &nome)
	fmt.Println("Bem-vindo,", nome)

}
