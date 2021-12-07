module github.com/amannthul/golang-example

go 1.16

replace (
	// FIXME: 由于 etcd 与 gRPC 的兼容问题，需要使用定制的 etcd 版本
	// https://github.com/etcd-io/etcd/issues/11721
	github.com/coreos/etcd => github.com/skyjia/etcd v3.3.22-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.29.1
)

require (
	github.com/docker/distribution v2.7.1+incompatible
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.12.1 // indirect
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/logger/logrus/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/service/v2 v2.9.1
	github.com/prometheus/client_golang v1.3.0 // indirect
	github.com/rs/xid v1.3.0
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
	go.uber.org/multierr v1.5.0 // indirect
	golang.org/x/crypto v0.0.0-20200728195943-123391ffb6de // indirect
	google.golang.org/genproto v0.0.0-20191230161307-f3c370f40bfb // indirect
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	gorm.io/driver/mysql v1.2.1
	gorm.io/gorm v1.22.4
)
