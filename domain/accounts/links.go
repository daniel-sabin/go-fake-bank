package accounts

import "strconv"

type Links struct {
	Self string
	Next string
}

func GetLinksFactory(r AccountsStore) func(page int) Links {
	return func(page int) Links {
		lastPageNumber := r.GetLastPageNumber()
		return getLinks(page, lastPageNumber)
	}
}

func getLinks(page int, lastPageNumber int) Links {
	self := getSelfPage(page)
	next := getNextPage(page, lastPageNumber)

	return Links{Self: self, Next: next}
}

func getSelfPage(currentPage int) string {
	return formatLink(currentPage)
}

func getNextPage(currentPage int, lastPageNumber int) string {
	next := ""

	nextPage := currentPage + 1
	if nextPage <= lastPageNumber {
		next = formatLink(nextPage)
	}

	return next
}

func formatLink(page int) string {
	return "/accounts?page=" + strconv.Itoa(page)
}
