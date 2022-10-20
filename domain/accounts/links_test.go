package accounts_test

import (
	"engineecore/demobank-server/domain/accounts"
	test "engineecore/demobank-server/utils/tests"
	"testing"
)

func TestLinks(t *testing.T) {
	t.Parallel()

	getLinksFor := accounts.GetLinksFactory(new(DumbStore))

	t.Run("get links for page 1", func(t *testing.T) {
		want := accounts.Links{Self: "/accounts?page=1", Next: "/accounts?page=2"}

		gotForPageOne := getLinksFor(1)

		test.AssertEquals(t, gotForPageOne, want)
	})

	t.Run("get links for page 2", func(t *testing.T) {
		want := accounts.Links{Self: "/accounts?page=2", Next: ""}

		got := getLinksFor(2)

		test.AssertEquals(t, got, want)
	})

	t.Run("get links for page 3", func(t *testing.T) {
		want := accounts.Links{Self: "/accounts?page=3", Next: ""}

		got := getLinksFor(3)

		test.AssertEquals(t, got, want)
	})
}
