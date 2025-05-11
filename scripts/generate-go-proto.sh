#!/bin/bash
set -e

# This script generates Go modules from .proto files.
# It's designed to be called by a GitHub Action.

# The root directory where .proto files are located.
# This can be overridden by setting the PROTO_DIR_ENV environment variable.
# For example, if checking out another branch's protos into "protobuf_checkout/proto_definitions",
# set PROTO_DIR_ENV="protobuf_checkout/proto_definitions".
ACTUAL_PROTO_DIR="${PROTO_DIR_ENV:-proto}" # Default to "proto" if PROTO_DIR_ENV is not set

# Output directory for generated Go files (relative to the repository root)
GO_OUT_DIR="go"

# Ensure MODULE_PREFIX is set (e.g., github.com/username/repository)
if [ -z "${MODULE_PREFIX}" ]; then
    echo "Error: MODULE_PREFIX is not set. Please set it to your GitHub repository path."
    exit 1
fi

# Check for necessary tools
if ! command -v protoc &> /dev/null; then
    echo "Error: protoc is not installed. Please install it to generate Go code."
    exit 1
fi
if ! command -v protoc-gen-go &> /dev/null; then
    echo "Error: protoc-gen-go is not installed. Please install it to generate Go code."
    exit 1
fi
if ! command -v protoc-gen-go-grpc &> /dev/null; then
    echo "Error: protoc-gen-go-grpc is not installed. Please install it to generate Go code."
    exit 1
fi

# Create the base output directory for Go files if it doesn't exist
mkdir -p "${GO_OUT_DIR}"

echo "Searching for .proto files in ${ACTUAL_PROTO_DIR}..."

# Find all .proto files recursively in the ACTUAL_PROTO_DIR
find "${ACTUAL_PROTO_DIR}" -name "*.proto" | while read proto_file_full_path; do
    # Get the path of the .proto file relative to ACTUAL_PROTO_DIR
    # e.g., if ACTUAL_PROTO_DIR is "temp_protos/proto" and proto_file_full_path is "temp_protos/proto/todofy/user.proto",
    # then proto_file_relative_path will be "todofy/user.proto"
    proto_file_relative_path="${proto_file_full_path#${ACTUAL_PROTO_DIR}/}"

    # Get the directory part of this relative path
    # e.g., "todofy" or "." if the file is at the root of ACTUAL_PROTO_DIR
    proto_file_rel_dir=$(dirname "${proto_file_relative_path}")

    echo "Processing ${proto_file_full_path} (relative: ${proto_file_relative_path})..."

    # Determine the output subdirectory for Go files
    # If proto_file_rel_dir is ".", use GO_OUT_DIR directly. Otherwise, append the subdirectory.
    go_out_subdir="${GO_OUT_DIR}"
    if [ "${proto_file_rel_dir}" != "." ]; then
        go_out_subdir="${GO_OUT_DIR}/${proto_file_rel_dir}"
    fi
    mkdir -p "${go_out_subdir}"

    echo "Generating Go code for ${proto_file_full_path} into ${go_out_subdir}..."

    # Generate Go files using protoc with gRPC plugin
    # --proto_path is set to ACTUAL_PROTO_DIR, where .proto files are located.
    # The input file to protoc is the path relative to ACTUAL_PROTO_DIR.
    # M flag maps the .proto file to the Go import path.
    protoc \
        --proto_path="${ACTUAL_PROTO_DIR}" \
        --go_out="${GO_OUT_DIR}" \
        --go_opt=paths=source_relative \
        --go-grpc_out="${GO_OUT_DIR}" \
        --go-grpc_opt=paths=source_relative \
        --go_opt=M${proto_file_relative_path}=${MODULE_PREFIX}/${go_out_subdir} \
        --go-grpc_opt=M${proto_file_relative_path}=${MODULE_PREFIX}/${go_out_subdir} \
        "${proto_file_relative_path}" # Pass the relative path to protoc

    # Handle cases where protoc might generate files in an unexpected subdirectory
    # This might happen if go_package option in .proto files causes confusion with paths=source_relative.
    # It checks if protoc created an extra directory named after the basename of ACTUAL_PROTO_DIR.
    # e.g. if ACTUAL_PROTO_DIR is "temp/proto", it checks for "go/proto/todofy"
    actual_proto_dir_basename=$(basename "${ACTUAL_PROTO_DIR}")
    incorrect_generated_path_segment=""
    if [ "${proto_file_rel_dir}" != "." ]; then
        incorrect_generated_path_segment="${GO_OUT_DIR}/${actual_proto_dir_basename}/${proto_file_rel_dir}"
    else
        incorrect_generated_path_segment="${GO_OUT_DIR}/${actual_proto_dir_basename}"
    fi

    if [ -d "${incorrect_generated_path_segment}" ]; then
        echo "Moving files from potentially incorrect path ${incorrect_generated_path_segment} to ${go_out_subdir}"
        # Ensure target directory exists (it should, but good practice)
        mkdir -p "${go_out_subdir}"
        # Move all contents from the incorrect path to the correct subdirectory
        mv ${incorrect_generated_path_segment}/* "${go_out_subdir}/" 2>/dev/null || true
        # Clean up the incorrect directory structure (if empty)
        rmdir -p "$(dirname "${incorrect_generated_path_segment}")/${actual_proto_dir_basename}" 2>/dev/null || true
    fi

    # Create go.mod file for each subdirectory if it doesn't exist
    if [ ! -f "${go_out_subdir}/go.mod" ]; then
        module_path="${MODULE_PREFIX}/${go_out_subdir}"
        # Clean up potential trailing slashes or "/./" from module_path
        module_path=$(echo "${module_path}" | sed 's#/\./#/#g; s#/$##')

        echo "Creating Go module in ${go_out_subdir} with module path ${module_path}..."
        pushd "${go_out_subdir}" > /dev/null
        go mod init "${module_path}"
        go mod tidy # Add dependencies based on imports in generated .go files
        popd > /dev/null
    else
        echo "Go module already exists in ${go_out_subdir}. Running go mod tidy..."
        pushd "${go_out_subdir}" > /dev/null
        go mod tidy
        popd > /dev/null
    fi
done

echo "Go code generation completed successfully!"
