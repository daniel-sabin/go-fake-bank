package server_test

import (
	"engineecore/demobank-server/infra/repository"
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
	as := repository.NewInMemoryAccountsStore()
	ts := repository.NewInMemoryTransactionsStore()
	server := server.NewServer(&DumbStore{Exist: &keyExist}, as, ts, nil)

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

	t.Run("Get accounts", func(t *testing.T) {
		// Given
		request, _ := http.NewRequest(http.MethodGet, "/accounts", nil)
		response := httptest.NewRecorder()
		want := "{\"accounts\":[{\"acc_number\":\"0000001\",\"amount\":50,\"currency\":\"EUR\"}"

		// When
		server.ServeHTTP(response, request)

		// Then
		test.AssertStatus(t, response.Code, http.StatusOK)
		test.AssertResponseBodyContains(t, response.Body.String(), want)
	})

	t.Run("Get transactions", func(t *testing.T) {
		// Given
		request, _ := http.NewRequest(http.MethodGet, "/accounts/0000001/transactions", nil)
		response := httptest.NewRecorder()
		want := "{\"transactions\":[{\"id\":1,\"label\":\"Label 1\",\"sign\":\"CDT\",\"amount\":50,\"currency\":\"EUR\"}"

		// When
		server.ServeHTTP(response, request)

		// Then
		test.AssertStatus(t, response.Code, http.StatusOK)
		test.AssertResponseBodyContains(t, response.Body.String(), want)
	})

	t.Run("Invalid transactions URL", func(t *testing.T) {
		// Given
		request, _ := http.NewRequest(http.MethodGet, "/accounts/0000001", nil)
		response := httptest.NewRecorder()
		want := "404 page not found"

		// When
		server.ServeHTTP(response, request)

		// Then
		test.AssertStatus(t, response.Code, http.StatusNotFound)
		test.AssertResponseBodyContains(t, response.Body.String(), want)
	})
}
