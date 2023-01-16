package server

import (
	"log"
	"net/http"
	"strconv"

	"engineecore/demobank-server/domain/accounts"
	"engineecore/demobank-server/domain/security"
	viewmodel "engineecore/demobank-server/infra/view_model"
	_ "engineecore/demobank-server/statik" // path to generated statik.go

	"github.com/go-chi/chi/v5"
	"github.com/rakyll/statik/fs"
)

func storeApiKeys(i security.ApiKeyStore, apiKeys []string) {
	store := security.ApiKeyStoreFactory(i)
	for _, key := range apiKeys {
		store(key)
	}
}

func NewServer(i security.ApiKeyStore, as accounts.AccountsStore, k []string) *chi.Mux {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	storeApiKeys(i, k)

	r := chi.NewRouter()

	r.Get("/health", handleHealthCheck)
	r.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", http.FileServer(statikFS)))
	r.Get("/applications", handleApplicationsFactory(i))
	r.Route("/accounts", func(r chi.Router) {
		r.Get("/", handleAccountsFactory(as))
		r.Get("/{accountId}/transactions", handleTransactionsFactory())
	})

	return r
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func handleApplicationsFactory(i security.ApiKeyStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		isKeyAllowed := security.IsKeyAllowedFactory(i)
		allowed, _ := isKeyAllowed(r.Header.Get("x-api-key"))
		if !allowed {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.Write([]byte("{applications: ok}"))
		}
	}
}

func handleAccountsFactory(as accounts.AccountsStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
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

func handleTransactionsFactory() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		accountNumber := chi.URLParam(r, "accountId")
		w.Write([]byte(accountNumber))
	}
}
