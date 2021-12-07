// GENERATED CODE -- DO NOT EDIT!

// package: example.report.v1
// file: example/report/v1/report_api.proto

import * as example_report_v1_report_api_pb from "../../../example/report/v1/report_api_pb";
import * as grpc from "grpc";

interface IReportAPIService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  sayHello: grpc.MethodDefinition<example_report_v1_report_api_pb.SayHelloRequest, example_report_v1_report_api_pb.SayHelloResponse>;
}

export const ReportAPIService: IReportAPIService;

export interface IReportAPIServer extends grpc.UntypedServiceImplementation {
  sayHello: grpc.handleUnaryCall<example_report_v1_report_api_pb.SayHelloRequest, example_report_v1_report_api_pb.SayHelloResponse>;
}

export class ReportAPIClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  sayHello(argument: example_report_v1_report_api_pb.SayHelloRequest, callback: grpc.requestCallback<example_report_v1_report_api_pb.SayHelloResponse>): grpc.ClientUnaryCall;
  sayHello(argument: example_report_v1_report_api_pb.SayHelloRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<example_report_v1_report_api_pb.SayHelloResponse>): grpc.ClientUnaryCall;
  sayHello(argument: example_report_v1_report_api_pb.SayHelloRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<example_report_v1_report_api_pb.SayHelloResponse>): grpc.ClientUnaryCall;
}
