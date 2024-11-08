package library

import (
	"encoding/json"
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

func Response400(w http.ResponseWriter, message string) {
	badResponse := model.Response{
		Success:    false,
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}

	response, _ := json.MarshalIndent(badResponse, "", " ")
	jsonResponse, _ := w.Write(response)
	json.NewEncoder(w).Encode(jsonResponse)
}

func Response404(w http.ResponseWriter, message string) {
	badResponse := model.Response{
		Success:    false,
		StatusCode: 404,
		Message:    message,
	}

	response, _ := json.MarshalIndent(badResponse, "", " ")
	jsonResponse, _ := w.Write(response)
	json.NewEncoder(w).Encode(jsonResponse)
}

func Response500(w http.ResponseWriter, message string) {
	badResponse := model.Response{
		Success:    false,
		StatusCode: 500,
		Message:    message,
	}

	response, _ := json.MarshalIndent(badResponse, "", " ")
	jsonResponse, _ := w.Write(response)
	json.NewEncoder(w).Encode(jsonResponse)
}

func Response200(w http.ResponseWriter, data any) {
	successResponse := model.Response{
		Success:    true,
		StatusCode: 200,
		Data:       data,
	}
	response, _ := json.MarshalIndent(successResponse, "", " ")
	jsonResponse, _ := w.Write(response)
	json.NewEncoder(w).Encode(jsonResponse)
}

func ResponseDelete200(w http.ResponseWriter, message string) {
	successResponse := model.Response{
		Success:    true,
		StatusCode: 200,
		Message:    message,
	}
	response, _ := json.MarshalIndent(successResponse, "", " ")
	jsonResponse, _ := w.Write(response)
	json.NewEncoder(w).Encode(jsonResponse)
}

func Response201(w http.ResponseWriter, message string, data any) {
	successResponse := model.Response{
		Success:    true,
		StatusCode: 201,
		Message:    message,
		Data:       data,
	}
	response, _ := json.MarshalIndent(successResponse, "", " ")
	jsonResponse, _ := w.Write(response)
	json.NewEncoder(w).Encode(jsonResponse)
}
