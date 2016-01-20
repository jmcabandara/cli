package models

import (
	"fmt"
	"net/url"
	"strings"
)

type Route struct {
	Guid   string
	Host   string
	Port   int
	Domain DomainFields
	Path   string

	Space           SpaceFields
	Apps            []ApplicationFields
	ServiceInstance ServiceInstanceFields
}

func (r Route) URL() string {
	return urlStringFromParts(r.Host, r.Domain.Name, r.Path, r.Port)
}

func urlStringFromParts(hostName, domainName, path string, port int) string {
	var host string
	if hostName != "" {
		if port == 0 {
			host = fmt.Sprintf("%s.%s", hostName, domainName)
		} else {
			host = fmt.Sprintf("%s.%s:%d", hostName, domainName, port)
		}
	} else {
		if port == 0 {
			host = domainName
		} else {
			host = fmt.Sprintf("%s:%d", domainName, port)
		}
	}

	u := url.URL{
		Host: host,
		Path: path,
	}

	return strings.TrimPrefix(u.String(), "//") // remove the empty scheme
}
