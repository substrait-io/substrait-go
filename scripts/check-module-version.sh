#!/usr/bin/env bash
# Validates that the go.mod module path major version matches the latest git tag.
# This prevents releasing a new major version without updating the module path.

set -euo pipefail

# Extract major version from go.mod module path (e.g., "v7" from ".../substrait-go/v7")
MODULE_PATH=$(go mod edit -json | jq -r '.Module.Path')
MODULE_MAJOR=$(echo "$MODULE_PATH" | grep -oE '/v[0-9]+$' | tr -d '/v' || echo "1")

# Get latest tag's major version
LATEST_TAG=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
TAG_MAJOR=$(echo "$LATEST_TAG" | sed 's/^v//' | cut -d. -f1)

if [ "$MODULE_MAJOR" != "$TAG_MAJOR" ]; then
    echo "ERROR: go.mod module path (v$MODULE_MAJOR) does not match latest git tag (v$TAG_MAJOR)"
    echo ""
    echo "Module path: $MODULE_PATH"
    echo "Latest tag:  $LATEST_TAG"
    echo ""
    echo "To fix: update go.mod and all imports to use v$TAG_MAJOR"
    echo "See: https://go.dev/doc/modules/major-version"
    exit 1
fi

echo "OK: module version (v$MODULE_MAJOR) matches latest tag (v$TAG_MAJOR)"
