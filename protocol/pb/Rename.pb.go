// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: Rename.proto

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

// CheckAndRenameArgs
// path: the file to be renamed
// newName: the file's new name
type CheckAndRenameArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path    string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	NewName string `protobuf:"bytes,2,opt,name=newName,proto3" json:"newName,omitempty"`
}

func (x *CheckAndRenameArgs) Reset() {
	*x = CheckAndRenameArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Rename_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAndRenameArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAndRenameArgs) ProtoMessage() {}

func (x *CheckAndRenameArgs) ProtoReflect() protoreflect.Message {
	mi := &file_Rename_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAndRenameArgs.ProtoReflect.Descriptor instead.
func (*CheckAndRenameArgs) Descriptor() ([]byte, []int) {
	return file_Rename_proto_rawDescGZIP(), []int{0}
}

func (x *CheckAndRenameArgs) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CheckAndRenameArgs) GetNewName() string {
	if x != nil {
		return x.NewName
	}
	return ""
}

type CheckAndRenameReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CheckAndRenameReply) Reset() {
	*x = CheckAndRenameReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Rename_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAndRenameReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAndRenameReply) ProtoMessage() {}

func (x *CheckAndRenameReply) ProtoReflect() protoreflect.Message {
	mi := &file_Rename_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAndRenameReply.ProtoReflect.Descriptor instead.
func (*CheckAndRenameReply) Descriptor() ([]byte, []int) {
	return file_Rename_proto_rawDescGZIP(), []int{1}
}

var File_Rename_proto protoreflect.FileDescriptor

var file_Rename_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x42, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e, 0x64, 0x52, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07,
	0x6e, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e,
	0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41,
	0x6e, 0x64, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x58, 0x0a,
	0x13, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e, 0x64,
	0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x17,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Rename_proto_rawDescOnce sync.Once
	file_Rename_proto_rawDescData = file_Rename_proto_rawDesc
)

func file_Rename_proto_rawDescGZIP() []byte {
	file_Rename_proto_rawDescOnce.Do(func() {
		file_Rename_proto_rawDescData = protoimpl.X.CompressGZIP(file_Rename_proto_rawDescData)
	})
	return file_Rename_proto_rawDescData
}

var file_Rename_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Rename_proto_goTypes = []interface{}{
	(*CheckAndRenameArgs)(nil),  // 0: pb.CheckAndRenameArgs
	(*CheckAndRenameReply)(nil), // 1: pb.CheckAndRenameReply
}
var file_Rename_proto_depIdxs = []int32{
	0, // 0: pb.MasterRenameService.CheckAndRename:input_type -> pb.CheckAndRenameArgs
	1, // 1: pb.MasterRenameService.CheckAndRename:output_type -> pb.CheckAndRenameReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Rename_proto_init() }
func file_Rename_proto_init() {
	if File_Rename_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Rename_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAndRenameArgs); i {
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
		file_Rename_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAndRenameReply); i {
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
			RawDescriptor: file_Rename_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Rename_proto_goTypes,
		DependencyIndexes: file_Rename_proto_depIdxs,
		MessageInfos:      file_Rename_proto_msgTypes,
	}.Build()
	File_Rename_proto = out.File
	file_Rename_proto_rawDesc = nil
	file_Rename_proto_goTypes = nil
	file_Rename_proto_depIdxs = nil
}
