package controllers

import (
	"encoding/json"
	"go-server/pkg/utils"
	"io"
	"net/http"
)

type RequestData struct {
	CardNumber string `json:"number"`
}
type ResponseData struct {
	Result bool `json:"result"`
}

func LunhCheck(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cant read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var requestData RequestData
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	valid := utils.Valid(requestData.CardNumber)
	responseData := ResponseData{Result: valid}

	jsonResponse, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "Cannot create response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(jsonResponse)
}
