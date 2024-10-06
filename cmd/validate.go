package cmd

import (
	"fmt"

	"github.com/MChorfa/TraceSync/internal/artifactmanager"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate <dataset>",
	Short: "Validate the quality and completeness of a dataset",
	Long:  `This command runs quality checks to ensure that the dataset is complete and valid for use.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		artifact := args[0]
		fmt.Printf("Validating dataset: %s\n", artifact)

		err := artifactmanager.ValidateArtifact(artifact)
		if err != nil {
			fmt.Printf("Validation failed: %v\n", err)
			return
		}

		fmt.Println("Validation successful.")
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
