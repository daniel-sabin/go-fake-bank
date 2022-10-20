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

func GetAccountsFactory(r AccountsStore) func(page int) []Account {
	return func(page int) []Account {
		pagesWithAccounts := r.GetPagesWithAccounts()
		return getAccounts(pagesWithAccounts, page)
	}
}

func getAccounts(pagesWithAccounts map[int][]Account, page int) []Account {
	accounts, ok := pagesWithAccounts[page]

	if !ok {
		return []Account{}
	}

	return accounts
}
