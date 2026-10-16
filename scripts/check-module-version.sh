#!/usr/bin/env bash
# Validate both module majors and codec's core requirement in a release PR.
# RELEASE_VERSION is the version proposed in .release-please-manifest.json.
#
# Usage:
#   RELEASE_VERSION=9.0.0 ./scripts/check-module-version.sh

set -euo pipefail

if [ -z "${RELEASE_VERSION:-}" ]; then
    echo "ERROR: RELEASE_VERSION not set"
    echo "Set RELEASE_VERSION to the version proposed by the release PR"
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

CORE_MODULE=$(go mod edit -json | jq -r '.Module.Path')
CORE_VERSION=$( (cd codec && go mod edit -json) | jq -r --arg core "$CORE_MODULE" \
    '.Require[] | select(.Path == $core) | .Version')
EXPECTED_CORE_VERSION="v${RELEASE_VERSION#v}"

if [ "$CORE_VERSION" != "$EXPECTED_CORE_VERSION" ]; then
    echo "ERROR: codec must require $CORE_MODULE $EXPECTED_CORE_VERSION, not $CORE_VERSION"
    exit 1
fi

echo "OK: codec requires core $EXPECTED_CORE_VERSION"
