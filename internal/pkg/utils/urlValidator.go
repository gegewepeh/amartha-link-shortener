package utils

import (
	"regexp"
)

func IsValidURL(url string) bool {
	// The regular expression to match a URL
	urlRegex := regexp.MustCompile(`^(http://|https://)?[a-zA-Z0-9\-_]+(\.[a-zA-Z]{2,})+(\/[a-zA-Z0-9\-._~:/?#[\]@!$&'()*+,;=%]+)?$`)

	return urlRegex.MatchString(url)
}