package apis

import (
	"api-firebird/config"
	"api-firebird/models"
	_struct "api-firebird/struct"
	"encoding/json"
	"net/http"
)

func GetNotasAprovadas(response http.ResponseWriter, r *http.Request) {
	
	db, err := config.Conexao()
	var Response _struct.ResponseData
	if err != nil {
		Response.Status = http.StatusInternalServerError
		Response.Message = err.Error()
		Response.Data = nil
		restponWithJson(response, http.StatusInternalServerError, Response)
	} else {
		_models := models.ModelGetData{DB: db}
		IsiData, err2 := _models.GetNotasAprovadas()
		if err2 != nil {
			Response.Status = http.StatusInternalServerError
			Response.Message = err2.Error()
			Response.Data = nil
			restponWithJson(response, http.StatusInternalServerError, Response)

		} else {
			Response.Status = http.StatusOK
			Response.Message = "Consulta todas as Notas Fiscais Aprovadas"
			Response.Data = IsiData
			restponWithJson(response, http.StatusOK, Response)

		}
	}
}

func GetNfeAprovada(response http.ResponseWriter, r *http.Request) {

	db, err := config.Conexao()
	var Response _struct.ResponseData
	if err != nil {
		Response.Status = http.StatusInternalServerError
		Response.Message = err.Error()
		Response.Data = nil
		restponWithJson(response, http.StatusInternalServerError, Response)
	} else {
		_models := models.ModelGetData{DB: db}
		IsiData, err2 := _models.GetNfeAprovada()
		if err2 != nil {
			Response.Status = http.StatusInternalServerError
			Response.Message = err2.Error()
			Response.Data = nil
			restponWithJson(response, http.StatusInternalServerError, Response)
		} else {
			Response.Status = http.StatusOK
			Response.Message = "Consulta apenas NF-e Aprovada"
			Response.Data = IsiData
			restponWithJson(response, http.StatusOK, Response)
		}
	}
}

func GetNfceAprovada(response http.ResponseWriter, r *http.Request) {
	db, err := config.Conexao()
	var Response _struct.ResponseData
	if err != nil {
		Response.Status = http.StatusInternalServerError
		Response.Message = err.Error()
		Response.Data = nil
		restponWithJson(response, http.StatusInternalServerError, Response)
	} else {
		_models := models.ModelGetData{DB: db}
		IsiData, err2 := _models.GetNfcAprovada()
		if err2 != nil {
			Response.Status = http.StatusInternalServerError
			Response.Message = err2.Error()
			Response.Data = nil
			restponWithJson(response, http.StatusInternalServerError, Response)
		} else {
			Response.Status = http.StatusOK
			Response.Message = "Consulta apenas NFC-e Aprovada"
			Response.Data = IsiData
			restponWithJson(response, http.StatusOK, Response)
		}
	}
}

func GetNotasCanceladas(response http.ResponseWriter, r *http.Request) {

	db, err := config.Conexao()

	var Response _struct.ResponseData

	if err != nil {
		Response.Status = http.StatusInternalServerError
		Response.Message = err.Error()
		Response.Data = nil
		restponWithJson(response, http.StatusInternalServerError, Response)
	} else {
		_models := models.ModelGetData{DB: db}
		IsiData, err2 := _models.GetNotasCanceladas()
		if err2 != nil {
			Response.Status = http.StatusInternalServerError
			Response.Message = err2.Error()
			Response.Data = nil
			restponWithJson(response, http.StatusInternalServerError, Response)
		} else {
			Response.Status = http.StatusOK
			Response.Message = "Consulta todas as Notas Fiscais Canceladas"
			Response.Data = IsiData
			restponWithJson(response, http.StatusOK, Response)
		}
	}
}

func restponWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
