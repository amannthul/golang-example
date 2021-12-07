#!/usr/bin/env bash
set -eo pipefail

CUR=$(dirname $0)
PKG=$1

if [ ! -z "${PKG}" ]; then
  # 指定 Package
  PKG_FILE_NAME=$(echo ${PKG} | sed -e "s/\//-/g")
  OUT_FILE=proto-doc_${PKG_FILE_NAME}.html
else
  # 全量
  OUT_FILE=full-proto-doc.html
fi

OUT_DIR=${CUR}/gen/doc/${PKG}

if [ ! -d "${OUT_DIR}" ]; then
  mkdir -p ${OUT_DIR}
fi

PROTO_PATH=${CUR}/idl/${PKG}
PROTO_INCLUDE_PATH_1="${CUR}/idl"
# PROTO_INCLUDE_PATH_2="${CUR}/3rd_party/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis"

echo "${PROTO_PATH}"

PROTO_FILES=$(buf ls-files | grep "^idl/${PKG}")

# Generate HTML document with pseudomuto/protoc-gen-doc
protoc \
  --doc_out=${OUT_DIR} \
  --doc_opt=html,${OUT_FILE} \
  --proto_path==${PROTO_INCLUDE_PATH_1} \
  ${PROTO_FILES}

# Generate HTML document with istio/tools/cmd/protoc-gen-docs
protoc \
  --docs_out=warnings=true:${OUT_DIR} \
  --proto_path==${PROTO_INCLUDE_PATH_1} \
  ${PROTO_FILES}
