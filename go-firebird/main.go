package main

import (
	"database/sql"
	"fmt"

	_ "github.com/nakagami/firebirdsql"
)

func main() {
	var n string
	conn, _ := sql.Open("firebirdsql", "SYSDBA:masterkey@localhost/ecosis/dados/econfe.eco")
	defer conn.Close()
	conn.QueryRow("SELECT ID FROM NFE_TBNFES WHERE GID = 1").Scan(&n)
	fmt.Println("A chave de acesso da nota fiscal é: ", n)
}
