package models

type DomainFields struct {
	Guid                   string
	Name                   string
	OwningOrganizationGuid string
	RouterGroupGuid        string
	Shared                 bool
}

func (model DomainFields) UrlForHostAndPath(host, path string, port int) string {
	return urlStringFromParts(host, model.Name, path, port)
}
