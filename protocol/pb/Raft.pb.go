// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: Raft.proto

package pb

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

type Join2ClusterArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Join2ClusterArgs) Reset() {
	*x = Join2ClusterArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Raft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Join2ClusterArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Join2ClusterArgs) ProtoMessage() {}

func (x *Join2ClusterArgs) ProtoReflect() protoreflect.Message {
	mi := &file_Raft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Join2ClusterArgs.ProtoReflect.Descriptor instead.
func (*Join2ClusterArgs) Descriptor() ([]byte, []int) {
	return file_Raft_proto_rawDescGZIP(), []int{0}
}

type Join2ClusterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Join2ClusterReply) Reset() {
	*x = Join2ClusterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Raft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Join2ClusterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Join2ClusterReply) ProtoMessage() {}

func (x *Join2ClusterReply) ProtoReflect() protoreflect.Message {
	mi := &file_Raft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Join2ClusterReply.ProtoReflect.Descriptor instead.
func (*Join2ClusterReply) Descriptor() ([]byte, []int) {
	return file_Raft_proto_rawDescGZIP(), []int{1}
}

var File_Raft_proto protoreflect.FileDescriptor

var file_Raft_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x52, 0x61, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x12, 0x0a, 0x10, 0x4a, 0x6f, 0x69, 0x6e, 0x32, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x41, 0x72, 0x67, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x4a, 0x6f, 0x69, 0x6e, 0x32, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x4a, 0x0a, 0x0b, 0x52, 0x61, 0x66,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x4a, 0x6f, 0x69, 0x6e,
	0x32, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4a, 0x6f,
	0x69, 0x6e, 0x32, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x15,
	0x2e, 0x70, 0x62, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x32, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Raft_proto_rawDescOnce sync.Once
	file_Raft_proto_rawDescData = file_Raft_proto_rawDesc
)

func file_Raft_proto_rawDescGZIP() []byte {
	file_Raft_proto_rawDescOnce.Do(func() {
		file_Raft_proto_rawDescData = protoimpl.X.CompressGZIP(file_Raft_proto_rawDescData)
	})
	return file_Raft_proto_rawDescData
}

var file_Raft_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Raft_proto_goTypes = []interface{}{
	(*Join2ClusterArgs)(nil),  // 0: pb.Join2ClusterArgs
	(*Join2ClusterReply)(nil), // 1: pb.Join2ClusterReply
}
var file_Raft_proto_depIdxs = []int32{
	0, // 0: pb.RaftService.Join2Cluster:input_type -> pb.Join2ClusterArgs
	1, // 1: pb.RaftService.Join2Cluster:output_type -> pb.Join2ClusterReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Raft_proto_init() }
func file_Raft_proto_init() {
	if File_Raft_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Raft_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Join2ClusterArgs); i {
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
		file_Raft_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Join2ClusterReply); i {
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
			RawDescriptor: file_Raft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Raft_proto_goTypes,
		DependencyIndexes: file_Raft_proto_depIdxs,
		MessageInfos:      file_Raft_proto_msgTypes,
	}.Build()
	File_Raft_proto = out.File
	file_Raft_proto_rawDesc = nil
	file_Raft_proto_goTypes = nil
	file_Raft_proto_depIdxs = nil
}
