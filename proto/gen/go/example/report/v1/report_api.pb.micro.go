// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: example/report/v1/report_api.proto

package reportv1

import (
	fmt "fmt"
	math "math"

	_ "github.com/amannthul/golang-example/proto/gen/go/example/annotation/v1"
	proto "github.com/golang/protobuf/proto"

	context "context"

	api "github.com/micro/go-micro/v2/api"

	client "github.com/micro/go-micro/v2/client"

	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ReportAPI service

func NewReportAPIEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ReportAPI service

type ReportAPIService interface {
	// SayHello Test.
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...client.CallOption) (*SayHelloResponse, error)
}

type reportAPIService struct {
	c    client.Client
	name string
}

func NewReportAPIService(name string, c client.Client) ReportAPIService {
	return &reportAPIService{
		c:    c,
		name: name,
	}
}

func (c *reportAPIService) SayHello(ctx context.Context, in *SayHelloRequest, opts ...client.CallOption) (*SayHelloResponse, error) {
	req := c.c.NewRequest(c.name, "ReportAPI.SayHello", in)
	out := new(SayHelloResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ReportAPI service

type ReportAPIHandler interface {
	// SayHello Test.
	SayHello(context.Context, *SayHelloRequest, *SayHelloResponse) error
}

func RegisterReportAPIHandler(s server.Server, hdlr ReportAPIHandler, opts ...server.HandlerOption) error {
	type reportAPI interface {
		SayHello(ctx context.Context, in *SayHelloRequest, out *SayHelloResponse) error
	}
	type ReportAPI struct {
		reportAPI
	}
	h := &reportAPIHandler{hdlr}
	return s.Handle(s.NewHandler(&ReportAPI{h}, opts...))
}

type reportAPIHandler struct {
	ReportAPIHandler
}

func (h *reportAPIHandler) SayHello(ctx context.Context, in *SayHelloRequest, out *SayHelloResponse) error {
	return h.ReportAPIHandler.SayHello(ctx, in, out)
}
