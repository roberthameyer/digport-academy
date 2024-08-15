package main

import "net/http"

func rotas() {
	rotas := Rotas()
	http.ListenAndServe(":8085", rotas)
}
