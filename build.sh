#!/bin/sh

set -e

(
  go build -o ./dist/mp-shell cmd/myshell/*.go
)