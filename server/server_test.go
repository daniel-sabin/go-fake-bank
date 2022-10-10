package server_test

import (
	"engineecore/demobank-server/server"
	"engineecore/demobank-server/test"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStartServer(t *testing.T) {
	// Before
	server := server.NewServer()

	t.Run("health check", func(t *testing.T) {
		// Given
		request, _ := http.NewRequest(http.MethodGet, "/health", nil)
		response := httptest.NewRecorder()

		// When
		server.ServeHTTP(response, request)

		// Then
		test.AssertStatus(t, response.Code, http.StatusOK)
		test.AssertResponseBody(t, response.Body.String(), "ok")
	})
}
