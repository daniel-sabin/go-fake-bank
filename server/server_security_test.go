package server_test

import (
	"engineecore/demobank-server/infra/repository"
	"engineecore/demobank-server/server"
	test "engineecore/demobank-server/utils/tests"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerSecurity(t *testing.T) {
	// Before
	ks := repository.NewInMemoryApiKeyStore()
	as := repository.NewInMemoryAccountsStore()
	ts := repository.NewInMemoryTransactionsStore()
	server := server.NewServer(ks, as, ts, []string{"fake-api-key"})

	resources := [...]string{
		"/applications",
		"/accounts",
		"/accounts/123/transactions",
	}

	t.Run("Forbidden, need an api-key", func(t *testing.T) {
		for _, resource := range resources {
			// Given
			request, _ := http.NewRequest(http.MethodGet, resource, nil)
			response := httptest.NewRecorder()
			// When
			server.ServeHTTP(response, request)

			// Then
			test.AssertStatus(t, response.Code, http.StatusUnauthorized)
		}
	})

	t.Run("Allowed with api-key ", func(t *testing.T) {
		for _, resource := range resources {
			// Given
			request, _ := http.NewRequest(http.MethodGet, resource, nil)
			request.Header.Add("x-api-key", "fake-api-key")
			response := httptest.NewRecorder()
			// When
			server.ServeHTTP(response, request)

			// Then
			test.AssertStatus(t, response.Code, http.StatusOK)
		}
	})

}
