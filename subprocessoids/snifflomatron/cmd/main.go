package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/EggLovej/Think-n-Link-FNDM/subprocessoids/snifflomatron/internal/handler"
)

func main() {
	_ = godotenv.Load()

	r := chi.NewRouter()
	r.Get("/stocks/{symbol}", handler.GetStockData)
	r.Get("/number", handler.GetNumber)

	port := ":8080"
	log.Printf("Starting data-gatherer on %s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}