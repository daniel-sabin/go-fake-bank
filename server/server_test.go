package server_test

import (
	"engineecore/demobank-server/server"
	test "engineecore/demobank-server/utils/tests"
	"net/http"
	"net/http/httptest"
	"testing"
)

type DumbStore struct {
	Exist *bool
}

func (i *DumbStore) Save(key string) {
}

func (i *DumbStore) Exists(key string) bool {
	return *i.Exist
}

func TestServer(t *testing.T) {
	keyExist := true

	// Before
	server := server.NewServer(&DumbStore{Exist: &keyExist}, nil)

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

	t.Run("swagger ui", func(t *testing.T) {
		// Given
		request, _ := http.NewRequest(http.MethodGet, "/swaggerui/", nil)
		response := httptest.NewRecorder()
		// When
		server.ServeHTTP(response, request)

		// Then
		test.AssertStatus(t, response.Code, http.StatusOK)
		test.AssertResponseBodyContains(t, response.Body.String(), "DOCTYPE")
	})

	t.Run("Get applications is forbidden, need an api-key", func(t *testing.T) {
		keyExist = false

		// Given
		request, _ := http.NewRequest(http.MethodGet, "/applications", nil)
		response := httptest.NewRecorder()
		// When
		server.ServeHTTP(response, request)

		// Then
		test.AssertStatus(t, response.Code, http.StatusUnauthorized)
	})

	t.Run("Get applications allowed ", func(t *testing.T) {
		keyExist = true

		// Given
		request, _ := http.NewRequest(http.MethodGet, "/applications", nil)
		request.Header.Add("x-api-key", "fake-key")
		response := httptest.NewRecorder()
		// When
		server.ServeHTTP(response, request)

		// Then
		test.AssertStatus(t, response.Code, http.StatusOK)
		test.AssertResponseBody(t, response.Body.String(), "{applications: ok}")

	})

}
