#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -o errexit
set -o nounset
set -o pipefail

# Set the CGO flags to use the portable version of BLST
#
# We use "export" here instead of just setting a bash variable because we need
# to pass this flag to all child processes spawned by the shell.
export CGO_CFLAGS="-O -D__BLST_PORTABLE__"

# Root directory
programsvm_PATH=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    cd .. && pwd
)

if [[ $# -eq 1 ]]; then
    BINARY_DIR=$(cd "$(dirname "$1")" && pwd)
    BINARY_FNAME=$(basename "$1")
    BINARY_PATH=$BINARY_DIR/$BINARY_FNAME
elif [[ $# -eq 0 ]]; then
    # Set default binary directory location
    name="pkEmJQuTUic3dxzg8EYnktwn4W7uCHofNcwiYo458vodAUbY7"
    BINARY_PATH=$programsvm_PATH/build/$name
else
    echo "Invalid arguments to build programsvm. Requires zero (default location) or one argument to specify binary location."
    exit 1
fi

cd "$programsvm_PATH"

echo "Building programsvm in $BINARY_PATH"
mkdir -p "$(dirname "$BINARY_PATH")"
go build -o "$BINARY_PATH" ./cmd/programsvm

CLI_PATH=$programsvm_PATH/build/programs-cli
echo "Building programs-cli in $CLI_PATH"
mkdir -p "$(dirname "$CLI_PATH")"
go build -o "$CLI_PATH" ./cmd/programs-cli
