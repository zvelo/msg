/**
 * @fileoverview gRPC Web JS generated client stub for zvelo.msg
 * @enhanceable
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!


goog.provide('proto.zvelo.msg.APIv1Client');

goog.require('grpc.web.GrpcWebClientBase');
goog.require('grpc.web.AbstractClientBase');
goog.require('grpc.web.ClientReadableStream');
goog.require('grpc.web.Error');
goog.require('proto.google.protobuf.Empty');
goog.require('proto.zvelo.msg.QueryReplies');
goog.require('proto.zvelo.msg.QueryRequests');
goog.require('proto.zvelo.msg.QueryResult');
goog.require('proto.zvelo.msg.RequestID');
goog.require('proto.zvelo.msg.Suggestion');



goog.scope(function() {

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @constructor
 * @struct
 * @final
 */
proto.zvelo.msg.APIv1Client =
    function(hostname, credentials, options) {
  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zvelo.msg.QueryRequests,
 *   !proto.zvelo.msg.QueryReplies>}
 */
const methodInfo_Query = new grpc.web.AbstractClientBase.MethodInfo(
  proto.zvelo.msg.QueryReplies,
  /** @param {!proto.zvelo.msg.QueryRequests} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.zvelo.msg.QueryReplies.deserializeBinary
);


/**
 * @param {!proto.zvelo.msg.QueryRequests} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.zvelo.msg.QueryReplies)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.zvelo.msg.QueryReplies>|undefined}
 *     The XHR Node Readable Stream
 */
proto.zvelo.msg.APIv1Client.prototype.query =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/zvelo.msg.APIv1/Query',
      request,
      metadata,
      methodInfo_Query,
      callback);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zvelo.msg.RequestID,
 *   !proto.zvelo.msg.QueryResult>}
 */
const methodInfo_Result = new grpc.web.AbstractClientBase.MethodInfo(
  proto.zvelo.msg.QueryResult,
  /** @param {!proto.zvelo.msg.RequestID} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.zvelo.msg.QueryResult.deserializeBinary
);


/**
 * @param {!proto.zvelo.msg.RequestID} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.zvelo.msg.QueryResult)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.zvelo.msg.QueryResult>|undefined}
 *     The XHR Node Readable Stream
 */
proto.zvelo.msg.APIv1Client.prototype.result =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/zvelo.msg.APIv1/Result',
      request,
      metadata,
      methodInfo_Result,
      callback);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zvelo.msg.Suggestion,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_Suggest = new grpc.web.AbstractClientBase.MethodInfo(
  proto.google.protobuf.Empty,
  /** @param {!proto.zvelo.msg.Suggestion} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.google.protobuf.Empty.deserializeBinary
);


/**
 * @param {!proto.zvelo.msg.Suggestion} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.zvelo.msg.APIv1Client.prototype.suggest =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/zvelo.msg.APIv1/Suggest',
      request,
      metadata,
      methodInfo_Suggest,
      callback);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.google.protobuf.Empty,
 *   !proto.zvelo.msg.QueryResult>}
 */
const methodInfo_Stream = new grpc.web.AbstractClientBase.MethodInfo(
  proto.zvelo.msg.QueryResult,
  /** @param {!proto.google.protobuf.Empty} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.zvelo.msg.QueryResult.deserializeBinary
);


/**
 * @param {!proto.google.protobuf.Empty} request The request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.zvelo.msg.QueryResult>}
 *     The XHR Node Readable Stream
 */
proto.zvelo.msg.APIv1Client.prototype.stream =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/zvelo.msg.APIv1/Stream',
      request,
      metadata,
      methodInfo_Stream);
};


}); // goog.scope

