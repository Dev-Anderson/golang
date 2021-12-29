# API com banco de dados Postgres
**Pacotes utilizados:**

> - database/sql
> - fmt
> - encoding/json
> - log
> - github.com/gorilla/mux
> - github.com/lib/pq

1 - Criar a base de dados com o seguinte comando: 

    create table "filmes" (
	    "filmeid" int, 
	    "filmenome" varchar(255)
    );
2 - Logo depois de criar a base de dados, será necessário criar a estruturada dos dados do banco de dados, para isso informar o seguinte código: 

    type Filme struct {
	    FilmeID int `json:filmeid`
	    FilmeNome string `json:filmenome`
    }
3 - Além da estrutura da base de dados, será criado a estrutura do retorno de cada requisição

    type JsonResponse struct {
	    Type string `json:type`
	    Message string `json:message`
	    Data []Filme `json:data`
    }
Veja que na estrutura do resultado da requisição foi repassado o campo 'Data' e esse campo vai receber o que está dentro da estrutura 'Filme'.

4 - Com a estrutura criada, foi adicionado o arquivo 'config.go' para repassar a conexão com o banco de dados, e para conectar com o banco de dados foi feito o seguinte código: 

    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
Caso apresente falha será exibido a mensagem 

    if er != nil {
	    panic(err)
    }
    
    
5 - Dentro do arquivo 'Api.go' foi adicionado todas as rotas que foram disponibilizadas para a API

- **Consulta Filmes**

> localhost:3000/filmes

Nesta consulta será retornado a lista com todos os Filmes dentro do banco de dados. 

- **Consulta Filme por ID**
> localhost:3000/{id}

Nesta consulta será retornado apenas um Filme. 

*Não foi tratado, caso o filme não exista.*

- **Adicionado Filme**

Nesta Rota para conseguir adicionar um filme novo, será necessário repassar os seguintes parâmetros

    filmeid=
    filmenome=
Os dois campos são obrigatórios. 

- **Deletando filme**

> localhost:3000/filmes/{id}

Nesta rota para conseguir deletar um filme especifico, será necessário repassar o parâmetro 'filmeid'. 


