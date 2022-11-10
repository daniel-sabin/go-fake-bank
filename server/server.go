package server

import (
	"log"
	"net/http"
	"strconv"

	"engineecore/demobank-server/domain/accounts"
	"engineecore/demobank-server/domain/security"
	"engineecore/demobank-server/domain/transactions"
	viewmodel "engineecore/demobank-server/infra/view_model"
	_ "engineecore/demobank-server/statik" // path to generated statik.go

	"github.com/rakyll/statik/fs"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)

type Server struct {
	http.Handler
}

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
) *Server {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	server := new(Server)
	storeApiKeys(i, k)

	router := http.NewServeMux()

	router.Handle("/health", http.HandlerFunc(handleHealthCheck))
	router.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", http.FileServer(statikFS)))
	router.Handle("/applications", http.HandlerFunc(handleApplicationsFactory(i)))
	router.Handle("/accounts", http.HandlerFunc(handleAccountsFactory(as)))
	router.Handle("/accounts/", http.HandlerFunc(handleTransactionsFactory(ts)))

	server.Handler = router

	return server
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func handleApplicationsFactory(i security.ApiKeyStore) HttpHandler {
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

func handleAccountsFactory(as accounts.AccountsStore) HttpHandler {
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

func handleTransactionsFactory(ts transactions.TransactionsStore) HttpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		pageFromUrl, _ := strconv.Atoi(query.Get("page"))
		accountNumber, err := transactions.ExtractAccountFromURL(r.URL.Path)

		if err != nil {
			sendNotFound(w)
			return
		}

		getTransactionsFor := transactions.GetTransactionsFactory(ts)
		getLinksFor := transactions.GetTransactionsLinksFactory(ts)

		transactionsForPage := getTransactionsFor(accountNumber, pageFromUrl)
		linksForPage := getLinksFor(accountNumber, pageFromUrl)

		response := viewmodel.GetTransactionsResponse(transactionsForPage, linksForPage)

		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.Write(response)
	}
}

func sendNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 page not found"))
}
