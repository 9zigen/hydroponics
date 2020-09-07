/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: state.proto */

/* Do not generate deprecated warnings for self */
#ifndef PROTOBUF_C__NO_DEPRECATED
#define PROTOBUF_C__NO_DEPRECATED
#endif

#include "state.pb-c.h"
void   hydroponics__state_task__init
                     (Hydroponics__StateTask         *message)
{
  static const Hydroponics__StateTask init_value = HYDROPONICS__STATE_TASK__INIT;
  *message = init_value;
}
size_t hydroponics__state_task__get_packed_size
                     (const Hydroponics__StateTask *message)
{
  assert(message->base.descriptor == &hydroponics__state_task__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t hydroponics__state_task__pack
                     (const Hydroponics__StateTask *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &hydroponics__state_task__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t hydroponics__state_task__pack_to_buffer
                     (const Hydroponics__StateTask *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &hydroponics__state_task__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Hydroponics__StateTask *
       hydroponics__state_task__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Hydroponics__StateTask *)
     protobuf_c_message_unpack (&hydroponics__state_task__descriptor,
                                allocator, len, data);
}
void   hydroponics__state_task__free_unpacked
                     (Hydroponics__StateTask *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &hydroponics__state_task__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   hydroponics__state_memory__init
                     (Hydroponics__StateMemory         *message)
{
  static const Hydroponics__StateMemory init_value = HYDROPONICS__STATE_MEMORY__INIT;
  *message = init_value;
}
size_t hydroponics__state_memory__get_packed_size
                     (const Hydroponics__StateMemory *message)
{
  assert(message->base.descriptor == &hydroponics__state_memory__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t hydroponics__state_memory__pack
                     (const Hydroponics__StateMemory *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &hydroponics__state_memory__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t hydroponics__state_memory__pack_to_buffer
                     (const Hydroponics__StateMemory *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &hydroponics__state_memory__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Hydroponics__StateMemory *
       hydroponics__state_memory__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Hydroponics__StateMemory *)
     protobuf_c_message_unpack (&hydroponics__state_memory__descriptor,
                                allocator, len, data);
}
void   hydroponics__state_memory__free_unpacked
                     (Hydroponics__StateMemory *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &hydroponics__state_memory__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   hydroponics__state_telemetry__value__init
                     (Hydroponics__StateTelemetry__Value         *message)
{
  static const Hydroponics__StateTelemetry__Value init_value = HYDROPONICS__STATE_TELEMETRY__VALUE__INIT;
  *message = init_value;
}
void   hydroponics__state_telemetry__init
                     (Hydroponics__StateTelemetry         *message)
{
  static const Hydroponics__StateTelemetry init_value = HYDROPONICS__STATE_TELEMETRY__INIT;
  *message = init_value;
}
size_t hydroponics__state_telemetry__get_packed_size
                     (const Hydroponics__StateTelemetry *message)
{
  assert(message->base.descriptor == &hydroponics__state_telemetry__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t hydroponics__state_telemetry__pack
                     (const Hydroponics__StateTelemetry *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &hydroponics__state_telemetry__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t hydroponics__state_telemetry__pack_to_buffer
                     (const Hydroponics__StateTelemetry *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &hydroponics__state_telemetry__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Hydroponics__StateTelemetry *
       hydroponics__state_telemetry__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Hydroponics__StateTelemetry *)
     protobuf_c_message_unpack (&hydroponics__state_telemetry__descriptor,
                                allocator, len, data);
}
void   hydroponics__state_telemetry__free_unpacked
                     (Hydroponics__StateTelemetry *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &hydroponics__state_telemetry__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   hydroponics__state_output__init
                     (Hydroponics__StateOutput         *message)
{
  static const Hydroponics__StateOutput init_value = HYDROPONICS__STATE_OUTPUT__INIT;
  *message = init_value;
}
size_t hydroponics__state_output__get_packed_size
                     (const Hydroponics__StateOutput *message)
{
  assert(message->base.descriptor == &hydroponics__state_output__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t hydroponics__state_output__pack
                     (const Hydroponics__StateOutput *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &hydroponics__state_output__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t hydroponics__state_output__pack_to_buffer
                     (const Hydroponics__StateOutput *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &hydroponics__state_output__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Hydroponics__StateOutput *
       hydroponics__state_output__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Hydroponics__StateOutput *)
     protobuf_c_message_unpack (&hydroponics__state_output__descriptor,
                                allocator, len, data);
}
void   hydroponics__state_output__free_unpacked
                     (Hydroponics__StateOutput *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &hydroponics__state_output__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   hydroponics__state__init
                     (Hydroponics__State         *message)
{
  static const Hydroponics__State init_value = HYDROPONICS__STATE__INIT;
  *message = init_value;
}
size_t hydroponics__state__get_packed_size
                     (const Hydroponics__State *message)
{
  assert(message->base.descriptor == &hydroponics__state__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t hydroponics__state__pack
                     (const Hydroponics__State *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &hydroponics__state__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t hydroponics__state__pack_to_buffer
                     (const Hydroponics__State *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &hydroponics__state__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Hydroponics__State *
       hydroponics__state__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Hydroponics__State *)
     protobuf_c_message_unpack (&hydroponics__state__descriptor,
                                allocator, len, data);
}
void   hydroponics__state__free_unpacked
                     (Hydroponics__State *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &hydroponics__state__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   hydroponics__states__init
                     (Hydroponics__States         *message)
{
  static const Hydroponics__States init_value = HYDROPONICS__STATES__INIT;
  *message = init_value;
}
size_t hydroponics__states__get_packed_size
                     (const Hydroponics__States *message)
{
  assert(message->base.descriptor == &hydroponics__states__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t hydroponics__states__pack
                     (const Hydroponics__States *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &hydroponics__states__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t hydroponics__states__pack_to_buffer
                     (const Hydroponics__States *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &hydroponics__states__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Hydroponics__States *
       hydroponics__states__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Hydroponics__States *)
     protobuf_c_message_unpack (&hydroponics__states__descriptor,
                                allocator, len, data);
}
void   hydroponics__states__free_unpacked
                     (Hydroponics__States *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &hydroponics__states__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
static const ProtobufCEnumValue hydroponics__state_task__state__enum_values_by_number[5] =
{
  { "RUNNING", "HYDROPONICS__STATE_TASK__STATE__RUNNING", 0 },
  { "READY", "HYDROPONICS__STATE_TASK__STATE__READY", 1 },
  { "BLOCKED", "HYDROPONICS__STATE_TASK__STATE__BLOCKED", 2 },
  { "SUSPENDED", "HYDROPONICS__STATE_TASK__STATE__SUSPENDED", 3 },
  { "DELETED", "HYDROPONICS__STATE_TASK__STATE__DELETED", 4 },
};
static const ProtobufCIntRange hydroponics__state_task__state__value_ranges[] = {
{0, 0},{0, 5}
};
static const ProtobufCEnumValueIndex hydroponics__state_task__state__enum_values_by_name[5] =
{
  { "BLOCKED", 2 },
  { "DELETED", 4 },
  { "READY", 1 },
  { "RUNNING", 0 },
  { "SUSPENDED", 3 },
};
const ProtobufCEnumDescriptor hydroponics__state_task__state__descriptor =
{
  PROTOBUF_C__ENUM_DESCRIPTOR_MAGIC,
  "hydroponics.StateTask.State",
  "State",
  "Hydroponics__StateTask__State",
  "hydroponics",
  5,
  hydroponics__state_task__state__enum_values_by_number,
  5,
  hydroponics__state_task__state__enum_values_by_name,
  1,
  hydroponics__state_task__state__value_ranges,
  NULL,NULL,NULL,NULL   /* reserved[1234] */
};
static const ProtobufCFieldDescriptor hydroponics__state_task__field_descriptors[6] =
{
  {
    "name",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Hydroponics__StateTask, name),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "state",
    2,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_ENUM,
    0,   /* quantifier_offset */
    offsetof(Hydroponics__StateTask, state),
    &hydroponics__state_task__state__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "priority",
    3,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Hydroponics__StateTask, priority),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "runtime",
    4,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT64,
    0,   /* quantifier_offset */
    offsetof(Hydroponics__StateTask, runtime),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "stats",
    5,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Hydroponics__StateTask, stats),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "highwater",
    6,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Hydroponics__StateTask, highwater),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned hydroponics__state_task__field_indices_by_name[] = {
  5,   /* field[5] = highwater */
  0,   /* field[0] = name */
  2,   /* field[2] = priority */
  3,   /* field[3] = runtime */
  1,   /* field[1] = state */
  4,   /* field[4] = stats */
};
static const ProtobufCIntRange hydroponics__state_task__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 6 }
};
const ProtobufCMessageDescriptor hydroponics__state_task__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "hydroponics.StateTask",
  "StateTask",
  "Hydroponics__StateTask",
  "hydroponics",
  sizeof(Hydroponics__StateTask),
  6,
  hydroponics__state_task__field_descriptors,
  hydroponics__state_task__field_indices_by_name,
  1,  hydroponics__state_task__number_ranges,
  (ProtobufCMessageInit) hydroponics__state_task__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCFieldDescriptor hydroponics__state_memory__field_descriptors[2] =
{
  {
    "min_free",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Hydroponics__StateMemory, min_free),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "free",
    2,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Hydroponics__StateMemory, free),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned hydroponics__state_memory__field_indices_by_name[] = {
  1,   /* field[1] = free */
  0,   /* field[0] = min_free */
};
static const ProtobufCIntRange hydroponics__state_memory__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 2 }
};
const ProtobufCMessageDescriptor hydroponics__state_memory__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "hydroponics.StateMemory",
  "StateMemory",
  "Hydroponics__StateMemory",
  "hydroponics",
  sizeof(Hydroponics__StateMemory),
  2,
  hydroponics__state_memory__field_descriptors,
  hydroponics__state_memory__field_indices_by_name,
  1,  hydroponics__state_memory__number_ranges,
  (ProtobufCMessageInit) hydroponics__state_memory__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCFieldDescriptor hydroponics__state_telemetry__value__field_descriptors[2] =
{
  {
    "type",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_ENUM,
    0,   /* quantifier_offset */
    offsetof(Hydroponics__StateTelemetry__Value, type),
    &hydroponics__state_telemetry__type__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "value",
    2,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_FLOAT,
    0,   /* quantifier_offset */
    offsetof(Hydroponics__StateTelemetry__Value, value),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned hydroponics__state_telemetry__value__field_indices_by_name[] = {
  0,   /* field[0] = type */
  1,   /* field[1] = value */
};
static const ProtobufCIntRange hydroponics__state_telemetry__value__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 2 }
};
const ProtobufCMessageDescriptor hydroponics__state_telemetry__value__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "hydroponics.StateTelemetry.Value",
  "Value",
  "Hydroponics__StateTelemetry__Value",
  "hydroponics",
  sizeof(Hydroponics__StateTelemetry__Value),
  2,
  hydroponics__state_telemetry__value__field_descriptors,
  hydroponics__state_telemetry__value__field_indices_by_name,
  1,  hydroponics__state_telemetry__value__number_ranges,
  (ProtobufCMessageInit) hydroponics__state_telemetry__value__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCEnumValue hydroponics__state_telemetry__type__enum_values_by_number[11] =
{
  { "UNKNOWN", "HYDROPONICS__STATE_TELEMETRY__TYPE__UNKNOWN", 0 },
  { "TEMP_INDOOR", "HYDROPONICS__STATE_TELEMETRY__TYPE__TEMP_INDOOR", 1 },
  { "TEMP_PROBE", "HYDROPONICS__STATE_TELEMETRY__TYPE__TEMP_PROBE", 2 },
  { "HUMIDITY", "HYDROPONICS__STATE_TELEMETRY__TYPE__HUMIDITY", 3 },
  { "PRESSURE", "HYDROPONICS__STATE_TELEMETRY__TYPE__PRESSURE", 4 },
  { "EC_A", "HYDROPONICS__STATE_TELEMETRY__TYPE__EC_A", 5 },
  { "EC_B", "HYDROPONICS__STATE_TELEMETRY__TYPE__EC_B", 6 },
  { "PH_A", "HYDROPONICS__STATE_TELEMETRY__TYPE__PH_A", 7 },
  { "PH_B", "HYDROPONICS__STATE_TELEMETRY__TYPE__PH_B", 8 },
  { "TANK_A", "HYDROPONICS__STATE_TELEMETRY__TYPE__TANK_A", 9 },
  { "TANK_B", "HYDROPONICS__STATE_TELEMETRY__TYPE__TANK_B", 10 },
};
static const ProtobufCIntRange hydroponics__state_telemetry__type__value_ranges[] = {
{0, 0},{0, 11}
};
static const ProtobufCEnumValueIndex hydroponics__state_telemetry__type__enum_values_by_name[11] =
{
  { "EC_A", 5 },
  { "EC_B", 6 },
  { "HUMIDITY", 3 },
  { "PH_A", 7 },
  { "PH_B", 8 },
  { "PRESSURE", 4 },
  { "TANK_A", 9 },
  { "TANK_B", 10 },
  { "TEMP_INDOOR", 1 },
  { "TEMP_PROBE", 2 },
  { "UNKNOWN", 0 },
};
const ProtobufCEnumDescriptor hydroponics__state_telemetry__type__descriptor =
{
  PROTOBUF_C__ENUM_DESCRIPTOR_MAGIC,
  "hydroponics.StateTelemetry.Type",
  "Type",
  "Hydroponics__StateTelemetry__Type",
  "hydroponics",
  11,
  hydroponics__state_telemetry__type__enum_values_by_number,
  11,
  hydroponics__state_telemetry__type__enum_values_by_name,
  1,
  hydroponics__state_telemetry__type__value_ranges,
  NULL,NULL,NULL,NULL   /* reserved[1234] */
};
static const ProtobufCFieldDescriptor hydroponics__state_telemetry__field_descriptors[1] =
{
  {
    "value",
    1,
    PROTOBUF_C_LABEL_REPEATED,
    PROTOBUF_C_TYPE_MESSAGE,
    offsetof(Hydroponics__StateTelemetry, n_value),
    offsetof(Hydroponics__StateTelemetry, value),
    &hydroponics__state_telemetry__value__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned hydroponics__state_telemetry__field_indices_by_name[] = {
  0,   /* field[0] = value */
};
static const ProtobufCIntRange hydroponics__state_telemetry__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 1 }
};
const ProtobufCMessageDescriptor hydroponics__state_telemetry__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "hydroponics.StateTelemetry",
  "StateTelemetry",
  "Hydroponics__StateTelemetry",
  "hydroponics",
  sizeof(Hydroponics__StateTelemetry),
  1,
  hydroponics__state_telemetry__field_descriptors,
  hydroponics__state_telemetry__field_indices_by_name,
  1,  hydroponics__state_telemetry__number_ranges,
  (ProtobufCMessageInit) hydroponics__state_telemetry__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCFieldDescriptor hydroponics__state_output__field_descriptors[2] =
{
  {
    "output",
    1,
    PROTOBUF_C_LABEL_REPEATED,
    PROTOBUF_C_TYPE_ENUM,
    offsetof(Hydroponics__StateOutput, n_output),
    offsetof(Hydroponics__StateOutput, output),
    &hydroponics__output__descriptor,
    NULL,
    0 | PROTOBUF_C_FIELD_FLAG_PACKED,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "state",
    2,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_ENUM,
    0,   /* quantifier_offset */
    offsetof(Hydroponics__StateOutput, state),
    &hydroponics__output_state__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned hydroponics__state_output__field_indices_by_name[] = {
  0,   /* field[0] = output */
  1,   /* field[1] = state */
};
static const ProtobufCIntRange hydroponics__state_output__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 2 }
};
const ProtobufCMessageDescriptor hydroponics__state_output__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "hydroponics.StateOutput",
  "StateOutput",
  "Hydroponics__StateOutput",
  "hydroponics",
  sizeof(Hydroponics__StateOutput),
  2,
  hydroponics__state_output__field_descriptors,
  hydroponics__state_output__field_indices_by_name,
  1,  hydroponics__state_output__number_ranges,
  (ProtobufCMessageInit) hydroponics__state_output__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCFieldDescriptor hydroponics__state__field_descriptors[5] =
{
  {
    "timestamp",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT64,
    0,   /* quantifier_offset */
    offsetof(Hydroponics__State, timestamp),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "task",
    2,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_MESSAGE,
    offsetof(Hydroponics__State, state_case),
    offsetof(Hydroponics__State, task),
    &hydroponics__state_task__descriptor,
    NULL,
    0 | PROTOBUF_C_FIELD_FLAG_ONEOF,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "memory",
    3,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_MESSAGE,
    offsetof(Hydroponics__State, state_case),
    offsetof(Hydroponics__State, memory),
    &hydroponics__state_memory__descriptor,
    NULL,
    0 | PROTOBUF_C_FIELD_FLAG_ONEOF,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "telemetry",
    4,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_MESSAGE,
    offsetof(Hydroponics__State, state_case),
    offsetof(Hydroponics__State, telemetry),
    &hydroponics__state_telemetry__descriptor,
    NULL,
    0 | PROTOBUF_C_FIELD_FLAG_ONEOF,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "output",
    5,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_MESSAGE,
    offsetof(Hydroponics__State, state_case),
    offsetof(Hydroponics__State, output),
    &hydroponics__state_output__descriptor,
    NULL,
    0 | PROTOBUF_C_FIELD_FLAG_ONEOF,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned hydroponics__state__field_indices_by_name[] = {
  2,   /* field[2] = memory */
  4,   /* field[4] = output */
  1,   /* field[1] = task */
  3,   /* field[3] = telemetry */
  0,   /* field[0] = timestamp */
};
static const ProtobufCIntRange hydroponics__state__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 5 }
};
const ProtobufCMessageDescriptor hydroponics__state__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "hydroponics.State",
  "State",
  "Hydroponics__State",
  "hydroponics",
  sizeof(Hydroponics__State),
  5,
  hydroponics__state__field_descriptors,
  hydroponics__state__field_indices_by_name,
  1,  hydroponics__state__number_ranges,
  (ProtobufCMessageInit) hydroponics__state__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCFieldDescriptor hydroponics__states__field_descriptors[1] =
{
  {
    "state",
    1,
    PROTOBUF_C_LABEL_REPEATED,
    PROTOBUF_C_TYPE_MESSAGE,
    offsetof(Hydroponics__States, n_state),
    offsetof(Hydroponics__States, state),
    &hydroponics__state__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned hydroponics__states__field_indices_by_name[] = {
  0,   /* field[0] = state */
};
static const ProtobufCIntRange hydroponics__states__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 1 }
};
const ProtobufCMessageDescriptor hydroponics__states__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "hydroponics.States",
  "States",
  "Hydroponics__States",
  "hydroponics",
  sizeof(Hydroponics__States),
  1,
  hydroponics__states__field_descriptors,
  hydroponics__states__field_indices_by_name,
  1,  hydroponics__states__number_ranges,
  (ProtobufCMessageInit) hydroponics__states__init,
  NULL,NULL,NULL    /* reserved[123] */
};
