// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: Shadow.proto

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

type OperationArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`          // Identifier
	Type     string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`          // Specify the Operation Type
	Src      string `protobuf:"bytes,3,opt,name=src,proto3" json:"src,omitempty"`            // Client `src` Parameter
	Des      string `protobuf:"bytes,4,opt,name=des,proto3" json:"des,omitempty"`            // Client `des` Parameter
	IsFile   bool   `protobuf:"varint,5,opt,name=isFile,proto3" json:"isFile,omitempty"`     // Available if Type is Add
	FileName string `protobuf:"bytes,6,opt,name=fileName,proto3" json:"fileName,omitempty"`  // Available if Type is Add and IsFile is True
	Size     int64  `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`         // Available if Type is Add
	IsFinish bool   `protobuf:"varint,8,opt,name=isFinish,proto3" json:"isFinish,omitempty"` // True when client notify the master
}

func (x *OperationArgs) Reset() {
	*x = OperationArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Shadow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationArgs) ProtoMessage() {}

func (x *OperationArgs) ProtoReflect() protoreflect.Message {
	mi := &file_Shadow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationArgs.ProtoReflect.Descriptor instead.
func (*OperationArgs) Descriptor() ([]byte, []int) {
	return file_Shadow_proto_rawDescGZIP(), []int{0}
}

func (x *OperationArgs) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *OperationArgs) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *OperationArgs) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

func (x *OperationArgs) GetDes() string {
	if x != nil {
		return x.Des
	}
	return ""
}

func (x *OperationArgs) GetIsFile() bool {
	if x != nil {
		return x.IsFile
	}
	return false
}

func (x *OperationArgs) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *OperationArgs) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *OperationArgs) GetIsFinish() bool {
	if x != nil {
		return x.IsFinish
	}
	return false
}

type OperationReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *OperationReply) Reset() {
	*x = OperationReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Shadow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationReply) ProtoMessage() {}

func (x *OperationReply) ProtoReflect() protoreflect.Message {
	mi := &file_Shadow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationReply.ProtoReflect.Descriptor instead.
func (*OperationReply) Descriptor() ([]byte, []int) {
	return file_Shadow_proto_rawDescGZIP(), []int{1}
}

func (x *OperationReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_Shadow_proto protoreflect.FileDescriptor

var file_Shadow_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x53, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0xbf, 0x01, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x72, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x72, 0x63, 0x12, 0x10,
	0x0a, 0x03, 0x64, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x65, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x69, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x22, 0x20, 0x0a, 0x0e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32, 0x4e, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36,
	0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x72,
	0x67, 0x73, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Shadow_proto_rawDescOnce sync.Once
	file_Shadow_proto_rawDescData = file_Shadow_proto_rawDesc
)

func file_Shadow_proto_rawDescGZIP() []byte {
	file_Shadow_proto_rawDescOnce.Do(func() {
		file_Shadow_proto_rawDescData = protoimpl.X.CompressGZIP(file_Shadow_proto_rawDescData)
	})
	return file_Shadow_proto_rawDescData
}

var file_Shadow_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Shadow_proto_goTypes = []interface{}{
	(*OperationArgs)(nil),  // 0: pb.OperationArgs
	(*OperationReply)(nil), // 1: pb.OperationReply
}
var file_Shadow_proto_depIdxs = []int32{
	0, // 0: pb.SendOperationService.SendOperation:input_type -> pb.OperationArgs
	1, // 1: pb.SendOperationService.SendOperation:output_type -> pb.OperationReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Shadow_proto_init() }
func file_Shadow_proto_init() {
	if File_Shadow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Shadow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationArgs); i {
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
		file_Shadow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationReply); i {
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
			RawDescriptor: file_Shadow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Shadow_proto_goTypes,
		DependencyIndexes: file_Shadow_proto_depIdxs,
		MessageInfos:      file_Shadow_proto_msgTypes,
	}.Build()
	File_Shadow_proto = out.File
	file_Shadow_proto_rawDesc = nil
	file_Shadow_proto_goTypes = nil
	file_Shadow_proto_depIdxs = nil
}
