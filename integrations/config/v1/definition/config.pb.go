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
// source: integrations/config/v1/definition/config.proto

package definition

import (
	definition "github.com/arangodb/kube-arangodb/integrations/shared/v1/definition"
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

// ConfigV1 Modules Call Response
type ConfigV1ModulesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of registered modules
	Modules []string `protobuf:"bytes,1,rep,name=modules,proto3" json:"modules,omitempty"`
}

func (x *ConfigV1ModulesResponse) Reset() {
	*x = ConfigV1ModulesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_config_v1_definition_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigV1ModulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigV1ModulesResponse) ProtoMessage() {}

func (x *ConfigV1ModulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_config_v1_definition_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigV1ModulesResponse.ProtoReflect.Descriptor instead.
func (*ConfigV1ModulesResponse) Descriptor() ([]byte, []int) {
	return file_integrations_config_v1_definition_config_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigV1ModulesResponse) GetModules() []string {
	if x != nil {
		return x.Modules
	}
	return nil
}

// ConfigV1 ModuleDetails Call Request
type ConfigV1ModuleDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the module
	Module string `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	// Define if checksum of module should be returned
	Checksum *bool `protobuf:"varint,2,opt,name=checksum,proto3,oneof" json:"checksum,omitempty"`
}

func (x *ConfigV1ModuleDetailsRequest) Reset() {
	*x = ConfigV1ModuleDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_config_v1_definition_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigV1ModuleDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigV1ModuleDetailsRequest) ProtoMessage() {}

func (x *ConfigV1ModuleDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_config_v1_definition_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigV1ModuleDetailsRequest.ProtoReflect.Descriptor instead.
func (*ConfigV1ModuleDetailsRequest) Descriptor() ([]byte, []int) {
	return file_integrations_config_v1_definition_config_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigV1ModuleDetailsRequest) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *ConfigV1ModuleDetailsRequest) GetChecksum() bool {
	if x != nil && x.Checksum != nil {
		return *x.Checksum
	}
	return false
}

// ConfigV1 ModuleDetails Call Response
type ConfigV1ModuleDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the module
	Module string `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	// List of the files
	Files []*ConfigV1File `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
	// Sha256Sum of the module (if requested)
	Checksum *string `protobuf:"bytes,3,opt,name=checksum,proto3,oneof" json:"checksum,omitempty"`
}

func (x *ConfigV1ModuleDetailsResponse) Reset() {
	*x = ConfigV1ModuleDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_config_v1_definition_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigV1ModuleDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigV1ModuleDetailsResponse) ProtoMessage() {}

func (x *ConfigV1ModuleDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_config_v1_definition_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigV1ModuleDetailsResponse.ProtoReflect.Descriptor instead.
func (*ConfigV1ModuleDetailsResponse) Descriptor() ([]byte, []int) {
	return file_integrations_config_v1_definition_config_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigV1ModuleDetailsResponse) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *ConfigV1ModuleDetailsResponse) GetFiles() []*ConfigV1File {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *ConfigV1ModuleDetailsResponse) GetChecksum() string {
	if x != nil && x.Checksum != nil {
		return *x.Checksum
	}
	return ""
}

// ConfigV1 ModuleDetails Call Request
type ConfigV1FileDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the module
	Module string `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	// Name of the file
	File string `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	// Define if checksum of module should be returned
	Checksum *bool `protobuf:"varint,3,opt,name=checksum,proto3,oneof" json:"checksum,omitempty"`
}

func (x *ConfigV1FileDetailsRequest) Reset() {
	*x = ConfigV1FileDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_config_v1_definition_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigV1FileDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigV1FileDetailsRequest) ProtoMessage() {}

func (x *ConfigV1FileDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_config_v1_definition_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigV1FileDetailsRequest.ProtoReflect.Descriptor instead.
func (*ConfigV1FileDetailsRequest) Descriptor() ([]byte, []int) {
	return file_integrations_config_v1_definition_config_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigV1FileDetailsRequest) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *ConfigV1FileDetailsRequest) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *ConfigV1FileDetailsRequest) GetChecksum() bool {
	if x != nil && x.Checksum != nil {
		return *x.Checksum
	}
	return false
}

// ConfigV1 ModuleDetails Call Response
type ConfigV1FileDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the module
	Module string `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	// Spec of the file
	File *ConfigV1File `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *ConfigV1FileDetailsResponse) Reset() {
	*x = ConfigV1FileDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_config_v1_definition_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigV1FileDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigV1FileDetailsResponse) ProtoMessage() {}

func (x *ConfigV1FileDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_config_v1_definition_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigV1FileDetailsResponse.ProtoReflect.Descriptor instead.
func (*ConfigV1FileDetailsResponse) Descriptor() ([]byte, []int) {
	return file_integrations_config_v1_definition_config_proto_rawDescGZIP(), []int{4}
}

func (x *ConfigV1FileDetailsResponse) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *ConfigV1FileDetailsResponse) GetFile() *ConfigV1File {
	if x != nil {
		return x.File
	}
	return nil
}

// Information about configuration file
type ConfigV1File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Relative path of the config file
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Size of the config file in bytes
	Size int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	// Sha256Sum of the file (if requested)
	Checksum *string `protobuf:"bytes,3,opt,name=checksum,proto3,oneof" json:"checksum,omitempty"`
	// Timestamp of the file creation
	Created *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	// Timestamp of the file update
	Updated *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *ConfigV1File) Reset() {
	*x = ConfigV1File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_config_v1_definition_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigV1File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigV1File) ProtoMessage() {}

func (x *ConfigV1File) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_config_v1_definition_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigV1File.ProtoReflect.Descriptor instead.
func (*ConfigV1File) Descriptor() ([]byte, []int) {
	return file_integrations_config_v1_definition_config_proto_rawDescGZIP(), []int{5}
}

func (x *ConfigV1File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ConfigV1File) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ConfigV1File) GetChecksum() string {
	if x != nil && x.Checksum != nil {
		return *x.Checksum
	}
	return ""
}

func (x *ConfigV1File) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *ConfigV1File) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

var File_integrations_config_v1_definition_config_proto protoreflect.FileDescriptor

var file_integrations_config_v1_definition_config_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x56, 0x31, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x64, 0x0a,
	0x1c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x31, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x22, 0x91, 0x01, 0x0a, 0x1d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x31,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2a, 0x0a,
	0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x31, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x08, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x22, 0x76, 0x0a, 0x1a, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x56, 0x31, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x1f, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x88,
	0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x22,
	0x5f, 0x0a, 0x1b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x31, 0x46, 0x69, 0x6c, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x56, 0x31, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x22, 0xd0, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x31, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x32, 0xfb, 0x01, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x31,
	0x12, 0x39, 0x0a, 0x07, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x0d, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x31, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0d, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x24, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x31, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x56, 0x31, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0b, 0x46, 0x69, 0x6c,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x31, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x31, 0x46, 0x69,
	0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x2d, 0x61, 0x72,
	0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_integrations_config_v1_definition_config_proto_rawDescOnce sync.Once
	file_integrations_config_v1_definition_config_proto_rawDescData = file_integrations_config_v1_definition_config_proto_rawDesc
)

func file_integrations_config_v1_definition_config_proto_rawDescGZIP() []byte {
	file_integrations_config_v1_definition_config_proto_rawDescOnce.Do(func() {
		file_integrations_config_v1_definition_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_integrations_config_v1_definition_config_proto_rawDescData)
	})
	return file_integrations_config_v1_definition_config_proto_rawDescData
}

var file_integrations_config_v1_definition_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_integrations_config_v1_definition_config_proto_goTypes = []interface{}{
	(*ConfigV1ModulesResponse)(nil),       // 0: config.ConfigV1ModulesResponse
	(*ConfigV1ModuleDetailsRequest)(nil),  // 1: config.ConfigV1ModuleDetailsRequest
	(*ConfigV1ModuleDetailsResponse)(nil), // 2: config.ConfigV1ModuleDetailsResponse
	(*ConfigV1FileDetailsRequest)(nil),    // 3: config.ConfigV1FileDetailsRequest
	(*ConfigV1FileDetailsResponse)(nil),   // 4: config.ConfigV1FileDetailsResponse
	(*ConfigV1File)(nil),                  // 5: config.ConfigV1File
	(*timestamppb.Timestamp)(nil),         // 6: google.protobuf.Timestamp
	(*definition.Empty)(nil),              // 7: shared.Empty
}
var file_integrations_config_v1_definition_config_proto_depIdxs = []int32{
	5, // 0: config.ConfigV1ModuleDetailsResponse.files:type_name -> config.ConfigV1File
	5, // 1: config.ConfigV1FileDetailsResponse.file:type_name -> config.ConfigV1File
	6, // 2: config.ConfigV1File.created:type_name -> google.protobuf.Timestamp
	6, // 3: config.ConfigV1File.updated:type_name -> google.protobuf.Timestamp
	7, // 4: config.ConfigV1.Modules:input_type -> shared.Empty
	1, // 5: config.ConfigV1.ModuleDetails:input_type -> config.ConfigV1ModuleDetailsRequest
	3, // 6: config.ConfigV1.FileDetails:input_type -> config.ConfigV1FileDetailsRequest
	0, // 7: config.ConfigV1.Modules:output_type -> config.ConfigV1ModulesResponse
	2, // 8: config.ConfigV1.ModuleDetails:output_type -> config.ConfigV1ModuleDetailsResponse
	4, // 9: config.ConfigV1.FileDetails:output_type -> config.ConfigV1FileDetailsResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_integrations_config_v1_definition_config_proto_init() }
func file_integrations_config_v1_definition_config_proto_init() {
	if File_integrations_config_v1_definition_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_integrations_config_v1_definition_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigV1ModulesResponse); i {
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
		file_integrations_config_v1_definition_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigV1ModuleDetailsRequest); i {
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
		file_integrations_config_v1_definition_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigV1ModuleDetailsResponse); i {
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
		file_integrations_config_v1_definition_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigV1FileDetailsRequest); i {
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
		file_integrations_config_v1_definition_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigV1FileDetailsResponse); i {
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
		file_integrations_config_v1_definition_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigV1File); i {
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
	file_integrations_config_v1_definition_config_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_integrations_config_v1_definition_config_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_integrations_config_v1_definition_config_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_integrations_config_v1_definition_config_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_integrations_config_v1_definition_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_integrations_config_v1_definition_config_proto_goTypes,
		DependencyIndexes: file_integrations_config_v1_definition_config_proto_depIdxs,
		MessageInfos:      file_integrations_config_v1_definition_config_proto_msgTypes,
	}.Build()
	File_integrations_config_v1_definition_config_proto = out.File
	file_integrations_config_v1_definition_config_proto_rawDesc = nil
	file_integrations_config_v1_definition_config_proto_goTypes = nil
	file_integrations_config_v1_definition_config_proto_depIdxs = nil
}
