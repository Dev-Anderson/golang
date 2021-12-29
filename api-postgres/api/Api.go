package api

import (
	"api/config"
	_struct "api/struct"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

func GetFilmes(w http.ResponseWriter, r *http.Request) {
	db := config.ConfigDB()

	printMessage("Consultando os filmes...")

	rows, err := db.Query("select * from filmes")

	if err != nil {
		panic(err)
	}

	var filmes []_struct.Filme

	for rows.Next() {
		var filmeID string
		var filmeNome string

		err = rows.Scan(&filmeID, &filmeNome)

		if err != nil {
			panic(err)
		}

		filmes = append(filmes, _struct.Filme{FilmeID: filmeID, FilmeNome: filmeNome})
	}

	var response = _struct.JsonResponse{Type: "sucesso", Data: filmes}

	json.NewEncoder(w).Encode(response)
}

func GetFilmeID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	db := config.ConfigDB()

	printMessage("Consultando o filme com o ID " + id)

	rows, err := db.Query("select * from filmes where filmeid = $1", id)

	if err != nil {
		panic(err)
	}

	var filmes []_struct.Filme

	for rows.Next() {
		var filmeID string
		var filmeNome string

		err = rows.Scan(&filmeID, &filmeNome)

		if err != nil {
			panic(err)
		}

		filmes = append(filmes, _struct.Filme{FilmeID: filmeID, FilmeNome: filmeNome})
	}

	var response = _struct.JsonResponse{Type: "sucesso", Message: "Consultando Filme", Data: filmes}

	json.NewEncoder(w).Encode(response)
}

func CreateFilmes(w http.ResponseWriter, r *http.Request) {
	filmeID := r.FormValue("filmeid")
	filmeNome := r.FormValue("filmenome")

	var response = _struct.JsonResponse{}

	if filmeID == "" || filmeNome == "" {
		response = _struct.JsonResponse{Type: "error", Message: "Informar os parâmetros filmeID e filmeNome"}
	} else {
		db := config.ConfigDB()

		printMessage("Inserindo o registro no banco de dados")

		fmt.Println("Inserindo o filme com o ID: " + filmeID + "com o nome: " + filmeNome)

		var lastInsertID int
		err := db.QueryRow("insert into filmes(filmeid, filmenome) values($1, $2) returning id;", filmeID, filmeNome).Scan(&lastInsertID)

		if err != nil {
			panic(err)
		}

		response = _struct.JsonResponse{Type: "sucesso", Message: "O Filme foi inserido com sucesso!"}
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteFilmes(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	filmeID := params["filmeid"]

	var response = _struct.JsonResponse{}

	if filmeID == "" {
		response = _struct.JsonResponse{Type: "erro", Message: "Não foi enviado o ID do filme"}
	} else {
		db := config.ConfigDB()

		printMessage("Deletando filme do banco de dados")

		_, err := db.Exec("delete from filmes where filmeid= $1", filmeID)

		if err != nil {
			panic(err)
		}

		response = _struct.JsonResponse{Type: "sucesso", Message: "O Filme foi deletado com sucesso!"}
	}

	json.NewEncoder(w).Encode(response)
}
