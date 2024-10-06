package artifactmanager

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

type ArtifactMetadata struct {
	Name      string            `yaml:"name"`
	Version   string            `yaml:"version"`
	CreatedAt time.Time         `yaml:"created_at"`
	UpdatedAt time.Time         `yaml:"updated_at"`
	Tags      map[string]string `yaml:"tags"`
	Lineage   []LineageEntry    `yaml:"lineage"`
}

type LineageEntry struct {
	Timestamp time.Time         `yaml:"timestamp"`
	Action    string            `yaml:"action"`
	Details   map[string]string `yaml:"details"`
}

func TagArtifact(artifactPath string, metadata map[string]string) error {
	metadataFilePath := filepath.Join(filepath.Dir(artifactPath), "ModelDescriptor.yaml")

	var artifactMetadata ArtifactMetadata
	if _, err := os.Stat(metadataFilePath); err == nil {
		// File exists, read existing metadata
		data, err := os.ReadFile(metadataFilePath)
		if err != nil {
			return fmt.Errorf("failed to read metadata file: %w", err)
		}
		if err := yaml.Unmarshal(data, &artifactMetadata); err != nil {
			return fmt.Errorf("failed to unmarshal metadata: %w", err)
		}
	} else {
		// File doesn't exist, initialize new metadata
		artifactMetadata = ArtifactMetadata{
			Name:      filepath.Base(artifactPath),
			Version:   "1.0",
			CreatedAt: time.Now(),
			Tags:      make(map[string]string),
		}
	}

	// Update metadata
	artifactMetadata.UpdatedAt = time.Now()
	for key, value := range metadata {
		artifactMetadata.Tags[key] = value
	}

	// Write updated metadata to file
	data, err := yaml.Marshal(artifactMetadata)
	if err != nil {
		return fmt.Errorf("failed to marshal metadata: %w", err)
	}
	if err := os.WriteFile(metadataFilePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write metadata file: %w", err)
	}

	return nil
}

func TrackLineage(artifactPath string, details map[string]string) error {
	metadataFilePath := filepath.Join(filepath.Dir(artifactPath), "ModelDescriptor.yaml")

	var artifactMetadata ArtifactMetadata
	if _, err := os.Stat(metadataFilePath); err == nil {
		// File exists, read existing metadata
		data, err := os.ReadFile(metadataFilePath)
		if err != nil {
			return fmt.Errorf("failed to read metadata file: %w", err)
		}
		if err := yaml.Unmarshal(data, &artifactMetadata); err != nil {
			return fmt.Errorf("failed to unmarshal metadata: %w", err)
		}
	} else {
		return errors.New("metadata file not found, please tag the artifact first")
	}

	// Add new lineage entry
	newEntry := LineageEntry{
		Timestamp: time.Now(),
		Action:    "Transformation",
		Details:   details,
	}
	artifactMetadata.Lineage = append(artifactMetadata.Lineage, newEntry)

	// Write updated metadata to file
	data, err := yaml.Marshal(artifactMetadata)
	if err != nil {
		return fmt.Errorf("failed to marshal metadata: %w", err)
	}
	if err := os.WriteFile(metadataFilePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write metadata file: %w", err)
	}

	return nil
}

func ValidateArtifact(artifactPath string) error {
	// Check if the artifact file exists
	if _, err := os.Stat(artifactPath); os.IsNotExist(err) {
		return fmt.Errorf("artifact file does not exist: %s", artifactPath)
	}

	// Check if the metadata file exists
	metadataFilePath := filepath.Join(filepath.Dir(artifactPath), "ModelDescriptor.yaml")
	if _, err := os.Stat(metadataFilePath); os.IsNotExist(err) {
		return fmt.Errorf("metadata file does not exist: %s", metadataFilePath)
	}

	// Read and parse the metadata file
	data, err := os.ReadFile(metadataFilePath)
	if err != nil {
		return fmt.Errorf("failed to read metadata file: %w", err)
	}

	var artifactMetadata ArtifactMetadata
	if err := yaml.Unmarshal(data, &artifactMetadata); err != nil {
		return fmt.Errorf("failed to unmarshal metadata: %w", err)
	}

	// Perform basic validation checks
	if artifactMetadata.Name == "" {
		return errors.New("artifact name is missing in metadata")
	}
	if artifactMetadata.Version == "" {
		return errors.New("artifact version is missing in metadata")
	}
	if artifactMetadata.CreatedAt.IsZero() {
		return errors.New("artifact creation date is missing in metadata")
	}

	// Add more specific validation checks as needed for your use case

	return nil
}

// ... existing helper functions ...

func GetArtifactMetadata(artifactPath string) (ArtifactMetadata, error) {
	metadataFilePath := filepath.Join(filepath.Dir(artifactPath), "ModelDescriptor.yaml")

	var artifactMetadata ArtifactMetadata
	if _, err := os.Stat(metadataFilePath); err == nil {
		// File exists, read existing metadata
		data, err := os.ReadFile(metadataFilePath)
		if err != nil {
			return ArtifactMetadata{}, fmt.Errorf("failed to read metadata file: %w", err)
		}
		if err := yaml.Unmarshal(data, &artifactMetadata); err != nil {
			return ArtifactMetadata{}, fmt.Errorf("failed to unmarshal metadata: %w", err)
		}
	} else {
		return ArtifactMetadata{}, errors.New("metadata file not found")
	}

	return artifactMetadata, nil
}
