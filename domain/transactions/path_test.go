package transactions_test

import (
	"engineecore/demobank-server/domain/transactions"
	"testing"
)

func TestPath(t *testing.T) {
	t.Parallel()

	t.Run("valid url returns account number", func(t *testing.T) {
		want := "1"

		got, _ := transactions.ExtractAccountFromURL("/accounts/1/transactions")

		if want != got {
			t.Errorf("invalid account number, got %s, want %s", got, want)
		}
	})

	t.Run("missing account in url", func(t *testing.T) {
		_, err := transactions.ExtractAccountFromURL("/accounts/")

		if err == nil {
			t.Errorf("must have an error")
		}
	})

	t.Run("invalid url", func(t *testing.T) {
		_, err := transactions.ExtractAccountFromURL("/accounts/1/transa")

		if err == nil {
			t.Errorf("must have an error")
		}
	})
}
