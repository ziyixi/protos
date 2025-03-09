#!/bin/bash
set -e

# Directory structure
PROTO_DIR="proto"
GO_OUT_DIR="go"

# raise error if MODULE_PREFIX is not set
if [ -z "${MODULE_PREFIX}" ]; then
    echo "Error: MODULE_PREFIX is not set. Please set it to your GitHub repository path."
    exit 1
fi
# Check if protoc is installed
if ! command -v protoc &> /dev/null; then
    echo "Error: protoc is not installed. Please install it to generate Go code."
    exit 1
fi
# Check if the Go plugin for protoc is installed
if ! command -v protoc-gen-go &> /dev/null; then
    echo "Error: protoc-gen-go is not installed. Please install it to generate Go code."
    exit 1
fi
# Check if the Go gRPC plugin for protoc is installed
if ! command -v protoc-gen-go-grpc &> /dev/null; then
    echo "Error: protoc-gen-go-grpc is not installed. Please install it to generate Go code."
    exit 1
fi

# Create the output directory if it doesn't exist
mkdir -p "${GO_OUT_DIR}"

# Find all proto files recursively
find "${PROTO_DIR}" -name "*.proto" | while read proto_file; do
    # Get the relative directory path (without the proto/ prefix)
    rel_dir=$(dirname "${proto_file}" | sed "s|^${PROTO_DIR}/||")
    
    # Create the output directory structure
    go_out_subdir="${GO_OUT_DIR}/${rel_dir}"
    mkdir -p "${go_out_subdir}"
    
    # Get the package name from the proto file
    package_name=$(grep "^package" "${proto_file}" | sed 's/package\s\+\([^;]\+\);/\1/')
    
    echo "Generating Go code for ${proto_file}..."
    
    # Generate go files using protoc with grpc plugin
    # Important: we need to map from proto/ to go/ correctly
    protoc \
        --proto_path=. \
        --go_out="${GO_OUT_DIR}" \
        --go_opt=paths=source_relative \
        --go-grpc_out="${GO_OUT_DIR}" \
        --go-grpc_opt=paths=source_relative \
        --go_opt=M${proto_file}=${MODULE_PREFIX}/${GO_OUT_DIR}/${rel_dir} \
        --go-grpc_opt=M${proto_file}=${MODULE_PREFIX}/${GO_OUT_DIR}/${rel_dir} \
        "${proto_file}"
    
    # Handle the case where protoc might generate files in proto/todofy instead of go/todofy
    # Check if files were generated in the wrong location
    incorrect_path="${GO_OUT_DIR}/${PROTO_DIR}/${rel_dir}"
    if [ -d "${incorrect_path}" ]; then
        echo "Moving files from incorrect path ${incorrect_path} to ${go_out_subdir}"
        # Move the generated files to the correct location
        mv ${incorrect_path}/* ${go_out_subdir}/ 2>/dev/null || true
        # Clean up empty directories
        rmdir -p ${incorrect_path} 2>/dev/null || true
    fi
    
    # Create go.mod file for each subdirectory if it doesn't exist
    if [ ! -f "${go_out_subdir}/go.mod" ]; then
        # Determine the module path based on your repository
        # The MODULE_PREFIX should be set to your GitHub repository path
        # e.g., github.com/username/repository
        module_path="${MODULE_PREFIX}/${GO_OUT_DIR}/${rel_dir}"
        
        echo "Creating Go module in ${go_out_subdir}..."
        pushd "${go_out_subdir}" > /dev/null
        go mod init "${module_path}"
        
        # Add go.mod dependencies based on imports in generated .go files
        # This is a basic implementation and might need adjustment based on your imports
        go mod tidy
        popd > /dev/null
    fi
done

echo "Go code generation completed successfully!"