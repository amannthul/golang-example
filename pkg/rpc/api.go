package rpc

import (
	"context"

	mlog "github.com/amannthul/golang-example/pkg/log"
	"github.com/sirupsen/logrus"
)

type APIService struct {
	name string
}

func (svc *APIService) GetLogger(ctx context.Context) *logrus.Entry {
	l := LoggerFromContext(ctx)
	if l == nil {
		l = logrus.NewEntry(mlog.StandardLogger())
	}

	return l
}

func (svc *APIService) Name() string {
	return svc.name
}
