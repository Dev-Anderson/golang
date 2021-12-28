Conectando com o banco de dados postgres

**Pacotes utilizados:** 

> - database/sql
> - fmt
> - github.com/lib/pq

1- Criando uma constante para repassar as configurações do banco de dados: 

    const (
	    host = "localhost"
	    port = 5432
	    user  = "postgres"
	    password = "postgres"
	    dbname = "banco"
    )
2- Abrindo o banco de dados e repassando os dados da constante: 

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s" + "password=%s dbname=%s sslmode-disable", host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)

2.1- Caso apresente falhas na conexão com o banco de dados, será exibido a mensagem de falha: 

    if err != nil {
	    panic(err)
    }
2.2- Fechando a conexão com o banco de dados

    defer db.Close()
3- Se tudo ocorrer certo, será exibido a mensagem "Conectado com sucesso"