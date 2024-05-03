/*
Copyright © 2024 Lawrence McDaniel <lawrence@querium.com>
*/
package delete

import (
	"fmt"

	"github.com/spf13/cobra"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Delete a chat history",
	Long: `Delete a chat history:

smarter delete chat -id

The Smarter API will permanently delete the chat history with the specified identifier.`,
	Run: func(cmd *cobra.Command, args []string) {

		body, err := GetAPI("chat")
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Response:", string(body))
		}

	},
}

func init() {
	DeleteCmd.AddCommand(chatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
