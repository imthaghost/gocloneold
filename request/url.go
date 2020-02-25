package request

import strutil "github.com/torden/go-strutil"

// Validate checks for a valid url
func Validate(domain string) bool {
	if !strutil.NewStringValidator().IsValidURL(domain) {
		return false
	}
	return true
}
