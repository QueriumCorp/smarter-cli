/*
Copyright © 2024 Lawrence McDaniel <lawrence@querium.com>
*/
package describe

import (
	"github.com/QueriumCorp/smarter-cli/cmd"

	"github.com/spf13/cobra"
)

func APIRequest(kind string, kwargs map[string]string) ([]byte, error) {

	return cmd.APIRequest("describe/"+kind+"/", kwargs)

}
func ConsoleOutput(bodyJson []byte) {
	cmd.ConsoleOutput(bodyJson)
}

// describeCmd represents the manifest command
var describeCmd = &cobra.Command{
	Use:   "describe <kind> <name>",
	Short: "Return a manifest for the resource kind",
	Long: `Returns a manifest for the resource kind. For example:

	smarter describe <kind> <name> > my-plugin.yaml

This will generate a manifest for the specified kind of resource and write it to my-plugin.yaml in the current working directory.`,
}

func init() {
	cmd.RootCmd.AddCommand(describeCmd)
}
