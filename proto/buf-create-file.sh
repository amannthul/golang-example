#!/usr/bin/env bash
set -eo pipefail

CUR=$(dirname ${0})

usage() {
  echo "usage: ${0} idl/examples/foo/v1/bar.proto"
}

check_flag_value_set() {
  if [ -z "${1}" ]; then
    usage
    exit 1
  fi
}

PROTO_FILE="${1}"

check_flag_value_set "${PROTO_FILE}"
# ${GO_MODULE_NAME} is defined in .envrc
check_flag_value_set "${GO_MODULE_NAME}"

PROTO_BASE_DIR="idl"
PROTO_DIR=$(dirname ${PROTO_FILE})

if [ ! -d "${PROTO_DIR}" ]; then
  # ${PROTO_DIR} does not exist
  mkdir -p "${PROTO_DIR}"
fi

prototool create ${PROTO_FILE}

# 修正 prototool 生成的 go_package 格式
GO_PKG_NAME="${GO_MODULE_NAME}"/$(realpath --relative-to="${PROTO_BASE_DIR}" "${PROTO_DIR}")
sed -i '' -e 's/option go_package = "/option go_package = "'${GO_PKG_NAME//\//\\\/}';/g' "${PROTO_FILE}"
