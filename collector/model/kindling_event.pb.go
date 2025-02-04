// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kindling_event.proto

package model

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Source int32

const (
	Source_SOURCE_UNKNOWN Source = 0
	Source_SYSCALL_ENTER  Source = 1
	Source_SYSCALL_EXIT   Source = 2
	Source_TRACEPOINT     Source = 3
	Source_KRPOBE         Source = 4
	Source_KRETPROBE      Source = 5
	Source_UPROBE         Source = 6
	Source_URETPROBE      Source = 7
)

var Source_name = map[int32]string{
	0: "SOURCE_UNKNOWN",
	1: "SYSCALL_ENTER",
	2: "SYSCALL_EXIT",
	3: "TRACEPOINT",
	4: "KRPOBE",
	5: "KRETPROBE",
	6: "UPROBE",
	7: "URETPROBE",
}

var Source_value = map[string]int32{
	"SOURCE_UNKNOWN": 0,
	"SYSCALL_ENTER":  1,
	"SYSCALL_EXIT":   2,
	"TRACEPOINT":     3,
	"KRPOBE":         4,
	"KRETPROBE":      5,
	"UPROBE":         6,
	"URETPROBE":      7,
}

func (x Source) String() string {
	return proto.EnumName(Source_name, int32(x))
}

func (Source) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_81bb5d1665ce2a0c, []int{0}
}

type Category int32

const (
	Category_CAT_NONE      Category = 0
	Category_CAT_OTHER     Category = 1
	Category_CAT_FILE      Category = 2
	Category_CAT_NET       Category = 3
	Category_CAT_IPC       Category = 4
	Category_CAT_WAIT      Category = 5
	Category_CAT_SIGNAL    Category = 6
	Category_CAT_SLEEP     Category = 7
	Category_CAT_TIME      Category = 8
	Category_CAT_PROCESS   Category = 9
	Category_CAT_SCHEDULER Category = 10
	Category_CAT_MEMORY    Category = 11
	Category_CAT_USER      Category = 12
	Category_CAT_SYSTEM    Category = 13
)

var Category_name = map[int32]string{
	0:  "CAT_NONE",
	1:  "CAT_OTHER",
	2:  "CAT_FILE",
	3:  "CAT_NET",
	4:  "CAT_IPC",
	5:  "CAT_WAIT",
	6:  "CAT_SIGNAL",
	7:  "CAT_SLEEP",
	8:  "CAT_TIME",
	9:  "CAT_PROCESS",
	10: "CAT_SCHEDULER",
	11: "CAT_MEMORY",
	12: "CAT_USER",
	13: "CAT_SYSTEM",
}

var Category_value = map[string]int32{
	"CAT_NONE":      0,
	"CAT_OTHER":     1,
	"CAT_FILE":      2,
	"CAT_NET":       3,
	"CAT_IPC":       4,
	"CAT_WAIT":      5,
	"CAT_SIGNAL":    6,
	"CAT_SLEEP":     7,
	"CAT_TIME":      8,
	"CAT_PROCESS":   9,
	"CAT_SCHEDULER": 10,
	"CAT_MEMORY":    11,
	"CAT_USER":      12,
	"CAT_SYSTEM":    13,
}

func (x Category) String() string {
	return proto.EnumName(Category_name, int32(x))
}

func (Category) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_81bb5d1665ce2a0c, []int{1}
}

type ValueType int32

const (
	ValueType_NONE    ValueType = 0
	ValueType_INT8    ValueType = 1
	ValueType_INT16   ValueType = 2
	ValueType_INT32   ValueType = 3
	ValueType_INT64   ValueType = 4
	ValueType_UINT8   ValueType = 5
	ValueType_UINT16  ValueType = 6
	ValueType_UINT32  ValueType = 7
	ValueType_UINT64  ValueType = 8
	ValueType_CHARBUF ValueType = 9
	ValueType_BYTEBUF ValueType = 10
	ValueType_FLOAT   ValueType = 11
	ValueType_DOUBLE  ValueType = 12
	ValueType_BOOL    ValueType = 13
)

var ValueType_name = map[int32]string{
	0:  "NONE",
	1:  "INT8",
	2:  "INT16",
	3:  "INT32",
	4:  "INT64",
	5:  "UINT8",
	6:  "UINT16",
	7:  "UINT32",
	8:  "UINT64",
	9:  "CHARBUF",
	10: "BYTEBUF",
	11: "FLOAT",
	12: "DOUBLE",
	13: "BOOL",
}

var ValueType_value = map[string]int32{
	"NONE":    0,
	"INT8":    1,
	"INT16":   2,
	"INT32":   3,
	"INT64":   4,
	"UINT8":   5,
	"UINT16":  6,
	"UINT32":  7,
	"UINT64":  8,
	"CHARBUF": 9,
	"BYTEBUF": 10,
	"FLOAT":   11,
	"DOUBLE":  12,
	"BOOL":    13,
}

func (x ValueType) String() string {
	return proto.EnumName(ValueType_name, int32(x))
}

func (ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_81bb5d1665ce2a0c, []int{2}
}

// File Descriptor type
type FDType int32

const (
	FDType_FD_UNKNOWN       FDType = 0
	FDType_FD_FILE          FDType = 1
	FDType_FD_DIRECTORY     FDType = 2
	FDType_FD_IPV4_SOCK     FDType = 3
	FDType_FD_IPV6_SOCK     FDType = 4
	FDType_FD_IPV4_SERVSOCK FDType = 5
	FDType_FD_IPV6_SERVSOCK FDType = 6
	FDType_FD_FIFO          FDType = 7
	FDType_FD_UNIX_SOCK     FDType = 8
	FDType_FD_EVENT         FDType = 9
	FDType_FD_UNSUPPORTED   FDType = 10
	FDType_FD_SIGNALFD      FDType = 11
	FDType_FD_EVENTPOLL     FDType = 12
	FDType_FD_INOTIFY       FDType = 13
	FDType_FD_TIMERFD       FDType = 14
	FDType_FD_NETLINK       FDType = 15
	FDType_FD_FILE_V2       FDType = 16
)

var FDType_name = map[int32]string{
	0:  "FD_UNKNOWN",
	1:  "FD_FILE",
	2:  "FD_DIRECTORY",
	3:  "FD_IPV4_SOCK",
	4:  "FD_IPV6_SOCK",
	5:  "FD_IPV4_SERVSOCK",
	6:  "FD_IPV6_SERVSOCK",
	7:  "FD_FIFO",
	8:  "FD_UNIX_SOCK",
	9:  "FD_EVENT",
	10: "FD_UNSUPPORTED",
	11: "FD_SIGNALFD",
	12: "FD_EVENTPOLL",
	13: "FD_INOTIFY",
	14: "FD_TIMERFD",
	15: "FD_NETLINK",
	16: "FD_FILE_V2",
}

var FDType_value = map[string]int32{
	"FD_UNKNOWN":       0,
	"FD_FILE":          1,
	"FD_DIRECTORY":     2,
	"FD_IPV4_SOCK":     3,
	"FD_IPV6_SOCK":     4,
	"FD_IPV4_SERVSOCK": 5,
	"FD_IPV6_SERVSOCK": 6,
	"FD_FIFO":          7,
	"FD_UNIX_SOCK":     8,
	"FD_EVENT":         9,
	"FD_UNSUPPORTED":   10,
	"FD_SIGNALFD":      11,
	"FD_EVENTPOLL":     12,
	"FD_INOTIFY":       13,
	"FD_TIMERFD":       14,
	"FD_NETLINK":       15,
	"FD_FILE_V2":       16,
}

func (x FDType) String() string {
	return proto.EnumName(FDType_name, int32(x))
}

func (FDType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_81bb5d1665ce2a0c, []int{3}
}

type L4Proto int32

const (
	L4Proto_UNKNOWN L4Proto = 0
	L4Proto_TCP     L4Proto = 1
	L4Proto_UDP     L4Proto = 2
	L4Proto_ICMP    L4Proto = 3
	L4Proto_RAW     L4Proto = 4
)

var L4Proto_name = map[int32]string{
	0: "UNKNOWN",
	1: "TCP",
	2: "UDP",
	3: "ICMP",
	4: "RAW",
}

var L4Proto_value = map[string]int32{
	"UNKNOWN": 0,
	"TCP":     1,
	"UDP":     2,
	"ICMP":    3,
	"RAW":     4,
}

func (x L4Proto) String() string {
	return proto.EnumName(L4Proto_name, int32(x))
}

func (L4Proto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_81bb5d1665ce2a0c, []int{4}
}

type KindlingEventList struct {
	KindlingEventList    []*KindlingEvent `protobuf:"bytes,1,rep,name=kindling_event_list,json=kindlingEventList,proto3" json:"kindling_event_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *KindlingEventList) Reset()         { *m = KindlingEventList{} }
func (m *KindlingEventList) String() string { return proto.CompactTextString(m) }
func (*KindlingEventList) ProtoMessage()    {}
func (*KindlingEventList) Descriptor() ([]byte, []int) {
	return fileDescriptor_81bb5d1665ce2a0c, []int{0}
}
func (m *KindlingEventList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KindlingEventList.Unmarshal(m, b)
}
func (m *KindlingEventList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KindlingEventList.Marshal(b, m, deterministic)
}
func (m *KindlingEventList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KindlingEventList.Merge(m, src)
}
func (m *KindlingEventList) XXX_Size() int {
	return xxx_messageInfo_KindlingEventList.Size(m)
}
func (m *KindlingEventList) XXX_DiscardUnknown() {
	xxx_messageInfo_KindlingEventList.DiscardUnknown(m)
}

var xxx_messageInfo_KindlingEventList proto.InternalMessageInfo

func (m *KindlingEventList) GetKindlingEventList() []*KindlingEvent {
	if m != nil {
		return m.KindlingEventList
	}
	return nil
}

type KindlingEvent struct {
	Source Source `protobuf:"varint,1,opt,name=source,proto3,enum=kindling.Source" json:"source,omitempty"`
	// Timestamp in nanoseconds at which the event were collected.
	Timestamp uint64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Name of Kindling Event
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Category of Kindling Event, enum
	Category Category `protobuf:"varint,4,opt,name=category,proto3,enum=kindling.Category" json:"category,omitempty"`
	// Native attributes of hook point, including arguments or return value.
	NativeAttributes *Property `protobuf:"bytes,5,opt,name=Native_attributes,json=NativeAttributes,proto3" json:"Native_attributes,omitempty"`
	// User-defined Attributions of Kindling Event, now including latency for syscall.
	UserAttributes []*KeyValue `protobuf:"bytes,6,rep,name=user_attributes,json=userAttributes,proto3" json:"user_attributes,omitempty"`
	// Context includes Thread information and Fd information.
	Ctx                  *Context `protobuf:"bytes,7,opt,name=ctx,proto3" json:"ctx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KindlingEvent) Reset()         { *m = KindlingEvent{} }
func (m *KindlingEvent) String() string { return proto.CompactTextString(m) }
func (*KindlingEvent) ProtoMessage()    {}
func (*KindlingEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_81bb5d1665ce2a0c, []int{1}
}
func (m *KindlingEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KindlingEvent.Unmarshal(m, b)
}
func (m *KindlingEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KindlingEvent.Marshal(b, m, deterministic)
}
func (m *KindlingEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KindlingEvent.Merge(m, src)
}
func (m *KindlingEvent) XXX_Size() int {
	return xxx_messageInfo_KindlingEvent.Size(m)
}
func (m *KindlingEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_KindlingEvent.DiscardUnknown(m)
}

var xxx_messageInfo_KindlingEvent proto.InternalMessageInfo

func (m *KindlingEvent) GetSource() Source {
	if m != nil {
		return m.Source
	}
	return Source_SOURCE_UNKNOWN
}

func (m *KindlingEvent) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *KindlingEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KindlingEvent) GetCategory() Category {
	if m != nil {
		return m.Category
	}
	return Category_CAT_NONE
}

func (m *KindlingEvent) GetNativeAttributes() *Property {
	if m != nil {
		return m.NativeAttributes
	}
	return nil
}

func (m *KindlingEvent) GetUserAttributes() []*KeyValue {
	if m != nil {
		return m.UserAttributes
	}
	return nil
}

func (m *KindlingEvent) GetCtx() *Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

type AnyValue struct {
	// Types that are valid to be assigned to Value:
	//	*AnyValue_StringValue
	//	*AnyValue_BoolValue
	//	*AnyValue_IntValue
	//	*AnyValue_UintValue
	//	*AnyValue_DoubleValue
	//	*AnyValue_BytesValue
	//	*AnyValue_ArrayValue
	Value                isAnyValue_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AnyValue) Reset()         { *m = AnyValue{} }
func (m *AnyValue) String() string { return proto.CompactTextString(m) }
func (*AnyValue) ProtoMessage()    {}
func (*AnyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_81bb5d1665ce2a0c, []int{2}
}
func (m *AnyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnyValue.Unmarshal(m, b)
}
func (m *AnyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnyValue.Marshal(b, m, deterministic)
}
func (m *AnyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnyValue.Merge(m, src)
}
func (m *AnyValue) XXX_Size() int {
	return xxx_messageInfo_AnyValue.Size(m)
}
func (m *AnyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_AnyValue.DiscardUnknown(m)
}

var xxx_messageInfo_AnyValue proto.InternalMessageInfo

type isAnyValue_Value interface {
	isAnyValue_Value()
}

type AnyValue_StringValue struct {
	StringValue string `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3,oneof" json:"string_value,omitempty"`
}
type AnyValue_BoolValue struct {
	BoolValue bool `protobuf:"varint,2,opt,name=bool_value,json=boolValue,proto3,oneof" json:"bool_value,omitempty"`
}
type AnyValue_IntValue struct {
	IntValue int64 `protobuf:"varint,3,opt,name=int_value,json=intValue,proto3,oneof" json:"int_value,omitempty"`
}
type AnyValue_UintValue struct {
	UintValue uint64 `protobuf:"varint,4,opt,name=uint_value,json=uintValue,proto3,oneof" json:"uint_value,omitempty"`
}
type AnyValue_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,5,opt,name=double_value,json=doubleValue,proto3,oneof" json:"double_value,omitempty"`
}
type AnyValue_BytesValue struct {
	BytesValue []byte `protobuf:"bytes,6,opt,name=bytes_value,json=bytesValue,proto3,oneof" json:"bytes_value,omitempty"`
}
type AnyValue_ArrayValue struct {
	ArrayValue *ArrayValue `protobuf:"bytes,7,opt,name=array_value,json=arrayValue,proto3,oneof" json:"array_value,omitempty"`
}

func (*AnyValue_StringValue) isAnyValue_Value() {}
func (*AnyValue_BoolValue) isAnyValue_Value()   {}
func (*AnyValue_IntValue) isAnyValue_Value()    {}
func (*AnyValue_UintValue) isAnyValue_Value()   {}
func (*AnyValue_DoubleValue) isAnyValue_Value() {}
func (*AnyValue_BytesValue) isAnyValue_Value()  {}
func (*AnyValue_ArrayValue) isAnyValue_Value()  {}

func (m *AnyValue) GetValue() isAnyValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *AnyValue) GetStringValue() string {
	if x, ok := m.GetValue().(*AnyValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *AnyValue) GetBoolValue() bool {
	if x, ok := m.GetValue().(*AnyValue_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *AnyValue) GetIntValue() int64 {
	if x, ok := m.GetValue().(*AnyValue_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *AnyValue) GetUintValue() uint64 {
	if x, ok := m.GetValue().(*AnyValue_UintValue); ok {
		return x.UintValue
	}
	return 0
}

func (m *AnyValue) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*AnyValue_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *AnyValue) GetBytesValue() []byte {
	if x, ok := m.GetValue().(*AnyValue_BytesValue); ok {
		return x.BytesValue
	}
	return nil
}

func (m *AnyValue) GetArrayValue() *ArrayValue {
	if x, ok := m.GetValue().(*AnyValue_ArrayValue); ok {
		return x.ArrayValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AnyValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AnyValue_StringValue)(nil),
		(*AnyValue_BoolValue)(nil),
		(*AnyValue_IntValue)(nil),
		(*AnyValue_UintValue)(nil),
		(*AnyValue_DoubleValue)(nil),
		(*AnyValue_BytesValue)(nil),
		(*AnyValue_ArrayValue)(nil),
	}
}

type ArrayValue struct {
	Values               []*AnyValue `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ArrayValue) Reset()         { *m = ArrayValue{} }
func (m *ArrayValue) String() string { return proto.CompactTextString(m) }
func (*ArrayValue) ProtoMessage()    {}
func (*ArrayValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_81bb5d1665ce2a0c, []int{3}
}
func (m *ArrayValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrayValue.Unmarshal(m, b)
}
func (m *ArrayValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrayValue.Marshal(b, m, deterministic)
}
func (m *ArrayValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrayValue.Merge(m, src)
}
func (m *ArrayValue) XXX_Size() int {
	return xxx_messageInfo_ArrayValue.Size(m)
}
func (m *ArrayValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrayValue.DiscardUnknown(m)
}

var xxx_messageInfo_ArrayValue proto.InternalMessageInfo

func (m *ArrayValue) GetValues() []*AnyValue {
	if m != nil {
		return m.Values
	}
	return nil
}

type KeyValue struct {
	Key                  string    `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                *AnyValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}
func (*KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_81bb5d1665ce2a0c, []int{4}
}
func (m *KeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValue.Unmarshal(m, b)
}
func (m *KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValue.Marshal(b, m, deterministic)
}
func (m *KeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValue.Merge(m, src)
}
func (m *KeyValue) XXX_Size() int {
	return xxx_messageInfo_KeyValue.Size(m)
}
func (m *KeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValue proto.InternalMessageInfo

func (m *KeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValue) GetValue() *AnyValue {
	if m != nil {
		return m.Value
	}
	return nil
}

type Property struct {
	// If type of syscall_enter, kprobe, uprobe, tracepoint
	Args []*KeyValue `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
	// If type of syscall_exit, kretprobe, uretprobe
	Ret                  []*KeyValue `protobuf:"bytes,2,rep,name=ret,proto3" json:"ret,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Property) Reset()         { *m = Property{} }
func (m *Property) String() string { return proto.CompactTextString(m) }
func (*Property) ProtoMessage()    {}
func (*Property) Descriptor() ([]byte, []int) {
	return fileDescriptor_81bb5d1665ce2a0c, []int{5}
}
func (m *Property) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Property.Unmarshal(m, b)
}
func (m *Property) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Property.Marshal(b, m, deterministic)
}
func (m *Property) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Property.Merge(m, src)
}
func (m *Property) XXX_Size() int {
	return xxx_messageInfo_Property.Size(m)
}
func (m *Property) XXX_DiscardUnknown() {
	xxx_messageInfo_Property.DiscardUnknown(m)
}

var xxx_messageInfo_Property proto.InternalMessageInfo

func (m *Property) GetArgs() []*KeyValue {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Property) GetRet() []*KeyValue {
	if m != nil {
		return m.Ret
	}
	return nil
}

type Pair struct {
	// Arguments' Name or Attributions' Name.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Type of Value.
	ValueType ValueType `protobuf:"varint,2,opt,name=value_type,json=valueType,proto3,enum=kindling.ValueType" json:"value_type,omitempty"`
	// Value of Key in bytes, should be converted according to ValueType.
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_81bb5d1665ce2a0c, []int{6}
}
func (m *Pair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pair.Unmarshal(m, b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
}
func (m *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(m, src)
}
func (m *Pair) XXX_Size() int {
	return xxx_messageInfo_Pair.Size(m)
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Pair) GetValueType() ValueType {
	if m != nil {
		return m.ValueType
	}
	return ValueType_NONE
}

func (m *Pair) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Context struct {
	// Thread information corresponding to Kindling Event, optional.
	ThreadInfo *Thread `protobuf:"bytes,1,opt,name=thread_info,json=threadInfo,proto3" json:"thread_info,omitempty"`
	// Fd information corresponding to Kindling Event, optional.
	FdInfo               *Fd      `protobuf:"bytes,2,opt,name=fd_info,json=fdInfo,proto3" json:"fd_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Context) Reset()         { *m = Context{} }
func (m *Context) String() string { return proto.CompactTextString(m) }
func (*Context) ProtoMessage()    {}
func (*Context) Descriptor() ([]byte, []int) {
	return fileDescriptor_81bb5d1665ce2a0c, []int{7}
}
func (m *Context) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Context.Unmarshal(m, b)
}
func (m *Context) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Context.Marshal(b, m, deterministic)
}
func (m *Context) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Context.Merge(m, src)
}
func (m *Context) XXX_Size() int {
	return xxx_messageInfo_Context.Size(m)
}
func (m *Context) XXX_DiscardUnknown() {
	xxx_messageInfo_Context.DiscardUnknown(m)
}

var xxx_messageInfo_Context proto.InternalMessageInfo

func (m *Context) GetThreadInfo() *Thread {
	if m != nil {
		return m.ThreadInfo
	}
	return nil
}

func (m *Context) GetFdInfo() *Fd {
	if m != nil {
		return m.FdInfo
	}
	return nil
}

type Thread struct {
	// Process id of thread.
	Pid uint32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	// Thread/task id of thread.
	Tid uint32 `protobuf:"varint,2,opt,name=tid,proto3" json:"tid,omitempty"`
	// User id of thread
	Uid uint32 `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	// Group id of thread
	Gid uint32 `protobuf:"varint,4,opt,name=gid,proto3" json:"gid,omitempty"`
	// Command of thread.
	Comm string `protobuf:"bytes,5,opt,name=comm,proto3" json:"comm,omitempty"`
	// ContainerId of thread
	ContainerId string `protobuf:"bytes,6,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	// ContainerName of thread
	ContainerName        string   `protobuf:"bytes,7,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Thread) Reset()         { *m = Thread{} }
func (m *Thread) String() string { return proto.CompactTextString(m) }
func (*Thread) ProtoMessage()    {}
func (*Thread) Descriptor() ([]byte, []int) {
	return fileDescriptor_81bb5d1665ce2a0c, []int{8}
}
func (m *Thread) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Thread.Unmarshal(m, b)
}
func (m *Thread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Thread.Marshal(b, m, deterministic)
}
func (m *Thread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Thread.Merge(m, src)
}
func (m *Thread) XXX_Size() int {
	return xxx_messageInfo_Thread.Size(m)
}
func (m *Thread) XXX_DiscardUnknown() {
	xxx_messageInfo_Thread.DiscardUnknown(m)
}

var xxx_messageInfo_Thread proto.InternalMessageInfo

func (m *Thread) GetPid() uint32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *Thread) GetTid() uint32 {
	if m != nil {
		return m.Tid
	}
	return 0
}

func (m *Thread) GetUid() uint32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Thread) GetGid() uint32 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func (m *Thread) GetComm() string {
	if m != nil {
		return m.Comm
	}
	return ""
}

func (m *Thread) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *Thread) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

type Fd struct {
	// FD number.
	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	// Type of FD in enum.
	TypeFd FDType `protobuf:"varint,2,opt,name=type_fd,json=typeFd,proto3,enum=kindling.FDType" json:"type_fd,omitempty"`
	// if FD is type of file
	Filename  string `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	Directory string `protobuf:"bytes,4,opt,name=directory,proto3" json:"directory,omitempty"`
	// if FD is type of ipv4 or ipv6
	Protocol L4Proto `protobuf:"varint,5,opt,name=protocol,proto3,enum=kindling.L4Proto" json:"protocol,omitempty"`
	// repeated for ipv6, client_ip[0] for ipv4
	Role  bool     `protobuf:"varint,6,opt,name=role,proto3" json:"role,omitempty"`
	Sip   []uint32 `protobuf:"varint,7,rep,packed,name=sip,proto3" json:"sip,omitempty"`
	Dip   []uint32 `protobuf:"varint,8,rep,packed,name=dip,proto3" json:"dip,omitempty"`
	Sport uint32   `protobuf:"varint,9,opt,name=sport,proto3" json:"sport,omitempty"`
	Dport uint32   `protobuf:"varint,10,opt,name=dport,proto3" json:"dport,omitempty"`
	// if FD is type of unix_sock
	// Source socket endpoint
	Source uint64 `protobuf:"varint,11,opt,name=source,proto3" json:"source,omitempty"`
	// Destination socket endpoint
	Destination          uint64   `protobuf:"varint,12,opt,name=destination,proto3" json:"destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Fd) Reset()         { *m = Fd{} }
func (m *Fd) String() string { return proto.CompactTextString(m) }
func (*Fd) ProtoMessage()    {}
func (*Fd) Descriptor() ([]byte, []int) {
	return fileDescriptor_81bb5d1665ce2a0c, []int{9}
}
func (m *Fd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fd.Unmarshal(m, b)
}
func (m *Fd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fd.Marshal(b, m, deterministic)
}
func (m *Fd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fd.Merge(m, src)
}
func (m *Fd) XXX_Size() int {
	return xxx_messageInfo_Fd.Size(m)
}
func (m *Fd) XXX_DiscardUnknown() {
	xxx_messageInfo_Fd.DiscardUnknown(m)
}

var xxx_messageInfo_Fd proto.InternalMessageInfo

func (m *Fd) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *Fd) GetTypeFd() FDType {
	if m != nil {
		return m.TypeFd
	}
	return FDType_FD_UNKNOWN
}

func (m *Fd) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Fd) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

func (m *Fd) GetProtocol() L4Proto {
	if m != nil {
		return m.Protocol
	}
	return L4Proto_UNKNOWN
}

func (m *Fd) GetRole() bool {
	if m != nil {
		return m.Role
	}
	return false
}

func (m *Fd) GetSip() []uint32 {
	if m != nil {
		return m.Sip
	}
	return nil
}

func (m *Fd) GetDip() []uint32 {
	if m != nil {
		return m.Dip
	}
	return nil
}

func (m *Fd) GetSport() uint32 {
	if m != nil {
		return m.Sport
	}
	return 0
}

func (m *Fd) GetDport() uint32 {
	if m != nil {
		return m.Dport
	}
	return 0
}

func (m *Fd) GetSource() uint64 {
	if m != nil {
		return m.Source
	}
	return 0
}

func (m *Fd) GetDestination() uint64 {
	if m != nil {
		return m.Destination
	}
	return 0
}

func init() {
	proto.RegisterEnum("kindling.Source", Source_name, Source_value)
	proto.RegisterEnum("kindling.Category", Category_name, Category_value)
	proto.RegisterEnum("kindling.ValueType", ValueType_name, ValueType_value)
	proto.RegisterEnum("kindling.FDType", FDType_name, FDType_value)
	proto.RegisterEnum("kindling.L4Proto", L4Proto_name, L4Proto_value)
	proto.RegisterType((*KindlingEventList)(nil), "kindling.KindlingEventList")
	proto.RegisterType((*KindlingEvent)(nil), "kindling.KindlingEvent")
	proto.RegisterType((*AnyValue)(nil), "kindling.AnyValue")
	proto.RegisterType((*ArrayValue)(nil), "kindling.ArrayValue")
	proto.RegisterType((*KeyValue)(nil), "kindling.KeyValue")
	proto.RegisterType((*Property)(nil), "kindling.Property")
	proto.RegisterType((*Pair)(nil), "kindling.Pair")
	proto.RegisterType((*Context)(nil), "kindling.Context")
	proto.RegisterType((*Thread)(nil), "kindling.Thread")
	proto.RegisterType((*Fd)(nil), "kindling.Fd")
}

func init() { proto.RegisterFile("kindling_event.proto", fileDescriptor_81bb5d1665ce2a0c) }

var fileDescriptor_81bb5d1665ce2a0c = []byte{
	// 1327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0x61, 0x6e, 0xdb, 0xc6,
	0x12, 0x36, 0x25, 0x59, 0x24, 0x87, 0x92, 0xb2, 0xde, 0x18, 0xef, 0x09, 0x0f, 0xef, 0xe1, 0x29,
	0x4a, 0x53, 0xa8, 0x06, 0xea, 0x22, 0x8e, 0xe1, 0x06, 0xc8, 0x8f, 0x42, 0x96, 0xc8, 0x98, 0x90,
	0x4c, 0x12, 0x2b, 0xca, 0x89, 0x8b, 0x02, 0x02, 0x2d, 0xd2, 0x2e, 0x11, 0x89, 0x14, 0x28, 0xca,
	0x88, 0x4e, 0xd0, 0x3b, 0xf4, 0x08, 0x2d, 0xd0, 0x23, 0xf4, 0x77, 0x0f, 0xd1, 0xc3, 0x14, 0xb3,
	0xbb, 0xa4, 0x94, 0x22, 0xfd, 0x23, 0xcc, 0x7c, 0xf3, 0xcd, 0xec, 0xf0, 0xdb, 0xd9, 0x11, 0x1c,
	0x7f, 0x88, 0x93, 0x70, 0x11, 0x27, 0x0f, 0xb3, 0xe8, 0x31, 0x4a, 0xf2, 0xd3, 0x55, 0x96, 0xe6,
	0x29, 0xd5, 0x0a, 0xb4, 0xfb, 0x03, 0x1c, 0x8d, 0xa4, 0x6d, 0x22, 0x61, 0x1c, 0xaf, 0x73, 0xfa,
	0x16, 0x9e, 0x7e, 0x9a, 0x36, 0x5b, 0xc4, 0xeb, 0xbc, 0xad, 0x74, 0xaa, 0x3d, 0xe3, 0xec, 0xdf,
	0xa7, 0x45, 0xec, 0xf4, 0x93, 0x4c, 0x76, 0xf4, 0xe1, 0xef, 0x85, 0xba, 0xbf, 0x57, 0xa0, 0xf9,
	0x09, 0x89, 0xf6, 0xa0, 0xbe, 0x4e, 0x37, 0xd9, 0x3c, 0x6a, 0x2b, 0x1d, 0xa5, 0xd7, 0x3a, 0x23,
	0xbb, 0x6a, 0x13, 0x8e, 0x33, 0x19, 0xa7, 0xff, 0x05, 0x3d, 0x8f, 0x97, 0xd1, 0x3a, 0x0f, 0x96,
	0xab, 0x76, 0xa5, 0xa3, 0xf4, 0x6a, 0x6c, 0x07, 0x50, 0x0a, 0xb5, 0x24, 0x58, 0x46, 0xed, 0x6a,
	0x47, 0xe9, 0xe9, 0x8c, 0xdb, 0xf4, 0x14, 0xb4, 0x79, 0x90, 0x47, 0x0f, 0x69, 0xb6, 0x6d, 0xd7,
	0x78, 0x75, 0xba, 0xab, 0x3e, 0x90, 0x11, 0x56, 0x72, 0xe8, 0x77, 0x70, 0xe4, 0x04, 0x79, 0xfc,
	0x18, 0xcd, 0x82, 0x3c, 0xcf, 0xe2, 0xbb, 0x4d, 0x1e, 0xad, 0xdb, 0x87, 0x1d, 0xa5, 0x67, 0xec,
	0x27, 0x7a, 0x59, 0xba, 0x8a, 0xb2, 0x7c, 0xcb, 0x88, 0x20, 0xf7, 0x4b, 0x2e, 0x7d, 0x03, 0x4f,
	0x36, 0xeb, 0x28, 0xdb, 0x4f, 0xaf, 0x73, 0x8d, 0xf6, 0xd2, 0x47, 0xd1, 0xf6, 0x26, 0x58, 0x6c,
	0x22, 0xd6, 0x42, 0xea, 0x5e, 0xf2, 0x73, 0xa8, 0xce, 0xf3, 0x8f, 0x6d, 0x95, 0x9f, 0x77, 0xb4,
	0xd7, 0x68, 0x9a, 0xe4, 0xd1, 0xc7, 0x9c, 0x61, 0xb4, 0xfb, 0x73, 0x05, 0xb4, 0x7e, 0x22, 0x2a,
	0xd0, 0xe7, 0xd0, 0x58, 0xe7, 0x19, 0x5e, 0xca, 0x23, 0xfa, 0x5c, 0x41, 0xfd, 0xea, 0x80, 0x19,
	0x02, 0x15, 0xa4, 0xff, 0x03, 0xdc, 0xa5, 0xe9, 0x42, 0x52, 0x50, 0x37, 0xed, 0xea, 0x80, 0xe9,
	0x88, 0x09, 0xc2, 0xff, 0x40, 0x8f, 0x93, 0x5c, 0xc6, 0x51, 0xbe, 0xea, 0xd5, 0x01, 0xd3, 0xe2,
	0x24, 0x2f, 0xf3, 0x37, 0xbb, 0x38, 0xca, 0x58, 0xc3, 0xfc, 0x4d, 0x49, 0x78, 0x0e, 0x8d, 0x30,
	0xdd, 0xdc, 0x2d, 0x22, 0x49, 0x41, 0xc1, 0x14, 0xec, 0x42, 0xa0, 0x82, 0xf4, 0x0c, 0x8c, 0xbb,
	0x6d, 0x1e, 0xad, 0x25, 0xa7, 0xde, 0x51, 0x7a, 0x8d, 0xab, 0x03, 0x06, 0x1c, 0x14, 0x94, 0x6f,
	0xc1, 0x08, 0xb2, 0x2c, 0xd8, 0x4a, 0x8a, 0xd0, 0xe1, 0x78, 0xa7, 0x43, 0x1f, 0x83, 0x9c, 0x8a,
	0x89, 0x41, 0xe9, 0x5d, 0xaa, 0x70, 0xc8, 0x53, 0xba, 0xaf, 0x01, 0x76, 0x24, 0x7a, 0x02, 0x75,
	0x0e, 0xaf, 0xe5, 0x9c, 0xee, 0xdd, 0x41, 0xa1, 0x20, 0x93, 0x8c, 0xae, 0x05, 0x5a, 0x71, 0x2f,
	0x94, 0x40, 0xf5, 0x43, 0xb4, 0x15, 0x62, 0x32, 0x34, 0x69, 0x4f, 0x1e, 0xc0, 0xd5, 0xfb, 0x7c,
	0x21, 0xd9, 0xc1, 0x7b, 0xd0, 0x8a, 0xf1, 0xa0, 0x5f, 0x42, 0x2d, 0xc8, 0x1e, 0x3e, 0x73, 0x7a,
	0x39, 0x01, 0x3c, 0x4e, 0xbf, 0x80, 0x6a, 0x16, 0xe5, 0xed, 0xca, 0x3f, 0xd2, 0x30, 0xdc, 0xbd,
	0x83, 0x9a, 0x17, 0xc4, 0xd9, 0x67, 0xba, 0x3b, 0x03, 0xe0, 0x87, 0xcf, 0xf2, 0xed, 0x4a, 0xb4,
	0xd8, 0x3a, 0x7b, 0xba, 0x2b, 0xc3, 0x6b, 0xf8, 0xdb, 0x55, 0xc4, 0xf4, 0xc7, 0xc2, 0xa4, 0xc7,
	0xc5, 0x17, 0xe1, 0x7d, 0x37, 0x8a, 0xee, 0xe7, 0xa0, 0xca, 0x61, 0xa3, 0x2f, 0xc1, 0xc8, 0x7f,
	0xcc, 0xa2, 0x20, 0x9c, 0xc5, 0xc9, 0x7d, 0xca, 0x8f, 0x33, 0xf6, 0xdf, 0xa6, 0xcf, 0x83, 0x0c,
	0x04, 0xc9, 0x4e, 0xee, 0x53, 0xfa, 0x02, 0xd4, 0x7b, 0x49, 0x17, 0x3a, 0x35, 0x76, 0x74, 0x2b,
	0x64, 0xf5, 0x7b, 0x4e, 0xeb, 0xfe, 0xa6, 0x40, 0x5d, 0x64, 0xe3, 0xb7, 0xac, 0xe2, 0x90, 0x17,
	0x6f, 0x32, 0x34, 0x11, 0xc9, 0xe3, 0x90, 0xe7, 0x37, 0x19, 0x9a, 0x88, 0x6c, 0xe2, 0x90, 0xf7,
	0xd9, 0x64, 0x68, 0x22, 0xf2, 0x10, 0x87, 0x7c, 0x12, 0x9b, 0x0c, 0x4d, 0x7c, 0xfb, 0xf3, 0x74,
	0xb9, 0xe4, 0x93, 0xa7, 0x33, 0x6e, 0xd3, 0x67, 0xd0, 0x98, 0xa7, 0x49, 0x1e, 0xc4, 0x49, 0x94,
	0xcd, 0xe2, 0x90, 0x4f, 0x9c, 0xce, 0x8c, 0x12, 0xb3, 0x43, 0xfa, 0x02, 0x5a, 0x3b, 0x0a, 0x5f,
	0x1e, 0x2a, 0x27, 0x35, 0x4b, 0xd4, 0x09, 0x96, 0x51, 0xf7, 0x8f, 0x0a, 0x54, 0x2c, 0x7e, 0x6c,
	0xb2, 0x59, 0xf2, 0x66, 0x0f, 0x19, 0x9a, 0xf4, 0x2b, 0x50, 0x51, 0xf2, 0xd9, 0x7d, 0x28, 0x55,
	0xdf, 0xd3, 0xc7, 0x1a, 0x72, 0xc9, 0xeb, 0x48, 0xb0, 0x42, 0xfa, 0x1f, 0xd0, 0xee, 0xe3, 0x45,
	0xb4, 0xb7, 0xa1, 0x4a, 0x1f, 0xf7, 0x5a, 0x18, 0x67, 0xd1, 0x3c, 0x2f, 0xd6, 0x94, 0xce, 0x76,
	0x00, 0xfd, 0x1a, 0x34, 0xbe, 0xa2, 0xe7, 0xe9, 0x82, 0x7f, 0x5f, 0x6b, 0x7f, 0x35, 0x8c, 0xcf,
	0x3d, 0x8c, 0xb1, 0x92, 0x82, 0x52, 0x64, 0xe9, 0x42, 0x3c, 0x30, 0x8d, 0x71, 0x1b, 0x3b, 0x5f,
	0xc7, 0xab, 0xb6, 0xda, 0xa9, 0xa2, 0x60, 0xeb, 0x78, 0x85, 0x48, 0x18, 0xaf, 0xda, 0x9a, 0x40,
	0xc2, 0x78, 0x85, 0x03, 0xb1, 0x5e, 0xa5, 0x59, 0xde, 0xd6, 0xb9, 0xac, 0xc2, 0x41, 0x34, 0xe4,
	0x28, 0x08, 0x94, 0x3b, 0xf4, 0x5f, 0xe5, 0xca, 0x36, 0xf8, 0x16, 0x2e, 0x16, 0x74, 0x07, 0x8c,
	0x30, 0x5a, 0xe7, 0x71, 0x12, 0xe4, 0x71, 0x9a, 0xb4, 0x1b, 0x3c, 0xb8, 0x0f, 0x9d, 0xfc, 0xa4,
	0x40, 0x5d, 0x6c, 0x75, 0x4a, 0xa1, 0x35, 0x71, 0xa7, 0x6c, 0x60, 0xce, 0xa6, 0xce, 0xc8, 0x71,
	0xdf, 0x39, 0xe4, 0x80, 0x1e, 0x41, 0x73, 0x72, 0x3b, 0x19, 0xf4, 0xc7, 0xe3, 0x99, 0xe9, 0xf8,
	0x26, 0x23, 0x0a, 0x25, 0xd0, 0x28, 0xa1, 0xf7, 0xb6, 0x4f, 0x2a, 0xb4, 0x05, 0xe0, 0xb3, 0xfe,
	0xc0, 0xf4, 0x5c, 0xdb, 0xf1, 0x49, 0x95, 0x02, 0xd4, 0x47, 0xcc, 0x73, 0x2f, 0x4d, 0x52, 0xa3,
	0x4d, 0xd0, 0x47, 0xcc, 0xf4, 0x3d, 0x86, 0xee, 0x21, 0x86, 0xa6, 0xc2, 0xae, 0x63, 0x68, 0x5a,
	0x86, 0xd4, 0x93, 0x3f, 0x15, 0xd0, 0x8a, 0x7f, 0x00, 0xda, 0x00, 0x6d, 0xd0, 0xf7, 0x67, 0x8e,
	0xeb, 0x98, 0xe4, 0x00, 0x99, 0xe8, 0xb9, 0xfe, 0x15, 0xef, 0x40, 0x06, 0x2d, 0x7b, 0x6c, 0x92,
	0x0a, 0x35, 0x40, 0xe5, 0x54, 0x13, 0x8f, 0x96, 0x8e, 0xed, 0x0d, 0x48, 0xad, 0xe0, 0xbd, 0xeb,
	0xdb, 0x3e, 0x39, 0xc4, 0x2e, 0xd1, 0x9b, 0xd8, 0x6f, 0x9d, 0xfe, 0x58, 0x1c, 0xcf, 0xfd, 0xb1,
	0x69, 0x7a, 0x44, 0x2d, 0xc8, 0xbe, 0x7d, 0x6d, 0x12, 0x8d, 0x3e, 0x01, 0x03, 0x3d, 0x8f, 0xb9,
	0x03, 0x73, 0x32, 0x21, 0x3a, 0x0a, 0xc1, 0xd9, 0x83, 0x2b, 0x73, 0x38, 0x1d, 0x9b, 0x8c, 0x40,
	0x51, 0xf0, 0xda, 0xbc, 0x76, 0xd9, 0x2d, 0x31, 0x8a, 0x0a, 0xd3, 0x89, 0xc9, 0x48, 0xa3, 0x3c,
	0xee, 0x76, 0xe2, 0x9b, 0xd7, 0xa4, 0x79, 0xf2, 0x8b, 0x02, 0x7a, 0xf9, 0xf0, 0xa9, 0x06, 0x35,
	0xf9, 0x6d, 0x1a, 0xd4, 0x6c, 0xc7, 0x7f, 0x4d, 0x14, 0xaa, 0xc3, 0xa1, 0xed, 0xf8, 0x2f, 0x2f,
	0x48, 0x45, 0x9a, 0xaf, 0xce, 0x48, 0x55, 0x9a, 0x17, 0xe7, 0xa4, 0x86, 0xe6, 0x94, 0x73, 0x85,
	0x8e, 0x82, 0x5c, 0x2f, 0xec, 0x57, 0x67, 0x44, 0x2d, 0xec, 0x8b, 0x73, 0xa2, 0x71, 0x2d, 0xae,
	0xfa, 0xec, 0x72, 0x6a, 0x11, 0x1d, 0x9d, 0xcb, 0x5b, 0xdf, 0x44, 0x07, 0xb0, 0x90, 0x35, 0x76,
	0xfb, 0x3e, 0x31, 0x30, 0x61, 0xe8, 0x4e, 0x2f, 0xc7, 0x26, 0x69, 0x60, 0x2b, 0x97, 0xae, 0x3b,
	0x26, 0xcd, 0x93, 0x5f, 0x2b, 0x50, 0x17, 0xef, 0x05, 0xbf, 0xc3, 0x1a, 0xee, 0x4d, 0x84, 0x01,
	0xaa, 0x35, 0x14, 0xda, 0xf3, 0x59, 0xb0, 0x86, 0xb3, 0xa1, 0xcd, 0xcc, 0x81, 0x8f, 0x22, 0x54,
	0x24, 0x62, 0x7b, 0x37, 0xe7, 0xb3, 0x89, 0x3b, 0x18, 0x91, 0xea, 0x0e, 0xb9, 0x10, 0x48, 0x8d,
	0x1e, 0x03, 0x29, 0x39, 0x26, 0xbb, 0xe1, 0xe8, 0xe1, 0x0e, 0xbd, 0xd8, 0xa1, 0xf5, 0xf2, 0x38,
	0xcb, 0x25, 0xaa, 0x2c, 0x35, 0x75, 0xec, 0xf7, 0xa2, 0x94, 0x86, 0x9a, 0x5b, 0xc3, 0x99, 0x79,
	0x63, 0x3a, 0x3e, 0xd1, 0x71, 0x82, 0x79, 0x7c, 0x32, 0xf5, 0x3c, 0x97, 0xf9, 0xe6, 0x90, 0x00,
	0xde, 0xa4, 0x35, 0x94, 0xb7, 0x6e, 0x0d, 0x89, 0x21, 0x8b, 0xf0, 0x14, 0xcf, 0x1d, 0x8f, 0xc5,
	0x55, 0xe1, 0xc9, 0x8e, 0xeb, 0xdb, 0xd6, 0x2d, 0x69, 0x4a, 0x1f, 0x27, 0x81, 0x59, 0x43, 0xd2,
	0x92, 0xbe, 0x63, 0xfa, 0x63, 0xdb, 0x19, 0x91, 0x27, 0xd2, 0x47, 0x09, 0x66, 0x37, 0x67, 0x84,
	0x9c, 0xbc, 0x01, 0x55, 0x3e, 0x7b, 0x6c, 0x77, 0x27, 0x95, 0x0a, 0x55, 0x7f, 0xe0, 0x11, 0x05,
	0x8d, 0xe9, 0xd0, 0x23, 0x15, 0x7e, 0xd9, 0x83, 0x6b, 0x8f, 0x54, 0x11, 0x62, 0xfd, 0x77, 0xa4,
	0x76, 0xd9, 0xfc, 0xde, 0x58, 0xa6, 0x61, 0xb4, 0xf8, 0x86, 0xff, 0xde, 0xd5, 0xf9, 0xde, 0x78,
	0xf5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x22, 0xba, 0x9f, 0x15, 0x0a, 0x00, 0x00,
}
