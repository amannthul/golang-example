// source: example/annotation/v1/annotation.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_protobuf_descriptor_pb = require('google-protobuf/google/protobuf/descriptor_pb.js');
goog.object.extend(proto, google_protobuf_descriptor_pb);
goog.exportSymbol('proto.example.annotation.v1.openRpc', null, global);
goog.exportSymbol('proto.example.annotation.v1.openService', null, global);

/**
 * A tuple of {field number, class constructor} for the extension
 * field named `openService`.
 * @type {!jspb.ExtensionFieldInfo<boolean>}
 */
proto.example.annotation.v1.openService = new jspb.ExtensionFieldInfo(
    52001,
    {openService: 0},
    null,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         null),
    0);

google_protobuf_descriptor_pb.ServiceOptions.extensionsBinary[52001] = new jspb.ExtensionFieldBinaryInfo(
    proto.example.annotation.v1.openService,
    jspb.BinaryReader.prototype.readBool,
    jspb.BinaryWriter.prototype.writeBool,
    undefined,
    undefined,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.ServiceOptions.extensions[52001] = proto.example.annotation.v1.openService;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `openRpc`.
 * @type {!jspb.ExtensionFieldInfo<boolean>}
 */
proto.example.annotation.v1.openRpc = new jspb.ExtensionFieldInfo(
    53001,
    {openRpc: 0},
    null,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         null),
    0);

google_protobuf_descriptor_pb.MethodOptions.extensionsBinary[53001] = new jspb.ExtensionFieldBinaryInfo(
    proto.example.annotation.v1.openRpc,
    jspb.BinaryReader.prototype.readBool,
    jspb.BinaryWriter.prototype.writeBool,
    undefined,
    undefined,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.MethodOptions.extensions[53001] = proto.example.annotation.v1.openRpc;

goog.object.extend(exports, proto.example.annotation.v1);
