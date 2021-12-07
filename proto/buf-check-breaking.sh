#!/usr/bin/env bash
set -eo pipefail

CUR=$(dirname $0)

cd ${CUR}

# 基于 image 进行比较
buf breaking --against buf-image.json.gz

# 基于 git 的 master 分支进行比较
# buf breaking --against .git#branch=master

# 基于 git 的 tag 进行比较
# buf breaking --against .git#tag=v1.0.0
