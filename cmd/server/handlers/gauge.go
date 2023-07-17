package handlers

import (
	"net/http"
	"strconv"
)

type GaugeHandler struct {
	MemStorage Repository
}

func NewGaugeHandler(ms Repository) *GaugeHandler {
	return &GaugeHandler{
		MemStorage: ms,
	}
}

func (g *GaugeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var metricName, gauge string

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST are allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		metricName, r.URL.Path = ShiftPath(r.URL.Path)
		if r.URL.Path != "/" {
			gauge, r.URL.Path = ShiftPath(r.URL.Path)
			t, err := strconv.ParseFloat(gauge, 64)
			if err != nil {
				http.Error(w, "BadRequest", http.StatusBadRequest)
				return
			}
			g.MemStorage.SaveGauges(metricName, t)
			w.Header().Set("content-type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			return
		}
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}
