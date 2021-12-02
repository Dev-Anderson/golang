# API + Banco de dados Postgres

## Rotas API 

> getFilmes 

Consulta todos os filmes no banco de dados 

> createFilme

Criar um filme no banco de dados, para isso será necessário repssar os seguintes parâmetros: 

```
filmeid = 
filmenome = 
```
Foi criado uma validação, caso não tenha informado os dois parâmetros é apresentando uma mensagem solicitando informar os dois parâmetros. 

> deleteFilme 

Rota para deletar um filme, neste rota é obrigatório infomrar o ID do Filme que será deletado. 

## O que foi usado

Neste projeto foi utilizado as seguintes bibliotecas

> mux
> pq
> log
> json
> fmt
> database/sql


ADICIONADO IMAGENS DOS RESULTADOS DA ROTA