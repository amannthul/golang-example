package main

import (
	"github.com/amannthul/golang-example/pkg/dbutils"
	mlog "github.com/amannthul/golang-example/pkg/log"
	"github.com/amannthul/golang-example/pkg/rpc/service"
	reportpb "github.com/amannthul/golang-example/proto/gen/go/example/report/v1"
	store "github.com/amannthul/golang-example/store/report"
	reportsvc "github.com/amannthul/golang-example/svc/report"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/server"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	// ServiceName 是本微服务的名称
	ServiceName = "example"
	// ServiceNamespace 是微服务的命名空间
	ServiceNamespace = "com.amannthul.service"
)

var (
	log = mlog.StandardLogger()

	// Following values will be set during build.
	// Do NOT manually modify them.

	// ProductVersion is current product version.
	ProductVersion = "(n/a)"
	// GitCommit is the git commit short hash
	GitCommit = "(n/a)"
	// GoVersion is go compiler version `go version`
	GoVersion = "(n/a)"
	// BuildTime is go build time
	BuildTime = "(n/a)"
)

func main() {
	// ServiceOptions
	opts := service.NewServiceOptions(ServiceNamespace, ServiceName)
	opts.ProductVersion = ProductVersion
	opts.GitCommit = GitCommit
	opts.GoVersion = GoVersion
	opts.BuildTime = BuildTime
	// 加入jwt的检测
	// authWrapper = wrapper.NewAuthWrapper(nil)
	// permissionWrapper = wrapper.NewPermissionWrapper(nil)
	// opts.PostServerHandlerWrappers = []server.HandlerWrapper{authWrapper.Auth, permissionWrapper.Auth}
	svc := service.CreateService(opts)
	err := service.RegisterServer(svc.Server(), serviceRegister(svc))
	die(err)

	// Run the service
	err = svc.Run()
	die(err)
}

func serviceRegister(service micro.Service) service.RegisterServerFunc {
	return func(srv server.Server) error {

		dbConfigKey := []string{
			"micro",
			"config",
			"amannthul",
			"com.platform.report.service.@global",
			"mysql"}

		// 获取数据库的配置信息
		dbConn := &dbutils.DBConnection{}
		err := config.Get(dbConfigKey...).Scan(dbConn)
		if err != nil {
			die(err)
		}
		// 连接数据库
		db, err := gorm.Open(mysql.Open(dbutils.GetDsn(dbConn)), &gorm.Config{})
		if err != nil {
			die(err)
		}
		// Register Report Store
		ucReport := store.NewReportStore(dbutils.NewConnection(db))

		reportAPI := reportsvc.NewReportAPIHandler(ucReport)

		if err := reportpb.RegisterReportAPIHandler(srv, reportAPI); err != nil {
			return err
		}
		log.Infof("Registered RPC service: %s", reportAPI.Name())

		return nil
	}
}

func die(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
