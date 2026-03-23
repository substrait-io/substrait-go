#!/usr/bin/env bash
# Validates that the go.mod module path major version matches the release version.
# This script is intended to be run as a preRelease hook in go-semantic-release,
# where RELEASE_VERSION is set to the computed next version (e.g., "8.0.0").
#
# Usage:
#   RELEASE_VERSION=8.0.0 ./check-module-version.sh

set -euo pipefail

if [ -z "${RELEASE_VERSION:-}" ]; then
    echo "RELEASE_VERSION not set, skipping check"
    exit 0
fi

# Extract major version from go.mod module path (e.g., "v8" from ".../substrait-go/v8")
MODULE_PATH=$(go mod edit -json | jq -r '.Module.Path')
MODULE_MAJOR=$(echo "$MODULE_PATH" | grep -oE '/v[0-9]+$' | tr -d '/v' || echo "1")

RELEASE_MAJOR=$(echo "$RELEASE_VERSION" | sed 's/^v//' | cut -d. -f1)

if [ "$MODULE_MAJOR" != "$RELEASE_MAJOR" ]; then
    echo "ERROR: go.mod module path (v$MODULE_MAJOR) does not match RELEASE_VERSION (v$RELEASE_MAJOR)"
    echo ""
    echo "Module path: $MODULE_PATH"
    echo "RELEASE_VERSION: $RELEASE_VERSION"
    echo ""
    echo "To fix: update go.mod and all imports to use v$RELEASE_MAJOR"
    echo "See: https://go.dev/doc/modules/major-version"
    exit 1
fi

echo "OK: module version (v$MODULE_MAJOR) matches RELEASE_VERSION (v$RELEASE_MAJOR)"
