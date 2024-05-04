/*
Copyright © 2024 Lawrence McDaniel <lawrence@querium.com>
*/
package delete

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// pluginCmd represents the plugin command
var pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Delete a Plugin",
	Long: `Delete a Plugin:

smarter delete plugin -name --dry-run

The Smarter API will permanently delete the Plugin with the specified name,
and dissassociate it from any ChatBots.`,
	Run: func(cmd *cobra.Command, args []string) {

		body, err := GetAPI("plugin")
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			bodyStr, err := json.Marshal(body)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Response:", string(bodyStr))
			}
		}

	},
}

func init() {
	DeleteCmd.AddCommand(pluginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pluginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pluginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
