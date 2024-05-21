/*
Copyright © 2024 Lawrence McDaniel <lawrence@querium.com>
*/
package undeploy

import (
	"github.com/QueriumCorp/smarter-cli/cmd"

	"github.com/spf13/cobra"
)

func APIRequest(kind string, kwargs map[string]string) ([]byte, error) {

	return cmd.APIRequest("undeploy/"+kind+"/", kwargs)

}
func ConsoleOutput(bodyJson []byte) {
	cmd.ConsoleOutput(bodyJson)
}
func ErrorOutput(err error) {
	cmd.ErrorOutput(err)
}

// UndeployCmd represents the get command
var UndeployCmd = &cobra.Command{
	Use:   "undeploy <kind> <name>",
	Short: "Undo a Smarter resource deployment.",
	Long: `Undo a Smarter resource deployment. For example:

smarter undeploy <kind> <name>

The Smarter API will undo the deployment of the resource.`,
}

func init() {
	cmd.RootCmd.AddCommand(UndeployCmd)
}