# Proto 协议定义

## 1 开发须知
```sh
make setup
```

> 参考 `buf-tools-setup.sh` 文件

#### 定期更新 Protobuf 相关工具

原则上每次开发新的 feture 提交 PR 前应当更新相关工具到最新版本

```sh
make update-tools
```

> 参考 `buf-tools-update.sh` 文件

### 编码规范

- Proto 编码规范遵循 [Uber Protobuf Style Guide V2](https://github.com/uber/prototool/blob/dev/style/README.md)
- Lint 将对编码规范进行审查

## 2 常见开发任务

### 常用脚本

```sh
# Build - 以下部分的全家桶，将依次执行下列操作
# 通常在提交PR前或发布版本时使用
make
# or
make build

# Lint 全部 proto 文件
make lint

# 检查 Break Changes
make break-check

# 清理生成的过期内容
make clean

# 生成代码
make gen-codes

# 生成文档
make gen-doc

# 生成 Proto Descriptor Set
make build-images
```

### 新建一个 proto 文件

需要添加新的 proto 文件时，使用 `buf-create-file.sh` 辅助脚本:

```sh
# 创建一个 idl/examples/foo/v1/bar.proto 文件
./buf-create-file.sh idl/examples/foo/v1/bar.proto
```

### 改名/移动一个 proto 文件

```sh
./buf-mv-file.sh \
	idl/examples/foo/v1/bar.proto \
	idl/examples/foo/v1/bar_new.proto
```

## 2 IDL 文件管理规范

### 概述

`idl`是所有 proto 定义文件的根目录，其下的子目录和文件将会参与到代码生成与文档生成过程。

本项目同时使用了 [**buf**](https://github.com/bufbuild/buf) 与 [**prototool**](https://github.com/uber/prototool) 两个 proto 管理工具进行 proto 文件管理

`buf.yaml` 配置文件中定义了根路径位置。

```yaml
build:
  roots:
    # 本项目 proto 根目录
    - idl
    # 引用 grpc-gateway 中的 googleapis
    - 3rd_party/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
```

`prototool.yaml`配置文件中也做了相关定义。需要特别注意的是，在`idl`目录中添加新的子目录时，必须修改本配置文件中对应位置：

```
