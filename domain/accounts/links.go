package accounts

import "strconv"

type Links struct {
	Self string
	Next string
}

func GetLinksFactory(r AccountsStore) func(page string) Links {
	return func(page string) Links {
		pageNumber := GetPageNumber(page)
		lastPageNumber := r.GetLastPageNumber()
		return getLinks(pageNumber, lastPageNumber)
	}
}

func getLinks(page int, lastPageNumber int) Links {
	self := getSelfPage(page)
	next := getNextPage(page, lastPageNumber)

	return Links{Self: self, Next: next}
}

func getSelfPage(currentPageNumber int) string {
	return "/accounts?page=" + strconv.Itoa(currentPageNumber)
}

func getNextPage(currentPageNumber int, lastPageNumber int) string {
	next := ""

	nextPage := currentPageNumber + 1
	if nextPage <= lastPageNumber {
		next = "/accounts?page=" + strconv.Itoa(nextPage)
	}

	return next
}
