/*
Copyright © 2024 Lawrence McDaniel <lawrence@querium.com>
*/
package get

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Retrieve your Account manifest",
	Long: `Retrieve your Account manifest:

smarter get account --json --yaml -n <10> --asc --desc

The Smarter API will your Account manifest.`,

	Run: func(cmd *cobra.Command, args []string) {

		name := viper.GetString("name")

		kwargs := map[string]string{
			"name": name,
		}
		bodyJson, err := APIRequest("Account", kwargs)
		if err != nil {
			ErrorOutput(err)
		} else {
			ConsoleOutput(bodyJson)
		}

	},
}

func init() {
	GetCmd.AddCommand(accountCmd)
}
