package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RenderJSON(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	data, err := json.Marshal(response)
	if err != nil {
		RenderJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		fmt.Println("failed to write response:", err)
	}
}

type JSONErrorResponse struct {
	Message string `json:"message"`
}

func RenderJSONError(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	res := &JSONErrorResponse{Message: message}
	RenderJSON(w, res)
}
