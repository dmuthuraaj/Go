package main

import (
	"testing"
)

var tests = []struct {
	name     string
	givenNum float64
	expected float64
	isErr    bool
}{
	{"valid-data", 4.0, 2.0, false},
	{"invalid_data", -8, 0, true},
}

func TestSqrt(t *testing.T) {
	for _, tt := range tests {
		got, err := sqrt(tt.givenNum)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not got one")
			}
		} else {
			if err != nil {
				t.Error("did not expected an error but got one", err.Error())
			}
		}
		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}
