package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//conexao com o banco de dados
const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "filmes"
)

//estruturas
type Filme struct {
	FilmeID   string `json:filmeid`
	FilmeNome string `json:filmenome`
}

type JsonResponse struct {
	Type    string  `json:type`
	Data    []Filme `json:data`
	Message string  `json:message`
}

//conexao com o banco de dados
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

//funcao para checar as falhas
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//funcao para escrever mensagem logo depois das requesicoes
func printMensagem(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

func getFilmes(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMensagem("Consultando filmes...")

	rows, err := db.Query("select * from filmes order by id")

	checkErr(err)

	var filmes []Filme

	for rows.Next() {
		var id int
		var filmeID string
		var filmeNome string

		err = rows.Scan(&id, &filmeID, &filmeNome)

		checkErr(err)

		filmes = append(filmes, Filme{FilmeID: filmeID, FilmeNome: filmeNome})
	}

	var response = JsonResponse{Type: "sucesso", Data: filmes}

	json.NewEncoder(w).Encode(response)
}

func createFilmes(w http.ResponseWriter, r *http.Request) {
	filmeID := r.FormValue("filmeid")
	filmeNome := r.FormValue("filmenome")

	var response = JsonResponse{}

	if filmeID == "" || filmeNome == "" {
		response = JsonResponse{Type: "error", Message: "Falta informar os parâmetros filmeid e filmenome"}
	} else {
		db := setupDB()

		printMensagem("Inserindo o filme no banco de dados")

		fmt.Println("Inserindo o filme com o ID: " + filmeID + " com o nome: " + filmeNome)

		var lastInsertID int
		err := db.QueryRow("insert into filmes (filmeid, filmenome) values($1, $2) returning id;", filmeID, filmeNome).Scan(&lastInsertID)

		checkErr(err)

		response = JsonResponse{Type: "Sucesso", Message: "O filme foi inserido com sucesso"}
	}

	json.NewEncoder(w).Encode(response)
}

func deleteFilme(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	filmeID := params["filmeid"]

	var response = JsonResponse{}

	if filmeID == "" {
		response = JsonResponse{Type: "Erro", Message: "Não foi enviado o ID do Filme"}
	} else {
		db := setupDB()

		printMensagem("Deletando o filme no banco de dados")

		_, err := db.Exec("delete from filmes where filmeid= $1", filmeID)

		checkErr(err)

		response = JsonResponse{Type: "Sucesso", Message: "O Filme foi deletado com sucesso!"}
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/filmes/", getFilmes).Methods("GET")
	router.HandleFunc("/filmes/", createFilmes).Methods("POST")
	router.HandleFunc("/filmes/{filmeid}", deleteFilme).Methods("DELETE")

	fmt.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
