//
// DISCLAIMER
//
// Copyright 2024 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: integrations/storage/v2/definition/storage.proto

package definition

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Defines Object Path/Key
type StorageV2Path struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines Object Path/Key
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *StorageV2Path) Reset() {
	*x = StorageV2Path{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageV2Path) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageV2Path) ProtoMessage() {}

func (x *StorageV2Path) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageV2Path.ProtoReflect.Descriptor instead.
func (*StorageV2Path) Descriptor() ([]byte, []int) {
	return file_integrations_storage_v2_definition_storage_proto_rawDescGZIP(), []int{0}
}

func (x *StorageV2Path) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// Defines Object Details
type StorageV2Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines Object Path/Key
	Path *StorageV2Path `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Defines Object Info
	Info *StorageV2ObjectInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *StorageV2Object) Reset() {
	*x = StorageV2Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageV2Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageV2Object) ProtoMessage() {}

func (x *StorageV2Object) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageV2Object.ProtoReflect.Descriptor instead.
func (*StorageV2Object) Descriptor() ([]byte, []int) {
	return file_integrations_storage_v2_definition_storage_proto_rawDescGZIP(), []int{1}
}

func (x *StorageV2Object) GetPath() *StorageV2Path {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *StorageV2Object) GetInfo() *StorageV2ObjectInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

// Defines Object Info
type StorageV2ObjectInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Size in bytes of the object
	Size uint64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// Timestamp of last update
	LastUpdated *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *StorageV2ObjectInfo) Reset() {
	*x = StorageV2ObjectInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageV2ObjectInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageV2ObjectInfo) ProtoMessage() {}

func (x *StorageV2ObjectInfo) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageV2ObjectInfo.ProtoReflect.Descriptor instead.
func (*StorageV2ObjectInfo) Descriptor() ([]byte, []int) {
	return file_integrations_storage_v2_definition_storage_proto_rawDescGZIP(), []int{2}
}

func (x *StorageV2ObjectInfo) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *StorageV2ObjectInfo) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

// StorageV2 ReadObject Request
type StorageV2ReadObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines Object Path/Key
	Path *StorageV2Path `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *StorageV2ReadObjectRequest) Reset() {
	*x = StorageV2ReadObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageV2ReadObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageV2ReadObjectRequest) ProtoMessage() {}

func (x *StorageV2ReadObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageV2ReadObjectRequest.ProtoReflect.Descriptor instead.
func (*StorageV2ReadObjectRequest) Descriptor() ([]byte, []int) {
	return file_integrations_storage_v2_definition_storage_proto_rawDescGZIP(), []int{3}
}

func (x *StorageV2ReadObjectRequest) GetPath() *StorageV2Path {
	if x != nil {
		return x.Path
	}
	return nil
}

// StorageV2 ReadObject Response
type StorageV2ReadObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Bytes of the object
	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *StorageV2ReadObjectResponse) Reset() {
	*x = StorageV2ReadObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageV2ReadObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageV2ReadObjectResponse) ProtoMessage() {}

func (x *StorageV2ReadObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageV2ReadObjectResponse.ProtoReflect.Descriptor instead.
func (*StorageV2ReadObjectResponse) Descriptor() ([]byte, []int) {
	return file_integrations_storage_v2_definition_storage_proto_rawDescGZIP(), []int{4}
}

func (x *StorageV2ReadObjectResponse) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

// StorageV2 WriteObject Request
type StorageV2WriteObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines Object Path/Key
	Path *StorageV2Path `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Bytes of the object
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *StorageV2WriteObjectRequest) Reset() {
	*x = StorageV2WriteObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageV2WriteObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageV2WriteObjectRequest) ProtoMessage() {}

func (x *StorageV2WriteObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageV2WriteObjectRequest.ProtoReflect.Descriptor instead.
func (*StorageV2WriteObjectRequest) Descriptor() ([]byte, []int) {
	return file_integrations_storage_v2_definition_storage_proto_rawDescGZIP(), []int{5}
}

func (x *StorageV2WriteObjectRequest) GetPath() *StorageV2Path {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *StorageV2WriteObjectRequest) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

// StorageV2 WriteObject Response
type StorageV2WriteObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Bytes Saved
	Bytes int64 `protobuf:"varint,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
	// Checksum (sha256) of the object
	Checksum string `protobuf:"bytes,2,opt,name=checksum,proto3" json:"checksum,omitempty"`
}

func (x *StorageV2WriteObjectResponse) Reset() {
	*x = StorageV2WriteObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageV2WriteObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageV2WriteObjectResponse) ProtoMessage() {}

func (x *StorageV2WriteObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageV2WriteObjectResponse.ProtoReflect.Descriptor instead.
func (*StorageV2WriteObjectResponse) Descriptor() ([]byte, []int) {
	return file_integrations_storage_v2_definition_storage_proto_rawDescGZIP(), []int{6}
}

func (x *StorageV2WriteObjectResponse) GetBytes() int64 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

func (x *StorageV2WriteObjectResponse) GetChecksum() string {
	if x != nil {
		return x.Checksum
	}
	return ""
}

// StorageV2 HeadObject Request
type StorageV2HeadObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines Object Path/Key
	Path *StorageV2Path `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *StorageV2HeadObjectRequest) Reset() {
	*x = StorageV2HeadObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageV2HeadObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageV2HeadObjectRequest) ProtoMessage() {}

func (x *StorageV2HeadObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageV2HeadObjectRequest.ProtoReflect.Descriptor instead.
func (*StorageV2HeadObjectRequest) Descriptor() ([]byte, []int) {
	return file_integrations_storage_v2_definition_storage_proto_rawDescGZIP(), []int{7}
}

func (x *StorageV2HeadObjectRequest) GetPath() *StorageV2Path {
	if x != nil {
		return x.Path
	}
	return nil
}

// StorageV2 HeadObject Response
type StorageV2HeadObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines Object Info
	Info *StorageV2ObjectInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *StorageV2HeadObjectResponse) Reset() {
	*x = StorageV2HeadObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageV2HeadObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageV2HeadObjectResponse) ProtoMessage() {}

func (x *StorageV2HeadObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageV2HeadObjectResponse.ProtoReflect.Descriptor instead.
func (*StorageV2HeadObjectResponse) Descriptor() ([]byte, []int) {
	return file_integrations_storage_v2_definition_storage_proto_rawDescGZIP(), []int{8}
}

func (x *StorageV2HeadObjectResponse) GetInfo() *StorageV2ObjectInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

// StorageV2 DeleteObject Request
type StorageV2DeleteObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines Object Path/Key
	Path *StorageV2Path `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *StorageV2DeleteObjectRequest) Reset() {
	*x = StorageV2DeleteObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageV2DeleteObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageV2DeleteObjectRequest) ProtoMessage() {}

func (x *StorageV2DeleteObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageV2DeleteObjectRequest.ProtoReflect.Descriptor instead.
func (*StorageV2DeleteObjectRequest) Descriptor() ([]byte, []int) {
	return file_integrations_storage_v2_definition_storage_proto_rawDescGZIP(), []int{9}
}

func (x *StorageV2DeleteObjectRequest) GetPath() *StorageV2Path {
	if x != nil {
		return x.Path
	}
	return nil
}

// StorageV2 DeleteObject Response
type StorageV2DeleteObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StorageV2DeleteObjectResponse) Reset() {
	*x = StorageV2DeleteObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageV2DeleteObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageV2DeleteObjectResponse) ProtoMessage() {}

func (x *StorageV2DeleteObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageV2DeleteObjectResponse.ProtoReflect.Descriptor instead.
func (*StorageV2DeleteObjectResponse) Descriptor() ([]byte, []int) {
	return file_integrations_storage_v2_definition_storage_proto_rawDescGZIP(), []int{10}
}

// StorageV2 ListObjects Request
type StorageV2ListObjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines Object Path/Key
	Path *StorageV2Path `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *StorageV2ListObjectsRequest) Reset() {
	*x = StorageV2ListObjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageV2ListObjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageV2ListObjectsRequest) ProtoMessage() {}

func (x *StorageV2ListObjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageV2ListObjectsRequest.ProtoReflect.Descriptor instead.
func (*StorageV2ListObjectsRequest) Descriptor() ([]byte, []int) {
	return file_integrations_storage_v2_definition_storage_proto_rawDescGZIP(), []int{11}
}

func (x *StorageV2ListObjectsRequest) GetPath() *StorageV2Path {
	if x != nil {
		return x.Path
	}
	return nil
}

// StorageV2 ListObjects Response
type StorageV2ListObjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of the objects
	Files []*StorageV2Object `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *StorageV2ListObjectsResponse) Reset() {
	*x = StorageV2ListObjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageV2ListObjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageV2ListObjectsResponse) ProtoMessage() {}

func (x *StorageV2ListObjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_storage_v2_definition_storage_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageV2ListObjectsResponse.ProtoReflect.Descriptor instead.
func (*StorageV2ListObjectsResponse) Descriptor() ([]byte, []int) {
	return file_integrations_storage_v2_definition_storage_proto_rawDescGZIP(), []int{12}
}

func (x *StorageV2ListObjectsResponse) GetFiles() []*StorageV2Object {
	if x != nil {
		return x.Files
	}
	return nil
}

var File_integrations_storage_v2_definition_storage_proto protoreflect.FileDescriptor

var file_integrations_storage_v2_definition_storage_proto_rawDesc = []byte{
	0x0a, 0x30, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a,
	0x0d, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x22, 0x71, 0x0a, 0x0f, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x50, 0x61, 0x74, 0x68, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x31, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x56, 0x32, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x68, 0x0a, 0x13, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x56, 0x32, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22,
	0x49, 0x0a, 0x1a, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x52, 0x65, 0x61, 0x64,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68,
	0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32,
	0x50, 0x61, 0x74, 0x68, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x33, 0x0a, 0x1b, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x22,
	0x60, 0x0a, 0x1b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73,
	0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56,
	0x32, 0x50, 0x61, 0x74, 0x68, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x22, 0x50, 0x0a, 0x1c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x22, 0x49, 0x0a, 0x1a, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32,
	0x48, 0x65, 0x61, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2b, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x56, 0x32, 0x50, 0x61, 0x74, 0x68, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x50,
	0x0a, 0x1b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x48, 0x65, 0x61, 0x64, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x68,
	0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x22, 0x4b, 0x0a, 0x1c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2b, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x56, 0x32, 0x50, 0x61, 0x74, 0x68, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x1f, 0x0a,
	0x1d, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4a,
	0x0a, 0x1b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68,
	0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32,
	0x50, 0x61, 0x74, 0x68, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x4f, 0x0a, 0x1c, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x68, 0x75, 0x74,
	0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x32, 0xe4, 0x03, 0x0a, 0x09,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x12, 0x5b, 0x0a, 0x0a, 0x52, 0x65, 0x61,
	0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x24, 0x2e, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f,
	0x77, 0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x52, 0x65, 0x61, 0x64,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x73, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x56, 0x32, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x5e, 0x0a, 0x0b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x25, 0x2e, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73,
	0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56,
	0x32, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x59, 0x0a, 0x0a, 0x48, 0x65, 0x61, 0x64, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x24, 0x2e, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x48, 0x65, 0x61, 0x64, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x68, 0x75,
	0x74, 0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x48,
	0x65, 0x61, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5f, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x26, 0x2e, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x68, 0x75, 0x74,
	0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5e, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x12, 0x25, 0x2e, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x68, 0x75, 0x74, 0x64,
	0x6f, 0x77, 0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x32, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x30, 0x01, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x2d, 0x61,
	0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x32, 0x2f,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_integrations_storage_v2_definition_storage_proto_rawDescOnce sync.Once
	file_integrations_storage_v2_definition_storage_proto_rawDescData = file_integrations_storage_v2_definition_storage_proto_rawDesc
)

func file_integrations_storage_v2_definition_storage_proto_rawDescGZIP() []byte {
	file_integrations_storage_v2_definition_storage_proto_rawDescOnce.Do(func() {
		file_integrations_storage_v2_definition_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_integrations_storage_v2_definition_storage_proto_rawDescData)
	})
	return file_integrations_storage_v2_definition_storage_proto_rawDescData
}

var file_integrations_storage_v2_definition_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_integrations_storage_v2_definition_storage_proto_goTypes = []interface{}{
	(*StorageV2Path)(nil),                 // 0: shutdown.StorageV2Path
	(*StorageV2Object)(nil),               // 1: shutdown.StorageV2Object
	(*StorageV2ObjectInfo)(nil),           // 2: shutdown.StorageV2ObjectInfo
	(*StorageV2ReadObjectRequest)(nil),    // 3: shutdown.StorageV2ReadObjectRequest
	(*StorageV2ReadObjectResponse)(nil),   // 4: shutdown.StorageV2ReadObjectResponse
	(*StorageV2WriteObjectRequest)(nil),   // 5: shutdown.StorageV2WriteObjectRequest
	(*StorageV2WriteObjectResponse)(nil),  // 6: shutdown.StorageV2WriteObjectResponse
	(*StorageV2HeadObjectRequest)(nil),    // 7: shutdown.StorageV2HeadObjectRequest
	(*StorageV2HeadObjectResponse)(nil),   // 8: shutdown.StorageV2HeadObjectResponse
	(*StorageV2DeleteObjectRequest)(nil),  // 9: shutdown.StorageV2DeleteObjectRequest
	(*StorageV2DeleteObjectResponse)(nil), // 10: shutdown.StorageV2DeleteObjectResponse
	(*StorageV2ListObjectsRequest)(nil),   // 11: shutdown.StorageV2ListObjectsRequest
	(*StorageV2ListObjectsResponse)(nil),  // 12: shutdown.StorageV2ListObjectsResponse
	(*timestamppb.Timestamp)(nil),         // 13: google.protobuf.Timestamp
}
var file_integrations_storage_v2_definition_storage_proto_depIdxs = []int32{
	0,  // 0: shutdown.StorageV2Object.path:type_name -> shutdown.StorageV2Path
	2,  // 1: shutdown.StorageV2Object.info:type_name -> shutdown.StorageV2ObjectInfo
	13, // 2: shutdown.StorageV2ObjectInfo.last_updated:type_name -> google.protobuf.Timestamp
	0,  // 3: shutdown.StorageV2ReadObjectRequest.path:type_name -> shutdown.StorageV2Path
	0,  // 4: shutdown.StorageV2WriteObjectRequest.path:type_name -> shutdown.StorageV2Path
	0,  // 5: shutdown.StorageV2HeadObjectRequest.path:type_name -> shutdown.StorageV2Path
	2,  // 6: shutdown.StorageV2HeadObjectResponse.info:type_name -> shutdown.StorageV2ObjectInfo
	0,  // 7: shutdown.StorageV2DeleteObjectRequest.path:type_name -> shutdown.StorageV2Path
	0,  // 8: shutdown.StorageV2ListObjectsRequest.path:type_name -> shutdown.StorageV2Path
	1,  // 9: shutdown.StorageV2ListObjectsResponse.files:type_name -> shutdown.StorageV2Object
	3,  // 10: shutdown.StorageV2.ReadObject:input_type -> shutdown.StorageV2ReadObjectRequest
	5,  // 11: shutdown.StorageV2.WriteObject:input_type -> shutdown.StorageV2WriteObjectRequest
	7,  // 12: shutdown.StorageV2.HeadObject:input_type -> shutdown.StorageV2HeadObjectRequest
	9,  // 13: shutdown.StorageV2.DeleteObject:input_type -> shutdown.StorageV2DeleteObjectRequest
	11, // 14: shutdown.StorageV2.ListObjects:input_type -> shutdown.StorageV2ListObjectsRequest
	4,  // 15: shutdown.StorageV2.ReadObject:output_type -> shutdown.StorageV2ReadObjectResponse
	6,  // 16: shutdown.StorageV2.WriteObject:output_type -> shutdown.StorageV2WriteObjectResponse
	8,  // 17: shutdown.StorageV2.HeadObject:output_type -> shutdown.StorageV2HeadObjectResponse
	10, // 18: shutdown.StorageV2.DeleteObject:output_type -> shutdown.StorageV2DeleteObjectResponse
	12, // 19: shutdown.StorageV2.ListObjects:output_type -> shutdown.StorageV2ListObjectsResponse
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_integrations_storage_v2_definition_storage_proto_init() }
func file_integrations_storage_v2_definition_storage_proto_init() {
	if File_integrations_storage_v2_definition_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_integrations_storage_v2_definition_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageV2Path); i {
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
		file_integrations_storage_v2_definition_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageV2Object); i {
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
		file_integrations_storage_v2_definition_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageV2ObjectInfo); i {
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
		file_integrations_storage_v2_definition_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageV2ReadObjectRequest); i {
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
		file_integrations_storage_v2_definition_storage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageV2ReadObjectResponse); i {
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
		file_integrations_storage_v2_definition_storage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageV2WriteObjectRequest); i {
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
		file_integrations_storage_v2_definition_storage_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageV2WriteObjectResponse); i {
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
		file_integrations_storage_v2_definition_storage_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageV2HeadObjectRequest); i {
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
		file_integrations_storage_v2_definition_storage_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageV2HeadObjectResponse); i {
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
		file_integrations_storage_v2_definition_storage_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageV2DeleteObjectRequest); i {
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
		file_integrations_storage_v2_definition_storage_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageV2DeleteObjectResponse); i {
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
		file_integrations_storage_v2_definition_storage_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageV2ListObjectsRequest); i {
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
		file_integrations_storage_v2_definition_storage_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageV2ListObjectsResponse); i {
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
			RawDescriptor: file_integrations_storage_v2_definition_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_integrations_storage_v2_definition_storage_proto_goTypes,
		DependencyIndexes: file_integrations_storage_v2_definition_storage_proto_depIdxs,
		MessageInfos:      file_integrations_storage_v2_definition_storage_proto_msgTypes,
	}.Build()
	File_integrations_storage_v2_definition_storage_proto = out.File
	file_integrations_storage_v2_definition_storage_proto_rawDesc = nil
	file_integrations_storage_v2_definition_storage_proto_goTypes = nil
	file_integrations_storage_v2_definition_storage_proto_depIdxs = nil
}
