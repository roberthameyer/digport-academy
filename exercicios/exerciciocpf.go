package exercicios

import (
	"fmt"

	"github.com/Nhanderu/brdoc"
)

func IsCPF() bool {
	var cpf string
	fmt.Println("Olá!")
	fmt.Println("Digite o seu CPF:")
	fmt.Scanln(&cpf)

	if brdoc.IsCPF(cpf) {
		fmt.Printf("Validado com sucesso!")
		return true
	} else {
		fmt.Printf("A entrada não corresponde a um CPF 👎🏻")
		return false
	}
}
