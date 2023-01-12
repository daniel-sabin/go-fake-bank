package server

import (
	"net/http"
	"strconv"

	"engineecore/demobank-server/domain/accounts"
	"engineecore/demobank-server/domain/security"
	viewmodel "engineecore/demobank-server/infra/view_model"
	_ "engineecore/demobank-server/statik" // path to generated statik.go

	"goyave.dev/goyave/v4"
)

type Server struct {
	http.Handler
}

func storeApiKeys(i security.ApiKeyStore, apiKeys []string) {
	store := security.ApiKeyStoreFactory(i)
	for _, key := range apiKeys {
		store(key)
	}
}

func RegisterRoutes(i security.ApiKeyStore, as accounts.AccountsStore, k []string) func(router *goyave.Router) {
	storeApiKeys(i, k)

	return func(router *goyave.Router) {
		router.Get("/health", handleHealthCheck)
		router.Get("/applications", handleApplicationsFactory(i))
		router.Get("/accounts", handleAccountsFactory(as))
		router.Get("/accounts/{accountId}/transactions", handleTransactionsFactory())
		router.Static("/swaggerui/", "server/swaggerui", false)
	}
}

func handleHealthCheck(w *goyave.Response, r *goyave.Request) {
	w.Write([]byte("ok"))
}

func handleApplicationsFactory(i security.ApiKeyStore) func(w *goyave.Response, r *goyave.Request) {
	return func(w *goyave.Response, r *goyave.Request) {
		isKeyAllowed := security.IsKeyAllowedFactory(i)
		allowed, _ := isKeyAllowed(r.Header().Get("x-api-key"))
		if !allowed {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.Write([]byte("{applications: ok}"))
		}
	}
}

func handleAccountsFactory(as accounts.AccountsStore) func(w *goyave.Response, r *goyave.Request) {
	return func(w *goyave.Response, r *goyave.Request) {
		query := r.Request().URL.Query()
		pageFromUrl, _ := strconv.Atoi(query.Get("page"))

		getAccountsFor := accounts.GetAccountsFactory(as)
		getLinksFor := accounts.GetAccountsLinksFactory(as)

		accountsForPage := getAccountsFor(pageFromUrl)
		linksForPage := getLinksFor(pageFromUrl)

		response := viewmodel.GetAccountsResponse(accountsForPage, linksForPage)

		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.Write(response)
	}
}

func handleTransactionsFactory() func(w *goyave.Response, r *goyave.Request) {
	return func(w *goyave.Response, r *goyave.Request) {
		accountNumber := r.Params["accountId"]
		w.Write([]byte(accountNumber))
	}
}
