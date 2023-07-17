package handlers

import (
	"github.com/nikolskayaos/practicum-metrics/server/storage"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateHandler(t *testing.T) {
	type expect struct {
		code        int
		response    string
		contentType string
	}

	tests := []struct {
		name   string
		path   string
		method string
		expect
	}{
		{
			name:   "",
			path:   "/update/counter/pollCount/100",
			method: http.MethodPost,
			expect: expect{
				code:        http.StatusOK,
				response:    "",
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:   "",
			path:   "/update/gauge/Alloc/111.54",
			method: http.MethodPost,
			expect: expect{
				code:        http.StatusOK,
				response:    "",
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:   "",
			path:   "/update/",
			method: http.MethodPost,
			expect: expect{
				code:        http.StatusBadRequest,
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:   "",
			path:   "/update/gauge",
			method: http.MethodPost,
			expect: expect{
				code:        http.StatusNotFound,
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:   "",
			path:   "/update/counter",
			method: http.MethodPost,
			expect: expect{
				code:        http.StatusNotFound,
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:   "",
			path:   "/",
			method: http.MethodPost,
			expect: expect{
				code:        http.StatusBadRequest,
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:   "",
			path:   "/update/counter/pollCount/100.5",
			method: http.MethodPost,
			expect: expect{
				code:        http.StatusBadRequest,
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:   "",
			path:   "/update/gauge/Alloc/val",
			method: http.MethodPost,
			expect: expect{
				code:        http.StatusBadRequest,
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:   "",
			path:   "/update/counter/pollCount/100",
			method: http.MethodGet,
			expect: expect{
				code:        http.StatusMethodNotAllowed,
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:   "",
			path:   "/update/gauge/Alloc/111.54",
			method: http.MethodGet,
			expect: expect{
				code:        http.StatusMethodNotAllowed,
				contentType: "text/plain; charset=utf-8",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(test.method, test.path, nil)
			w := httptest.NewRecorder()
			ms := storage.NewMemStorage()
			uh := UpdateHandler{
				GaugeHandler:   &GaugeHandler{MemStorage: ms},
				CounterHandler: &CounterHandler{MemStorage: ms},
			}
			uh.ServeHTTP(w, request)
			assert.Equal(t, test.expect.code, w.Code, "invalid status code")
			assert.Equal(t, test.expect.contentType, w.Header().Get("Content-Type"), "invalid Content-Type")
		})
	}
}
