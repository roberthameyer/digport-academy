package main

import "fmt"

type person struct {
	name string
	work string
	age  int
	pay  float64
}

func main() {

	p1 := person{"Jordana", "Analista de Dados", 26, 7900.32}

	p2 := person{"Gabriele", "Programadora", 22, 8976.56}

	p3 := person{"Anne", "Support Engineer", 34, 7983.89}

	fmt.Println(p2)
}
