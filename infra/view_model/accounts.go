package viewmodel

import (
	"encoding/json"
	"engineecore/demobank-server/domain/accounts"
	"engineecore/demobank-server/domain/enum"
	"engineecore/demobank-server/domain/links"
)

type AccountsResponse struct {
	Accounts []AccountResponse     `json:"accounts"`
	Links    LinksAccountsResponse `json:"links"`
}

type AccountResponse struct {
	Number   string        `json:"acc_number"`
	Amount   float32       `json:"amount"`
	Currency enum.Currency `json:"currency"`
}

type LinksAccountsResponse struct {
	Self string `json:"self"`
	Next string `json:"next"`
}

func GetAccountsResponse(accounts []accounts.Account, links links.Links) []byte {
	accountsForResponse := getAccountsForResponse(accounts)
	linksForResponse := getLinksForAccountsReponse(links)

	accountsResponse := AccountsResponse{
		Accounts: accountsForResponse,
		Links:    linksForResponse,
	}

	response, _ := json.Marshal(accountsResponse)

	return response
}

func getAccountsForResponse(accounts []accounts.Account) []AccountResponse {
	accountsForResponse := []AccountResponse{}

	for _, account := range accounts {
		accountForResponse := convertAccountForResponse(account)
		accountsForResponse = append(accountsForResponse, accountForResponse)
	}

	return accountsForResponse
}

func convertAccountForResponse(account accounts.Account) AccountResponse {
	return AccountResponse{
		Number:   account.Number,
		Amount:   account.Amount,
		Currency: account.Currency,
	}
}

func getLinksForAccountsReponse(links links.Links) LinksAccountsResponse {
	return LinksAccountsResponse{Self: links.Self, Next: links.Next}
}
