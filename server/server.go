package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"engineecore/demobank-server/domain/accounts"
	"engineecore/demobank-server/domain/security"
	"engineecore/demobank-server/domain/transactions"
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

func NewServer(
	i security.ApiKeyStore,
	as accounts.AccountsStore,
	ts transactions.TransactionsStore,
	k []string,
) *chi.Mux {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	storeApiKeys(i, k)

	r := chi.NewRouter()

	r.Get("/health", handleHealthCheck)
	r.Mount("/swaggerui/", http.StripPrefix("/swaggerui/", http.FileServer(statikFS)))

	// routes protected with x-api-key
	r.Group(func(r chi.Router) {
		r.Use(middlewareApiKeyAllowedFactory(i))
		r.Get("/applications", handleApplications)
		r.Post("/applications", handleCreateApplication)
		r.Route("/accounts", func(r chi.Router) {
			r.Get("/", handleAccountsFactory(as))
			r.Get("/{accountNumber}/transactions", handleTransactionsFactory(ts))
		})
	})

	return r
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func middlewareApiKeyAllowedFactory(i security.ApiKeyStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			isKeyAllowed := security.IsKeyAllowedFactory(i)
			allowed, _ := isKeyAllowed(r.Header.Get("x-api-key"))
			if !allowed {
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}

func handleApplications(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{applications: ok}"))
}

func handleCreateApplication(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var mcPostBody map[string]interface{}
	json.NewDecoder(r.Body).Decode(&mcPostBody)
	mcPostBody["question_response"] = "Hello world!"
	body, _ := json.Marshal(mcPostBody)

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
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

func handleTransactionsFactory(ts transactions.TransactionsStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		pageFromUrl, _ := strconv.Atoi(query.Get("page"))

		accountNumber := chi.URLParam(r, "accountNumber")

		getTransactionsFor := transactions.GetTransactionsFactory(ts)
		getLinksFor := transactions.GetTransactionsLinksFactory(ts)

		transactionsForPage := getTransactionsFor(accountNumber, pageFromUrl)
		linksForPage := getLinksFor(accountNumber, pageFromUrl)

		response := viewmodel.GetTransactionsResponse(transactionsForPage, linksForPage)

		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.Write(response)
	}
}
