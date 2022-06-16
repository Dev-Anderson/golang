package routes

import (
	"fmt"
	"net/http"
)

func Ola(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Olá a API está respondendo corretamente")
}
