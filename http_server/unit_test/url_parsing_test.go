package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFunc(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	resp := httptest.NewRecorder()
	ts := httptest.NewServer(http.HandlerFunc(GetFunc(resp, req)))
}
