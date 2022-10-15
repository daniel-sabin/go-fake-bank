package viewmodel

import (
	"encoding/json"
	"engineecore/demobank-server/domain/accounts"
	"engineecore/demobank-server/domain/enum"
)

type AccountsResponse struct {
	Accounts []AccountResponse `json:"accounts"`
	Links    LinksResponse     `json:"links"`
}

type AccountResponse struct {
	Number   string        `json:"acc_number"`
	Amount   float32       `json:"amount"`
	Currency enum.Currency `json:"currency"`
}

type LinksResponse struct {
	Self string `json:"self"`
	Next string `json:"next"`
}

func GetResponseFor(accounts []accounts.Account, links accounts.Links) []byte {
	accountsForResponse := getAccountsForResponse(accounts)
	linksForResponse := getLinksForReponse(links)

	accountsResponse := AccountsResponse{Accounts: accountsForResponse, Links: linksForResponse}
	response, _ := json.Marshal(accountsResponse)

	return response
}

func getAccountsForResponse(accounts []accounts.Account) []AccountResponse {
	var accountsForResponse []AccountResponse

	for _, account := range accounts {
		accountForResponse := AccountResponse{Number: account.Number, Amount: account.Amount, Currency: account.Currency}
		accountsForResponse = append(accountsForResponse, accountForResponse)
	}

	return accountsForResponse
}

func getLinksForReponse(links accounts.Links) LinksResponse {
	return LinksResponse{Self: links.Self, Next: links.Next}
}
