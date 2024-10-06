package lineage

import (
	"fmt"
	"time"
)

type LineageData struct {
	ArtifactID string
	Steps      []LineageStep
}

type LineageStep struct {
	Timestamp time.Time
	Action    string
	Details   map[string]string
}

func GetLineage(artifactID string) (string, error) {
	// In a real implementation, this would fetch data from a database or external service
	// For now, we'll return mock data
	lineage := LineageData{
		ArtifactID: artifactID,
		Steps: []LineageStep{
			{
				Timestamp: time.Now().Add(-24 * time.Hour),
				Action:    "Created",
				Details:   map[string]string{"creator": "user123", "version": "1.0"},
			},
			{
				Timestamp: time.Now().Add(-12 * time.Hour),
				Action:    "Validated",
				Details:   map[string]string{"validator": "user456", "result": "passed"},
			},
			{
				Timestamp: time.Now().Add(-6 * time.Hour),
				Action:    "Uploaded",
				Details:   map[string]string{"uploader": "user789", "destination": "s3://bucket/artifact"},
			},
		},
	}

	return formatLineage(lineage), nil
}

func formatLineage(lineage LineageData) string {
	output := fmt.Sprintf("Lineage for artifact %s:\n", lineage.ArtifactID)
	for _, step := range lineage.Steps {
		output += fmt.Sprintf("- %s: %s\n", step.Timestamp.Format(time.RFC3339), step.Action)
		for key, value := range step.Details {
			output += fmt.Sprintf("  %s: %s\n", key, value)
		}
	}
	return output
}
