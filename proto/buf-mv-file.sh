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

SRC_FILE=${1}
DST_FILE=${2}

check_flag_value_set "${SRC_FILE}"
if [ ! -f "${SRC_FILE}" ]; then
  # ${SRC_FILE} does not exist
  echo "${SRC_FILE} doesn't exist."
  exit 1
fi

check_flag_value_set "${DST_FILE}"

# 新建目标文件
./buf-create-file.sh "${DST_FILE}"

# 提取 SRC 文件中的内容
LINE=$(awk '/php_namespace/{ print NR; exit }' ${SRC_FILE})
check_flag_value_set "${LINE}"

sed -n "$((LINE + 1))"',$p' ${SRC_FILE} >>${DST_FILE}
echo "New file: ${DST_FILE}"

BAK_FILE="${SRC_FILE}.bak"
mv "${SRC_FILE}" "${BAK_FILE}"
echo "Backup file: ${BAK_FILE}"
