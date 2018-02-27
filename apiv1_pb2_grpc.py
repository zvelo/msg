# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from zvelo.msg import query_pb2 as zvelo_dot_msg_dot_query__pb2
from zvelo.msg import suggest_pb2 as zvelo_dot_msg_dot_suggest__pb2


class APIv1Stub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Query = channel.unary_unary(
        '/zvelo.msg.APIv1/Query',
        request_serializer=zvelo_dot_msg_dot_query__pb2.QueryRequests.SerializeToString,
        response_deserializer=zvelo_dot_msg_dot_query__pb2.QueryReplies.FromString,
        )
    self.Result = channel.unary_unary(
        '/zvelo.msg.APIv1/Result',
        request_serializer=zvelo_dot_msg_dot_query__pb2.RequestID.SerializeToString,
        response_deserializer=zvelo_dot_msg_dot_query__pb2.QueryResult.FromString,
        )
    self.Suggest = channel.unary_unary(
        '/zvelo.msg.APIv1/Suggest',
        request_serializer=zvelo_dot_msg_dot_suggest__pb2.Suggestion.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.Stream = channel.unary_stream(
        '/zvelo.msg.APIv1/Stream',
        request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
        response_deserializer=zvelo_dot_msg_dot_query__pb2.QueryResult.FromString,
        )


class APIv1Servicer(object):
  # missing associated documentation comment in .proto file
  pass

  def Query(self, request, context):
    """Create new query
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Result(self, request, context):
    """Results of active or unexpired query
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Suggest(self, request, context):
    """Suggest new datasets for a URL
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Stream(self, request, context):
    """Stream returns all QueryResult messages processed by zveloAPI
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_APIv1Servicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Query': grpc.unary_unary_rpc_method_handler(
          servicer.Query,
          request_deserializer=zvelo_dot_msg_dot_query__pb2.QueryRequests.FromString,
          response_serializer=zvelo_dot_msg_dot_query__pb2.QueryReplies.SerializeToString,
      ),
      'Result': grpc.unary_unary_rpc_method_handler(
          servicer.Result,
          request_deserializer=zvelo_dot_msg_dot_query__pb2.RequestID.FromString,
          response_serializer=zvelo_dot_msg_dot_query__pb2.QueryResult.SerializeToString,
      ),
      'Suggest': grpc.unary_unary_rpc_method_handler(
          servicer.Suggest,
          request_deserializer=zvelo_dot_msg_dot_suggest__pb2.Suggestion.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'Stream': grpc.unary_stream_rpc_method_handler(
          servicer.Stream,
          request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
          response_serializer=zvelo_dot_msg_dot_query__pb2.QueryResult.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'zvelo.msg.APIv1', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))