#!/usr/bin/env bash
set -eo pipefail

CUR=$(dirname $0)
PKG=$1

SRC=gen/ts-node/${PKG}
OUT=gen/ts-doc/${PKG}
cd ${CUR}
typedoc --out ${OUT} ${SRC}
