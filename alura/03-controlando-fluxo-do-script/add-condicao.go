package main

import "fmt"

func main() {
	nome := "Anderson"
	versao := 1.1
	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa está na veresão", versao)

	fmt.Prinln("1 - Iniciar Monitoramento")
	fmt.Prinln("2 - Exibir Logs")
	fmt.Prinln("0 - Sair do Programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo Logs...")
	} else if comando == 3 {
		fmt.Println("Saindo do programa...")
	} else {
		fmt.Println("Não conheço este comando")
	}
}
