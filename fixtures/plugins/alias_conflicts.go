/**
* This test plugin has an alias conlicts with the core command alias 'p'
**/

package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
)

type AliasConflicts struct {
}

func (c *AliasConflicts) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "conflict-cmd" || args[0] == "conflict-alias" {
		cmd()
	}
}

func (c *AliasConflicts) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "AliasConflicts",
		Commands: []plugin.Command{
			{
				Name:     "conflict-cmd",
				Alias:    "conflict-alias",
				HelpText: "help text for AliasConflicts",
			},
		},
	}
}

func cmd() {
	fmt.Println("You called AliasConflicts")
}

func main() {
	plugin.Start(new(AliasConflicts))
}