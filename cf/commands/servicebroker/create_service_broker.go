package servicebroker

import (
	"github.com/blang/semver"
	"github.com/cloudfoundry/cli/cf/api"
	"github.com/cloudfoundry/cli/cf/command_registry"
	"github.com/cloudfoundry/cli/cf/configuration/core_config"
	. "github.com/cloudfoundry/cli/cf/i18n"
	"github.com/cloudfoundry/cli/cf/requirements"
	"github.com/cloudfoundry/cli/cf/terminal"
	"github.com/cloudfoundry/cli/flags"
	"github.com/cloudfoundry/cli/flags/flag"
)

type CreateServiceBroker struct {
	ui                terminal.UI
	config            core_config.Reader
	serviceBrokerRepo api.ServiceBrokerRepository
}

func init() {
	command_registry.Register(&CreateServiceBroker{})
}

func (cmd *CreateServiceBroker) MetaData() command_registry.CommandMetadata {
	fs := make(map[string]flags.FlagSet)
	fs["space-scoped"] = &cliFlags.BoolFlag{Name: "space-scoped", Usage: T("Make the broker's service plans only visible within the targeted space")}

	return command_registry.CommandMetadata{
		Name:        "create-service-broker",
		Description: T("Create a service broker"),
		Usage:       T("CF_NAME create-service-broker SERVICE_BROKER USERNAME PASSWORD URL [--space-scoped]"),
		Flags:       fs,
	}
}

func (cmd *CreateServiceBroker) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
	if len(fc.Args()) != 4 {
		cmd.ui.Failed(T("Incorrect Usage. Requires SERVICE_BROKER, USERNAME, PASSWORD, URL as arguments\n\n") + command_registry.Commands.CommandUsage("create-service-broker"))
	}

	var reqs = []requirements.Requirement{
		requirementsFactory.NewLoginRequirement(),
	}

	if fc.IsSet("space-scoped") {
		requiredVersion, err := semver.Make("2.47.0")
		if err != nil {
			panic(err.Error())
		}

		reqs = append(
			reqs,
			requirementsFactory.NewTargetedSpaceRequirement(),
			requirementsFactory.NewMinAPIVersionRequirement("--space-scoped", requiredVersion),
		)
	}

	return reqs, nil
}

func (cmd *CreateServiceBroker) SetDependency(deps command_registry.Dependency, pluginCall bool) command_registry.Command {
	cmd.ui = deps.Ui
	cmd.config = deps.Config
	cmd.serviceBrokerRepo = deps.RepoLocator.GetServiceBrokerRepository()
	return cmd
}

func (cmd *CreateServiceBroker) Execute(c flags.FlagContext) {
	name := c.Args()[0]
	username := c.Args()[1]
	password := c.Args()[2]
	url := c.Args()[3]

	cmd.ui.Say(T("Creating service broker {{.Name}} as {{.Username}}...",
		map[string]interface{}{
			"Name":     terminal.EntityNameColor(name),
			"Username": terminal.EntityNameColor(cmd.config.Username())}))

	var err error
	if c.Bool("space-scoped") {
		err = cmd.serviceBrokerRepo.Create(name, url, username, password, cmd.config.SpaceFields().Guid)
	} else {
		err = cmd.serviceBrokerRepo.Create(name, url, username, password, "")
	}

	if err != nil {
		cmd.ui.Failed(err.Error())
	}

	cmd.ui.Ok()
}
