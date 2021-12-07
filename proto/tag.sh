#!/usr/bin/env bash
set -eo pipefail

CUR=$(dirname $0)
TAG_VERSION=$1

usage() {
  echo "usage: ${0} v1.0.1"
}

check_flag_value_set() {
  if [ -z "${1}" ]; then
    usage
    exit 1
  fi
}

check_flag_value_set ${TAG_VERSION}

# gen/go 版本
git tag gen/go/${TAG_VERSION}
