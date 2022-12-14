// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: Stat.proto

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

// CheckAndStatArgs
// path: file path
type CheckAndStatArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	IsLatest bool   `protobuf:"varint,2,opt,name=isLatest,proto3" json:"isLatest,omitempty"`
}

func (x *CheckAndStatArgs) Reset() {
	*x = CheckAndStatArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Stat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAndStatArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAndStatArgs) ProtoMessage() {}

func (x *CheckAndStatArgs) ProtoReflect() protoreflect.Message {
	mi := &file_Stat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAndStatArgs.ProtoReflect.Descriptor instead.
func (*CheckAndStatArgs) Descriptor() ([]byte, []int) {
	return file_Stat_proto_rawDescGZIP(), []int{0}
}

func (x *CheckAndStatArgs) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CheckAndStatArgs) GetIsLatest() bool {
	if x != nil {
		return x.IsLatest
	}
	return false
}

type CheckAndStatReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	IsFile   bool   `protobuf:"varint,2,opt,name=isFile,proto3" json:"isFile,omitempty"`
	Size     int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *CheckAndStatReply) Reset() {
	*x = CheckAndStatReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Stat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAndStatReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAndStatReply) ProtoMessage() {}

func (x *CheckAndStatReply) ProtoReflect() protoreflect.Message {
	mi := &file_Stat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAndStatReply.ProtoReflect.Descriptor instead.
func (*CheckAndStatReply) Descriptor() ([]byte, []int) {
	return file_Stat_proto_rawDescGZIP(), []int{1}
}

func (x *CheckAndStatReply) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *CheckAndStatReply) GetIsFile() bool {
	if x != nil {
		return x.IsFile
	}
	return false
}

func (x *CheckAndStatReply) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_Stat_proto protoreflect.FileDescriptor

var file_Stat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x42, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x4c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x22, 0x5b, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x32, 0x50, 0x0a, 0x11, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41,
	0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x41, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_Stat_proto_rawDescOnce sync.Once
	file_Stat_proto_rawDescData = file_Stat_proto_rawDesc
)

func file_Stat_proto_rawDescGZIP() []byte {
	file_Stat_proto_rawDescOnce.Do(func() {
		file_Stat_proto_rawDescData = protoimpl.X.CompressGZIP(file_Stat_proto_rawDescData)
	})
	return file_Stat_proto_rawDescData
}

var file_Stat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Stat_proto_goTypes = []interface{}{
	(*CheckAndStatArgs)(nil),  // 0: pb.CheckAndStatArgs
	(*CheckAndStatReply)(nil), // 1: pb.CheckAndStatReply
}
var file_Stat_proto_depIdxs = []int32{
	0, // 0: pb.MasterStatService.CheckAndStat:input_type -> pb.CheckAndStatArgs
	1, // 1: pb.MasterStatService.CheckAndStat:output_type -> pb.CheckAndStatReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Stat_proto_init() }
func file_Stat_proto_init() {
	if File_Stat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Stat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAndStatArgs); i {
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
		file_Stat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAndStatReply); i {
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
			RawDescriptor: file_Stat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Stat_proto_goTypes,
		DependencyIndexes: file_Stat_proto_depIdxs,
		MessageInfos:      file_Stat_proto_msgTypes,
	}.Build()
	File_Stat_proto = out.File
	file_Stat_proto_rawDesc = nil
	file_Stat_proto_goTypes = nil
	file_Stat_proto_depIdxs = nil
}
