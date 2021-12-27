# Bem vindo ao meu aprendizado GO!

Go é uma linguagem de programação criada pela Google e lançada em código aberto em novembro de 2009. É uma linguagem compilada e focada em produtividade e programação concorrente. 

Fonte: https://pt.wikipedia.org/wiki/Go_(linguagem_de_programa%C3%A7%C3%A3o)

# Projetos

**Logo abaixo ficará um link para cada projeto que estou fazendo para aprender Go.**



## [Conectando com o banco de dados postgres](https://github.com/Dev-Anderson/golang/tree/main/go-postgres)

**Pacotes utilizados**

> - database/sql
> - fmt
> - github.com/lib/pq

**1 -** Criando uma constante para repassar os dados para acesso ao banco de dados

    const  (
    host =  "localhost"
    port =  5432
    user =  "postgres"
    password =  "postgres"
    dbname =  "api"    
    )

**2 -** Atribuindo os dados da constante para uma nova variável

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
**3 -** Abrindo o banco de dados postgres

    sql.Open("postgres", psqlInfo)
**4 -** Verificando se ocorrêu algum tipo de falha na hora da conexão com o banco de dados

    if err !=  nil  {
    panic(err)
    }
**5 -** Se tudo deu certo, será exibido uma mensagem no console

## [Escrevendo dentro de um arquivo de texto](https://github.com/Dev-Anderson/golang/tree/main/escrevendo-arquivo-de-texto)


**Pacotes utilizados**

> - fmt 
> - os

**1 -** Comando para criar o arquivo de texto

    os.Create()
No caso do exemplo foi criado um arquivo de texto dentro da seguinte pasta (c:\temp\1.txt)
Foi criado 3 funções 

 - Criar arquivo de texto 
 - Escrever dentro do arquivo de texto 
 - Fechar o arquivo de texto

Em todas as funções foi utilizado recursos da biblioteca padrão "os". 

## [Montando uma API simples](https://github.com/Dev-Anderson/golang/tree/main/api-sem-banco)

**Pacotes utilizados**

> - net/http
> - github.com/gin-gonic/gin

O pacote GIN é vastamente utilizado para montar APIs

**1 -** Criado uma estrutura com os seguintes campos: 
| Campo | Tipo |
|--|--|
| ID | int |
| Modelo | string |
| Marca | string |
| Valor | float64 |

Todos os campos são retornados no formato JSON

**2 -** Alimentando a estrutura, para isso foi criado o seguinte código: 

    var carro =[]veiculo{
       {ID: 1, Modelo: "x", Marca: "coisa", Valor: 12.22},
    }
**3 -**  Se vier um status de OK (200), então retorna os dados da variável "carro"
**4 -**   Criando a rota "/carro", e falando que ela vai receber a função "getCarro"
**5 -**  Iniciando a API na porta 3000

## [Conectando com o banco de dados Firebird](https://github.com/Dev-Anderson/golang/tree/main/go-firebird)

**Pacotes utilizados:**

> - database/sql
> - fmt
> - github.com/nakagami/firebirdsql

Foi utilizado a biblioteca "github.com/nakagami/firebirdsql" para fazer a conexão com o banco de dados. 

**1 -** Criar uma variável para colocar o resultado da consulta
**2 -** Utilizado o seguinte código para fazer a conexão com o banco de dados: 

    sql.Open("firebirdsql",  "SYSDBA:masterkey@localhost/ecosis/dados/econfe.eco")
**3 -** Utilizado o "defer" para fechar a conexão com o banco de dados depois de utilizado
**4 -** Repassado a consulta e depois atribuindo esse resultado para a variável
**5 -** Exibindo o resultado no console. 

## [API com o banco de dados Firebird](https://github.com/Dev-Anderson/golang/tree/main/api-firebird)

**Pacotes utilizados:**
> - encoding/json
> - net/http
> - fmt
> - log
> - database/sql
> - github.com/nakagami/firebirdsql
> - github.com/gorilla/mux

**1-** Criando a conexão com o banco de dados para isso utilizei  a biblioteca 'github.com/nakagami/firebirdsql', criado a seguinte função: 

    func  Conexao() (db *sql.DB, err error) {
    db, err = sql.Open("firebirdsql", "SYSDBA:masterkey@localhost/ecosis/dados/econfe.eco")
    return
    }
**2 -** Criando uma estrutura para receber os retornos do banco de dados, para esse tipo de consulta, utilizei os seguintes campos: 

    type  StructData  struct {
    Numero int64  `json:"numero"`
    Serie int64  `json:"serie"`
    Id string  `json:"id"`
    }

**3 -** Criar um 'model' para capturar os dados do banco de dados e depois repassando para estrutura que foi criado anteriormente, Criado uma função que faz o seguinte 'select' no banco de dados: 

    SELECT NFE_TBNFESIDE.NUMERO, NFE_TBNFESIDE.SERIE, NFE_TBNFES.ID FROM NFE_TBNFES INNER JOIN NFE_TBNFESIDE ON NFE_TBNFES.GID = NFE_TBNFESIDE.GID WHERE NFE_TBNFES.NUMEROPROTOCOLO IS NOT NULL AND NFE_TBNFES.GID = 1

Se o retorno foi com sucesso, então ele vai jogar o resultado dessa consulta dentro da estrutura criada, caso apresenta falha será exibido a falha no console. 
**4 -** Arquivo principal 'main.go', desse arquivo foi utilizado o 'mux' para criar uma rota do tipo 'GET' para consulta da nota fiscal. 
