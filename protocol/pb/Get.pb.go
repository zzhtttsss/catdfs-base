// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: Get.proto

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

// CheckAndGetArgs
// path: target path to get
type CheckAndGetArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *CheckAndGetArgs) Reset() {
	*x = CheckAndGetArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Get_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAndGetArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAndGetArgs) ProtoMessage() {}

func (x *CheckAndGetArgs) ProtoReflect() protoreflect.Message {
	mi := &file_Get_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAndGetArgs.ProtoReflect.Descriptor instead.
func (*CheckAndGetArgs) Descriptor() ([]byte, []int) {
	return file_Get_proto_rawDescGZIP(), []int{0}
}

func (x *CheckAndGetArgs) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// CheckAndGetReply
// fileNodeId: file id stored in the directory tree
// chunkNum: the number of chunks the file will be cut into
// operationId: the get operation id stored in the edits
type CheckAndGetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileNodeId  string `protobuf:"bytes,1,opt,name=fileNodeId,proto3" json:"fileNodeId,omitempty"`
	ChunkNum    int32  `protobuf:"varint,2,opt,name=chunkNum,proto3" json:"chunkNum,omitempty"`
	OperationId string `protobuf:"bytes,3,opt,name=operationId,proto3" json:"operationId,omitempty"`
	FileSize    int64  `protobuf:"varint,4,opt,name=fileSize,proto3" json:"fileSize,omitempty"`
}

func (x *CheckAndGetReply) Reset() {
	*x = CheckAndGetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Get_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAndGetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAndGetReply) ProtoMessage() {}

func (x *CheckAndGetReply) ProtoReflect() protoreflect.Message {
	mi := &file_Get_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAndGetReply.ProtoReflect.Descriptor instead.
func (*CheckAndGetReply) Descriptor() ([]byte, []int) {
	return file_Get_proto_rawDescGZIP(), []int{1}
}

func (x *CheckAndGetReply) GetFileNodeId() string {
	if x != nil {
		return x.FileNodeId
	}
	return ""
}

func (x *CheckAndGetReply) GetChunkNum() int32 {
	if x != nil {
		return x.ChunkNum
	}
	return 0
}

func (x *CheckAndGetReply) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

func (x *CheckAndGetReply) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

// GetDataNodes4GetArgs
// fileNodeId: file id stored in the directory tree
// chunkIndex: the index of chunk
type GetDataNodes4GetArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileNodeId string `protobuf:"bytes,1,opt,name=fileNodeId,proto3" json:"fileNodeId,omitempty"`
	ChunkIndex int32  `protobuf:"varint,2,opt,name=chunkIndex,proto3" json:"chunkIndex,omitempty"`
}

func (x *GetDataNodes4GetArgs) Reset() {
	*x = GetDataNodes4GetArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Get_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataNodes4GetArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataNodes4GetArgs) ProtoMessage() {}

func (x *GetDataNodes4GetArgs) ProtoReflect() protoreflect.Message {
	mi := &file_Get_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataNodes4GetArgs.ProtoReflect.Descriptor instead.
func (*GetDataNodes4GetArgs) Descriptor() ([]byte, []int) {
	return file_Get_proto_rawDescGZIP(), []int{2}
}

func (x *GetDataNodes4GetArgs) GetFileNodeId() string {
	if x != nil {
		return x.FileNodeId
	}
	return ""
}

func (x *GetDataNodes4GetArgs) GetChunkIndex() int32 {
	if x != nil {
		return x.ChunkIndex
	}
	return 0
}

// GetDataNodes4GetReply
// primaryNodeId: id of the primary DataNode
// primaryNodeAddr: address of the primary DataNode
// chunkIndex: index of current chunk index of the file
// chunkSize: the size of current chunk
type GetDataNodes4GetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataNodeIds   []string `protobuf:"bytes,1,rep,name=dataNodeIds,proto3" json:"dataNodeIds,omitempty"`
	DataNodeAddrs []string `protobuf:"bytes,2,rep,name=dataNodeAddrs,proto3" json:"dataNodeAddrs,omitempty"`
	ChunkIndex    int32    `protobuf:"varint,3,opt,name=chunkIndex,proto3" json:"chunkIndex,omitempty"`
}

func (x *GetDataNodes4GetReply) Reset() {
	*x = GetDataNodes4GetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Get_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataNodes4GetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataNodes4GetReply) ProtoMessage() {}

func (x *GetDataNodes4GetReply) ProtoReflect() protoreflect.Message {
	mi := &file_Get_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataNodes4GetReply.ProtoReflect.Descriptor instead.
func (*GetDataNodes4GetReply) Descriptor() ([]byte, []int) {
	return file_Get_proto_rawDescGZIP(), []int{3}
}

func (x *GetDataNodes4GetReply) GetDataNodeIds() []string {
	if x != nil {
		return x.DataNodeIds
	}
	return nil
}

func (x *GetDataNodes4GetReply) GetDataNodeAddrs() []string {
	if x != nil {
		return x.DataNodeAddrs
	}
	return nil
}

func (x *GetDataNodes4GetReply) GetChunkIndex() int32 {
	if x != nil {
		return x.ChunkIndex
	}
	return 0
}

// SetupStream2DataNodeArgs sets up the stream with the datanode and then transfer data
// clientAddr: address of the client
// chunkId: fileNodeId+chunkIndex which represents a chunk
// dataNodeId: id of dataNode
type SetupStream2DataNodeArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientPort string `protobuf:"bytes,1,opt,name=clientPort,proto3" json:"clientPort,omitempty"`
	ChunkId    string `protobuf:"bytes,2,opt,name=chunkId,proto3" json:"chunkId,omitempty"`
	DataNodeId string `protobuf:"bytes,3,opt,name=dataNodeId,proto3" json:"dataNodeId,omitempty"`
}

func (x *SetupStream2DataNodeArgs) Reset() {
	*x = SetupStream2DataNodeArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Get_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetupStream2DataNodeArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetupStream2DataNodeArgs) ProtoMessage() {}

func (x *SetupStream2DataNodeArgs) ProtoReflect() protoreflect.Message {
	mi := &file_Get_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetupStream2DataNodeArgs.ProtoReflect.Descriptor instead.
func (*SetupStream2DataNodeArgs) Descriptor() ([]byte, []int) {
	return file_Get_proto_rawDescGZIP(), []int{4}
}

func (x *SetupStream2DataNodeArgs) GetClientPort() string {
	if x != nil {
		return x.ClientPort
	}
	return ""
}

func (x *SetupStream2DataNodeArgs) GetChunkId() string {
	if x != nil {
		return x.ChunkId
	}
	return ""
}

func (x *SetupStream2DataNodeArgs) GetDataNodeId() string {
	if x != nil {
		return x.DataNodeId
	}
	return ""
}

type SetupStream2DataNodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetupStream2DataNodeReply) Reset() {
	*x = SetupStream2DataNodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Get_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetupStream2DataNodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetupStream2DataNodeReply) ProtoMessage() {}

func (x *SetupStream2DataNodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_Get_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetupStream2DataNodeReply.ProtoReflect.Descriptor instead.
func (*SetupStream2DataNodeReply) Descriptor() ([]byte, []int) {
	return file_Get_proto_rawDescGZIP(), []int{5}
}

// ReleaseLease4GetArgs
// chunkId: id of chunk which has been stored by chunkserver and need to release its lease
type ReleaseLease4GetArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkId string `protobuf:"bytes,1,opt,name=chunkId,proto3" json:"chunkId,omitempty"`
}

func (x *ReleaseLease4GetArgs) Reset() {
	*x = ReleaseLease4GetArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Get_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseLease4GetArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseLease4GetArgs) ProtoMessage() {}

func (x *ReleaseLease4GetArgs) ProtoReflect() protoreflect.Message {
	mi := &file_Get_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseLease4GetArgs.ProtoReflect.Descriptor instead.
func (*ReleaseLease4GetArgs) Descriptor() ([]byte, []int) {
	return file_Get_proto_rawDescGZIP(), []int{6}
}

func (x *ReleaseLease4GetArgs) GetChunkId() string {
	if x != nil {
		return x.ChunkId
	}
	return ""
}

type ReleaseLease4GetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReleaseLease4GetReply) Reset() {
	*x = ReleaseLease4GetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Get_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseLease4GetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseLease4GetReply) ProtoMessage() {}

func (x *ReleaseLease4GetReply) ProtoReflect() protoreflect.Message {
	mi := &file_Get_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseLease4GetReply.ProtoReflect.Descriptor instead.
func (*ReleaseLease4GetReply) Descriptor() ([]byte, []int) {
	return file_Get_proto_rawDescGZIP(), []int{7}
}

//The same as "PieceOfChunk". Waiting to refine.
type Piece struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Piece []byte `protobuf:"bytes,1,opt,name=piece,proto3" json:"piece,omitempty"`
}

func (x *Piece) Reset() {
	*x = Piece{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Get_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Piece) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Piece) ProtoMessage() {}

func (x *Piece) ProtoReflect() protoreflect.Message {
	mi := &file_Get_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Piece.ProtoReflect.Descriptor instead.
func (*Piece) Descriptor() ([]byte, []int) {
	return file_Get_proto_rawDescGZIP(), []int{8}
}

func (x *Piece) GetPiece() []byte {
	if x != nil {
		return x.Piece
	}
	return nil
}

var File_Get_proto protoreflect.FileDescriptor

var file_Get_proto_rawDesc = []byte{
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0x25, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e, 0x64, 0x47, 0x65, 0x74, 0x41, 0x72,
	0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x8c, 0x01, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x41, 0x6e, 0x64, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x66,
	0x69, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x56, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x34, 0x47, 0x65, 0x74, 0x41, 0x72, 0x67, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x7f, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x34, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74,
	0x61, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61,
	0x4e, 0x6f, 0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0d, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x74,
	0x0a, 0x18, 0x53, 0x65, 0x74, 0x75, 0x70, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x32, 0x44, 0x61,
	0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x22, 0x1b, 0x0a, 0x19, 0x53, 0x65, 0x74, 0x75, 0x70, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x32, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x30, 0x0a, 0x14, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4c, 0x65, 0x61, 0x73,
	0x65, 0x34, 0x47, 0x65, 0x74, 0x41, 0x72, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4c, 0x65,
	0x61, 0x73, 0x65, 0x34, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x0a, 0x05,
	0x50, 0x69, 0x65, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x69, 0x65, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x69, 0x65, 0x63, 0x65, 0x32, 0xde, 0x01, 0x0a, 0x10,
	0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x38, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e, 0x64, 0x47, 0x65, 0x74, 0x12,
	0x13, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e, 0x64, 0x47, 0x65, 0x74,
	0x41, 0x72, 0x67, 0x73, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41,
	0x6e, 0x64, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x47, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x34, 0x47, 0x65, 0x74, 0x12, 0x18,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x34, 0x47, 0x65, 0x74, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x34, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x47, 0x0a, 0x10, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4c, 0x65,
	0x61, 0x73, 0x65, 0x34, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x34, 0x47, 0x65, 0x74, 0x41, 0x72, 0x67,
	0x73, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4c, 0x65,
	0x61, 0x73, 0x65, 0x34, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x50, 0x0a, 0x0b,
	0x53, 0x65, 0x74, 0x75, 0x70, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x41, 0x0a, 0x14, 0x53,
	0x65, 0x74, 0x75, 0x70, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x32, 0x44, 0x61, 0x74, 0x61, 0x4e,
	0x6f, 0x64, 0x65, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x32, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x72, 0x67,
	0x73, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x69, 0x65, 0x63, 0x65, 0x30, 0x01, 0x42, 0x05,
	0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Get_proto_rawDescOnce sync.Once
	file_Get_proto_rawDescData = file_Get_proto_rawDesc
)

func file_Get_proto_rawDescGZIP() []byte {
	file_Get_proto_rawDescOnce.Do(func() {
		file_Get_proto_rawDescData = protoimpl.X.CompressGZIP(file_Get_proto_rawDescData)
	})
	return file_Get_proto_rawDescData
}

var file_Get_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_Get_proto_goTypes = []interface{}{
	(*CheckAndGetArgs)(nil),           // 0: pb.CheckAndGetArgs
	(*CheckAndGetReply)(nil),          // 1: pb.CheckAndGetReply
	(*GetDataNodes4GetArgs)(nil),      // 2: pb.GetDataNodes4GetArgs
	(*GetDataNodes4GetReply)(nil),     // 3: pb.GetDataNodes4GetReply
	(*SetupStream2DataNodeArgs)(nil),  // 4: pb.SetupStream2DataNodeArgs
	(*SetupStream2DataNodeReply)(nil), // 5: pb.SetupStream2DataNodeReply
	(*ReleaseLease4GetArgs)(nil),      // 6: pb.ReleaseLease4GetArgs
	(*ReleaseLease4GetReply)(nil),     // 7: pb.ReleaseLease4GetReply
	(*Piece)(nil),                     // 8: pb.Piece
}
var file_Get_proto_depIdxs = []int32{
	0, // 0: pb.MasterGetService.CheckAndGet:input_type -> pb.CheckAndGetArgs
	2, // 1: pb.MasterGetService.GetDataNodes4Get:input_type -> pb.GetDataNodes4GetArgs
	6, // 2: pb.MasterGetService.ReleaseLease4Get:input_type -> pb.ReleaseLease4GetArgs
	4, // 3: pb.SetupStream.SetupStream2DataNode:input_type -> pb.SetupStream2DataNodeArgs
	1, // 4: pb.MasterGetService.CheckAndGet:output_type -> pb.CheckAndGetReply
	3, // 5: pb.MasterGetService.GetDataNodes4Get:output_type -> pb.GetDataNodes4GetReply
	7, // 6: pb.MasterGetService.ReleaseLease4Get:output_type -> pb.ReleaseLease4GetReply
	8, // 7: pb.SetupStream.SetupStream2DataNode:output_type -> pb.Piece
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Get_proto_init() }
func file_Get_proto_init() {
	if File_Get_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Get_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAndGetArgs); i {
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
		file_Get_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAndGetReply); i {
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
		file_Get_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataNodes4GetArgs); i {
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
		file_Get_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataNodes4GetReply); i {
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
		file_Get_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetupStream2DataNodeArgs); i {
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
		file_Get_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetupStream2DataNodeReply); i {
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
		file_Get_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseLease4GetArgs); i {
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
		file_Get_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseLease4GetReply); i {
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
		file_Get_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Piece); i {
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
			RawDescriptor: file_Get_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_Get_proto_goTypes,
		DependencyIndexes: file_Get_proto_depIdxs,
		MessageInfos:      file_Get_proto_msgTypes,
	}.Build()
	File_Get_proto = out.File
	file_Get_proto_rawDesc = nil
	file_Get_proto_goTypes = nil
	file_Get_proto_depIdxs = nil
}
