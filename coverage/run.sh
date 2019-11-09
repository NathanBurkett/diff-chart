#!/usr/bin/env bash
# Generate test coverage statistics for Go packages.
#
# Works around the fact that `go test -coverprofile` currently does not work
# with multiple packages, see https://code.google.com/p/go/issues/detail?id=6909
#
# Usage: coverage/run [--html]
#
#     --html      Additionally create HTML report and open it in browser
#

set -e

WORKDIR=coverage/output
TARGET="$WORKDIR/cover.out"
MODE=atomic

generate_cover_data() {
    rm -rf "$WORKDIR/*.out"
    rm -rf "$WORKDIR/coverage.txt"

    for PKG in "$@"; do
        f="$WORKDIR/$(echo $PKG | tr / -).cover"
        go test -count=1 -v -covermode="$MODE" -coverprofile="$f" "$PKG"
    done

    echo "mode: $MODE" >"$TARGET"
    grep -h -v "^mode:" "$WORKDIR"/*.cover >>"$TARGET"
    cat "$TARGET" >> "$WORKDIR/coverage.txt"
}

show_cover_report() {
    FLAG="$1"
    go tool cover -"$FLAG"="$TARGET"
}

case "$1" in
"")
	generate_cover_data $(go list ./...)
	show_cover_report func
    ;;
--html)
	generate_cover_data $(go list ./...)
    show_cover_report html ;;
*)
    echo >&2 "error: invalid option: $1"; exit 1 ;;
esac

