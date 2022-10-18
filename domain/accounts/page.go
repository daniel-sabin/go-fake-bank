package accounts

import "strconv"

func GetPageNumber(page string) int {
	var pageNumber int

	if page == "" {
		pageNumber = 1
	} else {
		pageNumber, _ = strconv.Atoi(page)
	}

	return pageNumber
}
