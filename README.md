# Protocol Buffer Code Generator

This repository contains a system for automatically generating language-specific code from Protocol Buffer (protobuf) files, including both regular proto and gRPC service definitions.

## Branching Strategy

This project utilizes a two-branch strategy:

-   **`protobuf` branch**: This branch is the source of truth for all `.proto` definition files. All changes and updates to your protobuf schemas should be made on this branch.
-   **`main` branch**: This branch contains the generated language-specific code (e.g., Go modules), the generation scripts, the GitHub Actions workflow, and this README. Generated code is automatically committed to this branch.

## Current Implementation

Currently, this project supports generating Go modules from protobuf files. Support for other languages can be added in the future.

## Directory Structure

### `protobuf` branch:
-   `proto/`: Contains all `.proto` files in potentially recursive directories (e.g., `proto/a/b/c.proto`). This is the primary directory for your schema definitions.

### `main` branch:
-   `go/`: The output directory for generated Go files, maintaining a clean structure based on the proto paths (e.g., `go/a/b/` would contain code generated from `proto/a/b/*.proto`).
-   `scripts/`: Contains the shell scripts for generating language-specific code (e.g., `generate-go-proto.sh`).
-   `.github/workflows/`: Contains the GitHub Actions workflow for automation.

## How It Works

1.  Modifications to `.proto` files are pushed to the `protobuf` branch.
2.  A GitHub Action (defined on the `main` branch) is triggered by pushes to the `proto/` directory on the `protobuf` branch.
3.  The GitHub Action checks out the `main` branch.
4.  It then checks out the `protobuf` branch into a temporary directory to access the `.proto` files.
5.  The generation script (`scripts/generate-go-proto.sh`) is executed:
    * It uses the root of the temporary `protobuf` checkout as the `--proto_path` base.
    * It finds all `.proto` files recursively within the `proto/` subdirectory of the temporary checkout.
    * For each proto file, it generates corresponding code using `protoc` with language-specific plugins.
    * Generated files are initially placed by `protoc` (due to `paths=source_relative`) into a path like `go/proto/a/b/` and then moved by the script to the final desired location (e.g., `go/a/b/`).
6.  If changes are detected in the generated code (e.g., in the `go/` directory on the `main` branch), the GitHub Action commits and pushes these changes back to the `main` branch.

### Go Specifics

-   Each subdirectory in the generated `go/` directory (e.g., `go/todofy`) is treated as a separate Go module.
-   Appropriate `go.mod` and `go.sum` files are created/updated in each of these subdirectories.

## Usage

### Automatic Execution via GitHub Actions

The primary way to use this system is through the included GitHub workflow:

1.  Make your changes to `.proto` files in your local clone of the `protobuf` branch.
2.  Commit and push these changes to the `protobuf` branch on GitHub.
3.  The GitHub Action will automatically:
    * Detect the changes.
    * Generate the corresponding Go code.
    * Commit and push the generated Go code to the `go/` directory on the `main` branch.

The workflow also allows for manual triggering if needed.

### Manual Execution (for local testing/development)

You can run the generation script locally, typically from the root of your `main` branch checkout. To simulate the CI environment where proto files are in a different location, you might need to manually check out the `protobuf` branch into a subdirectory.

```bash
# Ensure you are on the main branch
# git checkout main

# (Optional) Create a temporary directory for protos from the protobuf branch
# mkdir -p temp_protobuf_checkout
# git --work-tree=temp_protobuf_checkout checkout protobuf -- proto
# Or, if you have both branches locally and just want to test the script logic
# against protos in a known relative path:

# Set environment variables:
export MODULE_PREFIX="[github.com/yourusername/yourrepository](https://github.com/yourusername/yourrepository)"

# If your .proto files are in "./proto" relative to your current directory:
export PROTO_BASE_DIR_ENV="."
export PROTO_SUBDIR_NAME_ENV="proto"

# If you checked out the 'protobuf' branch into 'temp_protobuf_checkout'
# and your protos are in 'temp_protobuf_checkout/proto':
# export PROTO_BASE_DIR_ENV="temp_protobuf_checkout"
# export PROTO_SUBDIR_NAME_ENV="proto"

# Run the Go generation script
./scripts/generate-go-proto.sh