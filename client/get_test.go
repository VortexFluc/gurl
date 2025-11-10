package client

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {

	tests := []struct {
		name           string
		url            string
		serverHandler  http.HandlerFunc
		wantOutput     GetResult
		wantErr        bool
		wantErrContain string
		insecure       bool
	}{
		{
			name: "successful GET request",
			url:  "http://test-server/ok",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					t.Errorf("Expected GET method, got %s", r.Method)
				}
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"status": "ok"}`))
			},
			wantOutput: GetResult{false, []byte(`{"status": "ok"}`)},
			wantErr:    false,
		},
		{
			name: "failed GET request",
			url:  "http://test-server/error",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					t.Errorf("Expected GET method, got %s", r.Method)
				}
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"status": "error"}`))
			},
			wantOutput: GetResult{true, []byte(`{"status": "error"}`)},
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var serverURL string
			if tt.url != "" && strings.HasPrefix(tt.url, "http://test-server") {
				server := httptest.NewServer(tt.serverHandler)
				defer server.Close()
				serverURL = strings.Replace(tt.url, "http://test-server", server.URL, 1)
			} else if tt.url == "" {
				serverURL = ""
			} else {
				serverURL = tt.url
			}

			r, err := Get(tt.insecure, serverURL)

			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(*r, tt.wantOutput) {
				t.Errorf("Get() got = %v, want %v", r, tt.wantOutput)
			}
		})
	}
}
