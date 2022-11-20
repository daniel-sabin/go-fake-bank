package server

import (
	"log"
	"net/http"
	"strconv"

	"engineecore/demobank-server/domain/accounts"
	"engineecore/demobank-server/domain/security"
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

func NewServer(i security.ApiKeyStore, as accounts.AccountsStore, k []string) *Server {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	server := new(Server)
	storeApiKeys(i, k)

	router := new(Router)

	router.Handle("/health", http.HandlerFunc(handleHealthCheck))
	router.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", http.FileServer(statikFS)))
	router.Handle("/applications", http.HandlerFunc(handleApplicationsFactory(i)))
	router.Handle("/accounts", http.HandlerFunc(handleAccountsFactory(as)))
	router.HandleWithParam("/accounts/{param}/transactions", http.HandlerFunc(handleTransactionsFactory()))

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

func handleTransactionsFactory() HttpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		accountNumber := r.Context().Value(ParamFromUrl).(string)
		w.Write([]byte(accountNumber))
	}
}
