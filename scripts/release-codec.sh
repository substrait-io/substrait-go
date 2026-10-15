#!/usr/bin/env bash
# SPDX-License-Identifier: Apache-2.0
#
# Prepare and tag a codec release matching an already-published core release,
# in lockstep with core (#280). For the core vX.Y.Z just released, this pins
# codec to that exact core version, proves codec builds the way a consumer
# resolves it, and tags codec/vX.Y.Z. The work happens on a detached HEAD so
# main keeps its dev-time replace directive, and the caller pushes the tag.
# Re-running for an existing codec tag is a no-op.
#
# Usage: ./scripts/release-codec.sh v9.1.0

set -euo pipefail

version="${1:-}"
[[ "$version" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]] || { echo "usage: $0 vX.Y.Z"; exit 1; }

cd "$(git rev-parse --show-toplevel)"
core=$(go mod edit -json | jq -r .Module.Path)
# core and codec must already be on this release's major; reuse the release gate.
# GITHUB_ENV is cleared so the gate's version handoff only fires in the release hook.
RELEASE_VERSION="$version" GITHUB_ENV= bash scripts/check-module-version.sh

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
