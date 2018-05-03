/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.zvelo.msg.QueryReplies');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.zvelo.msg.QueryReply');


/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.zvelo.msg.QueryReplies = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.zvelo.msg.QueryReplies.repeatedFields_, null);
};
goog.inherits(proto.zvelo.msg.QueryReplies, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.zvelo.msg.QueryReplies.displayName = 'proto.zvelo.msg.QueryReplies';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.zvelo.msg.QueryReplies.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.zvelo.msg.QueryReplies.prototype.toObject = function(opt_includeInstance) {
  return proto.zvelo.msg.QueryReplies.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.zvelo.msg.QueryReplies} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.zvelo.msg.QueryReplies.toObject = function(includeInstance, msg) {
  var f, obj = {
    replyList: jspb.Message.toObjectList(msg.getReplyList(),
    proto.zvelo.msg.QueryReply.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.zvelo.msg.QueryReplies}
 */
proto.zvelo.msg.QueryReplies.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.zvelo.msg.QueryReplies;
  return proto.zvelo.msg.QueryReplies.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.zvelo.msg.QueryReplies} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.zvelo.msg.QueryReplies}
 */
proto.zvelo.msg.QueryReplies.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.zvelo.msg.QueryReply;
      reader.readMessage(value,proto.zvelo.msg.QueryReply.deserializeBinaryFromReader);
      msg.addReply(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.zvelo.msg.QueryReplies.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.zvelo.msg.QueryReplies.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.zvelo.msg.QueryReplies} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.zvelo.msg.QueryReplies.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getReplyList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.zvelo.msg.QueryReply.serializeBinaryToWriter
    );
  }
};


/**
 * repeated QueryReply reply = 1;
 * @return {!Array.<!proto.zvelo.msg.QueryReply>}
 */
proto.zvelo.msg.QueryReplies.prototype.getReplyList = function() {
  return /** @type{!Array.<!proto.zvelo.msg.QueryReply>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.zvelo.msg.QueryReply, 1));
};


/** @param {!Array.<!proto.zvelo.msg.QueryReply>} value */
proto.zvelo.msg.QueryReplies.prototype.setReplyList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.zvelo.msg.QueryReply=} opt_value
 * @param {number=} opt_index
 * @return {!proto.zvelo.msg.QueryReply}
 */
proto.zvelo.msg.QueryReplies.prototype.addReply = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.zvelo.msg.QueryReply, opt_index);
};


proto.zvelo.msg.QueryReplies.prototype.clearReplyList = function() {
  this.setReplyList([]);
};


