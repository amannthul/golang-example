package cid

import (
	"context"

	"github.com/micro/go-micro/v2/server"

	"github.com/amannthul/golang-example/pkg/micro/meta"
	"github.com/rs/xid"
)

// CidWrapper is a handler wrapper that generate a new cid if cid is not found from request
func CidWrapper(fn server.HandlerFunc) server.HandlerFunc {

	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		cid := meta.CidFromContext(ctx)
		// 如果没有找到 cid，则生成一个新的
		if cid == "" {
			cid = xid.New().String()
			ctx = meta.ContextWithCid(ctx, cid)
		}
		err := fn(ctx, req, rsp)
		return err
	}
}
