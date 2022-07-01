package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetNames(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	GetName(response, request)

	t.Run("return name", func(t *testing.T) {
		got := response.Body.String()
		want := "Dhana"
		if got != want {
			t.Errorf("exp:=%v,got=%v", want, got)
		}
	})
}
