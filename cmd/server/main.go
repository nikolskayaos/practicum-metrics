package main

import (
	"github.com/nikolskayaos/practicum-metrics/internal"
	"net/http"
	"path"
	"strconv"
	"strings"
)

func main() {
	ms := internal.NewMemStorage()
	gh := NewGaugeHandler(ms)
	ch := NewCounterHandler(ms)
	uh := NewUpdateHandler(gh, ch)
	err := http.ListenAndServe(":8080", uh)
	if err != nil {
		panic(err)
	}
}

type UpdateHandler struct {
	GaugeHandler   *GaugeHandler
	CounterHandler *CounterHandler
}

func NewUpdateHandler(gh *GaugeHandler, ch *CounterHandler) *UpdateHandler {
	return &UpdateHandler{
		GaugeHandler:   gh,
		CounterHandler: ch,
	}
}

func (h *UpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head, metricType string
	head, r.URL.Path = ShiftPath(r.URL.Path)
	if head == "update" {
		metricType, r.URL.Path = ShiftPath(r.URL.Path)
		switch metricType {
		case "gauge":
			h.GaugeHandler.ServeHTTP(w, r)
		case "counter":
			h.CounterHandler.ServeHTTP(w, r)
		default:
			http.Error(w, "BadRequest", http.StatusBadRequest)
		}
		return
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}

type GaugeHandler struct {
	MemStorage *internal.MemStorage
}

func NewGaugeHandler(ms *internal.MemStorage) *GaugeHandler {
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
			return
		}
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}

type CounterHandler struct {
	MemStorage *internal.MemStorage
}

func NewCounterHandler(ms *internal.MemStorage) *CounterHandler {
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
			return
		}
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}

func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
