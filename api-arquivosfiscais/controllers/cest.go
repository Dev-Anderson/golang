package controllers

import (
	"arquivos-fiscais/config"
	"arquivos-fiscais/data"
	"fmt"
)

func GetCestAll() []data.Cest {
	db := config.GetDb()
	defer db.Close()

	fmt.Println("Consulta de CEST")

	rows, err := db.Query(`
		select 
			id, 
			item, 
			cest, 
			ncm, 
			descricao, 
			ativo
		from cest`)
	if err != nil {
		panic(err)
	}

	var cest []data.Cest
	for rows.Next() {
		var Id int
		var Item string
		var Cest string
		var Ncm string
		var Descricao string
		var Ativo int

		err = rows.Scan(&Id, &Item, &Cest, &Ncm, &Descricao, &Ativo)

		if err != nil {
			panic(err)
		}

		cest = append(cest, data.Cest{Id: Id, Item: Item, Cest: Cest, Ncm: Ncm, Descricao: Descricao, Ativo: Ativo})
	}

	return cest
}
