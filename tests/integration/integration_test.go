package integration

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/MChorfa/TraceSync/internal/artifactmanager"
	"github.com/MChorfa/TraceSync/internal/compliance"
	"github.com/MChorfa/TraceSync/internal/storagemanager"
)

func TestArtifactUploadFlow(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "tracesync-integration-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a mock artifact file
	artifactPath := filepath.Join(tempDir, "test-artifact")
	if err := os.WriteFile(artifactPath, []byte("test artifact content"), 0644); err != nil {
		t.Fatalf("Failed to create test artifact: %v", err)
	}

	// Step 1: Tag the artifact
	metadata := map[string]string{
		"version": "1.0.0",
		"author":  "Integration Test",
	}
	err = artifactmanager.TagArtifact(artifactPath, metadata)
	if err != nil {
		t.Fatalf("TagArtifact failed: %v", err)
	}

	// Step 2: Validate the artifact
	err = artifactmanager.ValidateArtifact(artifactPath)
	if err != nil {
		t.Fatalf("ValidateArtifact failed: %v", err)
	}

	// Step 3: Generate SBOM
	err = compliance.GenerateSBOM(artifactPath)
	if err != nil {
		t.Fatalf("GenerateSBOM failed: %v", err)
	}

	// Step 4: Perform compliance check
	err = compliance.PerformComplianceCheck(artifactPath)
	if err != nil {
		t.Fatalf("PerformComplianceCheck failed: %v", err)
	}

	// Step 5: Encrypt the artifact
	encryptedPath, err := storagemanager.EncryptArtifact(artifactPath)
	if err != nil {
		t.Fatalf("EncryptArtifact failed: %v", err)
	}

	// Step 6: Upload the encrypted artifact
	err = storagemanager.UploadArtifact(encryptedPath, "minio")
	if err != nil {
		t.Fatalf("UploadArtifact failed: %v", err)
	}

	t.Log("Integration test completed successfully")
}
