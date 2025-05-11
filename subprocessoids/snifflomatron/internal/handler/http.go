package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/EggLovej/Think-n-Link-FNDM/subprocessoids/snifflomatron/internal/client"
)

func GetStockData(w http.ResponseWriter, r *http.Request) {
	symbol := chi.URLParam(r, "symbol")
	data, err := client.FetchDailyTimeSeries(symbol)
	if err != nil {
		http.Error(w, "Failed to fetch stock data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}