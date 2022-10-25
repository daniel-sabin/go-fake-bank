package links_test

import (
	"engineecore/demobank-server/domain/links"
	test "engineecore/demobank-server/utils/tests"
	"testing"
)

func TestLinks(t *testing.T) {
	t.Parallel()

	getLinks := links.GetLinksFactory("accounts")

	t.Run("get links for page 1", func(t *testing.T) {
		// Given
		currentPage := 1
		lastPage := 2

		// When
		want := links.Links{Self: "/accounts?page=1", Next: "/accounts?page=2"}
		got := getLinks(currentPage, lastPage)

		// Then
		test.AssertEquals(t, got, want)
	})

	t.Run("get links for page 2", func(t *testing.T) {
		// Given
		currentPage := 2
		lastPage := 2

		// When
		want := links.Links{Self: "/accounts?page=2", Next: ""}
		got := getLinks(currentPage, lastPage)

		// Then
		test.AssertEquals(t, got, want)
	})
}
