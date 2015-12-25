package securitygroup

import (
	"github.com/cloudfoundry/cli/cf/api/security_groups/defaults/running"
	"github.com/cloudfoundry/cli/cf/command_registry"
	"github.com/cloudfoundry/cli/cf/configuration/core_config"
	. "github.com/cloudfoundry/cli/cf/i18n"
	"github.com/cloudfoundry/cli/cf/requirements"
	"github.com/cloudfoundry/cli/cf/terminal"
	"github.com/simonleung8/flags"
)

type ListRunningSecurityGroups struct {
	ui                       terminal.UI
	runningSecurityGroupRepo running.RunningSecurityGroupsRepo
	configRepo               core_config.Reader
}

func (cmd *ListRunningSecurityGroups) MetaData() command_registry.CommandMetadata {
	return command_registry.CommandMetadata{
		Name:        "running-security-groups",
		Description: T("List security groups in the set of security groups for running applications"),
		Usage:       "CF_NAME running-security-groups",
	}
}

func (cmd *ListRunningSecurityGroups) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
	if len(fc.Args()) != 0 {
		cmd.ui.Failed(T("Incorrect Usage. No argument required\n\n") + command_registry.Commands.CommandUsage("running-security-groups"))
	}

	requirements := []requirements.Requirement{requirementsFactory.NewLoginRequirement()}
	return requirements, nil
}

func (cmd *ListRunningSecurityGroups) SetDependency(deps command_registry.Dependency, pluginCall bool) command_registry.Command {
	cmd.ui = deps.Ui
	cmd.configRepo = deps.Config
	cmd.runningSecurityGroupRepo = deps.RepoLocator.GetRunningSecurityGroupsRepository()
	return cmd
}

func (cmd *ListRunningSecurityGroups) Execute(context flags.FlagContext) {
	cmd.ui.Say(T("Acquiring running security groups as '{{.username}}'", map[string]interface{}{
		"username": terminal.EntityNameColor(cmd.configRepo.Username()),
	}))

	defaultSecurityGroupsFields, err := cmd.runningSecurityGroupRepo.List()
	if err != nil {
		cmd.ui.Failed(err.Error())
	}

	cmd.ui.Ok()
	cmd.ui.Say("")

	if len(defaultSecurityGroupsFields) > 0 {
		for _, value := range defaultSecurityGroupsFields {
			cmd.ui.Say(value.Name)
		}
	} else {
		cmd.ui.Say(T("No running security groups set"))
	}
}
