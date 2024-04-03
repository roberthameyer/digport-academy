package main

import "fmt"

func main() {
	contatos := map[string]string{
		"Robertha": "Seriado",
		"Ana":      "Música",
		"João":     "Leitura",
		"Alice":    "Culinária",
		"André":    "Jardinagem",
	}

	for nome, hobby := range contatos {
		fmt.Println("O hobby de", nome, "é", hobby)
	}
}
