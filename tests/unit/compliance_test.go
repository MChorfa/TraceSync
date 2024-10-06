package unit

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/MChorfa/TraceSync/internal/artifactmanager"
	"github.com/MChorfa/TraceSync/internal/compliance"
)

func TestGenerateSBOM(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "tracesync-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a mock artifact file
	artifactPath := filepath.Join(tempDir, "test-artifact")
	if err := os.WriteFile(artifactPath, []byte("test artifact content"), 0644); err != nil {
		t.Fatalf("Failed to create test artifact: %v", err)
	}

	// Create mock metadata
	metadata := artifactmanager.ArtifactMetadata{
		Name:    "test-artifact",
		Version: "1.0.0",
	}
	if err := artifactmanager.TagArtifact(artifactPath, map[string]string{"version": "1.0.0"}); err != nil {
		t.Fatalf("Failed to create test metadata: %v", err)
	}

	// Generate SBOM
	if err := compliance.GenerateSBOM(artifactPath); err != nil {
		t.Fatalf("GenerateSBOM failed: %v", err)
	}

	// Check if SBOM file was created
	sbomPath := filepath.Join(tempDir, "test-artifact-sbom.json")
	if _, err := os.Stat(sbomPath); os.IsNotExist(err) {
		t.Errorf("SBOM file was not created")
	}

	// TODO: Add more detailed checks on the content of the SBOM file
}

func TestPerformComplianceCheck(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "tracesync-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a mock artifact file
	artifactPath := filepath.Join(tempDir, "test-artifact")
	if err := os.WriteFile(artifactPath, []byte("test artifact content"), 0644); err != nil {
		t.Fatalf("Failed to create test artifact: %v", err)
	}

	// Test case 1: Compliance check should fail due to missing metadata
	err = compliance.PerformComplianceCheck(artifactPath)
	if err == nil {
		t.Errorf("Expected compliance check to fail, but it passed")
	}

	// Create mock metadata
	if err := artifactmanager.TagArtifact(artifactPath, map[string]string{"version": "1.0.0"}); err != nil {
		t.Fatalf("Failed to create test metadata: %v", err)
	}

	// Generate SBOM
	if err := compliance.GenerateSBOM(artifactPath); err != nil {
		t.Fatalf("GenerateSBOM failed: %v", err)
	}

	// Test case 2: Compliance check should pass
	err = compliance.PerformComplianceCheck(artifactPath)
	if err != nil {
		t.Errorf("Expected compliance check to pass, but it failed: %v", err)
	}
}
