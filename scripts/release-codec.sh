#!/usr/bin/env bash
# SPDX-License-Identifier: Apache-2.0
#
# Create codec/vX.Y.Z for an already-published core release vX.Y.Z.
#
# This script:
# - Checks that both module majors match the requested version.
# - Checks out the core release tag on detached HEAD.
# - Updates codec's core dependency and removes its local replace directive.
# - Tidies dependencies, builds, and tests codec.
# - Commits the dependency changes and tags the new commit.
#
# The release commit is not added to main. The caller must push the codec tag.
# HEAD remains detached. If the codec tag already exists, no changes are made.
#
# Usage: ./scripts/release-codec.sh v9.1.0

set -euo pipefail

version="${1:-}"
[[ "$version" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]] || { echo "usage: $0 vX.Y.Z"; exit 1; }

cd "$(git rev-parse --show-toplevel)"
core=$(go mod edit -json | jq -r .Module.Path)
# core and codec must already be on this release's major; reuse the release gate.
RELEASE_VERSION="$version" bash scripts/check-module-version.sh

git rev-parse -q --verify "refs/tags/$version" >/dev/null || { echo "core tag $version missing; release core first"; exit 1; }
git rev-parse -q --verify "refs/tags/codec/$version" >/dev/null && { echo "codec/$version exists; nothing to do"; exit 0; }

git checkout -q --detach "$version"

# Pin core to the tag and drop the dev replace, then prove codec builds the way a
# consumer resolves it. GOPROXY=direct because the core tag may be seconds old.
(cd codec
    go mod edit -require="$core@$version" -dropreplace="$core"
    GOPROXY=direct GOFLAGS=-mod=mod go mod tidy
    go build ./... && go test ./...)

git commit -q -m "chore(codec): release $version" codec/go.mod codec/go.sum
git tag -a "codec/$version" -m "codec $version"
echo "OK: tagged codec/$version"
