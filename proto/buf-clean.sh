#!/usr/bin/env bash
set -eo pipefail

CUR=$(dirname $0)

GEN_DIR=${CUR}/gen

. ${CUR}/util.sh

if [ -d "${GEN_DIR}" ]; then
    rm -rf ${GEN_DIR}/go
    rm -rf ${GEN_DIR}/ts-node
    rm -rf ${GEN_DIR}/js-grpc-web
    rm -rf ${GEN_DIR}/java
    rm -rf ${GEN_DIR}/doc
    rm -rf ${GEN_DIR}/ts-doc
    rm -rf ${GEN_DIR}/swift-ios
fi
