package main

import "fmt"

func main() {
	nome := "Anderson"
	idade := 27
	versao := 1.1

	fmt.Println("Ola sr.", nome, "sua idade é:", idade)
	fmt.Println("A versão dest programa é:", versao)

	fmt.Println("1 - Iniciar Monitoramente")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do programa")

	var comando int

	fmt.Scan("%d", &comando)

	fmt.Println("O comando digitado é:", comando)
}
