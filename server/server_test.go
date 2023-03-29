package server_test

import (
	"bytes"
	"encoding/json"
	"engineecore/demobank-server/infra/repository"
	"engineecore/demobank-server/server"
	test "engineecore/demobank-server/utils/tests"
	"net/http"
	"net/http/httptest"
	"testing"
)

type DumbStore struct {
}

func (i *DumbStore) Save(key string) {
}

func (i *DumbStore) Exists(key string) bool {
	return true
}

func TestServer(t *testing.T) {

	// Before
	as := repository.NewInMemoryAccountsStore()
	ts := repository.NewInMemoryTransactionsStore()
	server := server.NewServer(&DumbStore{}, as, ts, nil)

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

	t.Run("Get applications ", func(t *testing.T) {
		// Given
		request, _ := http.NewRequest(http.MethodGet, "/applications", nil)
		response := httptest.NewRecorder()
		// When
		server.ServeHTTP(response, request)

		// Then
		test.AssertStatus(t, response.Code, http.StatusOK)
		test.AssertResponseBody(t, response.Body.String(), "{applications: ok}")
	})

	t.Run("accounts", func(t *testing.T) {
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

	t.Run("transactions", func(t *testing.T) {
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

	t.Run("Create a new application - Set of client_id/client_secret", func(t *testing.T) {
		// Given
		mcPostBody := map[string]interface{}{
			"question_text": "Is this a test post for MutliQuestion?",
		}
		body, _ := json.Marshal(mcPostBody)
		request, _ := http.NewRequest(http.MethodPost, "/applications", bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")
		defer request.Body.Close()
		response := httptest.NewRecorder()

		// When
		server.ServeHTTP(response, request)

		// Then
		var m map[string]interface{}
		json.NewDecoder(response.Body).Decode(&m)

		test.AssertStatus(t, response.Code, http.StatusOK)
		test.AssertResponseBody(t, m["question_response"].(string), "Hello world!")
		test.AssertResponseBody(t, m["question_text"].(string), "Is this a test post for MutliQuestion?")
	})

}
