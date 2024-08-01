package pessoa

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

func ListaPessoa() []Pessoa {
	pessoas := []Pessoa{
		{Nome: "Robertha", Idade: 24},
		{Nome: "Gabriela", Idade: 22},
		{Nome: "Bruna", Idade: 20},
		{Nome: "Carol", Idade: 19},
		{Nome: "Larissa", Idade: 29},
		{Nome: "Julia", Idade: 12},
		{Nome: "Ana", Idade: 32},
		{Nome: "Rosana", Idade: 45},
	}
	return pessoas
}
