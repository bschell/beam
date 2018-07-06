// Code generated by protoc-gen-go. DO NOT EDIT.
// source: standard_window_fns.proto

package pipeline_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GlobalWindowsPayload_Enum int32

const (
	// TODO(BEAM-3595): Change this to beam:windowfn:global_windows:v1
	GlobalWindowsPayload_PROPERTIES GlobalWindowsPayload_Enum = 0
)

var GlobalWindowsPayload_Enum_name = map[int32]string{
	0: "PROPERTIES",
}
var GlobalWindowsPayload_Enum_value = map[string]int32{
	"PROPERTIES": 0,
}

func (x GlobalWindowsPayload_Enum) String() string {
	return proto.EnumName(GlobalWindowsPayload_Enum_name, int32(x))
}
func (GlobalWindowsPayload_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_standard_window_fns_da1ada993b818d10, []int{0, 0}
}

type FixedWindowsPayload_Enum int32

const (
	// TODO(BEAM-3595): Change this to beam:windowfn:fixed_windows:v1
	FixedWindowsPayload_PROPERTIES FixedWindowsPayload_Enum = 0
)

var FixedWindowsPayload_Enum_name = map[int32]string{
	0: "PROPERTIES",
}
var FixedWindowsPayload_Enum_value = map[string]int32{
	"PROPERTIES": 0,
}

func (x FixedWindowsPayload_Enum) String() string {
	return proto.EnumName(FixedWindowsPayload_Enum_name, int32(x))
}
func (FixedWindowsPayload_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_standard_window_fns_da1ada993b818d10, []int{1, 0}
}

type SlidingWindowsPayload_Enum int32

const (
	// TODO(BEAM-3595): Change this to beam:windowfn:sliding_windows:v1
	SlidingWindowsPayload_PROPERTIES SlidingWindowsPayload_Enum = 0
)

var SlidingWindowsPayload_Enum_name = map[int32]string{
	0: "PROPERTIES",
}
var SlidingWindowsPayload_Enum_value = map[string]int32{
	"PROPERTIES": 0,
}

func (x SlidingWindowsPayload_Enum) String() string {
	return proto.EnumName(SlidingWindowsPayload_Enum_name, int32(x))
}
func (SlidingWindowsPayload_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_standard_window_fns_da1ada993b818d10, []int{2, 0}
}

type SessionsPayload_Enum int32

const (
	// TODO(BEAM-3595): Change this to beam:windowfn:session_windows:v1
	SessionsPayload_PROPERTIES SessionsPayload_Enum = 0
)

var SessionsPayload_Enum_name = map[int32]string{
	0: "PROPERTIES",
}
var SessionsPayload_Enum_value = map[string]int32{
	"PROPERTIES": 0,
}

func (x SessionsPayload_Enum) String() string {
	return proto.EnumName(SessionsPayload_Enum_name, int32(x))
}
func (SessionsPayload_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_standard_window_fns_da1ada993b818d10, []int{3, 0}
}

type GlobalWindowsPayload struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GlobalWindowsPayload) Reset()         { *m = GlobalWindowsPayload{} }
func (m *GlobalWindowsPayload) String() string { return proto.CompactTextString(m) }
func (*GlobalWindowsPayload) ProtoMessage()    {}
func (*GlobalWindowsPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_standard_window_fns_da1ada993b818d10, []int{0}
}
func (m *GlobalWindowsPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlobalWindowsPayload.Unmarshal(m, b)
}
func (m *GlobalWindowsPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlobalWindowsPayload.Marshal(b, m, deterministic)
}
func (dst *GlobalWindowsPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalWindowsPayload.Merge(dst, src)
}
func (m *GlobalWindowsPayload) XXX_Size() int {
	return xxx_messageInfo_GlobalWindowsPayload.Size(m)
}
func (m *GlobalWindowsPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalWindowsPayload.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalWindowsPayload proto.InternalMessageInfo

type FixedWindowsPayload struct {
	Size                 *duration.Duration   `protobuf:"bytes,1,opt,name=size,proto3" json:"size,omitempty"`
	Offset               *timestamp.Timestamp `protobuf:"bytes,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FixedWindowsPayload) Reset()         { *m = FixedWindowsPayload{} }
func (m *FixedWindowsPayload) String() string { return proto.CompactTextString(m) }
func (*FixedWindowsPayload) ProtoMessage()    {}
func (*FixedWindowsPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_standard_window_fns_da1ada993b818d10, []int{1}
}
func (m *FixedWindowsPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FixedWindowsPayload.Unmarshal(m, b)
}
func (m *FixedWindowsPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FixedWindowsPayload.Marshal(b, m, deterministic)
}
func (dst *FixedWindowsPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FixedWindowsPayload.Merge(dst, src)
}
func (m *FixedWindowsPayload) XXX_Size() int {
	return xxx_messageInfo_FixedWindowsPayload.Size(m)
}
func (m *FixedWindowsPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_FixedWindowsPayload.DiscardUnknown(m)
}

var xxx_messageInfo_FixedWindowsPayload proto.InternalMessageInfo

func (m *FixedWindowsPayload) GetSize() *duration.Duration {
	if m != nil {
		return m.Size
	}
	return nil
}

func (m *FixedWindowsPayload) GetOffset() *timestamp.Timestamp {
	if m != nil {
		return m.Offset
	}
	return nil
}

type SlidingWindowsPayload struct {
	Size                 *duration.Duration   `protobuf:"bytes,1,opt,name=size,proto3" json:"size,omitempty"`
	Offset               *timestamp.Timestamp `protobuf:"bytes,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Period               *duration.Duration   `protobuf:"bytes,3,opt,name=period,proto3" json:"period,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SlidingWindowsPayload) Reset()         { *m = SlidingWindowsPayload{} }
func (m *SlidingWindowsPayload) String() string { return proto.CompactTextString(m) }
func (*SlidingWindowsPayload) ProtoMessage()    {}
func (*SlidingWindowsPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_standard_window_fns_da1ada993b818d10, []int{2}
}
func (m *SlidingWindowsPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SlidingWindowsPayload.Unmarshal(m, b)
}
func (m *SlidingWindowsPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SlidingWindowsPayload.Marshal(b, m, deterministic)
}
func (dst *SlidingWindowsPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlidingWindowsPayload.Merge(dst, src)
}
func (m *SlidingWindowsPayload) XXX_Size() int {
	return xxx_messageInfo_SlidingWindowsPayload.Size(m)
}
func (m *SlidingWindowsPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_SlidingWindowsPayload.DiscardUnknown(m)
}

var xxx_messageInfo_SlidingWindowsPayload proto.InternalMessageInfo

func (m *SlidingWindowsPayload) GetSize() *duration.Duration {
	if m != nil {
		return m.Size
	}
	return nil
}

func (m *SlidingWindowsPayload) GetOffset() *timestamp.Timestamp {
	if m != nil {
		return m.Offset
	}
	return nil
}

func (m *SlidingWindowsPayload) GetPeriod() *duration.Duration {
	if m != nil {
		return m.Period
	}
	return nil
}

type SessionsPayload struct {
	GapSize              *duration.Duration `protobuf:"bytes,1,opt,name=gap_size,json=gapSize,proto3" json:"gap_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SessionsPayload) Reset()         { *m = SessionsPayload{} }
func (m *SessionsPayload) String() string { return proto.CompactTextString(m) }
func (*SessionsPayload) ProtoMessage()    {}
func (*SessionsPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_standard_window_fns_da1ada993b818d10, []int{3}
}
func (m *SessionsPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionsPayload.Unmarshal(m, b)
}
func (m *SessionsPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionsPayload.Marshal(b, m, deterministic)
}
func (dst *SessionsPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionsPayload.Merge(dst, src)
}
func (m *SessionsPayload) XXX_Size() int {
	return xxx_messageInfo_SessionsPayload.Size(m)
}
func (m *SessionsPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionsPayload.DiscardUnknown(m)
}

var xxx_messageInfo_SessionsPayload proto.InternalMessageInfo

func (m *SessionsPayload) GetGapSize() *duration.Duration {
	if m != nil {
		return m.GapSize
	}
	return nil
}

func init() {
	proto.RegisterType((*GlobalWindowsPayload)(nil), "org.apache.beam.model.pipeline.v1.GlobalWindowsPayload")
	proto.RegisterType((*FixedWindowsPayload)(nil), "org.apache.beam.model.pipeline.v1.FixedWindowsPayload")
	proto.RegisterType((*SlidingWindowsPayload)(nil), "org.apache.beam.model.pipeline.v1.SlidingWindowsPayload")
	proto.RegisterType((*SessionsPayload)(nil), "org.apache.beam.model.pipeline.v1.SessionsPayload")
	proto.RegisterEnum("org.apache.beam.model.pipeline.v1.GlobalWindowsPayload_Enum", GlobalWindowsPayload_Enum_name, GlobalWindowsPayload_Enum_value)
	proto.RegisterEnum("org.apache.beam.model.pipeline.v1.FixedWindowsPayload_Enum", FixedWindowsPayload_Enum_name, FixedWindowsPayload_Enum_value)
	proto.RegisterEnum("org.apache.beam.model.pipeline.v1.SlidingWindowsPayload_Enum", SlidingWindowsPayload_Enum_name, SlidingWindowsPayload_Enum_value)
	proto.RegisterEnum("org.apache.beam.model.pipeline.v1.SessionsPayload_Enum", SessionsPayload_Enum_name, SessionsPayload_Enum_value)
}

func init() {
	proto.RegisterFile("standard_window_fns.proto", fileDescriptor_standard_window_fns_da1ada993b818d10)
}

var fileDescriptor_standard_window_fns_da1ada993b818d10 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x31, 0x4f, 0xdb, 0x40,
	0x14, 0xc7, 0xeb, 0x36, 0x4d, 0xab, 0xcb, 0xd0, 0xd6, 0x6d, 0xa4, 0xc4, 0x43, 0x9b, 0x78, 0x68,
	0xb3, 0xf4, 0x52, 0xa7, 0x55, 0x41, 0x19, 0x00, 0x05, 0x12, 0xc4, 0x44, 0x64, 0x47, 0x8a, 0xc4,
	0x62, 0x9d, 0xb9, 0xb3, 0x39, 0xc9, 0xbe, 0x3b, 0xf9, 0xec, 0x04, 0xf2, 0x0d, 0xf8, 0x1a, 0x7c,
	0x06, 0x06, 0x66, 0xbe, 0x10, 0x3b, 0x13, 0xca, 0xf9, 0x82, 0xe4, 0x30, 0x04, 0x16, 0x46, 0xfb,
	0xfd, 0xdf, 0x7b, 0xbf, 0x9f, 0x4e, 0x0f, 0x34, 0x65, 0x86, 0x18, 0x46, 0x29, 0xf6, 0xe7, 0x94,
	0x61, 0x3e, 0xf7, 0x43, 0x26, 0xa1, 0x48, 0x79, 0xc6, 0xcd, 0x36, 0x4f, 0x23, 0x88, 0x04, 0x3a,
	0x3d, 0x23, 0x30, 0x20, 0x28, 0x81, 0x09, 0xc7, 0x24, 0x86, 0x82, 0x0a, 0x12, 0x53, 0x46, 0xe0,
	0xcc, 0xb1, 0xea, 0xcb, 0xff, 0x7e, 0x9a, 0x33, 0x46, 0x52, 0x1f, 0x09, 0x5a, 0x74, 0x5a, 0xdf,
	0x23, 0xce, 0xa3, 0x98, 0x74, 0xd5, 0x57, 0x90, 0x87, 0x5d, 0x9c, 0xa7, 0x28, 0xa3, 0x9c, 0xe9,
	0xfa, 0x8f, 0xf5, 0x7a, 0x46, 0x13, 0x22, 0x33, 0x94, 0x88, 0x22, 0x60, 0x4f, 0xc1, 0xb7, 0xc3,
	0x98, 0x07, 0x28, 0x9e, 0x2a, 0x28, 0x39, 0x46, 0x17, 0x31, 0x47, 0xd8, 0xde, 0x05, 0x95, 0x21,
	0xcb, 0x13, 0x73, 0x0b, 0x80, 0xb1, 0x7b, 0x3c, 0x1e, 0xba, 0x93, 0xa3, 0xa1, 0xf7, 0xf9, 0x8d,
	0xf5, 0xeb, 0xea, 0xfa, 0xfe, 0xf6, 0x7d, 0x7b, 0x49, 0xd3, 0x2f, 0x3c, 0x42, 0xd6, 0x8f, 0xd4,
	0x04, 0xed, 0x25, 0xfb, 0xb3, 0x3f, 0xd0, 0xb1, 0x6f, 0x0c, 0xf0, 0x75, 0x44, 0xcf, 0x09, 0x2e,
	0x0f, 0x36, 0x7f, 0x83, 0x8a, 0xa4, 0x0b, 0xd2, 0x30, 0x5a, 0x46, 0xa7, 0xd6, 0x6b, 0xc2, 0x02,
	0x10, 0xae, 0x00, 0xe1, 0x81, 0x16, 0x70, 0x55, 0xcc, 0xec, 0x81, 0x2a, 0x0f, 0x43, 0x49, 0xb2,
	0xc6, 0x5b, 0xd5, 0x60, 0x3d, 0x69, 0x98, 0xac, 0x8c, 0x5c, 0x9d, 0xb4, 0x77, 0x34, 0xfb, 0xff,
	0x35, 0xf6, 0x9f, 0x8a, 0xbd, 0x55, 0x66, 0x0f, 0x97, 0x8c, 0x65, 0xf4, 0x3b, 0x03, 0xd4, 0xbd,
	0x98, 0x62, 0xca, 0xa2, 0x57, 0x87, 0x37, 0x1d, 0x50, 0x15, 0x24, 0xa5, 0x1c, 0x37, 0xde, 0x6d,
	0x5a, 0xa2, 0x83, 0xf6, 0x9e, 0xf6, 0xdd, 0x5e, 0xf3, 0xed, 0x28, 0x5f, 0xbb, 0xec, 0x2b, 0x0b,
	0xb1, 0xb2, 0xf1, 0xa5, 0x01, 0x3e, 0x79, 0x44, 0x4a, 0xca, 0xd9, 0xa3, 0xeb, 0x3f, 0xf0, 0x31,
	0x42, 0xc2, 0x7f, 0x9e, 0xef, 0x87, 0x08, 0x09, 0x8f, 0x2e, 0xc8, 0x0b, 0x59, 0x8a, 0x95, 0x25,
	0x96, 0xc1, 0x3e, 0xd8, 0x7c, 0x0e, 0x83, 0x2f, 0x9e, 0x3e, 0xa6, 0xe2, 0x81, 0x46, 0x4c, 0x9e,
	0xd4, 0x56, 0x75, 0x7f, 0xe6, 0x04, 0x55, 0x85, 0xf8, 0xf7, 0x21, 0x00, 0x00, 0xff, 0xff, 0x00,
	0x17, 0x94, 0x23, 0x75, 0x03, 0x00, 0x00,
}
