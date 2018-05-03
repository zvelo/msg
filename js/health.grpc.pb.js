/**
 * @fileoverview gRPC Web JS generated client stub for grpc.health.v1
 * @enhanceable
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!


goog.provide('proto.grpc.health.v1.HealthClient');

goog.require('grpc.web.GrpcWebClientBase');
goog.require('grpc.web.AbstractClientBase');
goog.require('grpc.web.ClientReadableStream');
goog.require('grpc.web.Error');
goog.require('proto.grpc.health.v1.HealthCheckRequest');
goog.require('proto.grpc.health.v1.HealthCheckResponse');



goog.scope(function() {

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @constructor
 * @struct
 * @final
 */
proto.grpc.health.v1.HealthClient =
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
 *   !proto.grpc.health.v1.HealthCheckRequest,
 *   !proto.grpc.health.v1.HealthCheckResponse>}
 */
const methodInfo_Check = new grpc.web.AbstractClientBase.MethodInfo(
  proto.grpc.health.v1.HealthCheckResponse,
  /** @param {!proto.grpc.health.v1.HealthCheckRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.grpc.health.v1.HealthCheckResponse.deserializeBinary
);


/**
 * @param {!proto.grpc.health.v1.HealthCheckRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.grpc.health.v1.HealthCheckResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.grpc.health.v1.HealthCheckResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.grpc.health.v1.HealthClient.prototype.check =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/grpc.health.v1.Health/Check',
      request,
      metadata,
      methodInfo_Check,
      callback);
};


}); // goog.scope

