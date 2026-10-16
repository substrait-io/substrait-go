#!/usr/bin/env bash
# Check both module majors before publishing core, so a codec major-version
# mismatch cannot leave only core published (#280).
#
# Run as a preRelease hook in go-semantic-release, where RELEASE_VERSION is the
# computed next version (e.g., "9.0.0"). On success it also publishes
# RELEASE_VERSION to $GITHUB_ENV so later CI steps (the codec release) use the
# computed version rather than rediscovering it from tags.
#
# Usage:
#   RELEASE_VERSION=9.0.0 ./scripts/check-module-version.sh

set -euo pipefail

if [ -z "${RELEASE_VERSION:-}" ]; then
    echo "ERROR: RELEASE_VERSION not set"
    echo "This script should be run as a preRelease hook in go-semantic-release"
    exit 1
fi

RELEASE_MAJOR=$(echo "$RELEASE_VERSION" | sed 's/^v//' | cut -d. -f1)

# Compare a module's go.mod path major (e.g., "v9" from ".../substrait-go/v9")
# against the release major. A v1 module has no suffix, so it maps to major 1.
check_module() {
    local name="$1" dir="$2" path major
    path=$( (cd "$dir" && go mod edit -json) | jq -r '.Module.Path')
    major=$(echo "$path" | grep -oE '/v[0-9]+$' | tr -d '/v' || echo "1")

    if [ "$major" != "$RELEASE_MAJOR" ]; then
        echo "ERROR: $name go.mod module path (v$major) does not match RELEASE_VERSION (v$RELEASE_MAJOR)"
        echo ""
        echo "Module path: $path"
        echo "RELEASE_VERSION: $RELEASE_VERSION"
        echo ""
        echo "To fix: update $name go.mod and all imports to use v$RELEASE_MAJOR"
        echo "See: https://go.dev/doc/modules/major-version"
        exit 1
    fi
    echo "OK: $name module version (v$major) matches RELEASE_VERSION (v$RELEASE_MAJOR)"
}

check_module core .
check_module codec codec

if [ -n "${GITHUB_ENV:-}" ]; then
    echo "RELEASE_VERSION=$RELEASE_VERSION" >> "$GITHUB_ENV"
fi
