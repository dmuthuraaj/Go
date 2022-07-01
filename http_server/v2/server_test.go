package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubNameStore struct {
	ages map[string]int
}

func (s *StubNameStore) GetPersonAge(name string) int {
	age := s.ages[name]
	return age
}

func TestGetNames(t *testing.T) {
	store := StubNameStore{
		map[string]int{
			"Dhana": 22,
			"Muthu": 23,
		},
	}
	server := &GetNameServer{&store}

	tests := []struct {
		name               string
		person             string
		expectedHttpStatus int
		expectedAge        string
	}{
		{
			name:               "return Dhana's age",
			person:             "Dhana",
			expectedHttpStatus: http.StatusOK,
			expectedAge:        "22",
		},
		{
			name:               "return Muthu's age",
			person:             "Muthu",
			expectedHttpStatus: http.StatusOK,
			expectedAge:        "23",
		},
		{
			name:               "return 404 on missing person",
			person:             "raaj",
			expectedHttpStatus: http.StatusNotFound,
			expectedAge:        "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := newGetAgeRequest(tt.person)
			response := httptest.NewRecorder()
			server.ServeHTTP(response, request)

			assertStatus(t, response.Code, tt.expectedHttpStatus)
			assertResponseBody(t, response.Body.String(), tt.expectedAge)
		})
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func newGetAgeRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/name/%s", name), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
