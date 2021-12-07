package report

import (
	reportpb "github.com/amannthul/golang-example/proto/gen/go/example/report/v1"
	"github.com/amannthul/golang-example/svc/report/domain"
)

type ReportAPIHandler struct {
	store domain.ReportRepository
}

var _ reportpb.ReportAPIHandler = (*ReportAPIHandler)(nil)

func (svc *ReportAPIHandler) Name() string {
	const name = "ReportAPI"
	return name
}

func NewReportAPIHandler(store domain.ReportRepository) *ReportAPIHandler {
	return &ReportAPIHandler{
		store: store,
	}
}
