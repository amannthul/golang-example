#!/usr/bin/env bash
set -eo pipefail

CUR=$(dirname $0)

checkVersion() {
    local name=$1
    local actual=$2
    local expected=$3
    if [[ "$actual" != "$expected" ]]; then
        echo "Failed: $1 should be $actual, but got $expected"
    else
        echo "Passed: $1 $actual"
    fi
}

# golang
VER=$(go version | awk '{print $3}')
checkVersion "go" "${VER}" "go1.14.2"

# protoc
VER=$(protoc --version | awk '{print $2}')
checkVersion "protoc" "${VER}" "3.11.4"

# buf
VER=$(buf --version 2>&1)
checkVersion "buf" "${VER}" "0.11.0"
