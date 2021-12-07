package store

import (
	"time"

	"github.com/amannthul/golang-example/svc/report/domain"
	"gorm.io/gorm"
)

type Report struct {
	ReportID    string          `gorm:"primary_key;column:report_id"` // 脉搏波算法计算的 ID
	DeveloperID string          `gorm:"column:developer_id"`          // 测量开发者ID
	AppID       string          `gorm:"column:app_id"`                // 受试者使用的appid
	SubjectId   string          `gorm:"column:subject_id"`            // 用户 ID
	Rev         int32           `gorm:"column:rev"`
	CreatedAt   time.Time       // 创建时间
	UpdatedAt   time.Time       // 更新时间
	DeletedAt   *gorm.DeletedAt // 删除时间
}

func (r Report) TableName() string {
	return "report"
}

// domain->db
func (r *Report) FromDomainReport(d domain.ReportIntf) {
	if r == nil || d == nil {
		return
	}

	r.ReportID = d.GetReportID()
	r.DeveloperID = d.GetDeveloperID()
	r.AppID = d.GetAppID()
	r.SubjectId = d.GetSubjectId()
	r.Rev = d.GetRev()
}

// db->domain
func (r *Report) ToDomainReport() domain.ReportIntf {
	if r == nil {
		return nil
	}

	p := domain.Report{
		ReportID:    r.ReportID,
		DeveloperID: r.DeveloperID,
		AppID:       r.AppID,
		SubjectId:   r.SubjectId,
		Rev:         r.Rev,
	}
	return &p
}
