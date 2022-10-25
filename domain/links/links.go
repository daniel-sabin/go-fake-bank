package links

import (
	"fmt"
)

type Links struct {
	Self string
	Next string
}

func GetLinksFactory(base string) func(pageFromUrl int, lastPageNumber int) Links {
	formatLink := func(page int) string {
		return fmt.Sprintf("/%s?page=%d", base, page)
	}

	getSelfPage := func(page int) string {
		return formatLink(page)
	}

	getNextPage := func(currentPage int, lastPageNumber int) string {
		next := ""

		nextPage := currentPage + 1
		if nextPage <= lastPageNumber {
			next = formatLink(nextPage)
		}

		return next
	}

	return func(pageFromUrl int, lastPageNumber int) Links {
		var page int

		if pageFromUrl == 0 {
			page = 1
		} else {
			page = pageFromUrl
		}

		self := getSelfPage(page)
		next := getNextPage(page, lastPageNumber)

		return Links{Self: self, Next: next}
	}
}
