/**
 * @fileoverview gRPC-Web generated client stub for example.report.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var example_annotation_v1_annotation_pb = require('../../../example/annotation/v1/annotation_pb.js')
const proto = {};
proto.example = {};
proto.example.report = {};
proto.example.report.v1 = require('./report_api_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.example.report.v1.ReportAPIClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.example.report.v1.ReportAPIPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.example.report.v1.SayHelloRequest,
 *   !proto.example.report.v1.SayHelloResponse>}
 */
const methodDescriptor_ReportAPI_SayHello = new grpc.web.MethodDescriptor(
  '/example.report.v1.ReportAPI/SayHello',
  grpc.web.MethodType.UNARY,
  proto.example.report.v1.SayHelloRequest,
  proto.example.report.v1.SayHelloResponse,
  /**
   * @param {!proto.example.report.v1.SayHelloRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.example.report.v1.SayHelloResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.example.report.v1.SayHelloRequest,
 *   !proto.example.report.v1.SayHelloResponse>}
 */
const methodInfo_ReportAPI_SayHello = new grpc.web.AbstractClientBase.MethodInfo(
  proto.example.report.v1.SayHelloResponse,
  /**
   * @param {!proto.example.report.v1.SayHelloRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.example.report.v1.SayHelloResponse.deserializeBinary
);


/**
 * @param {!proto.example.report.v1.SayHelloRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.example.report.v1.SayHelloResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.example.report.v1.SayHelloResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.example.report.v1.ReportAPIClient.prototype.sayHello =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/example.report.v1.ReportAPI/SayHello',
      request,
      metadata || {},
      methodDescriptor_ReportAPI_SayHello,
      callback);
};


/**
 * @param {!proto.example.report.v1.SayHelloRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.example.report.v1.SayHelloResponse>}
 *     Promise that resolves to the response
 */
proto.example.report.v1.ReportAPIPromiseClient.prototype.sayHello =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/example.report.v1.ReportAPI/SayHello',
      request,
      metadata || {},
      methodDescriptor_ReportAPI_SayHello);
};


module.exports = proto.example.report.v1;

