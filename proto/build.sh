#!/usr/bin/env bash
set -eo pipefail

CUR=$(dirname $0)
GEN_OUT=${CUR}/gen

. ${CUR}/util.sh

info "[1/8] Linting proto files..."
${CUR}/buf-check-lint.sh

info "[2/8] Checking breaking changes..."
${CUR}/buf-check-breaking.sh

info "[3/8] Cleaning outdated outputs..."
${CUR}/buf-clean.sh

info "[4/8] Generating codes..."
${CUR}/buf-gen-codes.sh 1>/dev/null

info "[5/8] Generating documents..."
# 为不同 package 单独生成 proto 文档
# FIXME: 目前生成文档有问题，只能生成最后一个 proto 文档
# ${CUR}/buf-gen-doc.sh examples
${CUR}/buf-gen-doc.sh example/report
# 全部 package 文档 all-in-one
# ${CUR}/buf-gen-doc.sh
# 生成 ts 文档
# 为不同 package 单独生成 ts 文档
${CUR}/buf-gen-ts-doc.sh example/report 1>/dev/null

info "[6/8] Format codes..."
make format

info "[7/8] Generating buf images (FileDescriptorSet)..."
${CUR}/buf-build-image.sh

info "[8/8] Build is done. 🍻🍻🍻"
