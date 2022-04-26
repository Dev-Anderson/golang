package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "api"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	insertPedido := `insert into pedido (gid) values (133)`
	_, e := db.Exec(insertPedido)
	CheckError(e)

	insertPedido2 := `insert into pedido (gid) values($1)`
	_, e = db.Exec(insertPedido2, "333")
	CheckError(e)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
