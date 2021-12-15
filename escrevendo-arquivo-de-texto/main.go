package main

import (
	"fmt"
	"os"
)

func main() {
	f := criandoArquivo("C:/temp/1.txt")
	defer fechandoArquivo(f)
	escrevendoArquivo(f)
}

func criandoArquivo(p string) *os.File {
	fmt.Println("Criando arquivo")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}

	return f
}

func escrevendoArquivo(f *os.File) {
	fmt.Println("Escrevendo dentro do arquivo")
}

func fechandoArquivo(f *os.File) {
	fmt.Println("Fechando o arquivo")
	err := f.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
