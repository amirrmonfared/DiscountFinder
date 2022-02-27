package crawler

import "net/url"

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
