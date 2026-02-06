package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlers(t *testing.T) {
	tests := []struct {
		name           string
		handler        http.HandlerFunc
		expectedStatus int
		expectedBody   string
	}{
		{"Home Handler", home, http.StatusOK, "Welcome to the Shapes API"},
		{"Health Handler", health, http.StatusOK, "Server is running"},
		{"About Handler", about, http.StatusOK, "Jeimy Vasquez"},
		{"Greeting Handler", greeting, http.StatusOK, "Hello! Thanks for visiting the API"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			rr := httptest.NewRecorder()

			tt.handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("got %v, expected %v", rr.Code, tt.expectedStatus)
			}

			if !strings.Contains(rr.Body.String(), tt.expectedBody) {
				t.Errorf("got %v, expected to contain %v", rr.Body.String(), tt.expectedBody)
			}
		})
	}
}
