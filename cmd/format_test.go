package cmd

import (
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		name       string
		data       []byte
		formatType string
		want       string
	}{
		{
			name:       "format string",
			data:       []byte(`{"hello": "world"}`),
			formatType: "string",
			want:       `{"hello": "world"}`,
		},
		{
			name:       "format json",
			data:       []byte(`{"hello": "world"}`),
			formatType: "json",
			want:       "{\n\t\"hello\": \"world\"\n}",
		},
		{
			name:       "format raw",
			data:       []byte(`{"hello": "world"}`),
			formatType: "raw",
			want:       fmt.Sprint([]byte(`{"hello": "world"}`)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := Format(tt.data, tt.formatType)
			if err != nil {
				t.Fatal(err)
			}

			if r != tt.want {
				t.Errorf("got %q, want %q", r, tt.want)
			}
		})
	}
}
