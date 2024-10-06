package unit

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/MChorfa/TraceSync/internal/storagemanager"
)

func TestEncryptArtifact(t *testing.T) {
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

	// Encrypt the artifact
	encryptedPath, err := storagemanager.EncryptArtifact(artifactPath)
	if err != nil {
		t.Fatalf("EncryptArtifact failed: %v", err)
	}

	// Check if the encrypted file was created
	if _, err := os.Stat(encryptedPath); os.IsNotExist(err) {
		t.Errorf("Encrypted file was not created")
	}

	// Check if the encrypted file is different from the original
	originalContent, err := os.ReadFile(artifactPath)
	if err != nil {
		t.Fatalf("Failed to read original artifact: %v", err)
	}
	encryptedContent, err := os.ReadFile(encryptedPath)
	if err != nil {
		t.Fatalf("Failed to read encrypted artifact: %v", err)
	}
	if string(originalContent) == string(encryptedContent) {
		t.Errorf("Encrypted content is identical to original content")
	}
}

func TestUploadArtifact(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "tracesync-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a mock encrypted artifact file
	encryptedArtifactPath := filepath.Join(tempDir, "test-artifact.enc")
	if err := os.WriteFile(encryptedArtifactPath, []byte("encrypted content"), 0644); err != nil {
		t.Fatalf("Failed to create test encrypted artifact: %v", err)
	}

	// Test uploading to different backends
	backends := []string{"aws", "gcs", "minio"}
	for _, backend := range backends {
		err := storagemanager.UploadArtifact(encryptedArtifactPath, backend)
		if err != nil {
			t.Errorf("UploadArtifact failed for backend %s: %v", backend, err)
		}
	}

	// Test uploading to an unsupported backend
	err = storagemanager.UploadArtifact(encryptedArtifactPath, "unsupported")
	if err == nil {
		t.Errorf("Expected error for unsupported backend, but got nil")
	}
}

func TestSwitchBackend(t *testing.T) {
	testCases := []struct {
		env      string
		expected string
	}{
		{"production", "aws"},
		{"staging", "gcs"},
		{"development", "minio"},
		{"", "minio"},
	}

	for _, tc := range testCases {
		result := storagemanager.SwitchBackend(tc.env)
		if result != tc.expected {
			t.Errorf("SwitchBackend(%s) = %s; want %s", tc.env, result, tc.expected)
		}
	}
}
