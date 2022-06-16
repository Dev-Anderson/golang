package routes

import (
	"arquivos-fiscais/controllers"
	"arquivos-fiscais/data"
	"encoding/json"
	"net/http"
)

func GetCestAll(w http.ResponseWriter, r *http.Request) {
	controllers.GetCestAll()
	var response = data.ResponseData{Data: controllers.GetCestAll()}
	json.NewEncoder(w).Encode(response)
}
