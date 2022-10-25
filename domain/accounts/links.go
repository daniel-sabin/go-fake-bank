package accounts

import (
	"engineecore/demobank-server/domain/links"
)

func GetAccountsLinksFactory(r AccountsStore) func(pageFromUrl int) links.Links {
	return func(pageFromUrl int) links.Links {
		lastPageNumber := r.GetLastPageNumber()
		getLinks := links.GetLinksFactory("accounts")

		return getLinks(pageFromUrl, lastPageNumber)
	}
}
