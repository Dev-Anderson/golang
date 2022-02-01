package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
	Carro Carro
}

type Carro struct {
	Fabricante string
	Modelo     string
	Ano        int
}

func main() {

	vw := Carro{
		Fabricante: "VW",
		Modelo:     "Polo",
		Ano:        2007,
	}

	Anderson := Pessoa{
		Nome:  "Anderson",
		Idade: 31,
		Carro: vw,
	}

	fmt.Println(Anderson)
}
