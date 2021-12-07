package report

import (
	"context"
	gerr "errors"

	"github.com/amannthul/golang-example/pkg/micro/errors"
	"github.com/amannthul/golang-example/pkg/micro/errors/codes"
	pb "github.com/amannthul/golang-example/proto/gen/go/example/report/v1"
)

// SayHello 获取常客列表 (按首字母排序).
func (c *ReportAPIHandler) SayHello(ctx context.Context, req *pb.SayHelloRequest, rsp *pb.SayHelloResponse) error {

	// 验证request
	err := validateSayHelloRequest(req)
	if err != nil {
		return errors.Error(codes.InvalidRequest, err.Error())
	}

	// TODO: FIX
	return nil
}

// 验证request
func validateSayHelloRequest(req *pb.SayHelloRequest) error {
	if req.GetMessage() == "" {
		return gerr.New("message should not be empty")
	}
	return nil
}
