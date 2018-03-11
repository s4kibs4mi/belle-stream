package api

import (
	"net/http"
	"encoding/json"
)

type response struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Error   error  `json:"error,omitempty"`
}

func (resp *response) ServeJSON(w http.ResponseWriter) {
	w.WriteHeader(resp.Status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
