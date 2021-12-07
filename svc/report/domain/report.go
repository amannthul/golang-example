package domain

type Report struct {
	ReportID    string
	DeveloperID string
	AppID       string
	SubjectId   string
	Rev         int32
}

type ReportMapper interface {
	ToDomainReportMapper
	FromDomainReportMapper
}

type ToDomainReportMapper interface {
	ToDomainReport() ReportIntf
}

type FromDomainReportMapper interface {
	FromDomainReport(ReportIntf)
}

type ReportIntf interface {
	GetReportID() string
	GetDeveloperID() string
	GetAppID() string
	GetSubjectId() string
	GetRev() int32
}

var _ ReportIntf = (*Report)(nil)

func (r *Report) GetReportID() string {
	if r == nil {
		return ""
	}
	return r.ReportID
}

func (r *Report) GetDeveloperID() string {
	if r == nil {
		return ""
	}
	return r.DeveloperID
}

func (r *Report) GetAppID() string {
	if r == nil {
		return ""
	}
	return r.AppID
}

func (r *Report) GetSubjectId() string {
	if r == nil {
		return ""
	}
	return r.SubjectId
}

func (r *Report) GetRev() int32 {
	if r == nil {
		return 0
	}
	return r.Rev
}
