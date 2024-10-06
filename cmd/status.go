package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status <artifact>",
	Short: "Check the status of an ongoing or completed artifact transfer",
	Long:  `This command checks the status of an ongoing or completed transfer and displays the results.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		artifact := args[0]
		fmt.Printf("Checking status of artifact: %s\n", artifact)
		// Implement status checking logic here
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
