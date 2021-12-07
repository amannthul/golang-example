import * as grpcWeb from 'grpc-web';

import * as example_report_v1_report_api_pb from '../../../example/report/v1/report_api_pb';


export class ReportAPIClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  sayHello(
    request: example_report_v1_report_api_pb.SayHelloRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: example_report_v1_report_api_pb.SayHelloResponse) => void
  ): grpcWeb.ClientReadableStream<example_report_v1_report_api_pb.SayHelloResponse>;

}

export class ReportAPIPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  sayHello(
    request: example_report_v1_report_api_pb.SayHelloRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<example_report_v1_report_api_pb.SayHelloResponse>;

}

