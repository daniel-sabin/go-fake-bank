package transactions

import (
	"errors"
	"regexp"
)

func ExtractAccountFromURL(url string) (string, error) {
	r, _ := regexp.Compile("accounts/(.*?)/transactions")
	match := r.FindStringSubmatch(url)

	if len(match) == 0 {
		return "", errors.New("invalid url")
	}

	return match[1], nil
}
