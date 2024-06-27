package exercicios

import "fmt"

func Despesas() {
	// Cria a lista de despesas
	contasapagar := []string{ mercado, aluguel, eletricidade, condomínio, faculdade}
}

// Loop para imprimir todas as despesas
fmt.Println("Aqui estão todas as despesas:", contasapagar)
for _, contasapagar := range contasapagar {
	fmt.Println("-", contasapagar)
}

// Verificar se um item específico está na lista
var item string
fmt.Println("\nDigite o item que deseja verificar na lista de despesas:")
fmt.Scanln(&item)
encontrado := false
for _, despesa := range contasapagar {
	if despesa == item {
		encontrado = true
		break
	}
}
if encontrado {
	fmt.Println(item, "está na lista de despesas.")
} else {
	fmt.Println(item, "não está na lista de despesas.")
}

// Formatar uma string para exibir o total de itens na lista
fmt.Printf("\nTotal de itens na lista de despesas: %d\n", len(despesas))
}