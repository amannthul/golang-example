#!/usr/bin/env bash
set -eo pipefail

CUR=$(dirname $0)

cd ${CUR}
buf lint

# Spell check
misspell -error gen/go/**/*
