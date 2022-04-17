package cmd

import (
	"fmt"

	"github.com/el10savio/goROT/rot"

	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "decrypt a rot13 encoded message to plain text",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sentence, _ := cmd.Flags().GetString("message")
		fmt.Println(rot.Decrypt(sentence))
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
	decryptCmd.Flags().String("message", "", "Message to decrypt")
}
