package repository

import (
	"dfe-admin/config"
	"dfe-admin/data"
	"fmt"
)

func GetPedido() []data.Pedido {
	db := config.ConnectDB()
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM PEDIDO`)

	if err != nil {
		panic(err)
	}

	var pedido []data.Pedido
	for rows.Next() {
		var Gid int
		
		err = rows.Scan(&Gid)

		if err != nil {
			panic(err)
		}

		pedido = append(pedido, data.Pedido{Gid: Gid})
	}

	return pedido
}

func CreatePedido(pedido data.Pedido) error {
	db := config.ConnectDB()
	defer db.Close()

	fmt.Println(pedido)

	insert := `INSERT INTO PEDIDO2 (GID) VALUES(?)`

	_, err := db.Exec(insert, pedido.Gid)

	if err != nil {
		panic(err)
	}

	return nil
}
