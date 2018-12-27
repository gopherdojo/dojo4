package downloader_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gopherdojo/dojo4/kadai3-2/iwata/downloader"
)

func TestNew(t *testing.T) {
	got := downloader.New(2, "tmp")
	if got == nil {
		t.Error("New() = nil, should not return nil")
	}
}

func TestGetContentLength(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name          string
		responseError bool
		acceptRanges  string
		contentLength int
		wantErr       bool
	}{
		{
			"return error when response error",
			true,
			"",
			0,
			true,
		},
		{
			"return error when invalid Accept-Ranges header",
			false,
			"byte",
			0,
			true,
		},
		{
			"return error when invalid Content-Length",
			false,
			"bytes",
			0,
			true,
		},
		{
			"return length",
			false,
			"bytes",
			1024,
			false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if tt.responseError {
					http.Error(w, "Server Error", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Accept-Ranges", tt.acceptRanges)
				_, err := w.Write(make([]byte, tt.contentLength))
				if err != nil {
					t.Fatalf("Failed to write response: %v", err)
				}

			}))
			defer server.Close()

			got, err := downloader.GetContentLength(ctx, server.URL)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetContentLength() error = %v, wantErr %v", err, tt.wantErr)
			}
			if int(got) != tt.contentLength {
				t.Errorf("GetContentLength() = %v, want %v", got, tt.contentLength)
			}
		})
	}
}
