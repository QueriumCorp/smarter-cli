/*
Copyright © 2024 Lawrence McDaniel <lawrence@querium.com>
*/
package delete

import (
	"github.com/spf13/cobra"
)

var apikeyCmd = &cobra.Command{
	Use:   "apikey <name>",
	Short: "Delete a SmarterAuthToken",
	Long: `Deletes a SmarterAuthToken:

smarter delete apikey <name> [flags]

The Smarter API will permanently delete the SmarterAuthToken with the specified name.`,
	Run: func(cmd *cobra.Command, args []string) {

		kwargs := map[string]string{
			"name": args[0],
		}
		_, err := APIRequest("SmarterAuthToken", kwargs)
		if err != nil {
			ErrorOutput(err)
		} else {
			ConsoleOutput()
		}

	},
}

func init() {
	deleteCmd.AddCommand(apikeyCmd)
}
