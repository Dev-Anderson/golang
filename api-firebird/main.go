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

	router.HandleFunc("/restapi/consultanota", apis.GetNotaFiscal).Methods("GET")

	fmt.Println("API rodando na porta 3000")

	err := http.ListenAndServe(":3000", router)

	if err != nil {
		log.Println(err)
	}
}
