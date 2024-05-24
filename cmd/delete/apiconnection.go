/*
Copyright © 2024 Lawrence McDaniel <lawrence@querium.com>
*/
package delete

import (
	"github.com/spf13/cobra"
)

var apiConnectionCmd = &cobra.Command{
	Use:   "apiconnection <name>",
	Short: "Delete a PluginDataApiConnection",
	Long: `Deletes a PluginDataApiConnection:

smarter delete apiconnection <name>  [flags]

The Smarter API will permanently delete the PluginDataApiConnection with the specified name.`,
	Run: func(cmd *cobra.Command, args []string) {

		kwargs := map[string]string{
			"name": args[0],
		}
		_, err := APIRequest("PluginDataApiConnection", kwargs)
		if err != nil {
			ErrorOutput(err)
		} else {
			ConsoleOutput()
		}

	},
}

func init() {
	deleteCmd.AddCommand(apiConnectionCmd)
}
