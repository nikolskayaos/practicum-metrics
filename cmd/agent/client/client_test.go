package client

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientSendGaugeMetric(t *testing.T) {
	type args struct {
		name  string
		value float64
	}

	tests := []struct {
		name string
		args
	}{
		{
			name: "request without error",
			args: args{
				name:  "Alloc",
				value: 1.3659236339758827,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

			}))
			client := NewClient(server.URL)
			err := client.SendGaugeMetric(test.args.name, test.args.value)
			assert.NoError(t, err)
			defer server.Close()
		})
	}
}

func TestClientSendCounterMetric(t *testing.T) {
	type args struct {
		name  string
		value int64
	}

	tests := []struct {
		name string
		args
	}{
		{
			name: "request without error",
			args: args{
				name:  "PollCount",
				value: 144,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

			}))
			client := NewClient(server.URL)
			err := client.SendCounterMetric(test.args.name, test.args.value)
			assert.NoError(t, err)
		})
	}
}
