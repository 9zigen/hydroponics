// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: state.proto

package hydroponics

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Task state that directly matches task.h eTaskState enumeration.
type StateTask_State int32

const (
	// A task is querying the state of itself, so must be running.
	StateTask_RUNNING StateTask_State = 0
	// The task being queried is in a read or pending ready list.
	StateTask_READY StateTask_State = 1
	// The task being queried is in the Blocked state.
	StateTask_BLOCKED StateTask_State = 2
	// The task being queried is in the Suspended state, or is in the Blocked state with an infinite time out.
	StateTask_SUSPENDED StateTask_State = 3
	// The task being queried has been deleted, but its TCB has not yet been freed.
	StateTask_DELETED StateTask_State = 4
)

// Enum value maps for StateTask_State.
var (
	StateTask_State_name = map[int32]string{
		0: "RUNNING",
		1: "READY",
		2: "BLOCKED",
		3: "SUSPENDED",
		4: "DELETED",
	}
	StateTask_State_value = map[string]int32{
		"RUNNING":   0,
		"READY":     1,
		"BLOCKED":   2,
		"SUSPENDED": 3,
		"DELETED":   4,
	}
)

func (x StateTask_State) Enum() *StateTask_State {
	p := new(StateTask_State)
	*p = x
	return p
}

func (x StateTask_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StateTask_State) Descriptor() protoreflect.EnumDescriptor {
	return file_state_proto_enumTypes[0].Descriptor()
}

func (StateTask_State) Type() protoreflect.EnumType {
	return &file_state_proto_enumTypes[0]
}

func (x StateTask_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StateTask_State.Descriptor instead.
func (StateTask_State) EnumDescriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{0, 0}
}

type StateTelemetry_Type int32

const (
	StateTelemetry_UNKNOWN     StateTelemetry_Type = 0
	StateTelemetry_TEMP_INDOOR StateTelemetry_Type = 1
	StateTelemetry_TEMP_PROBE  StateTelemetry_Type = 2
	StateTelemetry_HUMIDITY    StateTelemetry_Type = 3
	StateTelemetry_PRESSURE    StateTelemetry_Type = 4
	StateTelemetry_EC_A        StateTelemetry_Type = 5
	StateTelemetry_EC_B        StateTelemetry_Type = 6
	StateTelemetry_PH_A        StateTelemetry_Type = 7
	StateTelemetry_PH_B        StateTelemetry_Type = 8
	StateTelemetry_TANK_A      StateTelemetry_Type = 9
	StateTelemetry_TANK_B      StateTelemetry_Type = 10
)

// Enum value maps for StateTelemetry_Type.
var (
	StateTelemetry_Type_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "TEMP_INDOOR",
		2:  "TEMP_PROBE",
		3:  "HUMIDITY",
		4:  "PRESSURE",
		5:  "EC_A",
		6:  "EC_B",
		7:  "PH_A",
		8:  "PH_B",
		9:  "TANK_A",
		10: "TANK_B",
	}
	StateTelemetry_Type_value = map[string]int32{
		"UNKNOWN":     0,
		"TEMP_INDOOR": 1,
		"TEMP_PROBE":  2,
		"HUMIDITY":    3,
		"PRESSURE":    4,
		"EC_A":        5,
		"EC_B":        6,
		"PH_A":        7,
		"PH_B":        8,
		"TANK_A":      9,
		"TANK_B":      10,
	}
)

func (x StateTelemetry_Type) Enum() *StateTelemetry_Type {
	p := new(StateTelemetry_Type)
	*p = x
	return p
}

func (x StateTelemetry_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StateTelemetry_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_state_proto_enumTypes[1].Descriptor()
}

func (StateTelemetry_Type) Type() protoreflect.EnumType {
	return &file_state_proto_enumTypes[1]
}

func (x StateTelemetry_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StateTelemetry_Type.Descriptor instead.
func (StateTelemetry_Type) EnumDescriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{2, 0}
}

type StateTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	State     StateTask_State `protobuf:"varint,2,opt,name=state,proto3,enum=hydroponics.StateTask_State" json:"state,omitempty"`
	Priority  uint32          `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
	Runtime   uint64          `protobuf:"varint,4,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Stats     uint32          `protobuf:"varint,5,opt,name=stats,proto3" json:"stats,omitempty"`
	Highwater uint32          `protobuf:"varint,6,opt,name=highwater,proto3" json:"highwater,omitempty"`
}

func (x *StateTask) Reset() {
	*x = StateTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateTask) ProtoMessage() {}

func (x *StateTask) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateTask.ProtoReflect.Descriptor instead.
func (*StateTask) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{0}
}

func (x *StateTask) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StateTask) GetState() StateTask_State {
	if x != nil {
		return x.State
	}
	return StateTask_RUNNING
}

func (x *StateTask) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *StateTask) GetRuntime() uint64 {
	if x != nil {
		return x.Runtime
	}
	return 0
}

func (x *StateTask) GetStats() uint32 {
	if x != nil {
		return x.Stats
	}
	return 0
}

func (x *StateTask) GetHighwater() uint32 {
	if x != nil {
		return x.Highwater
	}
	return 0
}

type StateMemory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinFree uint32 `protobuf:"varint,1,opt,name=min_free,json=minFree,proto3" json:"min_free,omitempty"`
	Free    uint32 `protobuf:"varint,2,opt,name=free,proto3" json:"free,omitempty"`
}

func (x *StateMemory) Reset() {
	*x = StateMemory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateMemory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateMemory) ProtoMessage() {}

func (x *StateMemory) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateMemory.ProtoReflect.Descriptor instead.
func (*StateMemory) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{1}
}

func (x *StateMemory) GetMinFree() uint32 {
	if x != nil {
		return x.MinFree
	}
	return 0
}

func (x *StateMemory) GetFree() uint32 {
	if x != nil {
		return x.Free
	}
	return 0
}

type StateTelemetry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []*StateTelemetry_Value `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *StateTelemetry) Reset() {
	*x = StateTelemetry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateTelemetry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateTelemetry) ProtoMessage() {}

func (x *StateTelemetry) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateTelemetry.ProtoReflect.Descriptor instead.
func (*StateTelemetry) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{2}
}

func (x *StateTelemetry) GetValue() []*StateTelemetry_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type StateOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output []Output    `protobuf:"varint,1,rep,packed,name=output,proto3,enum=hydroponics.Output" json:"output,omitempty"`
	State  OutputState `protobuf:"varint,2,opt,name=state,proto3,enum=hydroponics.OutputState" json:"state,omitempty"`
}

func (x *StateOutput) Reset() {
	*x = StateOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateOutput) ProtoMessage() {}

func (x *StateOutput) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateOutput.ProtoReflect.Descriptor instead.
func (*StateOutput) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{3}
}

func (x *StateOutput) GetOutput() []Output {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *StateOutput) GetState() OutputState {
	if x != nil {
		return x.State
	}
	return OutputState_OFF
}

type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp uint64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are assignable to State:
	//	*State_Task
	//	*State_Memory
	//	*State_Telemetry
	//	*State_Output
	State isState_State `protobuf_oneof:"state"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{4}
}

func (x *State) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (m *State) GetState() isState_State {
	if m != nil {
		return m.State
	}
	return nil
}

func (x *State) GetTask() *StateTask {
	if x, ok := x.GetState().(*State_Task); ok {
		return x.Task
	}
	return nil
}

func (x *State) GetMemory() *StateMemory {
	if x, ok := x.GetState().(*State_Memory); ok {
		return x.Memory
	}
	return nil
}

func (x *State) GetTelemetry() *StateTelemetry {
	if x, ok := x.GetState().(*State_Telemetry); ok {
		return x.Telemetry
	}
	return nil
}

func (x *State) GetOutput() *StateOutput {
	if x, ok := x.GetState().(*State_Output); ok {
		return x.Output
	}
	return nil
}

type isState_State interface {
	isState_State()
}

type State_Task struct {
	Task *StateTask `protobuf:"bytes,2,opt,name=task,proto3,oneof"`
}

type State_Memory struct {
	Memory *StateMemory `protobuf:"bytes,3,opt,name=memory,proto3,oneof"`
}

type State_Telemetry struct {
	Telemetry *StateTelemetry `protobuf:"bytes,4,opt,name=telemetry,proto3,oneof"`
}

type State_Output struct {
	Output *StateOutput `protobuf:"bytes,5,opt,name=output,proto3,oneof"`
}

func (*State_Task) isState_State() {}

func (*State_Memory) isState_State() {}

func (*State_Telemetry) isState_State() {}

func (*State_Output) isState_State() {}

type States struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State []*State `protobuf:"bytes,1,rep,name=state,proto3" json:"state,omitempty"`
}

func (x *States) Reset() {
	*x = States{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *States) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*States) ProtoMessage() {}

func (x *States) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use States.ProtoReflect.Descriptor instead.
func (*States) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{5}
}

func (x *States) GetState() []*State {
	if x != nil {
		return x.State
	}
	return nil
}

type StateTelemetry_Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  StateTelemetry_Type `protobuf:"varint,1,opt,name=type,proto3,enum=hydroponics.StateTelemetry_Type" json:"type,omitempty"`
	Value float32             `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *StateTelemetry_Value) Reset() {
	*x = StateTelemetry_Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateTelemetry_Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateTelemetry_Value) ProtoMessage() {}

func (x *StateTelemetry_Value) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateTelemetry_Value.ProtoReflect.Descriptor instead.
func (*StateTelemetry_Value) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{2, 0}
}

func (x *StateTelemetry_Value) GetType() StateTelemetry_Type {
	if x != nil {
		return x.Type
	}
	return StateTelemetry_UNKNOWN
}

func (x *StateTelemetry_Value) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_state_proto protoreflect.FileDescriptor

var file_state_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68,
	0x79, 0x64, 0x72, 0x6f, 0x70, 0x6f, 0x6e, 0x69, 0x63, 0x73, 0x1a, 0x0c, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x02, 0x0a, 0x09, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x68, 0x79, 0x64, 0x72,
	0x6f, 0x70, 0x6f, 0x6e, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x69,
	0x67, 0x68, 0x77, 0x61, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x68,
	0x69, 0x67, 0x68, 0x77, 0x61, 0x74, 0x65, 0x72, 0x22, 0x48, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4c, 0x4f,
	0x43, 0x4b, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e,
	0x44, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44,
	0x10, 0x04, 0x22, 0x3c, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x46, 0x72, 0x65, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x72, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x66, 0x72, 0x65, 0x65,
	0x22, 0xb1, 0x02, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x12, 0x37, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x68, 0x79, 0x64, 0x72, 0x6f, 0x70, 0x6f, 0x6e, 0x69, 0x63, 0x73,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x53, 0x0a, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x68, 0x79, 0x64, 0x72, 0x6f, 0x70, 0x6f, 0x6e, 0x69, 0x63,
	0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x90, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x45, 0x4d, 0x50, 0x5f,
	0x49, 0x4e, 0x44, 0x4f, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x45, 0x4d, 0x50,
	0x5f, 0x50, 0x52, 0x4f, 0x42, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x48, 0x55, 0x4d, 0x49,
	0x44, 0x49, 0x54, 0x59, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x45, 0x53, 0x53, 0x55,
	0x52, 0x45, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x43, 0x5f, 0x41, 0x10, 0x05, 0x12, 0x08,
	0x0a, 0x04, 0x45, 0x43, 0x5f, 0x42, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x48, 0x5f, 0x41,
	0x10, 0x07, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x48, 0x5f, 0x42, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06,
	0x54, 0x41, 0x4e, 0x4b, 0x5f, 0x41, 0x10, 0x09, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x41, 0x4e, 0x4b,
	0x5f, 0x42, 0x10, 0x0a, 0x22, 0x6a, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x2b, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x68, 0x79, 0x64, 0x72, 0x6f, 0x70, 0x6f, 0x6e, 0x69, 0x63,
	0x73, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x2e, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x18, 0x2e, 0x68, 0x79, 0x64, 0x72, 0x6f, 0x70, 0x6f, 0x6e, 0x69, 0x63, 0x73, 0x2e, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x22, 0x81, 0x02, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68, 0x79, 0x64, 0x72, 0x6f, 0x70, 0x6f,
	0x6e, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x48, 0x00,
	0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x32, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x68, 0x79, 0x64, 0x72, 0x6f, 0x70, 0x6f,
	0x6e, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x48, 0x00, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x09, 0x74, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x68, 0x79, 0x64, 0x72, 0x6f, 0x70, 0x6f, 0x6e, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x48, 0x00, 0x52, 0x09, 0x74, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x68, 0x79, 0x64, 0x72, 0x6f, 0x70,
	0x6f, 0x6e, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x22, 0x32, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x28,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x68, 0x79, 0x64, 0x72, 0x6f, 0x70, 0x6f, 0x6e, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x43, 0x0a, 0x1d, 0x70, 0x74, 0x2e, 0x73,
	0x6f, 0x62, 0x72, 0x69, 0x6e, 0x68, 0x6f, 0x2e, 0x68, 0x79, 0x64, 0x72, 0x6f, 0x70, 0x6f, 0x6e,
	0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x73, 0x6f, 0x62, 0x72, 0x69, 0x6e, 0x68,
	0x6f, 0x2f, 0x68, 0x79, 0x64, 0x72, 0x6f, 0x70, 0x6f, 0x6e, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_state_proto_rawDescOnce sync.Once
	file_state_proto_rawDescData = file_state_proto_rawDesc
)

func file_state_proto_rawDescGZIP() []byte {
	file_state_proto_rawDescOnce.Do(func() {
		file_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_state_proto_rawDescData)
	})
	return file_state_proto_rawDescData
}

var file_state_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_state_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_state_proto_goTypes = []interface{}{
	(StateTask_State)(0),         // 0: hydroponics.StateTask.State
	(StateTelemetry_Type)(0),     // 1: hydroponics.StateTelemetry.Type
	(*StateTask)(nil),            // 2: hydroponics.StateTask
	(*StateMemory)(nil),          // 3: hydroponics.StateMemory
	(*StateTelemetry)(nil),       // 4: hydroponics.StateTelemetry
	(*StateOutput)(nil),          // 5: hydroponics.StateOutput
	(*State)(nil),                // 6: hydroponics.State
	(*States)(nil),               // 7: hydroponics.States
	(*StateTelemetry_Value)(nil), // 8: hydroponics.StateTelemetry.Value
	(Output)(0),                  // 9: hydroponics.Output
	(OutputState)(0),             // 10: hydroponics.OutputState
}
var file_state_proto_depIdxs = []int32{
	0,  // 0: hydroponics.StateTask.state:type_name -> hydroponics.StateTask.State
	8,  // 1: hydroponics.StateTelemetry.value:type_name -> hydroponics.StateTelemetry.Value
	9,  // 2: hydroponics.StateOutput.output:type_name -> hydroponics.Output
	10, // 3: hydroponics.StateOutput.state:type_name -> hydroponics.OutputState
	2,  // 4: hydroponics.State.task:type_name -> hydroponics.StateTask
	3,  // 5: hydroponics.State.memory:type_name -> hydroponics.StateMemory
	4,  // 6: hydroponics.State.telemetry:type_name -> hydroponics.StateTelemetry
	5,  // 7: hydroponics.State.output:type_name -> hydroponics.StateOutput
	6,  // 8: hydroponics.States.state:type_name -> hydroponics.State
	1,  // 9: hydroponics.StateTelemetry.Value.type:type_name -> hydroponics.StateTelemetry.Type
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_state_proto_init() }
func file_state_proto_init() {
	if File_state_proto != nil {
		return
	}
	file_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateTask); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateMemory); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateTelemetry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateOutput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_state_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_state_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*States); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_state_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateTelemetry_Value); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_state_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*State_Task)(nil),
		(*State_Memory)(nil),
		(*State_Telemetry)(nil),
		(*State_Output)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_state_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_state_proto_goTypes,
		DependencyIndexes: file_state_proto_depIdxs,
		EnumInfos:         file_state_proto_enumTypes,
		MessageInfos:      file_state_proto_msgTypes,
	}.Build()
	File_state_proto = out.File
	file_state_proto_rawDesc = nil
	file_state_proto_goTypes = nil
	file_state_proto_depIdxs = nil
}