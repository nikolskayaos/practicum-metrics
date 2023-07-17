package main

import (
	"github.com/nikolskayaos/practicum-metrics/server/handlers"
	"github.com/nikolskayaos/practicum-metrics/server/storage"
	"net/http"
)

func main() {
	ms := storage.NewMemStorage()
	gh := handlers.NewGaugeHandler(ms)
	ch := handlers.NewCounterHandler(ms)
	uh := handlers.NewUpdateHandler(gh, ch)

	err := http.ListenAndServe(":8080", uh)
	if err != nil {
		panic(err)
	}
}
