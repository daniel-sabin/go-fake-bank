package accounts

import "engineecore/demobank-server/domain/enum"

type AccountsStore interface {
	GetPagesWithAccounts() map[int][]Account
	GetLastPageNumber() int
}

type Account struct {
	Number   string
	Amount   float32
	Currency enum.Currency
}

func GetAccountsFactory(r AccountsStore) func(pageFromUrl int) []Account {
	return func(pageFromUrl int) []Account {
		pagesWithAccounts := r.GetPagesWithAccounts()
		return getAccounts(pagesWithAccounts, pageFromUrl)
	}
}

func getAccounts(pagesWithAccounts map[int][]Account, pageFromUrl int) []Account {
	var page int

	if pageFromUrl == 0 {
		page = 1
	} else {
		page = pageFromUrl
	}

	accounts, ok := pagesWithAccounts[page]

	if !ok {
		return []Account{}
	}

	return accounts
}
