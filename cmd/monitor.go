package cmd

import (
	"fmt"

	"github.com/MChorfa/TraceSync/internal/lineage"
	"github.com/MChorfa/TraceSync/internal/telemetry"
	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor [artifact]",
	Short: "Monitor data lineage and quality of artifacts",
	Long:  `This command monitors the data lineage, quality, and other metrics for datasets or models.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			artifact := args[0]
			fmt.Printf("Monitoring artifact: %s\n", artifact)
			monitorArtifact(artifact)
		} else {
			fmt.Println("Monitoring all artifacts...")
			monitorAllArtifacts()
		}
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
	monitorCmd.Flags().BoolP("lineage", "l", false, "Show data lineage")
	monitorCmd.Flags().BoolP("quality", "q", false, "Show quality metrics")
}

func monitorArtifact(artifact string) {
	showLineage, _ := monitorCmd.Flags().GetBool("lineage")
	showQuality, _ := monitorCmd.Flags().GetBool("quality")

	if showLineage {
		lineageData, err := lineage.GetLineage(artifact)
		if err != nil {
			fmt.Printf("Error retrieving lineage data: %v\n", err)
		} else {
			fmt.Printf("Lineage for %s:\n%s\n", artifact, lineageData)
		}
	}

	if showQuality {
		qualityMetrics, err := telemetry.GetQualityMetrics(artifact)
		if err != nil {
			fmt.Printf("Error retrieving quality metrics: %v\n", err)
		} else {
			fmt.Printf("Quality metrics for %s:\n%s\n", artifact, qualityMetrics)
		}
	}

	if !showLineage && !showQuality {
		fmt.Println("Please specify --lineage or --quality flag to view specific information.")
	}
}

func monitorAllArtifacts() {
	artifacts, err := telemetry.GetAllArtifacts()
	if err != nil {
		fmt.Printf("Error retrieving artifacts: %v\n", err)
		return
	}

	for _, artifact := range artifacts {
		fmt.Printf("Monitoring artifact: %s\n", artifact)
		monitorArtifact(artifact)
		fmt.Println("---")
	}
}
