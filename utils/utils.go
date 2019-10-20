package utils

import "net/url"

// IsRequestURL check if the string rawurl, assuming
// it was received in an HTTP request, is a valid
// URL confirm to RFC 3986
func IsRequestURL(URL string) bool {
	u, err := url.ParseRequestURI(URL)
	if err != nil {
		return false //Couldn't even parse the rawurl
	}
	if len(u.Scheme) == 0 {
		return false //No Scheme found
	}
	return true
}

func GetHost(URL string) string {
	u, err := url.ParseRequestURI(URL)
	if err != nil {
		return ""
	}

	return u.Host
}