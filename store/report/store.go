package store

import (
	"github.com/amannthul/golang-example/pkg/dbutils"
	"github.com/amannthul/golang-example/svc/report/domain"
)

// ReportStore 实现 Domain 层 ReportStore 接口
type ReportStore struct {
	// 数据库连接
	*dbutils.Connection
}

func NewReportStore(db *dbutils.Connection) *ReportStore {
	return &ReportStore{
		Connection: db,
	}
}

// 实现 Domain 行为
var _ domain.ReportRepository = (*ReportStore)(nil)
