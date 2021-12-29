package main

import (
	"api/api"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/filmes", api.GetFilmes).Methods("GET")
	router.HandleFunc("/filmes/{id}", api.GetFilmeID).Methods("GET")
	router.HandleFunc("/filmes", api.CreateFilmes).Methods("POST")
	router.HandleFunc("/filmes", api.DeleteFilmes).Methods("DELETE")

	fmt.Println("Servidor rodando na porta 3000")

	log.Fatal(http.ListenAndServe(":3000", router))
}
