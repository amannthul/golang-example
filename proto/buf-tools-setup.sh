#!/usr/bin/env bash
set -eo pipefail
set -x

CUR=$(dirname $0)

cd ${CUR}

# GO Proxy
export GOPROXY=https://goproxy.io,direct
export GOPRIVATE=github.com/amannthul/*

function get_latest_release() {
    local repo=$1
    curl --silent "https://api.github.com/repos/${repo}/tags" | jq -r '.[0].name'
}

function download_grpc_web() {
    local version=$(get_latest_release "grpc/grpc-web")
    if [ -z "${version}" ]; then
        # 如果 GitHub API 访问出问题则使用 package.json 里面定义的版本号
        version=$(jq -r '.dependencies."grpc-web" | ltrimstr("^")' package.json)
    fi
    local sys_name=$(uname -s | awk '{print tolower($0)}')
    local hw_name=$(uname -m)
    local url="https://github.com/grpc/grpc-web/releases/download/${version}/protoc-gen-grpc-web-${version}-${sys_name}-${hw_name}"

    local out=bin/protoc-gen-grpc-web
    curl -L ${url} -o ${out}
    chmod +x ${out}
}

brew upgrade
brew install protobuf
brew install prototool
# echo "Downloading golangci-lint"
# curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s
brew install golangci/tap/golangci-lint
# Buf - https://buf.build/docs/installation#via-homebrew
brew tap bufbuild/buf
brew install buf
brew install direnv
brew install clang-format

pushd ${CUR}/tools
# go install github.com/golang/protobuf/{proto,protoc-gen-go}
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# FIXME: etcd 与 grpc 1.30+ 版本不兼容，此处grpc降级到 1.29, protoc-gen-go-grpc 工具降级到 v1.0.0
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.0.0
go install github.com/micro/micro/v2/cmd/protoc-gen-micro@latest
go install github.com/client9/misspell/cmd/misspell@latest
go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest
go install \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest
go install istio.io/tools/cmd/protoc-gen-docs@latest@latest
# go install github.com/fullstorydev/grpcui/cmd/grpcui
# GO111MODULE=on go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
# GO111MODULE=on go install google.golang.org/grpc
# GO111MODULE=on go install github.com/gogo/protobuf/proto
# GO111MODULE=on go install github.com/gogo/protobuf/jsonpb
# GO111MODULE=on go install github.com/gogo/protobuf/gogoproto
# GO111MODULE=on go install github.com/gogo/protobuf/protoc-gen-gogo
# GO111MODULE=on go install github.com/gogo/protobuf/protoc-gen-gofast
# GO111MODULE=on go install github.com/gogo/protobuf/protoc-gen-gogofast
# GO111MODULE=on go install github.com/gogo/protobuf/protoc-gen-gogofaster
# GO111MODULE=on go install github.com/gogo/protobuf/protoc-gen-gogoslick
go mod tidy
popd

# echo "Installing protoc-gen-doc"
# docker pull pseudomuto/protoc-gen-doc

echo "Installing npm packages"
yarn
ncu -u
npm install -g typescript
# npm i
yarn

echo "Downloading protoc-gen-grpc-web"
download_grpc_web
