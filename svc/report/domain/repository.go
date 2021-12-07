package domain

import (
	"github.com/amannthul/golang-example/pkg/dbutils"
)

// Domain Repository
type ReportRepository interface {
	dbutils.Tx
}
