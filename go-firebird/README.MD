# Conexão com o banco de dados Firebird
**Pacotes utilizados:**

> - database/sql
> - fmt
> - github.com/nakagami/firebirdsql

1 - Dentro da função 'main' foi conectado com o banco de dados Firebird, repassando um string com a conexão com o banco, conforme código abaixo:

    var n string 
    conn, _ := sql.Open("firebirdsql", "SYSDBA:masterkey@localhost/ecosis/dados/econfe.eco")

2 - Fechando a conexão com o banco de dados, com o seguinte código: 

    defer conn.Close()

3- Realizando uma consulta no banco de dados

    conn.QueryRow("SELECT ID FROM NFE_TBNFES WHERE GID = 1").Scan(&n)"

4- Exibindo o resultado no console 

    fmt.prinln("A chave da nota fiscal", n)
