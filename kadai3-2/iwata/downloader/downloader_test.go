package downloader_test

import (
	"context"
	"fmt"
	"io/ioutil"
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

func TestChunkRequest_Do(t *testing.T) {
	ctx := context.Background()

	t.Run("successfully", func(t *testing.T) {
		r := &downloader.ChunkRequest{Start: 10, End: 50}
		wantBody := "OK"

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			meth := req.Method
			if meth != "GET" {
				t.Errorf("ChunkRequest.Do() sent %s request, want GET request", meth)
			}
			got := req.Header.Get("Range")
			want := fmt.Sprintf("bytes=%d-%d", r.Start, r.End)
			if got != want {
				t.Errorf("ChunkRequest.Do() requested with Range %s, want %s", got, want)
			}
			_, err := w.Write([]byte(wantBody))
			if err != nil {
				t.Fatalf("Failed to write response: %v", err)
			}

		}))
		defer server.Close()

		res, err := r.Do(ctx, server.URL)
		if err != nil {
			t.Errorf("ChunkRequest.Do() error = %v, but not want any errors", err)
		}

		got, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Failed to read Body: %v", err)
		}
		defer res.Body.Close()
		if string(got) != wantBody {
			t.Errorf("ChunkRequest.Do() body = %s, want = %s", string(got), wantBody)
		}
	})
}
