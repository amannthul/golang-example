// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: example/report/v1/report_api.proto

package com.example.report.v1;

public final class ReportApiProto {
  private ReportApiProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_example_report_v1_SayHelloRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_example_report_v1_SayHelloRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_example_report_v1_SayHelloResponse_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_example_report_v1_SayHelloResponse_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\"example/report/v1/report_api.proto\022\021ex" +
      "ample.report.v1\032&example/annotation/v1/a" +
      "nnotation.proto\"\"\n\017SayHelloRequest\022\017\n\007me" +
      "ssage\030\001 \001(\t\"\036\n\020SayHelloResponse\022\n\n\002ok\030\001 " +
      "\001(\0102l\n\tReportAPI\022Y\n\010SayHello\022\".example.r" +
      "eport.v1.SayHelloRequest\032#.example.repor" +
      "t.v1.SayHelloResponse\"\004\310\360\031\001\032\004\210\262\031\001B\244\001\n\025co" +
      "m.example.report.v1B\016ReportApiProtoP\001ZKg" +
      "ithub.com/amannthul/golang-example/proto" +
      "/gen/go/example/report/v1;reportv1\242\002\003ERX" +
      "\252\002\021Example.Report.V1\312\002\021Example\\Report\\V1" +
      "b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.example.annotation.v1.AnnotationProto.getDescriptor(),
        });
    internal_static_example_report_v1_SayHelloRequest_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_example_report_v1_SayHelloRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_example_report_v1_SayHelloRequest_descriptor,
        new java.lang.String[] { "Message", });
    internal_static_example_report_v1_SayHelloResponse_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_example_report_v1_SayHelloResponse_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_example_report_v1_SayHelloResponse_descriptor,
        new java.lang.String[] { "Ok", });
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(com.example.annotation.v1.AnnotationProto.openRpc);
    registry.add(com.example.annotation.v1.AnnotationProto.openService);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
    com.example.annotation.v1.AnnotationProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}