package requirements

import (
	"github.com/cloudfoundry/cli/cf/api/organizations"
	"github.com/cloudfoundry/cli/cf/models"
	"github.com/cloudfoundry/cli/cf/terminal"
)

//go:generate counterfeiter -o fakes/fake_organization_requirement.go . Organization
type Organization interface {
	Requirement
	SetOrganizationName(string)
	GetOrganization() models.Organization
}

type organization struct {
	name    string
	ui      terminal.UI
	orgRepo organizations.OrganizationRepository
	org     models.Organization
}

func NewOrganizationRequirement(name string, ui terminal.UI, sR organizations.OrganizationRepository) *organization {
	return &organization{
		name:    name,
		ui:      ui,
		orgRepo: sR,
	}
}

func (r *organization) Execute() (success bool) {
	var apiErr error
	r.org, apiErr = r.orgRepo.FindByName(r.name)

	if apiErr != nil {
		r.ui.Failed(apiErr.Error())
		return false
	}

	return true
}

func (r *organization) SetOrganizationName(name string) {
	r.name = name
}

func (r *organization) GetOrganization() models.Organization {
	return r.org
}
