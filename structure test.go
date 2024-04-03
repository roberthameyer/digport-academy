package main

import "fmt"

type Geolocation struct{
	latitude float64
	longitude float64
}

type Airport struct{
	name string
	code string
	address string
	avarageFlightsPerDay int
	location Geolocation
}

a := Airport "Aeroporto Internacional de São Paulo", "GRU", "Rod. Hélio Smidt, s/nº - Aeroporto, Guarulhos - SP, 07190-100, Brazil", 1651} 
fmt.Printf("%+v\n", a)