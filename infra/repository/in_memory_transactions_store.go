package repository

import (
	"engineecore/demobank-server/domain/enum"
	"engineecore/demobank-server/domain/transactions"
)

var allTransactions = map[int]transactions.Transaction{
	1: {Id: 1, Amount: 50, Label: "Label 1", Sign: enum.SIGN_CREDIT, Currency: enum.CURRENCY_EURO},
	2: {Id: 2, Amount: 200, Label: "Label 2", Sign: enum.SIGN_DEBIT, Currency: enum.CURRENCY_EURO},
	3: {Id: 3, Amount: 32, Label: "Label 3", Sign: enum.SIGN_DEBIT, Currency: enum.CURRENCY_EURO},
	4: {Id: 4, Amount: 15, Label: "Label 4", Sign: enum.SIGN_CREDIT, Currency: enum.CURRENCY_EURO},
	5: {Id: 5, Amount: 156, Label: "Label 5", Sign: enum.SIGN_DEBIT, Currency: enum.CURRENCY_EURO},
	6: {Id: 6, Amount: 59, Label: "Label 6", Sign: enum.SIGN_CREDIT, Currency: enum.CURRENCY_EURO},
	7: {Id: 7, Amount: 33, Label: "Label 7", Sign: enum.SIGN_CREDIT, Currency: enum.CURRENCY_EURO},
	8: {Id: 8, Amount: 98, Label: "Label 8", Sign: enum.SIGN_CREDIT, Currency: enum.CURRENCY_EURO},
	9: {Id: 9, Amount: 579, Label: "Label 9", Sign: enum.SIGN_CREDIT, Currency: enum.CURRENCY_EURO},
}

var transactionsForAccounts = transactions.Transactions{
	"0000001": {
		1: {allTransactions[1], allTransactions[2], allTransactions[3]},
		2: {allTransactions[4], allTransactions[5], allTransactions[6]},
		3: {allTransactions[3], allTransactions[2], allTransactions[6]},
	},
	"0000002": {
		1: {allTransactions[4], allTransactions[5], allTransactions[6]},
		2: {allTransactions[1], allTransactions[2], allTransactions[3]},
		3: {allTransactions[4], allTransactions[5]},
	},
	"0000003": {
		1: {allTransactions[2], allTransactions[8], allTransactions[6]},
		2: {allTransactions[4], allTransactions[2], allTransactions[8]},
		3: {allTransactions[6], allTransactions[4], allTransactions[2]},
	},
	"0000005": {
		1: {allTransactions[1], allTransactions[7], allTransactions[3]},
		2: {allTransactions[3], allTransactions[9], allTransactions[5]},
		3: {allTransactions[5], allTransactions[1], allTransactions[7]},
	},
	"0000006": {
		1: {allTransactions[4], allTransactions[2], allTransactions[1]},
		2: {allTransactions[9], allTransactions[4], allTransactions[9]},
		3: {allTransactions[8], allTransactions[3], allTransactions[6]},
	},
	"0000007": {
		1: {allTransactions[3], allTransactions[6], allTransactions[9]},
		2: {allTransactions[2], allTransactions[5], allTransactions[8]},
		3: {allTransactions[1], allTransactions[4], allTransactions[7]},
	},
	"0000009": {
		1: {allTransactions[1], allTransactions[4], allTransactions[6]},
		2: {allTransactions[2], allTransactions[5], allTransactions[3]},
		3: {allTransactions[3], allTransactions[6], allTransactions[3]},
	},
}

type inMemoryTransactionsStore struct {
}

func (*inMemoryTransactionsStore) GetPagesWithTransactions() transactions.Transactions {
	return transactionsForAccounts
}

func (*inMemoryTransactionsStore) GetLastPageNumberFor(account string) int {
	return len(transactionsForAccounts[account])
}

func NewInMemoryTransactionsStore() transactions.TransactionsStore {
	return &inMemoryTransactionsStore{}
}
