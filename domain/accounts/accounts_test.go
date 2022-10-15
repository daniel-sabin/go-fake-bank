package accounts_test

import (
	"engineecore/demobank-server/domain/accounts"
	"engineecore/demobank-server/domain/enum"
	test "engineecore/demobank-server/utils/tests"
	"testing"
)

type DumbStore struct {
}

func (i *DumbStore) GetPagesWithAccounts() map[int][]accounts.Account {
	return map[int][]accounts.Account{
		1: {accounts.Account{Number: "0000001", Amount: 50, Currency: enum.CURRENCY_EURO}},
		2: {accounts.Account{Number: "0000002", Amount: 200, Currency: enum.CURRENCY_EURO}},
	}
}

func (i *DumbStore) GetLastPageNumber() int {
	return 2
}

func TestAccounts(t *testing.T) {
	t.Parallel()

	getAccountsFor := accounts.GetAccountsFactory(new(DumbStore))

	t.Run("get accounts for page 1", func(t *testing.T) {
		want := []accounts.Account{
			{Number: "0000001", Amount: 50, Currency: enum.CURRENCY_EURO},
		}

		gotForNoPage := getAccountsFor("")
		gotForPageOne := getAccountsFor("1")

		test.AssertEquals(t, gotForNoPage, want)
		test.AssertEquals(t, gotForPageOne, want)
	})

	t.Run("get accounts for page 2", func(t *testing.T) {
		want := []accounts.Account{
			{Number: "0000002", Amount: 200, Currency: enum.CURRENCY_EURO},
		}

		got := getAccountsFor("2")

		test.AssertEquals(t, got, want)
	})

	t.Run("get accounts response for page 3", func(t *testing.T) {
		want := []accounts.Account{}

		got := getAccountsFor("5")

		test.AssertEquals(t, got, want)
	})
}
