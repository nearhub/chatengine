// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mtproto_meta.proto

package brpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MTProtoMeta struct {
	AuthKeyId            *int64   `protobuf:"varint,1,opt,name=auth_key_id,json=authKeyId" json:"auth_key_id,omitempty"`
	SessionId            *int64   `protobuf:"varint,2,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	MessageId            *int64   `protobuf:"varint,3,opt,name=message_id,json=messageId" json:"message_id,omitempty"`
	Layer                *int32   `protobuf:"varint,4,opt,name=layer" json:"layer,omitempty"`
	UserId               *int32   `protobuf:"varint,5,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	AccessHash           *int64   `protobuf:"varint,6,opt,name=access_hash,json=accessHash" json:"access_hash,omitempty"`
	ServerId             *int32   `protobuf:"varint,7,opt,name=server_id,json=serverId" json:"server_id,omitempty"`
	ClientConnId         *int64   `protobuf:"varint,8,opt,name=client_conn_id,json=clientConnId" json:"client_conn_id,omitempty"`
	ClientAddr           *string  `protobuf:"bytes,9,opt,name=client_addr,json=clientAddr" json:"client_addr,omitempty"`
	From                 *string  `protobuf:"bytes,10,opt,name=from" json:"from,omitempty"`
	ReceiveTime          *int64   `protobuf:"varint,11,opt,name=receive_time,json=receiveTime" json:"receive_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MTProtoMeta) Reset()         { *m = MTProtoMeta{} }
func (m *MTProtoMeta) String() string { return proto.CompactTextString(m) }
func (*MTProtoMeta) ProtoMessage()    {}
func (*MTProtoMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_mtproto_meta_0c7f8c31c82a7c80, []int{0}
}
func (m *MTProtoMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MTProtoMeta.Unmarshal(m, b)
}
func (m *MTProtoMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MTProtoMeta.Marshal(b, m, deterministic)
}
func (dst *MTProtoMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MTProtoMeta.Merge(dst, src)
}
func (m *MTProtoMeta) XXX_Size() int {
	return xxx_messageInfo_MTProtoMeta.Size(m)
}
func (m *MTProtoMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_MTProtoMeta.DiscardUnknown(m)
}

var xxx_messageInfo_MTProtoMeta proto.InternalMessageInfo

func (m *MTProtoMeta) GetAuthKeyId() int64 {
	if m != nil && m.AuthKeyId != nil {
		return *m.AuthKeyId
	}
	return 0
}

func (m *MTProtoMeta) GetSessionId() int64 {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return 0
}

func (m *MTProtoMeta) GetMessageId() int64 {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return 0
}

func (m *MTProtoMeta) GetLayer() int32 {
	if m != nil && m.Layer != nil {
		return *m.Layer
	}
	return 0
}

func (m *MTProtoMeta) GetUserId() int32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *MTProtoMeta) GetAccessHash() int64 {
	if m != nil && m.AccessHash != nil {
		return *m.AccessHash
	}
	return 0
}

func (m *MTProtoMeta) GetServerId() int32 {
	if m != nil && m.ServerId != nil {
		return *m.ServerId
	}
	return 0
}

func (m *MTProtoMeta) GetClientConnId() int64 {
	if m != nil && m.ClientConnId != nil {
		return *m.ClientConnId
	}
	return 0
}

func (m *MTProtoMeta) GetClientAddr() string {
	if m != nil && m.ClientAddr != nil {
		return *m.ClientAddr
	}
	return ""
}

func (m *MTProtoMeta) GetFrom() string {
	if m != nil && m.From != nil {
		return *m.From
	}
	return ""
}

func (m *MTProtoMeta) GetReceiveTime() int64 {
	if m != nil && m.ReceiveTime != nil {
		return *m.ReceiveTime
	}
	return 0
}

func init() {
	proto.RegisterType((*MTProtoMeta)(nil), "brpc.MTProtoMeta")
}

func init() { proto.RegisterFile("mtproto_meta.proto", fileDescriptor_mtproto_meta_0c7f8c31c82a7c80) }

var fileDescriptor_mtproto_meta_0c7f8c31c82a7c80 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x91, 0xcd, 0x6a, 0xeb, 0x30,
	0x10, 0x85, 0x71, 0xfe, 0x3d, 0x0e, 0x77, 0xa1, 0x5b, 0xa8, 0x69, 0x69, 0xeb, 0x96, 0x2e, 0xb2,
	0x4a, 0x9e, 0xa1, 0xe9, 0xa6, 0xa6, 0x84, 0x96, 0x90, 0xbd, 0x51, 0xa4, 0x49, 0x2c, 0x1a, 0x49,
	0x46, 0xa3, 0x04, 0xf2, 0xa2, 0x7d, 0x9e, 0x22, 0xc9, 0xbb, 0x99, 0xef, 0x9b, 0x73, 0x6c, 0x10,
	0x30, 0xed, 0x3b, 0x67, 0xbd, 0x6d, 0x34, 0x7a, 0xbe, 0x8c, 0x23, 0x1b, 0xed, 0x5d, 0x27, 0xee,
	0xaa, 0xa3, 0xb5, 0xc7, 0x13, 0xae, 0x22, 0xdb, 0x9f, 0x0f, 0x2b, 0x89, 0x24, 0x9c, 0xea, 0xbc,
	0x75, 0xe9, 0xee, 0xe5, 0x77, 0x00, 0xc5, 0x66, 0xf7, 0x1d, 0xe6, 0x0d, 0x7a, 0xce, 0x1e, 0xa1,
	0xe0, 0x67, 0xdf, 0x36, 0x3f, 0x78, 0x6d, 0x94, 0x2c, 0xb3, 0x2a, 0x5b, 0x0c, 0xb7, 0x79, 0x40,
	0x9f, 0x78, 0xad, 0x25, 0x7b, 0x00, 0x20, 0x24, 0x52, 0xd6, 0x04, 0x3d, 0x48, 0xba, 0x27, 0x49,
	0x6b, 0x24, 0xe2, 0x47, 0x0c, 0x7a, 0x98, 0x74, 0x4f, 0x6a, 0xc9, 0x6e, 0x60, 0x7c, 0xe2, 0x57,
	0x74, 0xe5, 0xa8, 0xca, 0x16, 0xe3, 0x6d, 0x5a, 0xd8, 0x2d, 0x4c, 0xcf, 0x84, 0x2e, 0x24, 0xc6,
	0x91, 0x4f, 0xc2, 0x5a, 0x4b, 0xf6, 0x04, 0x05, 0x17, 0x02, 0x89, 0x9a, 0x96, 0x53, 0x5b, 0x4e,
	0x62, 0x1d, 0x24, 0xf4, 0xc1, 0xa9, 0x65, 0xf7, 0x90, 0x13, 0xba, 0x4b, 0xca, 0x4e, 0x63, 0x76,
	0x96, 0x40, 0x2d, 0xd9, 0x2b, 0xfc, 0x13, 0x27, 0x85, 0xc6, 0x37, 0xc2, 0x9a, 0xf8, 0xbb, 0xb3,
	0x58, 0x30, 0x4f, 0xf4, 0xdd, 0x1a, 0x93, 0xbe, 0xd1, 0x5f, 0x71, 0x29, 0x5d, 0x99, 0x57, 0xd9,
	0x22, 0xdf, 0x42, 0x42, 0x6f, 0x52, 0x3a, 0xc6, 0x60, 0x74, 0x70, 0x56, 0x97, 0x10, 0x4d, 0x9c,
	0xd9, 0x33, 0xcc, 0x1d, 0x0a, 0x54, 0x17, 0x6c, 0xbc, 0xd2, 0x58, 0x16, 0xb1, 0xb8, 0xe8, 0xd9,
	0x4e, 0x69, 0x5c, 0xff, 0x87, 0x99, 0xb0, 0x7a, 0x19, 0x9e, 0x61, 0x3d, 0xfd, 0xea, 0xbc, 0xb2,
	0x86, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xaf, 0x65, 0x39, 0xaa, 0x01, 0x00, 0x00,
}