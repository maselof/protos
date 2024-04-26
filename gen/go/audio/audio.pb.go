// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: audio/audio.proto

package audiov1

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

type AudioRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AudioId       int64  `protobuf:"varint,1,opt,name=audio_id,json=audioId,proto3" json:"audio_id,omitempty"`
	AudioName     string `protobuf:"bytes,2,opt,name=audio_name,json=audioName,proto3" json:"audio_name,omitempty"`
	PerformerName string `protobuf:"bytes,3,opt,name=performer_name,json=performerName,proto3" json:"performer_name,omitempty"`
}

func (x *AudioRequest) Reset() {
	*x = AudioRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioRequest) ProtoMessage() {}

func (x *AudioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioRequest.ProtoReflect.Descriptor instead.
func (*AudioRequest) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{0}
}

func (x *AudioRequest) GetAudioId() int64 {
	if x != nil {
		return x.AudioId
	}
	return 0
}

func (x *AudioRequest) GetAudioName() string {
	if x != nil {
		return x.AudioName
	}
	return ""
}

func (x *AudioRequest) GetPerformerName() string {
	if x != nil {
		return x.PerformerName
	}
	return ""
}

type AudioFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AudioChunk []byte `protobuf:"bytes,1,opt,name=audio_chunk,json=audioChunk,proto3" json:"audio_chunk,omitempty"`
}

func (x *AudioFileRequest) Reset() {
	*x = AudioFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioFileRequest) ProtoMessage() {}

func (x *AudioFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioFileRequest.ProtoReflect.Descriptor instead.
func (*AudioFileRequest) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{1}
}

func (x *AudioFileRequest) GetAudioChunk() []byte {
	if x != nil {
		return x.AudioChunk
	}
	return nil
}

type AudioIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AudioId int64 `protobuf:"varint,1,opt,name=audio_id,json=audioId,proto3" json:"audio_id,omitempty"`
}

func (x *AudioIDRequest) Reset() {
	*x = AudioIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioIDRequest) ProtoMessage() {}

func (x *AudioIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioIDRequest.ProtoReflect.Descriptor instead.
func (*AudioIDRequest) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{2}
}

func (x *AudioIDRequest) GetAudioId() int64 {
	if x != nil {
		return x.AudioId
	}
	return 0
}

type AudioSearchNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AudioSearchNameRequest) Reset() {
	*x = AudioSearchNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioSearchNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioSearchNameRequest) ProtoMessage() {}

func (x *AudioSearchNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioSearchNameRequest.ProtoReflect.Descriptor instead.
func (*AudioSearchNameRequest) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{3}
}

func (x *AudioSearchNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AudioStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AudioChunk []byte `protobuf:"bytes,1,opt,name=audio_chunk,json=audioChunk,proto3" json:"audio_chunk,omitempty"`
}

func (x *AudioStreamResponse) Reset() {
	*x = AudioStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioStreamResponse) ProtoMessage() {}

func (x *AudioStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioStreamResponse.ProtoReflect.Descriptor instead.
func (*AudioStreamResponse) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{4}
}

func (x *AudioStreamResponse) GetAudioChunk() []byte {
	if x != nil {
		return x.AudioChunk
	}
	return nil
}

type AudioListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*AudioRequest `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *AudioListResponse) Reset() {
	*x = AudioListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioListResponse) ProtoMessage() {}

func (x *AudioListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioListResponse.ProtoReflect.Descriptor instead.
func (*AudioListResponse) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{5}
}

func (x *AudioListResponse) GetItems() []*AudioRequest {
	if x != nil {
		return x.Items
	}
	return nil
}

type AudioResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AudioId       int64  `protobuf:"varint,1,opt,name=audio_id,json=audioId,proto3" json:"audio_id,omitempty"`
	AudioName     string `protobuf:"bytes,2,opt,name=audio_name,json=audioName,proto3" json:"audio_name,omitempty"`
	PerformerName string `protobuf:"bytes,3,opt,name=performer_name,json=performerName,proto3" json:"performer_name,omitempty"`
}

func (x *AudioResponse) Reset() {
	*x = AudioResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioResponse) ProtoMessage() {}

func (x *AudioResponse) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioResponse.ProtoReflect.Descriptor instead.
func (*AudioResponse) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{6}
}

func (x *AudioResponse) GetAudioId() int64 {
	if x != nil {
		return x.AudioId
	}
	return 0
}

func (x *AudioResponse) GetAudioName() string {
	if x != nil {
		return x.AudioName
	}
	return ""
}

func (x *AudioResponse) GetPerformerName() string {
	if x != nil {
		return x.PerformerName
	}
	return ""
}

type UserRegRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password    string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	IsPerformer bool   `protobuf:"varint,3,opt,name=is_performer,json=isPerformer,proto3" json:"is_performer,omitempty"`
}

func (x *UserRegRequest) Reset() {
	*x = UserRegRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegRequest) ProtoMessage() {}

func (x *UserRegRequest) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegRequest.ProtoReflect.Descriptor instead.
func (*UserRegRequest) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{7}
}

func (x *UserRegRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserRegRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserRegRequest) GetIsPerformer() bool {
	if x != nil {
		return x.IsPerformer
	}
	return false
}

type UserLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserLogRequest) Reset() {
	*x = UserLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLogRequest) ProtoMessage() {}

func (x *UserLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLogRequest.ProtoReflect.Descriptor instead.
func (*UserLogRequest) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{8}
}

func (x *UserLogRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserLogRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	IsPerformer bool   `protobuf:"varint,3,opt,name=is_performer,json=isPerformer,proto3" json:"is_performer,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{9}
}

func (x *UserResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserResponse) GetIsPerformer() bool {
	if x != nil {
		return x.IsPerformer
	}
	return false
}

var File_audio_audio_proto protoreflect.FileDescriptor

var file_audio_audio_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x22, 0x6f, 0x0a, 0x0c, 0x41, 0x75,
	0x64, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x75,
	0x64, 0x69, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x75,
	0x64, 0x69, 0x6f, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x64, 0x69, 0x6f,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x65,
	0x72, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x10, 0x41,
	0x75, 0x64, 0x69, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x22, 0x2b, 0x0a, 0x0e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x49, 0x64, 0x22, 0x2c, 0x0a,
	0x16, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x13, 0x41,
	0x75, 0x64, 0x69, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x22, 0x3e, 0x0a, 0x11, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2e,
	0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x70, 0x0a, 0x0d, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x6b, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d,
	0x65, 0x72, 0x22, 0x48, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x4d, 0x0a, 0x0c,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x70,
	0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x69, 0x73, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x32, 0x8d, 0x03, 0x0a, 0x0c,
	0x41, 0x75, 0x64, 0x69, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0b,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x15, 0x2e, 0x6d, 0x75,
	0x73, 0x69, 0x63, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01,
	0x12, 0x38, 0x0a, 0x09, 0x4c, 0x69, 0x6b, 0x65, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x15, 0x2e,
	0x6d, 0x75, 0x73, 0x69, 0x63, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2e, 0x41, 0x75, 0x64,
	0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x15, 0x2e, 0x6d, 0x75,
	0x73, 0x69, 0x63, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01,
	0x12, 0x46, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x12,
	0x1d, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x41,
	0x75, 0x64, 0x69, 0x6f, 0x12, 0x13, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2e, 0x41, 0x75, 0x64,
	0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x75, 0x73, 0x69,
	0x63, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3a, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x13,
	0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2e, 0x41, 0x75, 0x64, 0x69,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x32, 0xb6, 0x01, 0x0a, 0x05,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x3a, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6d,
	0x75, 0x73, 0x69, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x37, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15,
	0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x6d, 0x61, 0x73, 0x65, 0x6c, 0x6f, 0x66, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_audio_audio_proto_rawDescOnce sync.Once
	file_audio_audio_proto_rawDescData = file_audio_audio_proto_rawDesc
)

func file_audio_audio_proto_rawDescGZIP() []byte {
	file_audio_audio_proto_rawDescOnce.Do(func() {
		file_audio_audio_proto_rawDescData = protoimpl.X.CompressGZIP(file_audio_audio_proto_rawDescData)
	})
	return file_audio_audio_proto_rawDescData
}

var file_audio_audio_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_audio_audio_proto_goTypes = []interface{}{
	(*AudioRequest)(nil),           // 0: music.AudioRequest
	(*AudioFileRequest)(nil),       // 1: music.AudioFileRequest
	(*AudioIDRequest)(nil),         // 2: music.AudioIDRequest
	(*AudioSearchNameRequest)(nil), // 3: music.AudioSearchNameRequest
	(*AudioStreamResponse)(nil),    // 4: music.AudioStreamResponse
	(*AudioListResponse)(nil),      // 5: music.AudioListResponse
	(*AudioResponse)(nil),          // 6: music.AudioResponse
	(*UserRegRequest)(nil),         // 7: music.UserRegRequest
	(*UserLogRequest)(nil),         // 8: music.UserLogRequest
	(*UserResponse)(nil),           // 9: music.UserResponse
}
var file_audio_audio_proto_depIdxs = []int32{
	0,  // 0: music.AudioListResponse.items:type_name -> music.AudioRequest
	2,  // 1: music.AudioService.StreamAudio:input_type -> music.AudioIDRequest
	2,  // 2: music.AudioService.LikeAudio:input_type -> music.AudioIDRequest
	2,  // 3: music.AudioService.DownloadAudio:input_type -> music.AudioIDRequest
	3,  // 4: music.AudioService.SearchAudio:input_type -> music.AudioSearchNameRequest
	0,  // 5: music.AudioService.AddAudio:input_type -> music.AudioRequest
	0,  // 6: music.AudioService.UploadAudio:input_type -> music.AudioRequest
	7,  // 7: music.Users.RegisterUser:input_type -> music.UserRegRequest
	8,  // 8: music.Users.LoginUser:input_type -> music.UserLogRequest
	8,  // 9: music.Users.DeleteUser:input_type -> music.UserLogRequest
	4,  // 10: music.AudioService.StreamAudio:output_type -> music.AudioStreamResponse
	6,  // 11: music.AudioService.LikeAudio:output_type -> music.AudioResponse
	4,  // 12: music.AudioService.DownloadAudio:output_type -> music.AudioStreamResponse
	5,  // 13: music.AudioService.SearchAudio:output_type -> music.AudioListResponse
	6,  // 14: music.AudioService.AddAudio:output_type -> music.AudioResponse
	6,  // 15: music.AudioService.UploadAudio:output_type -> music.AudioResponse
	9,  // 16: music.Users.RegisterUser:output_type -> music.UserResponse
	9,  // 17: music.Users.LoginUser:output_type -> music.UserResponse
	9,  // 18: music.Users.DeleteUser:output_type -> music.UserResponse
	10, // [10:19] is the sub-list for method output_type
	1,  // [1:10] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_audio_audio_proto_init() }
func file_audio_audio_proto_init() {
	if File_audio_audio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_audio_audio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioRequest); i {
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
		file_audio_audio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioFileRequest); i {
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
		file_audio_audio_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioIDRequest); i {
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
		file_audio_audio_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioSearchNameRequest); i {
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
		file_audio_audio_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioStreamResponse); i {
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
		file_audio_audio_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioListResponse); i {
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
		file_audio_audio_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioResponse); i {
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
		file_audio_audio_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegRequest); i {
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
		file_audio_audio_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLogRequest); i {
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
		file_audio_audio_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
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
			RawDescriptor: file_audio_audio_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_audio_audio_proto_goTypes,
		DependencyIndexes: file_audio_audio_proto_depIdxs,
		MessageInfos:      file_audio_audio_proto_msgTypes,
	}.Build()
	File_audio_audio_proto = out.File
	file_audio_audio_proto_rawDesc = nil
	file_audio_audio_proto_goTypes = nil
	file_audio_audio_proto_depIdxs = nil
}
