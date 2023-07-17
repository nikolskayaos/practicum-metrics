package handlers

import (
	"net/http"
	"strconv"
)

type CounterHandler struct {
	MemStorage Repository
}

func NewCounterHandler(ms Repository) *CounterHandler {
	return &CounterHandler{
		MemStorage: ms,
	}
}

func (c *CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var metricName, counter string

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST are allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		metricName, r.URL.Path = ShiftPath(r.URL.Path)
		if r.URL.Path != "/" {
			counter, r.URL.Path = ShiftPath(r.URL.Path)
			t, err := strconv.ParseInt(counter, 10, 64)
			if err != nil {
				http.Error(w, "BadRequest", http.StatusBadRequest)
				return
			}
			c.MemStorage.SaveCounters(metricName, t)
			w.Header().Set("content-type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			return
		}
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}
