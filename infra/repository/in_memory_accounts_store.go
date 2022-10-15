package repository

import (
	"engineecore/demobank-server/domain/accounts"
	"engineecore/demobank-server/domain/enum"
)

type inMemoryAccountsStore struct {
}

var allAccounts = map[string]accounts.Account{
	"0000001": {Number: "0000001", Amount: 50, Currency: enum.CURRENCY_EURO},
	"0000002": {Number: "0000002", Amount: 200, Currency: enum.CURRENCY_EURO},
	"0000003": {Number: "0000003", Amount: 240, Currency: enum.CURRENCY_EURO},
	"0000004": {Number: "0000004", Amount: 100, Currency: enum.CURRENCY_EURO},
	"0000005": {Number: "0000005", Amount: 150, Currency: enum.CURRENCY_EURO},
	"0000006": {Number: "0000006", Amount: 10, Currency: enum.CURRENCY_EURO},
	"0000007": {Number: "0000007", Amount: 90, Currency: enum.CURRENCY_EURO},
	"0000008": {Number: "0000008", Amount: 500, Currency: enum.CURRENCY_EURO},
	"0000009": {Number: "0000009", Amount: 210, Currency: enum.CURRENCY_EURO},
}

var pagesWithAccounts = map[int][]accounts.Account{
	1: {allAccounts["0000001"], allAccounts["0000002"], allAccounts["0000003"]},
	2: {allAccounts["0000004"], allAccounts["0000005"], allAccounts["0000006"]},
	3: {allAccounts["0000007"], allAccounts["0000002"], allAccounts["0000004"]},
	4: {allAccounts["0000008"], allAccounts["0000009"], allAccounts["0000001"]},
}

func (*inMemoryAccountsStore) GetPagesWithAccounts() map[int][]accounts.Account {
	return pagesWithAccounts
}

func (*inMemoryAccountsStore) GetLastPageNumber() int {
	return len(pagesWithAccounts)
}

func NewInMemoryAccountsStore() accounts.AccountsStore {
	return &inMemoryAccountsStore{}
}
