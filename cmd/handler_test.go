package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// BenchmarkVersionWriter benchmarks the version handler.
func BenchmarkVersionWriter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/", nil)
		version(w, r)
	}
}
