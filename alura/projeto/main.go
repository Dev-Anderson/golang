package main

import "fmt"

func main() {
	nome := "Anderson"
	versao := 1.1

	fmt.Println("Olá sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - Iniciar o Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("O valor da variável é", comando)

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
