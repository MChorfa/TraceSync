package compliance

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/MChorfa/TraceSync/internal/artifactmanager"
)

type SBOM struct {
	Name        string      `json:"name"`
	Version     string      `json:"version"`
	Description string      `json:"description"`
	Components  []Component `json:"components"`
	CreatedAt   time.Time   `json:"created_at"`
}

type Component struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"`
	License string `json:"license"`
}

func GenerateSBOM(artifactPath string) error {
	// Read artifact metadata
	metadata, err := artifactmanager.GetArtifactMetadata(artifactPath)
	if err != nil {
		return fmt.Errorf("failed to read artifact metadata: %w", err)
	}

	// Create SBOM structure
	sbom := SBOM{
		Name:        metadata.Name,
		Version:     metadata.Version,
		Description: fmt.Sprintf("SBOM for %s", metadata.Name),
		Components:  []Component{}, // You would populate this based on your artifact's dependencies
		CreatedAt:   time.Now(),
	}

	// For demonstration, we'll add a dummy component
	sbom.Components = append(sbom.Components, Component{
		Name:    "example-dependency",
		Version: "1.0.0",
		Type:    "library",
		License: "MIT",
	})

	// Convert SBOM to JSON
	sbomJSON, err := json.MarshalIndent(sbom, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal SBOM to JSON: %w", err)
	}

	// Write SBOM to file
	sbomPath := filepath.Join(filepath.Dir(artifactPath), fmt.Sprintf("%s-sbom.json", metadata.Name))
	if err := os.WriteFile(sbomPath, sbomJSON, 0644); err != nil {
		return fmt.Errorf("failed to write SBOM file: %w", err)
	}

	fmt.Printf("SBOM generated and saved to: %s\n", sbomPath)
	return nil
}

func PerformComplianceCheck(artifactPath string) error {
	// Read artifact metadata
	metadata, err := artifactmanager.GetArtifactMetadata(artifactPath)
	if err != nil {
		return fmt.Errorf("failed to read artifact metadata: %w", err)
	}

	// Perform compliance checks
	issues := []string{}

	// Check for required metadata fields
	if metadata.Name == "" {
		issues = append(issues, "Artifact name is missing")
	}
	if metadata.Version == "" {
		issues = append(issues, "Artifact version is missing")
	}

	// Check for SBOM existence
	sbomPath := filepath.Join(filepath.Dir(artifactPath), fmt.Sprintf("%s-sbom.json", metadata.Name))
	if _, err := os.Stat(sbomPath); os.IsNotExist(err) {
		issues = append(issues, "SBOM file is missing")
	}

	// Add more compliance checks as needed

	// Report compliance status
	if len(issues) > 0 {
		fmt.Println("Compliance check failed. Issues found:")
		for _, issue := range issues {
			fmt.Printf("- %s\n", issue)
		}
		return fmt.Errorf("compliance check failed")
	}

	fmt.Println("Compliance check passed successfully.")
	return nil
}
