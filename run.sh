#!/bin/sh

set -e

temp_file=$(mktemp)

go build -o "${temp_file}" main.go

exec "${temp_file}" "$@"
