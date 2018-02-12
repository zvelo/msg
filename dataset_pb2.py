# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: zvelo/msg/dataset.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from zvelo.msg import category_pb2 as zvelo_dot_msg_dot_category__pb2
from zvelo.msg import status_pb2 as zvelo_dot_msg_dot_status__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='zvelo/msg/dataset.proto',
  package='zvelo.msg',
  syntax='proto3',
  serialized_pb=_b('\n\x17zvelo/msg/dataset.proto\x12\tzvelo.msg\x1a\x18zvelo/msg/category.proto\x1a\x16zvelo/msg/status.proto\"\xc6\x04\n\x07\x44\x61taSet\x12\x39\n\x0e\x63\x61tegorization\x18\x01 \x01(\x0b\x32!.zvelo.msg.DataSet.Categorization\x12/\n\tmalicious\x18\x05 \x01(\x0b\x32\x1c.zvelo.msg.DataSet.Malicious\x12%\n\x04\x65\x63ho\x18\x06 \x01(\x0b\x32\x17.zvelo.msg.DataSet.Echo\x12-\n\x08language\x18\x08 \x01(\x0b\x32\x1b.zvelo.msg.DataSet.Language\x1a\x62\n\x0e\x43\x61tegorization\x12\"\n\x05value\x18\x03 \x03(\x0e\x32\x13.zvelo.msg.Category\x12 \n\x05\x65rror\x18\x04 \x01(\x0b\x32\x11.zvelo.msg.StatusJ\x04\x08\x01\x10\x02J\x04\x08\x02\x10\x03\x1a\x46\n\x08Language\x12\x0c\n\x04\x63ode\x18\x02 \x01(\t\x12 \n\x05\x65rror\x18\x04 \x01(\x0b\x32\x11.zvelo.msg.StatusJ\x04\x08\x01\x10\x02J\x04\x08\x03\x10\x04\x1a~\n\tMalicious\x12 \n\x05\x65rror\x18\x08 \x01(\x0b\x32\x11.zvelo.msg.Status\x12%\n\x08\x63\x61tegory\x18\t \x03(\x0e\x32\x13.zvelo.msg.CategoryJ\x04\x08\x01\x10\x02J\x04\x08\x02\x10\x03J\x04\x08\x03\x10\x04J\x04\x08\x04\x10\x05J\x04\x08\x05\x10\x06J\x04\x08\x06\x10\x07J\x04\x08\x07\x10\x08\x1a\x35\n\x04\x45\x63ho\x12\x0b\n\x03url\x18\x01 \x01(\t\x12 \n\x05\x65rror\x18\x02 \x01(\x0b\x32\x11.zvelo.msg.StatusJ\x04\x08\x02\x10\x03J\x04\x08\x03\x10\x04J\x04\x08\x04\x10\x05J\x04\x08\x07\x10\x08*H\n\x0b\x44\x61taSetType\x12\x12\n\x0e\x43\x41TEGORIZATION\x10\x00\x12\r\n\tMALICIOUS\x10\x04\x12\x08\n\x04\x45\x43HO\x10\x05\x12\x0c\n\x08LANGUAGE\x10\x07\x42\x0eZ\x0czvelo.io/msgb\x06proto3')
  ,
  dependencies=[zvelo_dot_msg_dot_category__pb2.DESCRIPTOR,zvelo_dot_msg_dot_status__pb2.DESCRIPTOR,])

_DATASETTYPE = _descriptor.EnumDescriptor(
  name='DataSetType',
  full_name='zvelo.msg.DataSetType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='CATEGORIZATION', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MALICIOUS', index=1, number=4,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ECHO', index=2, number=5,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='LANGUAGE', index=3, number=7,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=673,
  serialized_end=745,
)
_sym_db.RegisterEnumDescriptor(_DATASETTYPE)

DataSetType = enum_type_wrapper.EnumTypeWrapper(_DATASETTYPE)
CATEGORIZATION = 0
MALICIOUS = 4
ECHO = 5
LANGUAGE = 7



_DATASET_CATEGORIZATION = _descriptor.Descriptor(
  name='Categorization',
  full_name='zvelo.msg.DataSet.Categorization',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='zvelo.msg.DataSet.Categorization.value', index=0,
      number=3, type=14, cpp_type=8, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='error', full_name='zvelo.msg.DataSet.Categorization.error', index=1,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=294,
  serialized_end=392,
)

_DATASET_LANGUAGE = _descriptor.Descriptor(
  name='Language',
  full_name='zvelo.msg.DataSet.Language',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='zvelo.msg.DataSet.Language.code', index=0,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='error', full_name='zvelo.msg.DataSet.Language.error', index=1,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=394,
  serialized_end=464,
)

_DATASET_MALICIOUS = _descriptor.Descriptor(
  name='Malicious',
  full_name='zvelo.msg.DataSet.Malicious',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='error', full_name='zvelo.msg.DataSet.Malicious.error', index=0,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='category', full_name='zvelo.msg.DataSet.Malicious.category', index=1,
      number=9, type=14, cpp_type=8, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=466,
  serialized_end=592,
)

_DATASET_ECHO = _descriptor.Descriptor(
  name='Echo',
  full_name='zvelo.msg.DataSet.Echo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='url', full_name='zvelo.msg.DataSet.Echo.url', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='error', full_name='zvelo.msg.DataSet.Echo.error', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=594,
  serialized_end=647,
)

_DATASET = _descriptor.Descriptor(
  name='DataSet',
  full_name='zvelo.msg.DataSet',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='categorization', full_name='zvelo.msg.DataSet.categorization', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='malicious', full_name='zvelo.msg.DataSet.malicious', index=1,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='echo', full_name='zvelo.msg.DataSet.echo', index=2,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='language', full_name='zvelo.msg.DataSet.language', index=3,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[_DATASET_CATEGORIZATION, _DATASET_LANGUAGE, _DATASET_MALICIOUS, _DATASET_ECHO, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=89,
  serialized_end=671,
)

_DATASET_CATEGORIZATION.fields_by_name['value'].enum_type = zvelo_dot_msg_dot_category__pb2._CATEGORY
_DATASET_CATEGORIZATION.fields_by_name['error'].message_type = zvelo_dot_msg_dot_status__pb2._STATUS
_DATASET_CATEGORIZATION.containing_type = _DATASET
_DATASET_LANGUAGE.fields_by_name['error'].message_type = zvelo_dot_msg_dot_status__pb2._STATUS
_DATASET_LANGUAGE.containing_type = _DATASET
_DATASET_MALICIOUS.fields_by_name['error'].message_type = zvelo_dot_msg_dot_status__pb2._STATUS
_DATASET_MALICIOUS.fields_by_name['category'].enum_type = zvelo_dot_msg_dot_category__pb2._CATEGORY
_DATASET_MALICIOUS.containing_type = _DATASET
_DATASET_ECHO.fields_by_name['error'].message_type = zvelo_dot_msg_dot_status__pb2._STATUS
_DATASET_ECHO.containing_type = _DATASET
_DATASET.fields_by_name['categorization'].message_type = _DATASET_CATEGORIZATION
_DATASET.fields_by_name['malicious'].message_type = _DATASET_MALICIOUS
_DATASET.fields_by_name['echo'].message_type = _DATASET_ECHO
_DATASET.fields_by_name['language'].message_type = _DATASET_LANGUAGE
DESCRIPTOR.message_types_by_name['DataSet'] = _DATASET
DESCRIPTOR.enum_types_by_name['DataSetType'] = _DATASETTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

DataSet = _reflection.GeneratedProtocolMessageType('DataSet', (_message.Message,), dict(

  Categorization = _reflection.GeneratedProtocolMessageType('Categorization', (_message.Message,), dict(
    DESCRIPTOR = _DATASET_CATEGORIZATION,
    __module__ = 'zvelo.msg.dataset_pb2'
    # @@protoc_insertion_point(class_scope:zvelo.msg.DataSet.Categorization)
    ))
  ,

  Language = _reflection.GeneratedProtocolMessageType('Language', (_message.Message,), dict(
    DESCRIPTOR = _DATASET_LANGUAGE,
    __module__ = 'zvelo.msg.dataset_pb2'
    # @@protoc_insertion_point(class_scope:zvelo.msg.DataSet.Language)
    ))
  ,

  Malicious = _reflection.GeneratedProtocolMessageType('Malicious', (_message.Message,), dict(
    DESCRIPTOR = _DATASET_MALICIOUS,
    __module__ = 'zvelo.msg.dataset_pb2'
    # @@protoc_insertion_point(class_scope:zvelo.msg.DataSet.Malicious)
    ))
  ,

  Echo = _reflection.GeneratedProtocolMessageType('Echo', (_message.Message,), dict(
    DESCRIPTOR = _DATASET_ECHO,
    __module__ = 'zvelo.msg.dataset_pb2'
    # @@protoc_insertion_point(class_scope:zvelo.msg.DataSet.Echo)
    ))
  ,
  DESCRIPTOR = _DATASET,
  __module__ = 'zvelo.msg.dataset_pb2'
  # @@protoc_insertion_point(class_scope:zvelo.msg.DataSet)
  ))
_sym_db.RegisterMessage(DataSet)
_sym_db.RegisterMessage(DataSet.Categorization)
_sym_db.RegisterMessage(DataSet.Language)
_sym_db.RegisterMessage(DataSet.Malicious)
_sym_db.RegisterMessage(DataSet.Echo)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z\014zvelo.io/msg'))
# @@protoc_insertion_point(module_scope)
