package viewmodel

import (
	"encoding/json"
	"engineecore/demobank-server/domain/enum"
	"engineecore/demobank-server/domain/links"
	"engineecore/demobank-server/domain/transactions"
)

type TransactionsResponse struct {
	Transactions []TransactionResponse     `json:"transactions"`
	Links        LinksTransactionsResponse `json:"links"`
}

type TransactionResponse struct {
	Id       int           `json:"id"`
	Label    string        `json:"label"`
	Sign     enum.Sign     `json:"sign"`
	Amount   float32       `json:"amount"`
	Currency enum.Currency `json:"currency"`
}

type LinksTransactionsResponse struct {
	Self string `json:"self"`
	Next string `json:"next"`
}

func GetTransactionsResponse(transactions []transactions.Transaction, links links.Links) []byte {
	transactionsForResponse := getTransactionsForResponse(transactions)
	linksForResponse := getLinksForTransactionsReponse(links)

	transactionsResponse := TransactionsResponse{
		Transactions: transactionsForResponse,
		Links:        linksForResponse,
	}

	response, _ := json.Marshal(transactionsResponse)

	return response
}

func getTransactionsForResponse(transactions []transactions.Transaction) []TransactionResponse {
	transactionsForResponse := []TransactionResponse{}

	for _, transaction := range transactions {
		transactionForResponse := convertTransactionForResponse(transaction)
		transactionsForResponse = append(transactionsForResponse, transactionForResponse)
	}

	return transactionsForResponse
}

func convertTransactionForResponse(transaction transactions.Transaction) TransactionResponse {
	return TransactionResponse{
		Id:       transaction.Id,
		Label:    transaction.Label,
		Sign:     transaction.Sign,
		Amount:   transaction.Amount,
		Currency: transaction.Currency,
	}
}

func getLinksForTransactionsReponse(links links.Links) LinksTransactionsResponse {
	return LinksTransactionsResponse{Self: links.Self, Next: links.Next}
}
