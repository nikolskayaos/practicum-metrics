package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShiftPath(t *testing.T) {
	type expect struct {
		head string
		tail string
	}
	tests := []struct {
		name   string
		path   string
		expect expect
	}{
		{
			name: "1",
			path: "/",
			expect: expect{
				head: "",
				tail: "/",
			},
		},
		{
			name: "2",
			path: "//update/",
			expect: expect{
				head: "update",
				tail: "/",
			},
		},
		{
			name: "3",
			path: "/update",
			expect: expect{
				head: "update",
				tail: "/",
			},
		},
		{
			name: "4",
			path: "/update/gauge/123",
			expect: expect{
				head: "update",
				tail: "/gauge/123",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			head, tail := ShiftPath(test.path)
			assert.Equal(t, test.expect.head, head, "")
			assert.Equal(t, test.expect.tail, tail, "")
		})
	}
}
