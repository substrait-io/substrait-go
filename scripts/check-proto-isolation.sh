#!/usr/bin/env bash
# SPDX-License-Identifier: Apache-2.0
#
# Guards issue #280: substrait-protobuf must stay isolated to the wire/ package.
#
# Core domain packages (types, expr, plan, extensions, ...) must not import the
# generated protobuf types, directly or transitively, so a consumer that only
# builds and inspects plans does not drag substrait-protobuf into its dependency
# tree. All encode/decode lives in wire/ (and later, the codec module).
#
# This is the package-level form of the issue's "Done when" check
# (`go mod graph | grep substrait-protobuf`). That module-level check only comes
# back empty once wire/ moves into its own module; until then this guards the
# in-core isolation.

set -euo pipefail

proto='github.com/substrait-io/substrait-protobuf'
non_wire=$(go list ./... | grep -v '/wire')
fail=0

# 1. No non-wire package may reach substrait-protobuf in its import closure.
if go list -deps $non_wire 2>/dev/null | grep -q "$proto"; then
    echo "substrait-protobuf reachable from a non-wire package:"
    for pkg in $non_wire; do
        if go list -deps "$pkg" 2>/dev/null | grep -q "$proto"; then
            echo "  $pkg"
        fi
    done
    fail=1
fi

# 2. No source outside wire/ (or the codec module) may reference it, tests included.
if leaks=$(git grep -l "$proto" -- '*.go' ':(exclude)wire/' ':(exclude)codec/'); then
    echo "substrait-protobuf referenced in source outside wire/:"
    echo "$leaks" | sed 's/^/  /'
    fail=1
fi

if [ "$fail" -ne 0 ]; then
    echo
    echo "substrait-protobuf must live only in wire/ (issue #280)."
    exit 1
fi

echo "OK: substrait-protobuf is isolated to wire/"
