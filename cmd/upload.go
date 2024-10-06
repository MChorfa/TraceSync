package cmd

import (
	"fmt"
	"os"

	"github.com/MChorfa/TraceSync/internal/artifactmanager"
	"github.com/MChorfa/TraceSync/internal/compliance"
	"github.com/MChorfa/TraceSync/internal/storagemanager"
	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload <artifact>",
	Short: "Upload an artifact to TraceSync",
	Long:  `This command uploads artifacts (datasets, models, reports) to TraceSync's storage system.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		artifact := args[0]
		fmt.Printf("Uploading artifact: %s\n", artifact)

		// Validate artifact
		if err := artifactmanager.ValidateArtifact(artifact); err != nil {
			fmt.Printf("Artifact validation failed: %v\n", err)
			return
		}

		// Generate SBOM
		if err := compliance.GenerateSBOM(artifact); err != nil {
			fmt.Printf("SBOM generation failed: %v\n", err)
			return
		}

		// Perform compliance check
		if err := compliance.PerformComplianceCheck(artifact); err != nil {
			fmt.Printf("Compliance check failed: %v\n", err)
			return
		}

		// Encrypt artifact
		encryptedArtifact, err := storagemanager.EncryptArtifact(artifact)
		if err != nil {
			fmt.Printf("Artifact encryption failed: %v\n", err)
			return
		}

		// Determine storage backend
		backend := storagemanager.SwitchBackend(os.Getenv("TRACESYNC_ENV"))

		// Upload encrypted artifact
		if err := storagemanager.UploadArtifact(encryptedArtifact, backend); err != nil {
			fmt.Printf("Artifact upload failed: %v\n", err)
			return
		}

		fmt.Println("Artifact uploaded successfully.")
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	uploadCmd.Flags().StringToStringP("metadata", "m", nil, "Metadata key-value pairs")
	uploadCmd.Flags().StringToStringP("lineage", "l", nil, "Data lineage information")
}
