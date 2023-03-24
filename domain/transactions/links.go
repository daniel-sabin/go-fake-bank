package transactions

import (
	"engineecore/demobank-server/domain/links"
	"fmt"
)

func GetTransactionsLinksFactory(r TransactionsStore) func(accountNumber string, pageFromUrl int) links.Links {
	return func(accountNumber string, pageFromUrl int) links.Links {
		lastPageNumber := r.GetLastPageNumberFor(accountNumber)
		base := fmt.Sprintf("accounts/%s/transactions", accountNumber)
		getLinks := links.GetLinksFactory(base)

		return getLinks(pageFromUrl, lastPageNumber)
	}
}
