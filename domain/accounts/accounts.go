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

func GetAccountsFactory(r AccountsStore) func(page string) []Account {
	return func(page string) []Account {
		pageNumber := GetPageNumber(page)
		pagesWithAccounts := r.GetPagesWithAccounts()
		return getAccounts(pagesWithAccounts, pageNumber)
	}
}

func getAccounts(pagesWithAccounts map[int][]Account, page int) []Account {
	accounts, ok := pagesWithAccounts[page]

	if !ok {
		return []Account{}
	}

	return accounts
}
