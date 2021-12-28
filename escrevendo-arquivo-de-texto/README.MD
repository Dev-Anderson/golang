# Escrevendo arquivo de texto

**Pacotes utilizados:** 

> - bufio
> - fmt
> - log
> - os

1- Criando uma função para ler o arquivo de texto
1.1- Abrindo o arquivo

    arquivo, err := os.Open(caminhoDoArquivo)

1.2- Retornando falha 

    if err != nil {
	    return nil, err
    }
1.3- Se apresentar será fechado o arquivo

    defer arquivo.close()

1.4- Criando um 'scanner' que lê cada linha do arquivo

    var linhas []string
    scanner := bufio.NewScanner(arquivo)
    for scanner.Scan() {
	    linhas = append(linhas, scanner.Text())
    }
1.5- Retornado as linhas lidas e uma falha se ocorrer erro na leitura do arquivo

    return linhas, scanner.Err() 

2- Criando função para escrever dentro do arquivo de texto
2.1- Criando o arquivo de texto

    arquivo, err := os.Create(caminhoDoArquivo)
2.2- Caso apresente falhas, será retornado a falha

    if err != nil {
	    return err
    }
    defer arquivo.Close()
2.3- Criando um escritor responsável por escrever cada linha do 'slice' no arquivo de texto

    escritor := bufio.NewWriter(arquivo)
    for _, linha := range linhas{
	    fmt.Fprintln(escritor, linha)
    }
2.4- Caso a função 'flush' retornar um erro ele será retorno aqui também

    return escritor.Flush()
3- Dentro da função 'main' foi criado uma variável com o nome 'conteudo' para escrever o que deve estar dentro do arquivo de texto 
4- Logo depois repassado o caminho do arquivo que será escrito
5- Exibindo as falhas caso seja apresentando