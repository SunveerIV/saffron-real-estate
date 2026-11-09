package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorBody struct {
	Error ErrorDetails `json:"error"`
}

type ErrorDetails struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func JSON(
	w http.ResponseWriter,
	status int,
	value any,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if value == nil {
		return
	}

	err := json.NewEncoder(w).Encode(value)
	if err != nil {
		log.Printf("encode JSON response: %v", err)
	}
}

func Error(
	w http.ResponseWriter,
	status int,
	code string,
	message string,
) {
	JSON(w, status, ErrorBody{
		Error: ErrorDetails{
			Code:    code,
			Message: message,
		},
	})
}
