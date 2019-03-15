package main

import "strings"

func formatHostname(s string) string {
	tr := strings.Split(s, ".")
	str := strings.Replace(tr[0], "-", "_", 99)
	return str
}

func (a authResponse) getEndpoint(s string) (string, error) {
	var err error
	var res string
	for _, endp := range a.Access.ServiceCatalog {
		if endp.Type == s {
			res := endp.Endpoints[0].PublicURL
			return res, err
		}

	}
	return res, err
}
