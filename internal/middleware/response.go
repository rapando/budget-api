package middleware

import (
	"encoding/json"
	"net/http"
)

var UnauthorizedResponse = map[string]string{
	"message": "Unauthorized",
}

var ErrorResponse = map[string]string{
	"message": "An error occurred",
}

func Response(w http.ResponseWriter, code int, payload any) {
	var payloadJSON, _ = json.Marshal(payload)
	w.Header().Set("content-type", "application/json")
	w.Header().Set("connection", "close")
	w.WriteHeader(code)
	_, _ = w.Write(payloadJSON)

}
