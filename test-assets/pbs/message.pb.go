// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package pbs

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Message_Type int32

const (
	Message_CREATE_COLLECTION Message_Type = 0
	Message_CREATE_SCHEMA     Message_Type = 1
	Message_DELETE_COLLECTION Message_Type = 2
	Message_CREATE_REPLAY     Message_Type = 3
	Message_DELETE_REPLAY     Message_Type = 4
	Message_UPDATE_REPLAY     Message_Type = 5
	Message_PAUSE_REPLAY      Message_Type = 6
	Message_RESUME_REPLAY     Message_Type = 7
)

var Message_Type_name = map[int32]string{
	0: "CREATE_COLLECTION",
	1: "CREATE_SCHEMA",
	2: "DELETE_COLLECTION",
	3: "CREATE_REPLAY",
	4: "DELETE_REPLAY",
	5: "UPDATE_REPLAY",
	6: "PAUSE_REPLAY",
	7: "RESUME_REPLAY",
}

var Message_Type_value = map[string]int32{
	"CREATE_COLLECTION": 0,
	"CREATE_SCHEMA":     1,
	"DELETE_COLLECTION": 2,
	"CREATE_REPLAY":     3,
	"DELETE_REPLAY":     4,
	"UPDATE_REPLAY":     5,
	"PAUSE_REPLAY":      6,
	"RESUME_REPLAY":     7,
}

func (x Message_Type) String() string {
	return proto.EnumName(Message_Type_name, int32(x))
}

func (Message_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0}
}

type Message struct {
	// What kind of a message is this?
	Type Message_Type `protobuf:"varint,1,opt,name=type,proto3,enum=events.Message_Type" json:"type,omitempty"`
	// Who does this message pertain to?
	Team *Team `protobuf:"bytes,2,opt,name=team,proto3" json:"team,omitempty"`
	// Contains event debug info
	Info *Info `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*Message_Collect
	//	*Message_Replay
	//	*Message_Schema
	Event                isMessage_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetType() Message_Type {
	if m != nil {
		return m.Type
	}
	return Message_CREATE_COLLECTION
}

func (m *Message) GetTeam() *Team {
	if m != nil {
		return m.Team
	}
	return nil
}

func (m *Message) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

type isMessage_Event interface {
	isMessage_Event()
}

type Message_Collect struct {
	Collect *Collect `protobuf:"bytes,100,opt,name=collect,proto3,oneof"`
}

type Message_Replay struct {
	Replay *Replay `protobuf:"bytes,101,opt,name=replay,proto3,oneof"`
}

type Message_Schema struct {
	Schema *Schema `protobuf:"bytes,102,opt,name=schema,proto3,oneof"`
}

func (*Message_Collect) isMessage_Event() {}

func (*Message_Replay) isMessage_Event() {}

func (*Message_Schema) isMessage_Event() {}

func (m *Message) GetEvent() isMessage_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Message) GetCollect() *Collect {
	if x, ok := m.GetEvent().(*Message_Collect); ok {
		return x.Collect
	}
	return nil
}

func (m *Message) GetReplay() *Replay {
	if x, ok := m.GetEvent().(*Message_Replay); ok {
		return x.Replay
	}
	return nil
}

func (m *Message) GetSchema() *Schema {
	if x, ok := m.GetEvent().(*Message_Schema); ok {
		return x.Schema
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message_Collect)(nil),
		(*Message_Replay)(nil),
		(*Message_Schema)(nil),
	}
}

func init() {
	proto.RegisterEnum("events.Message_Type", Message_Type_name, Message_Type_value)
	proto.RegisterType((*Message)(nil), "events.Message")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0xd1, 0x5d, 0xab, 0xa2, 0x40,
	0x1c, 0x06, 0xf0, 0x2c, 0x53, 0x98, 0xad, 0xd6, 0x86, 0x5d, 0x90, 0xae, 0xa2, 0x2b, 0x61, 0x77,
	0x15, 0xda, 0x4f, 0x60, 0x36, 0x50, 0x60, 0x5b, 0x8c, 0x76, 0xb1, 0xe7, 0xe6, 0xa0, 0x9e, 0x7f,
	0x2f, 0xa0, 0x29, 0xce, 0x74, 0xa0, 0x0f, 0x74, 0x2e, 0xce, 0xb7, 0x3c, 0xcc, 0x8c, 0x46, 0xdd,
	0xe9, 0xf3, 0xfc, 0x98, 0xf1, 0x41, 0x34, 0x2c, 0x80, 0xb1, 0xe4, 0x08, 0x6e, 0x55, 0x97, 0xbc,
	0xc4, 0x06, 0xbc, 0xc3, 0x85, 0xb3, 0xc9, 0x30, 0x2b, 0xf3, 0x1c, 0x32, 0xae, 0xe2, 0xc9, 0xa0,
	0x86, 0x2a, 0x4f, 0x6e, 0xcd, 0x1b, 0x3a, 0x5f, 0x0e, 0x65, 0xdb, 0xb0, 0xec, 0x04, 0x45, 0xd2,
	0x36, 0x1c, 0x92, 0x42, 0x3d, 0xcf, 0x3e, 0x7b, 0xc8, 0xdc, 0xa8, 0xc3, 0xb1, 0x83, 0x74, 0x7e,
	0xab, 0xc0, 0xd6, 0xa6, 0x9a, 0x33, 0x9a, 0xff, 0x70, 0xd5, 0x2d, 0x6e, 0x53, 0xbb, 0xf1, 0xad,
	0x02, 0x2a, 0x05, 0x9e, 0x22, 0x5d, 0x9c, 0x61, 0x77, 0xa7, 0x9a, 0xf3, 0x6d, 0x3e, 0x68, 0x65,
	0x0c, 0x49, 0x41, 0x65, 0x23, 0x84, 0xb8, 0xdf, 0xee, 0x3d, 0x8b, 0xf5, 0xe5, 0x50, 0x52, 0xd9,
	0xe0, 0x5f, 0xc8, 0x6c, 0x3e, 0xdf, 0x7e, 0x93, 0xe8, 0x7b, 0x8b, 0x02, 0x15, 0xaf, 0x3a, 0xb4,
	0x15, 0xd8, 0x41, 0x86, 0x1a, 0x67, 0x83, 0xb4, 0xa3, 0xd6, 0x52, 0x99, 0xae, 0x3a, 0xb4, 0xe9,
	0x85, 0x54, 0x63, 0xed, 0xc3, 0xb3, 0x8c, 0x64, 0x2a, 0xa4, 0xea, 0x67, 0x1f, 0x1a, 0xd2, 0xc5,
	0x26, 0xfc, 0x13, 0x8d, 0x03, 0x4a, 0xfc, 0x98, 0xbc, 0x06, 0xdb, 0x30, 0x24, 0x41, 0xbc, 0xde,
	0xfe, 0xb3, 0x3a, 0x78, 0x8c, 0x86, 0x4d, 0x1c, 0x05, 0x2b, 0xb2, 0xf1, 0x2d, 0x4d, 0xc8, 0x25,
	0x09, 0xc9, 0xb3, 0xec, 0x3e, 0x48, 0x4a, 0x76, 0xa1, 0xff, 0xdf, 0xea, 0x89, 0xa8, 0x91, 0x4d,
	0xa4, 0x8b, 0x68, 0xbf, 0x5b, 0x3e, 0xa8, 0x3e, 0xb6, 0xd0, 0x60, 0xe7, 0xef, 0xa3, 0x7b, 0x62,
	0x08, 0x44, 0x49, 0xb4, 0xdf, 0xdc, 0x23, 0x73, 0x61, 0xa2, 0xbe, 0x9c, 0xb0, 0x70, 0x5f, 0x7e,
	0x1f, 0xcf, 0xfc, 0x74, 0x4d, 0xdd, 0xac, 0x2c, 0xbc, 0x34, 0xe1, 0xd9, 0x29, 0x2b, 0xeb, 0xca,
	0xab, 0xf2, 0x6b, 0x91, 0x42, 0xed, 0x71, 0x60, 0xfc, 0x4f, 0xc2, 0x18, 0x70, 0xe6, 0x55, 0x29,
	0x4b, 0x0d, 0xf9, 0x8b, 0xff, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x75, 0x3f, 0x24, 0x3e,
	0x02, 0x00, 0x00,
}
