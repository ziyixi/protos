#!/bin/bash
set -e

# This script generates Go modules from .proto files, accommodating linter-enforced
# root-relative import paths like "proto/path/to/file.proto".

# PROTO_BASE_DIR_ENV: Path to the directory that will serve as the base for --proto_path.
# In CI, this is typically the root of the checkout of the 'protobuf' branch (e.g., 'temp_protobuf_checkout').
# Locally, this would typically be '.' (the current project root).
PROTOC_BASE_DIR="${PROTO_BASE_DIR_ENV:-.}"

# PROTO_SUBDIR_NAME_ENV: The name of the subdirectory within PROTOC_BASE_DIR that holds the .proto files.
# This is usually "proto".
PROTO_SUBDIR_NAME="${PROTO_SUBDIR_NAME_ENV:-proto}"

# GO_OUT_DIR: The root directory for the final generated Go packages (e.g., "go").
GO_OUT_DIR="go"

# MODULE_PREFIX: Go module prefix for your repository (e.g., "github.com/youruser/yourrepo").
if [ -z "${MODULE_PREFIX}" ]; then
    echo "Error: MODULE_PREFIX is not set. Please set it to your GitHub repository path."
    exit 1
fi

# --- Tool Installation Checks ---
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
# --- End Tool Installation Checks ---

# Create the base output directory for Go files if it doesn't exist
mkdir -p "${GO_OUT_DIR}"

# PROTO_FILES_SEARCH_ROOT: The absolute or relative path to the directory where .proto files are located.
# e.g., "temp_protobuf_checkout/proto" in CI, or "./proto" locally.
PROTO_FILES_SEARCH_ROOT="${PROTOC_BASE_DIR}/${PROTO_SUBDIR_NAME}"

echo "Using --proto_path=${PROTOC_BASE_DIR}"
echo "Searching for .proto files in ${PROTO_FILES_SEARCH_ROOT}..."

# Find all .proto files recursively
find "${PROTO_FILES_SEARCH_ROOT}" -name "*.proto" | while read proto_file_full_path; do
    # proto_file_full_path: Full path to the .proto file.
    # e.g., temp_protobuf_checkout/proto/todofy/database.proto (CI)
    # or ./proto/todofy/database.proto (local, if PROTOC_BASE_DIR=".")

    # protoc_input_path: Path of the .proto file relative to PROTOC_BASE_DIR.
    # This is what's passed to protoc for compilation and used in M-flags.
    # e.g., "proto/todofy/database.proto"
    protoc_input_path="${proto_file_full_path#${PROTOC_BASE_DIR}/}"
    protoc_input_path="${protoc_input_path#./}" # Clean up leading "./" if PROTOC_BASE_DIR was "."

    # path_inside_proto_root: Path of the .proto file relative to PROTO_FILES_SEARCH_ROOT.
    # Used to determine the subdirectory structure within GO_OUT_DIR.
    # e.g., "todofy/database.proto" or "database.proto" (if at the root of PROTO_FILES_SEARCH_ROOT)
    path_inside_proto_root="${proto_file_full_path#${PROTO_FILES_SEARCH_ROOT}/}"
    path_inside_proto_root="${path_inside_proto_root#./}" # Clean up leading "./"

    # go_output_rel_dir: The relative directory structure for the Go package.
    # e.g., "todofy" or "." (if the .proto file is at the root of PROTO_FILES_SEARCH_ROOT)
    go_output_rel_dir=$(dirname "${path_inside_proto_root}")

    # final_go_out_subdir: The final target directory for the generated Go files for this package.
    # e.g., "go/todofy" or "go" (if go_output_rel_dir is ".")
    final_go_out_subdir="${GO_OUT_DIR}"
    if [ "${go_output_rel_dir}" != "." ]; then
        final_go_out_subdir="${GO_OUT_DIR}/${go_output_rel_dir}"
    fi
    mkdir -p "${final_go_out_subdir}" # Ensure it exists

    echo "Processing ${protoc_input_path}..."
    echo "  Final Go output directory for this package: ${final_go_out_subdir}"

    # protoc command execution
    # --proto_path is PROTOC_BASE_DIR (e.g., "." or "temp_protobuf_checkout")
    # --go_out is GO_OUT_DIR (e.g., "go")
    # paths=source_relative will cause output to mirror protoc_input_path structure under GO_OUT_DIR.
    # e.g., if protoc_input_path is "proto/todofy/file.proto", output is in "go/proto/todofy/"
    protoc \
        --proto_path="${PROTOC_BASE_DIR}" \
        --go_out="${GO_OUT_DIR}" \
        --go_opt=paths=source_relative \
        --go-grpc_out="${GO_OUT_DIR}" \
        --go-grpc_opt=paths=source_relative \
        --go_opt=M${protoc_input_path}=${MODULE_PREFIX}/${final_go_out_subdir} \
        --go-grpc_opt=M${protoc_input_path}=${MODULE_PREFIX}/${final_go_out_subdir} \
        "${protoc_input_path}"

    # Determine where protoc actually generated files with paths=source_relative.
    # This will be GO_OUT_DIR / (directory part of protoc_input_path).
    # e.g., "go/proto/todofy" if protoc_input_path was "proto/todofy/file.proto"
    # e.g., "go/proto" if protoc_input_path was "proto/file.proto"
    protoc_generated_files_dir="${GO_OUT_DIR}/$(dirname "${protoc_input_path}")"

    # If protoc generated files into an intermediate path (e.g., "go/proto/todofy")
    # and this is different from the final desired path (e.g., "go/todofy"), move them.
    if [ -d "${protoc_generated_files_dir}" ] && [ "${protoc_generated_files_dir}" != "${final_go_out_subdir}" ]; then
        echo "  Moving generated files from ${protoc_generated_files_dir} to ${final_go_out_subdir}..."
        # Move all files from the source directory to the target.
        # Using find to handle cases where there might be multiple generated files (*.pb.go, *.pb.gw.go etc.)
        # -maxdepth 1 ensures we only move files directly within protoc_generated_files_dir, not from its subdirs (if any).
        find "${protoc_generated_files_dir}" -maxdepth 1 -type f -exec mv -t "${final_go_out_subdir}/" {} +
        
        # Attempt to clean up the (now hopefully empty) directory structure protoc created.
        # e.g., remove "go/proto/todofy", then try to remove "go/proto" if it's empty.
        if [ -d "${protoc_generated_files_dir}" ]; then # Check if it still exists
             rmdir "${protoc_generated_files_dir}" 2>/dev/null || echo "  Note: Directory ${protoc_generated_files_dir} not empty after move, or already removed."
        fi
        
        # Attempt to remove the parent of protoc_generated_files_dir if it's not GO_OUT_DIR itself
        # and if it's part of the "proto/" prefix path.
        # e.g., if protoc_generated_files_dir was "go/proto/todofy", its parent is "go/proto".
        protoc_intermediate_parent_dir=$(dirname "${protoc_generated_files_dir}")
        if [ "${protoc_intermediate_parent_dir}" != "${GO_OUT_DIR}" ] && [ -d "${protoc_intermediate_parent_dir}" ] && [[ "${protoc_intermediate_parent_dir}" == "${GO_OUT_DIR}/${PROTO_SUBDIR_NAME}"* ]]; then
             rmdir "${protoc_intermediate_parent_dir}" 2>/dev/null || echo "  Note: Directory ${protoc_intermediate_parent_dir} not empty or already removed."
        fi
    elif [ "${protoc_generated_files_dir}" == "${final_go_out_subdir}" ]; then
        echo "  Generated files are already in the final target directory: ${final_go_out_subdir}"
    else
        echo "  Warning: Protoc generated files directory ${protoc_generated_files_dir} not found. Skipping move."
    fi

    # Create or update go.mod in the final_go_out_subdir for the package.
    module_path_for_gomod="${MODULE_PREFIX}/${final_go_out_subdir}"
    # Clean the module path: remove "/./", "//", and trailing "/".
    module_path_for_gomod=$(echo "${module_path_for_gomod}" | sed -e 's#/\./#/#g' -e 's#//#/#g' -e 's#/$##')

    if [ ! -f "${final_go_out_subdir}/go.mod" ]; then
        echo "  Creating Go module in ${final_go_out_subdir} (module ${module_path_for_gomod})..."
        pushd "${final_go_out_subdir}" > /dev/null
        go mod init "${module_path_for_gomod}"
        go mod tidy # To fetch dependencies like google/protobuf if needed by generated code.
        popd > /dev/null
    else
        echo "  Go module already exists in ${final_go_out_subdir}. Running go mod tidy..."
        pushd "${final_go_out_subdir}" > /dev/null
        go mod tidy
        popd > /dev/null
    fi
done

echo "Go code generation completed successfully!"
