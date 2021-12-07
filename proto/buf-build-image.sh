#!/usr/bin/env bash
set -eo pipefail

CUR=$(dirname $0)

buf build -o buf-image.json
buf build -o buf-image.bin
buf build -o buf-image.json.gz
