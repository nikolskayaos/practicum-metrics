package handlers

import (
	"github.com/nikolskayaos/practicum-metrics/cmd/server/storage"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStoreGaugeMetrics(t *testing.T) {
	type expectGauges struct {
		key string
		val float64
	}

	tests := []struct {
		name         string
		path         string
		gauges       map[string]float64
		expectGauges expectGauges
	}{
		{
			name:         "empty storage",
			path:         "/Alloc/111.54",
			gauges:       make(map[string]float64),
			expectGauges: expectGauges{key: "Alloc", val: 111.54},
		},
		{
			name:         "the metric is already in storage",
			path:         "/Alloc/12",
			gauges:       map[string]float64{"Alloc": 111.54},
			expectGauges: expectGauges{key: "Alloc", val: 12},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, test.path, nil)
			w := httptest.NewRecorder()
			ms := &storage.MemStorage{
				Gauges: test.gauges,
			}
			gh := NewGaugeHandler(ms)

			gh.ServeHTTP(w, request)

			res, ok := ms.Gauges[test.expectGauges.key]
			if !ok {
				t.Errorf("metric wasn't stored")
			}
			assert.Equal(t, test.expectGauges.val, res)
		})
	}
}
