package service

import mlog "github.com/amannthul/golang-example/pkg/log"

var (
	// log is the package global logger
	log = mlog.StandardLogger()
)

func Logger() *mlog.Logger {
	return mlog.StandardLogger()
}
