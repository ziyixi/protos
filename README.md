# Protocol Buffer Code Generator

This repository contains a system for automatically generating language-specific code from Protocol Buffer (protobuf) files, including both regular proto and gRPC service definitions.

## Current Implementation

Currently, this project supports generating Go modules from protobuf files. Support for other languages can be added in the future.

## Directory Structure

- `proto/`: Contains all `.proto` files in potentially recursive directories (e.g., `proto/a/b/c.proto`)
- `go/`: The output directory for generated Go files, maintaining the same structure (e.g., `go/a/b/c.go`)
- `scripts/`: Contains the shell scripts for generating language-specific code

## How It Works

1. The system finds all `.proto` files recursively in the `proto/` directory
2. For each proto file, it generates corresponding code using `protoc` with language-specific plugins
3. Each language maintains its own directory structure and package conventions

### Go Specifics

- Each subdirectory in the generated `go/` directory is treated as a separate Go module
- Appropriate `go.mod` files are created in each subdirectory

## Usage

### Manual Execution

```bash
# Set your GitHub repository path
export MODULE_PREFIX="github.com/yourusername/yourrepository"

# Run the Go generation script
./scripts/generate-go-proto.sh
```

### Automatic Execution via GitHub Actions

The included GitHub workflow automatically runs when:
- Changes are pushed to the `proto/` directory on the main branch
- A pull request affecting the `proto/` directory is created
- The workflow is manually triggered

## Import Paths

Proto files are expected to use imports relative to the repository root, not the `proto/` directory. For example:

```protobuf
// In a proto file
import "proto/todofy/large_language_model.proto";
```

## Generated Packages

Each language has its own import conventions, for example:

### Go

```go
import "github.com/ziyixi/protos/go/todofy"
```

## Dependencies

To run locally, you need:
- Protobuf compiler (`protoc`)
- Language-specific protobuf plugins:

### Go
- Go 1.20 or later
- `protoc-gen-go`
- `protoc-gen-go-grpc`

## Extending to Other Languages

This framework can be extended to support other languages by:

1. Creating new generator scripts in the `scripts/` directory
2. Configuring the appropriate protoc plugins for your target language
3. Updating the GitHub workflow to include the new language generation step

## Customization

If you need to customize the generation process, modify the language-specific scripts in the `scripts/` directory.