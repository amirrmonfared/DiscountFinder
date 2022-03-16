package crawler

import (
	"net/url"
)

// IsSameDomain is trying to determine if the URL is in the same domain as the base URL.
func IsSameDomain(href, baseUrl string) bool {
	uri, err := url.Parse(href)
	if err != nil {
		return false
	}
	parentUri, err := url.Parse(baseUrl)
	if err != nil {
		return false
	}

	if uri.Host != parentUri.Host {
		return false
	}

	return true
}

// ToFixedURL converts that base url into a string with the help of the ResolveReference function from net/url.
func ToFixedURL(href, baseUrl string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return "cannot parse uri"
	}

	base, err := url.Parse(baseUrl)
	if err != nil {
		return "cannot parse base url"
	}

	toFixedUri := base.ResolveReference(uri)

	return toFixedUri.String()
}
