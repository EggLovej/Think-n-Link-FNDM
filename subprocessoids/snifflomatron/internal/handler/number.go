package handler

import (
	"encoding/json"
	"net/http"
)

func GetNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"number": 62})
}
