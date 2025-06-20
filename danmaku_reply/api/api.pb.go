// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: reply/api/api.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PingReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PingReq) Reset() {
	*x = PingReq{}
	mi := &file_reply_api_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReq) ProtoMessage() {}

func (x *PingReq) ProtoReflect() protoreflect.Message {
	mi := &file_reply_api_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReq.ProtoReflect.Descriptor instead.
func (*PingReq) Descriptor() ([]byte, []int) {
	return file_reply_api_api_proto_rawDescGZIP(), []int{0}
}

type PingRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PingRes) Reset() {
	*x = PingRes{}
	mi := &file_reply_api_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRes) ProtoMessage() {}

func (x *PingRes) ProtoReflect() protoreflect.Message {
	mi := &file_reply_api_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRes.ProtoReflect.Descriptor instead.
func (*PingRes) Descriptor() ([]byte, []int) {
	return file_reply_api_api_proto_rawDescGZIP(), []int{1}
}

type ReplyByGroupIDReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GroupID       int64                  `protobuf:"varint,1,opt,name=GroupID,proto3" json:"GroupID,omitempty"`
	Question      string                 `protobuf:"bytes,2,opt,name=Question,proto3" json:"Question,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReplyByGroupIDReq) Reset() {
	*x = ReplyByGroupIDReq{}
	mi := &file_reply_api_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplyByGroupIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyByGroupIDReq) ProtoMessage() {}

func (x *ReplyByGroupIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_reply_api_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyByGroupIDReq.ProtoReflect.Descriptor instead.
func (*ReplyByGroupIDReq) Descriptor() ([]byte, []int) {
	return file_reply_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *ReplyByGroupIDReq) GetGroupID() int64 {
	if x != nil {
		return x.GroupID
	}
	return 0
}

func (x *ReplyByGroupIDReq) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

type ReplyByGroupIDRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Reply         string                 `protobuf:"bytes,1,opt,name=Reply,proto3" json:"Reply,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReplyByGroupIDRes) Reset() {
	*x = ReplyByGroupIDRes{}
	mi := &file_reply_api_api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplyByGroupIDRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyByGroupIDRes) ProtoMessage() {}

func (x *ReplyByGroupIDRes) ProtoReflect() protoreflect.Message {
	mi := &file_reply_api_api_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyByGroupIDRes.ProtoReflect.Descriptor instead.
func (*ReplyByGroupIDRes) Descriptor() ([]byte, []int) {
	return file_reply_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *ReplyByGroupIDRes) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

var File_reply_api_api_proto protoreflect.FileDescriptor

const file_reply_api_api_proto_rawDesc = "" +
	"\n" +
	"\x13reply/api/api.proto\"\t\n" +
	"\aPingReq\"\t\n" +
	"\aPingRes\"I\n" +
	"\x11ReplyByGroupIDReq\x12\x18\n" +
	"\aGroupID\x18\x01 \x01(\x03R\aGroupID\x12\x1a\n" +
	"\bQuestion\x18\x02 \x01(\tR\bQuestion\")\n" +
	"\x11ReplyByGroupIDRes\x12\x14\n" +
	"\x05Reply\x18\x01 \x01(\tR\x05Reply2c\n" +
	"\vReplyServer\x12\x1a\n" +
	"\x04Ping\x12\b.PingReq\x1a\b.PingRes\x128\n" +
	"\x0eReplyByGroupID\x12\x12.ReplyByGroupIDReq\x1a\x12.ReplyByGroupIDResB\vZ\treply/apib\x06proto3"

var (
	file_reply_api_api_proto_rawDescOnce sync.Once
	file_reply_api_api_proto_rawDescData []byte
)

func file_reply_api_api_proto_rawDescGZIP() []byte {
	file_reply_api_api_proto_rawDescOnce.Do(func() {
		file_reply_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_reply_api_api_proto_rawDesc), len(file_reply_api_api_proto_rawDesc)))
	})
	return file_reply_api_api_proto_rawDescData
}

var file_reply_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_reply_api_api_proto_goTypes = []any{
	(*PingReq)(nil),           // 0: PingReq
	(*PingRes)(nil),           // 1: PingRes
	(*ReplyByGroupIDReq)(nil), // 2: ReplyByGroupIDReq
	(*ReplyByGroupIDRes)(nil), // 3: ReplyByGroupIDRes
}
var file_reply_api_api_proto_depIdxs = []int32{
	0, // 0: ReplyServer.Ping:input_type -> PingReq
	2, // 1: ReplyServer.ReplyByGroupID:input_type -> ReplyByGroupIDReq
	1, // 2: ReplyServer.Ping:output_type -> PingRes
	3, // 3: ReplyServer.ReplyByGroupID:output_type -> ReplyByGroupIDRes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_reply_api_api_proto_init() }
func file_reply_api_api_proto_init() {
	if File_reply_api_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_reply_api_api_proto_rawDesc), len(file_reply_api_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reply_api_api_proto_goTypes,
		DependencyIndexes: file_reply_api_api_proto_depIdxs,
		MessageInfos:      file_reply_api_api_proto_msgTypes,
	}.Build()
	File_reply_api_api_proto = out.File
	file_reply_api_api_proto_goTypes = nil
	file_reply_api_api_proto_depIdxs = nil
}
