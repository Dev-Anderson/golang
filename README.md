# golang

Abaixo uma lista do que foi utilizado em cada projeto. 


## GO + Postgres 

Conectando com o banco de dados postgresSQL, para realizar esse procedimento foi seguido o seguinte passo a passo: 

1. Iniciar o 'go.mod' para isso utilizar o seguinte comando: 
```
go mod init "nome do projeto"
```
2. Download do pacote que faz a conexão com o banco de dados, rodar o seguinte comando: 
```
got get github.com/lib/pq
```
3. Criar o arquivo 'main.go' arquivo padrão de toda aplicação
4. Criar uma constante com os dados do banco de dados
5. Dentro da função 'main' repassar os valores da constante para abrir uma conexão com o banco de dados
6. Utilizado o 'panic' caso aconteça alguma falha o projeto será paralizado, se ocorrer tudo certo será exibido a mensagem: 
'Conexao com sucesso'


## GO + API Sem banco de dados

Neste projeto foi utilizado o GIN para criar uma API, foi seguindo o seguinte passo a passo neste projeto: 
1. Iniciar o 'go.mod'
2. Instalar o GIN, com o seguinte comando:
```
go get github.com/gin-gonic/gin
```
3. Criar uma estrutura que será o resultado do 'get'
4. Inserir os dados dentro da estrutura 
5. Montar uma consulta/get repassando o que foi inserido dentro da estrutura 
6. Dentro da função 'main' 
7. Criar a rota do 'get', repassar o GET que foi criado no passo 5
8. Iniciar a aplicação na porta 3000

### Resultado do GET
https://github.com/Dev-Anderson/golang/blob/main/api-sem-banco/Screenshot_2.png
