package main

import (
	"arquivos-fiscais/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/cest", routes.GetCestAll).Methods("GET")
	router.HandleFunc("/", routes.Ola).Methods("GET")

	fmt.Println("Api rodando na porta 3000")

	log.Fatal(http.ListenAndServe(":3000", router))
}
