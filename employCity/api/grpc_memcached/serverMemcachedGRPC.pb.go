// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: api/grpc_memcached/serverMemcachedGRPC.proto

package grpc_memcached

import (
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

type SetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Body string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *SetRequest) Reset() {
	*x = SetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_memcached_serverMemcachedGRPC_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRequest) ProtoMessage() {}

func (x *SetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_memcached_serverMemcachedGRPC_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRequest.ProtoReflect.Descriptor instead.
func (*SetRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_memcached_serverMemcachedGRPC_proto_rawDescGZIP(), []int{0}
}

func (x *SetRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SetRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type DeleteGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGetRequest) Reset() {
	*x = DeleteGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_memcached_serverMemcachedGRPC_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGetRequest) ProtoMessage() {}

func (x *DeleteGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_memcached_serverMemcachedGRPC_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGetRequest.ProtoReflect.Descriptor instead.
func (*DeleteGetRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_memcached_serverMemcachedGRPC_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteGetRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_memcached_serverMemcachedGRPC_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_memcached_serverMemcachedGRPC_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_api_grpc_memcached_serverMemcachedGRPC_proto_rawDescGZIP(), []int{2}
}

func (x *Reply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Reply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ReplyGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Body string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *ReplyGet) Reset() {
	*x = ReplyGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_memcached_serverMemcachedGRPC_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyGet) ProtoMessage() {}

func (x *ReplyGet) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_memcached_serverMemcachedGRPC_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyGet.ProtoReflect.Descriptor instead.
func (*ReplyGet) Descriptor() ([]byte, []int) {
	return file_api_grpc_memcached_serverMemcachedGRPC_proto_rawDescGZIP(), []int{3}
}

func (x *ReplyGet) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReplyGet) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_api_grpc_memcached_serverMemcachedGRPC_proto protoreflect.FileDescriptor

var file_api_grpc_memcached_serverMemcachedGRPC_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x6d, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x6d, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x64, 0x47, 0x52, 0x50, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30,
	0x0a, 0x0a, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x22, 0x22, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2e, 0x0a, 0x08, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x47, 0x65,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x32, 0x74, 0x0a, 0x06, 0x4d, 0x79, 0x47, 0x52, 0x50, 0x43, 0x12,
	0x25, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x47, 0x65, 0x74, 0x22, 0x00, 0x12, 0x1c, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x0b, 0x2e,
	0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x06, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x2e,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x6d, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_memcached_serverMemcachedGRPC_proto_rawDescOnce sync.Once
	file_api_grpc_memcached_serverMemcachedGRPC_proto_rawDescData = file_api_grpc_memcached_serverMemcachedGRPC_proto_rawDesc
)

func file_api_grpc_memcached_serverMemcachedGRPC_proto_rawDescGZIP() []byte {
	file_api_grpc_memcached_serverMemcachedGRPC_proto_rawDescOnce.Do(func() {
		file_api_grpc_memcached_serverMemcachedGRPC_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_memcached_serverMemcachedGRPC_proto_rawDescData)
	})
	return file_api_grpc_memcached_serverMemcachedGRPC_proto_rawDescData
}

var file_api_grpc_memcached_serverMemcachedGRPC_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_grpc_memcached_serverMemcachedGRPC_proto_goTypes = []interface{}{
	(*SetRequest)(nil),       // 0: SetRequest
	(*DeleteGetRequest)(nil), // 1: DeleteGetRequest
	(*Reply)(nil),            // 2: Reply
	(*ReplyGet)(nil),         // 3: ReplyGet
}
var file_api_grpc_memcached_serverMemcachedGRPC_proto_depIdxs = []int32{
	1, // 0: MyGRPC.Get:input_type -> DeleteGetRequest
	0, // 1: MyGRPC.Set:input_type -> SetRequest
	1, // 2: MyGRPC.Delete:input_type -> DeleteGetRequest
	3, // 3: MyGRPC.Get:output_type -> ReplyGet
	2, // 4: MyGRPC.Set:output_type -> Reply
	2, // 5: MyGRPC.Delete:output_type -> Reply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_grpc_memcached_serverMemcachedGRPC_proto_init() }
func file_api_grpc_memcached_serverMemcachedGRPC_proto_init() {
	if File_api_grpc_memcached_serverMemcachedGRPC_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_memcached_serverMemcachedGRPC_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRequest); i {
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
		file_api_grpc_memcached_serverMemcachedGRPC_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGetRequest); i {
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
		file_api_grpc_memcached_serverMemcachedGRPC_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
		file_api_grpc_memcached_serverMemcachedGRPC_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyGet); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_grpc_memcached_serverMemcachedGRPC_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_memcached_serverMemcachedGRPC_proto_goTypes,
		DependencyIndexes: file_api_grpc_memcached_serverMemcachedGRPC_proto_depIdxs,
		MessageInfos:      file_api_grpc_memcached_serverMemcachedGRPC_proto_msgTypes,
	}.Build()
	File_api_grpc_memcached_serverMemcachedGRPC_proto = out.File
	file_api_grpc_memcached_serverMemcachedGRPC_proto_rawDesc = nil
	file_api_grpc_memcached_serverMemcachedGRPC_proto_goTypes = nil
	file_api_grpc_memcached_serverMemcachedGRPC_proto_depIdxs = nil
}
