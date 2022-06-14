package shared

import (
	"encoding/json"
	"net/http"
)

// JsonResponse function for Marshal and format response using Json
func BuildJSONResponse[T any](res http.ResponseWriter, resp T, httpCode int) {
	msg, _ := json.Marshal(resp)
	res.Header().Set("Content-Type", "application-json; charset=utf-8")
	res.WriteHeader(httpCode)
	res.Write(msg)
}

// EmptyJSON type
type EmptyJSON struct{}

// Response type
type Response[T any] struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}
