package main

import (
	"api-firebird/apis"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/v1/notaaprovada", apis.GetNotasAprovadas).Methods("GET")
	router.HandleFunc("/v1/notacancelada", apis.GetNotasCanceladas).Methods("GET")

	fmt.Println("API rodando na porta 3000")

	err := http.ListenAndServe(":3000", router)

	if err != nil {
		log.Println(err)
	}
}
