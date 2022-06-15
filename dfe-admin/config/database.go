package config

import (
	"database/sql"
	"fmt"

	_ "github.com/nakagami/firebirdsql"
)

const (
	banco = "/ecosis/dados/ecodados.eco"
)

func ConnectDB() (db *sql.DB) {
	caminho := fmt.Sprintf("SYSDBA:masterkey@localhost%s", banco)
	db, err := sql.Open("firebirdsql", caminho)

	if err != nil {
		panic(err)
	}

	return db
}
