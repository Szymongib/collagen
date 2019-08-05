#!/usr/bin/env bash
set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

echo "Downloading dependencies..."
go mod download
echo "Dependencies downloaded"

echo "Testing..."
go test ${DIR}/internal/... ${DIR}/cmd/...
echo "Testing done"

echo "Installing..."
go install ${DIR}
echo "Installation finished"

echo "Generating code in examples..."
go generate ${DIR}/examples/doggo/...
echo "Examples code generated"

echo "Testing examples..."
go test ${DIR}/examples/...
echo "Examples tested"

echo "Done!"
