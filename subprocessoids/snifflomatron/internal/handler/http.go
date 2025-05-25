package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/EggLovej/Think-n-Link-FNDM/subprocessoids/snifflomatron/internal/client"
	"github.com/go-chi/chi/v5"
)

func GetStockData(w http.ResponseWriter, r *http.Request) {
	symbol := chi.URLParam(r, "symbol")
	log.Printf("Fetching stock data for: %s", symbol)

	data, err := client.FetchDailyTimeSeries(symbol)
	if err != nil {
		log.Printf("Error fetching stock data for %s: %v", symbol, err)
		http.Error(w, "Failed to fetch stock data", http.StatusInternalServerError)
		return
	}

	log.Printf("Fetched %d days of data for %s", len(data), symbol)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}
