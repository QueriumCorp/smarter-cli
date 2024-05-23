/*
Copyright © 2024 Lawrence McDaniel <lawrence@querium.com>
*/
package get

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// apiConnectionCmd represents the sqlconnection command
var apiConnectionCmd = &cobra.Command{
	Use:   "apiconnection",
	Short: "Retrieve a list of PluginDataApiConnections",
	Long: `Retrieve a list of PluginDataApiConnections:

smarter get sqlconnection [flags]

The Smarter API will return a list of PluginDataApiConnection in the specified format,
or a manifest for a specific PluginDataApiConnection.`,
	Run: func(cmd *cobra.Command, args []string) {

		name := viper.GetString("name")

		kwargs := map[string]string{
			"name": name,
		}

		bodyJson, err := APIRequest("PluginDataApiConnection", kwargs)
		if err != nil {
			ErrorOutput(err)
		} else {
			ConsoleOutput(bodyJson)
		}

	},
}

func init() {
	GetCmd.AddCommand(apiConnectionCmd)

	apiConnectionCmd.Flags().StringP("name", "n", "", "Name of the PluginDataApiConnection")

	if err := viper.BindPFlag("name", apiConnectionCmd.Flags().Lookup("name")); err != nil {
		log.Fatalf("Error binding flag: %v", err)
	}
}
