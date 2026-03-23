#!/usr/bin/env bash
# Validates that the go.mod module path major version matches the expected version.
#
# Usage:
#   ./check-module-version.sh           # Compares against latest git tag
#   RELEASE_VERSION=8.0.0 ./check-module-version.sh  # Compares against provided version (for pre-release hooks)

set -euo pipefail

# Extract major version from go.mod module path (e.g., "v7" from ".../substrait-go/v7")
MODULE_PATH=$(go mod edit -json | jq -r '.Module.Path')
MODULE_MAJOR=$(echo "$MODULE_PATH" | grep -oE '/v[0-9]+$' | tr -d '/v' || echo "1")

# Get expected major version - either from RELEASE_VERSION env var or latest git tag
if [ -n "${RELEASE_VERSION:-}" ]; then
    EXPECTED_VERSION="$RELEASE_VERSION"
    SOURCE="RELEASE_VERSION"
else
    EXPECTED_VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
    SOURCE="latest git tag"
fi

EXPECTED_MAJOR=$(echo "$EXPECTED_VERSION" | sed 's/^v//' | cut -d. -f1)

if [ "$MODULE_MAJOR" != "$EXPECTED_MAJOR" ]; then
    echo "ERROR: go.mod module path (v$MODULE_MAJOR) does not match $SOURCE (v$EXPECTED_MAJOR)"
    echo ""
    echo "Module path: $MODULE_PATH"
    echo "$SOURCE:  $EXPECTED_VERSION"
    echo ""
    echo "To fix: update go.mod and all imports to use v$EXPECTED_MAJOR"
    echo "See: https://go.dev/doc/modules/major-version"
    exit 1
fi

echo "OK: module version (v$MODULE_MAJOR) matches $SOURCE (v$EXPECTED_MAJOR)"
