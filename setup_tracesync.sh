#!/bin/bash
echo "Initializing TraceSync project structure..."

# Create the folder structure
mkdir -p cmd/commands
mkdir -p internal/artifactmanager
mkdir -p internal/compliance
mkdir -p internal/lineage
mkdir -p internal/metadatamanager
mkdir -p internal/storagemanager
mkdir -p internal/telemetry
mkdir -p internal/utils
mkdir -p tests/integration
mkdir -p tests/unit

# Touch key files
touch cmd/commands/upload.go
touch cmd/commands/validate.go
touch cmd/commands/monitor.go
touch cmd/commands/status.go
touch cmd/commands/metadata.go
touch internal/artifactmanager/artifact.go
touch internal/compliance/sbom.go
touch internal/lineage/lineage.go
touch internal/metadatamanager/metadata.go
touch internal/storagemanager/storage.go
touch internal/telemetry/telemetry.go
touch internal/utils/utils.go
touch tests/unit/artifactmanager_test.go

echo "TraceSync project structure created."
