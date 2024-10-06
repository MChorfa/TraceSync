package unit

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/MChorfa/TraceSync/internal/artifactmanager"
)

func TestTagArtifact(t *testing.T) {
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

	// Test tagging the artifact
	metadata := map[string]string{
		"version": "1.0.0",
		"author":  "Test Author",
	}
	err = artifactmanager.TagArtifact(artifactPath, metadata)
	if err != nil {
		t.Fatalf("TagArtifact failed: %v", err)
	}

	// Verify that the metadata file was created
	metadataPath := filepath.Join(tempDir, "ModelDescriptor.yaml")
	if _, err := os.Stat(metadataPath); os.IsNotExist(err) {
		t.Errorf("Metadata file was not created")
	}

	// Verify the content of the metadata file
	storedMetadata, err := artifactmanager.GetArtifactMetadata(artifactPath)
	if err != nil {
		t.Fatalf("Failed to read metadata: %v", err)
	}

	if storedMetadata.Name != "test-artifact" {
		t.Errorf("Expected artifact name 'test-artifact', got '%s'", storedMetadata.Name)
	}
	if storedMetadata.Version != "1.0.0" {
		t.Errorf("Expected version '1.0.0', got '%s'", storedMetadata.Version)
	}
	if storedMetadata.Tags["author"] != "Test Author" {
		t.Errorf("Expected author 'Test Author', got '%s'", storedMetadata.Tags["author"])
	}
}

func TestValidateArtifact(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "tracesync-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Test case 1: Non-existent artifact
	err = artifactmanager.ValidateArtifact(filepath.Join(tempDir, "non-existent-artifact"))
	if err == nil {
		t.Errorf("Expected error for non-existent artifact, but got nil")
	}

	// Create a mock artifact file
	artifactPath := filepath.Join(tempDir, "test-artifact")
	if err := os.WriteFile(artifactPath, []byte("test artifact content"), 0644); err != nil {
		t.Fatalf("Failed to create test artifact: %v", err)
	}

	// Test case 2: Artifact without metadata
	err = artifactmanager.ValidateArtifact(artifactPath)
	if err == nil {
		t.Errorf("Expected error for artifact without metadata, but got nil")
	}

	// Create metadata for the artifact
	metadata := map[string]string{
		"version": "1.0.0",
		"author":  "Test Author",
	}
	err = artifactmanager.TagArtifact(artifactPath, metadata)
	if err != nil {
		t.Fatalf("Failed to tag artifact: %v", err)
	}

	// Test case 3: Valid artifact with metadata
	err = artifactmanager.ValidateArtifact(artifactPath)
	if err != nil {
		t.Errorf("Expected no error for valid artifact, but got: %v", err)
	}
}
