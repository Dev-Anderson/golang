package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"digital-bank/helpers"
	"digital-bank/vulnerableDB"

	"github.com/gorilla/mux"
)

type Login struct {
	Username string
	Password string
}

type Response struct {
	Data []vulnerableDB.User
}

type ErrResponse struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	login := vulnerableDB.VulnerableLogin(formattedBody.Username, formattedBody.Password)

	if len(login) > 0 {
		resp := Response{Data: login}
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func ola(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Olá teste com sucesso!")
}

func StartApi() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/login", ola).Methods("GET")
	fmt.Println("App is working on port :8888")
	log.Fatal(http.ListenAndServe(":8888", router))

}
