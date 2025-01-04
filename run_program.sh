#!/bin/sh
set -e # Exit early if any commands fail
(
  cd "$(dirname "$0")"
  go build -o /tmp/go-lsm-db cmd/*.go
)

exec /tmp/go-lsm-db "$@"

