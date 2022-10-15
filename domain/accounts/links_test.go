package accounts_test

import (
	"engineecore/demobank-server/domain/accounts"
	"engineecore/demobank-server/infra/repository"
	test "engineecore/demobank-server/utils/tests"
	"testing"
)

func TestLinks(t *testing.T) {
	t.Parallel()

	store := repository.NewInMemoryAccountsStore()
	getLinksFor := accounts.GetLinksFactory(store)

	t.Run("get links for page 1", func(t *testing.T) {
		want := accounts.Links{Self: "/accounts?page=1", Next: "/accounts?page=2"}

		gotForNoPage := getLinksFor("")
		gotForPageOne := getLinksFor("1")

		test.AssertEquals(t, gotForNoPage, want)
		test.AssertEquals(t, gotForPageOne, want)
	})

	t.Run("get links for page 2", func(t *testing.T) {
		want := accounts.Links{Self: "/accounts?page=2", Next: "/accounts?page=3"}

		got := getLinksFor("2")

		test.AssertEquals(t, got, want)
	})

	t.Run("get links for page 3", func(t *testing.T) {
		want := accounts.Links{Self: "/accounts?page=3", Next: "/accounts?page=4"}

		got := getLinksFor("3")

		test.AssertEquals(t, got, want)
	})

	t.Run("get links for page 4", func(t *testing.T) {
		want := accounts.Links{Self: "/accounts?page=4", Next: ""}

		got := getLinksFor("4")

		test.AssertEquals(t, got, want)
	})

	t.Run("get links for page 5", func(t *testing.T) {
		want := accounts.Links{Self: "/accounts?page=5", Next: ""}

		got := getLinksFor("5")

		test.AssertEquals(t, got, want)
	})
}
