package storagemanager

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
)

// EncryptArtifact encrypts the given artifact and returns the path to the encrypted file
func EncryptArtifact(artifactPath string) (string, error) {
	// Read the artifact file
	plaintext, err := os.ReadFile(artifactPath)
	if err != nil {
		return "", fmt.Errorf("failed to read artifact file: %w", err)
	}

	// Generate a random 32-byte key for AES-256
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return "", fmt.Errorf("failed to generate encryption key: %w", err)
	}

	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %w", err)
	}

	// Generate a random nonce
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("failed to generate nonce: %w", err)
	}

	// Encrypt the plaintext
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %w", err)
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	// Combine nonce and ciphertext
	encryptedData := append(nonce, ciphertext...)

	// Encode the encrypted data as base64
	encodedData := base64.StdEncoding.EncodeToString(encryptedData)

	// Write the encrypted data to a new file
	encryptedPath := artifactPath + ".enc"
	if err := os.WriteFile(encryptedPath, []byte(encodedData), 0644); err != nil {
		return "", fmt.Errorf("failed to write encrypted file: %w", err)
	}

	return encryptedPath, nil
}

// UploadArtifact uploads the encrypted artifact to the specified storage backend
func UploadArtifact(encryptedArtifactPath, backend string) error {
	switch backend {
	case "aws":
		return uploadToAWS(encryptedArtifactPath)
	case "gcs":
		return uploadToGCS(encryptedArtifactPath)
	case "minio":
		return uploadToMinIO(encryptedArtifactPath)
	default:
		return errors.New("unsupported storage backend")
	}
}

// SwitchBackend determines the appropriate storage backend based on the environment
func SwitchBackend(env string) string {
	switch env {
	case "production":
		return "aws"
	case "staging":
		return "gcs"
	default:
		return "minio"
	}
}

// Helper functions for different storage backends
func uploadToAWS(filePath string) error {
	// Implement AWS S3 upload logic here
	fmt.Printf("Uploading %s to AWS S3\n", filePath)
	return nil
}

func uploadToGCS(filePath string) error {
	// Implement Google Cloud Storage upload logic here
	fmt.Printf("Uploading %s to Google Cloud Storage\n", filePath)
	return nil
}

func uploadToMinIO(filePath string) error {
	// Implement MinIO upload logic here
	fmt.Printf("Uploading %s to MinIO\n", filePath)
	return nil
}
