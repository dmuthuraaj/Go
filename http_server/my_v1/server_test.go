package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubDataStore struct {
	NumberOfChild map[string]int
}

func (s *StubDataStore) GetNumberOfChild(name string) int {
	numofchild := s.NumberOfChild[name]
	return numofchild
}

func TestGetNumberOfChild(t *testing.T) {
	store := StubDataStore{
		map[string]int{
			"Rajesh": 2,
			"Sajith": 3,
		},
	}
	server := &DataServer{&store}

	tests := []struct {
		name                  string
		person                string
		expectedHttpStatus    int
		expectednumberofchild string
	}{
		{
			name:                  "return Rajesh's no.of.childs",
			person:                "Rajesh",
			expectedHttpStatus:    http.StatusOK,
			expectednumberofchild: "2",
		},
		{
			name:                  "return Sajith's no.of.childs",
			person:                "Sajith",
			expectedHttpStatus:    http.StatusOK,
			expectednumberofchild: "3",
		},
		{
			name:                  "return 404 for unknown persons",
			person:                "Muthu",
			expectedHttpStatus:    http.StatusNotFound,
			expectednumberofchild: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := newGetNumberOfChildRequest(tt.person)
			response := httptest.NewRecorder()
			server.ServeHTTP(response, request)
			assertstatus(t, response.Code, tt.expectedHttpStatus)
			assertResponseBody(t, response.Body.String(), tt.expectednumberofchild)
		})
	}
}

func assertstatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected:%v but got:%v", want, got)
	}
}

func newGetNumberOfChildRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/name/%s", name), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()

}
