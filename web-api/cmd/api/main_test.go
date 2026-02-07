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
		{
			"Home",
			home,
			http.StatusOK,
			"Welcome to the Shapes API",
		},
		{
			"Health",
			health,
			http.StatusOK,
			"Server is running",
		},
		{
			"About",
			about,
			http.StatusOK,
			"Allen",
		},
		{
			"Time",
			timeHandler,
			http.StatusOK,
			"",
		},
		{
			"Random",
			random,
			http.StatusOK,
			"",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			req := httptest.NewRequest("GET", "/", nil)
			rr := httptest.NewRecorder()

			tt.handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected %v, got %v", tt.expectedStatus, rr.Code)
			}

			if tt.expectedBody != "" {
				if !strings.Contains(rr.Body.String(), tt.expectedBody) {
					t.Errorf("Unexpected body: %v", rr.Body.String())
				}
			}

		})
	}
}
