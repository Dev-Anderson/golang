package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func escreverTexto(linhas []string, caminhoDoArquivo string) error {
	arquivo, err := os.Create(caminhoDoArquivo)

	if err != nil {
		return err
	}

	defer arquivo.Close()

	escritor := bufio.NewWriter(arquivo)
	for _, linha := range linhas {
		fmt.Fprintln(escritor, linha)
	}

	return escritor.Flush()
}

func lerTexto(caminhoDoArquivo string) ([]string, error) {
	arquivo, err := os.Open(caminhoDoArquivo)

	if err != nil {
		return nil, err
	}

	defer arquivo.Close()

	var linhas []string
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}

	return linhas, scanner.Err()
}

func main() {
	var conteudo []string
	conteudo = append(conteudo, "Batatinha Frita 1")
	conteudo = append(conteudo, "Batatinha Frita 2")
	conteudo = append(conteudo, "Batatinha Frita 3")

	err := escreverTexto(conteudo, "c:/temp/batatinhafrita.txt")

	if err != nil {
		log.Fatalf("Erro:", err)
	}

	conteudo = nil
	conteudo, err = lerTexto("c:/temp/batatinhafrita.txt")

	if err != nil {
		log.Fatalf("Erro:", err)
	}

	for indice, linha := range conteudo {
		fmt.Println(indice, linha)
	}
}
