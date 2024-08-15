package main

import (
	"fmt"
)

func hello() {

	const hello = "Hello"
	name := []string{hello, "Ana", "Robertha", "Carol"}
	newName := name[2]

	fmt.Println(hello, ",", newName)

}
