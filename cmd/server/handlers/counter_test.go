package handlers

import (
	"github.com/nikolskayaos/practicum-metrics/server/storage"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStoreCounterMetrics(t *testing.T) {
	type expectCounters struct {
		key string
		val int64
	}

	tests := []struct {
		name           string
		path           string
		counters       map[string]int64
		expectCounters expectCounters
	}{
		{
			name:     "empty storage",
			path:     "/pollCount/100",
			counters: make(map[string]int64),
			expectCounters: expectCounters{
				key: "pollCount",
				val: 100,
			},
		},
		{
			name:     "the metric is already in storage",
			path:     "/pollCount/45",
			counters: map[string]int64{"pollCount": 45},
			expectCounters: expectCounters{
				key: "pollCount",
				val: 90,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, test.path, nil)
			w := httptest.NewRecorder()
			ms := &storage.MemStorage{
				Counters: test.counters,
			}
			ch := NewCounterHandler(ms)
			ch.ServeHTTP(w, request)
			res, ok := ms.Counters[test.expectCounters.key]
			if !ok {
				t.Errorf("metric wasn't stored")
			}
			assert.Equal(t, test.expectCounters.val, res)

		})
	}
}
