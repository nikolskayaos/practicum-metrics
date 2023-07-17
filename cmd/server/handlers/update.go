package handlers

import "net/http"

type Repository interface {
	SaveGauges(name string, gauge float64)
	SaveCounters(name string, count int64)
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
	http.Error(w, "Not Found", http.StatusBadRequest)
}
