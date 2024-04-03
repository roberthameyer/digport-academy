package main

import "fmt"

var lista = map[string]int{"Batata": 3, "Cenoura": 2}

func main() {
	qtdRemovida, err := removerDaListaCompras()
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	} else {
		fmt.Println(qtdRemovida, "produtos foram removidos da lista")
	}
}

func removerDaListaCompras() (int, error) {
	produto := ""
	var removidos int
	if qtd, ok := lista[produto]; ok {
		removidos = qtd
		delete(lista, produto)
	} else {
		return 0, fmt.Errorf("O produto não foi encontrado na lista")
	}
	return removidos, nil
}
