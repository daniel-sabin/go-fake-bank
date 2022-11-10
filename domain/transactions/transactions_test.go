package transactions_test

import (
	"engineecore/demobank-server/domain/enum"
	"engineecore/demobank-server/domain/transactions"
	test "engineecore/demobank-server/utils/tests"
	"testing"
)

type DumbStore struct{}

var allTransactions = transactions.Transactions{
	"account": {
		1: {
			{Id: 1, Amount: 50, Label: "Label 1", Sign: enum.SIGN_CREDIT, Currency: enum.CURRENCY_EURO},
		},
		2: {
			{Id: 2, Amount: 200, Label: "Label 2", Sign: enum.SIGN_DEBIT, Currency: enum.CURRENCY_EURO},
		},
	},
}

func (i *DumbStore) GetPagesWithTransactions() transactions.Transactions {
	return allTransactions
}

func (i *DumbStore) GetLastPageNumberFor(account string) int {
	return len(allTransactions[account])
}

func TestTransactions(t *testing.T) {
	t.Parallel()

	getTransactionsFor := transactions.GetTransactionsFactory(new(DumbStore))

	t.Run("get transactions for account and page 1", func(t *testing.T) {
		want := []transactions.Transaction{
			{Id: 1, Amount: 50, Label: "Label 1", Sign: enum.SIGN_CREDIT, Currency: enum.CURRENCY_EURO},
		}

		got := getTransactionsFor("account", 1)

		test.AssertEquals(t, got, want)
	})

	t.Run("get transactions for account and page 2", func(t *testing.T) {
		want := []transactions.Transaction{
			{Id: 2, Amount: 200, Label: "Label 2", Sign: enum.SIGN_DEBIT, Currency: enum.CURRENCY_EURO},
		}

		got := getTransactionsFor("account", 2)

		test.AssertEquals(t, got, want)
	})

	t.Run("get transactions for account and inexistant page", func(t *testing.T) {
		want := []transactions.Transaction{}

		got := getTransactionsFor("account", 5)

		test.AssertEquals(t, got, want)
	})

	t.Run("get transactions for inexistant account", func(t *testing.T) {
		want := []transactions.Transaction{}

		got := getTransactionsFor("inexistant", 1)

		test.AssertEquals(t, got, want)
	})
}
