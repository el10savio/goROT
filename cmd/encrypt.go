package cmd

import (
	"fmt"

	"github.com/el10savio/goROT/rot"
	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "encrypt a plain text sentence to rot13",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sentence, _ := cmd.Flags().GetString("message")
		fmt.Println(rot.Encrypt(sentence))
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
	encryptCmd.Flags().String("message", "", "Message to encrypt")
}
