# TraceSync

## Overview

TraceSync is a CLI tool for managing and synchronizing artifacts, tracking metadata, and ensuring compliance across CI/CD pipelines. 

## Features

- Upload artifacts (datasets, models, SBOMs)
- Manage metadata and track data lineage
- Ensure compliance with TraceGuard's security and provenance standards
- Generate Software Bill of Materials (SBOM)
- Encrypt artifacts for secure storage
- Support for multiple storage backends (AWS S3, Google Cloud Storage, MinIO)

# Pre-requisites

- Go 1.16 or later
- Cobra CLI:
- ```bash
  go get -u github.com/spf13/cobra
  ```
- ```bash
    go get -u github.com/spf13/viper
    ```
- ```bash
    go get -u github.com/spf13/pflag
    ```
- ```bash
    go get -u  github.com/spf13/cobra-cli@latest
    ```

    cobra-cli init tracesync

## Installation

Run the following command to install the CLI:

```bash
go install github.com/MChorfa/tracesync@latest
```

## Usage

Here are the available commands and their options:

### Upload an artifact

```bash
tracesync upload /path/to/artifact
```

### Validate an artifact

```bash
tracesync validate /path/to/artifact
```

### Monitor artifact lineage and quality

```bash
tracesync monitor /path/to/artifact
```

### View or update artifact metadata

```bash
tracesync metadata /path/to/artifact
```

### Check artifact status

```bash
tracesync status /path/to/artifact
```

## Running Tests

```bash
go test ./tests/unit
```

## Development

To add a new command, use the following command:

```bash
cobra-cli add cmd/yourcommand.go
```

## Build

```bash
go build -o tracesync main.go
```

## Run

```bash
go run main.go upload /path/to/artifact | validate /path/to/artifact | monitor /path/to/artifact | metadata /path/to/artifact | status /path/to/artifact
```

## Contributing

Please read CONTRIBUTING.md for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the Apache License, Version 2.0 - see the LICENSE file for details.

## CI/CD Pipeline

This project uses Dagger for CI/CD. To run the pipeline locally:

1. Install Dagger CLI:
   ```bash
   curl -L https://dl.dagger.io/dagger/install.sh | sh
   ```

2. Run the pipeline:
   ```bash
   go run ci/pipeline.go
   ```

This will build the project, run tests, perform a security scan, and output the binary.
