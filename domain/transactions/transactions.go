package transactions

import "engineecore/demobank-server/domain/enum"

type Transactions map[string]map[int][]Transaction

type TransactionsStore interface {
	GetAllTransactions() Transactions
	GetLastPageNumberFor(string) int
}

type Transaction struct {
	Id       int
	Label    string
	Sign     enum.Sign
	Amount   float32
	Currency enum.Currency
}

func GetTransactionsFactory(r TransactionsStore) func(accountNumber string, pageFromUrl int) []Transaction {
	return func(accountNumber string, pageFromUrl int) []Transaction {
		transactions := r.GetAllTransactions()
		pagesWithTransactions, ok := transactions[accountNumber]

		if !ok {
			return emptyTransaction()
		}

		return getTransactions(pagesWithTransactions, pageFromUrl)
	}
}

func getTransactions(pagesWithTransactions map[int][]Transaction, pageFromUrl int) []Transaction {
	var page int

	if pageFromUrl == 0 {
		page = 1
	} else {
		page = pageFromUrl
	}

	transactions, ok := pagesWithTransactions[page]

	if !ok {
		return emptyTransaction()
	}

	return transactions
}

func emptyTransaction() []Transaction {
	return []Transaction{}
}
