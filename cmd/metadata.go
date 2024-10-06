package cmd

import (
	"fmt"

	"github.com/MChorfa/TraceSync/internal/artifactmanager"
	"github.com/spf13/cobra"
)

var metadataCmd = &cobra.Command{
	Use:   "metadata <artifact>",
	Short: "Manage metadata tagging for artifacts",
	Long:  `This command allows you to add metadata tags to artifacts for classification and easier retrieval.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		artifact := args[0]
		fmt.Printf("Managing metadata for artifact: %s\n", artifact)

		metadata, _ := cmd.Flags().GetStringToString("add")
		if len(metadata) > 0 {
			err := artifactmanager.TagArtifact(artifact, metadata)
			if err != nil {
				fmt.Printf("Error adding metadata: %v\n", err)
				return
			}
			fmt.Println("Metadata added successfully.")
		}

		// Implement logic to display current metadata
	},
}

func init() {
	rootCmd.AddCommand(metadataCmd)
	metadataCmd.Flags().StringToStringP("add", "a", nil, "Add metadata key-value pairs")
}
