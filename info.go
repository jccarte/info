package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/cf/configuration/config_helpers"
	"github.com/cloudfoundry/cli/cf/configuration/core_config"
	"github.com/cloudfoundry/cli/plugin"
	"github.com/mitchellh/colorstring"
)

func fatalIf(err error) {
	if err != nil {
		fmt.Fprintln(os.Stdout, "error:", err)
		os.Exit(1)
	}
}

func main() {
	plugin.Start(&InfoPlugin{})
}

type InfoPlugin struct{}

func (plugin InfoPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	confRepo := core_config.NewRepositoryFromFilepath(config_helpers.DefaultFilePath(), fatalIf)
	fmt.Println(colorstring.Color("Current User Info"))
	fmt.Println(colorstring.Color("User: [bold][cyan]" + confRepo.UserEmail()))
	fmt.Println(colorstring.Color("Org: [bold][cyan]" + confRepo.OrganizationFields().Name))
	fmt.Println(colorstring.Color("Space: [bold][cyan]" + confRepo.SpaceFields().Name))
	fmt.Println(colorstring.Color("API Version: [bold][cyan]" + confRepo.ApiVersion()))
	fmt.Println(colorstring.Color("API Endpoint: [bold][cyan]" + confRepo.ApiEndpoint()))
}

func (InfoPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "Info",
		Commands: []plugin.Command{
			{
				Name:     "info",
				HelpText: "Print out info about current state",
			},
		},
	}
}
