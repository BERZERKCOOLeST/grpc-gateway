// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: internal/descriptor/openapiconfig/openapiconfig.proto

package openapiconfig

import (
	options "github.com/BERZERKCOOLeST/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

// OpenAPIFileOption represents OpenAPI options on a file
type OpenAPIFileOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File   string           `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	Option *options.Swagger `protobuf:"bytes,2,opt,name=option,proto3" json:"option,omitempty"`
}

func (x *OpenAPIFileOption) Reset() {
	*x = OpenAPIFileOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAPIFileOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAPIFileOption) ProtoMessage() {}

func (x *OpenAPIFileOption) ProtoReflect() protoreflect.Message {
	mi := &file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAPIFileOption.ProtoReflect.Descriptor instead.
func (*OpenAPIFileOption) Descriptor() ([]byte, []int) {
	return file_internal_descriptor_openapiconfig_openapiconfig_proto_rawDescGZIP(), []int{0}
}

func (x *OpenAPIFileOption) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *OpenAPIFileOption) GetOption() *options.Swagger {
	if x != nil {
		return x.Option
	}
	return nil
}

// OpenAPIMethodOption represents OpenAPI options on a method
type OpenAPIMethodOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method string             `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Option *options.Operation `protobuf:"bytes,2,opt,name=option,proto3" json:"option,omitempty"`
}

func (x *OpenAPIMethodOption) Reset() {
	*x = OpenAPIMethodOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAPIMethodOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAPIMethodOption) ProtoMessage() {}

func (x *OpenAPIMethodOption) ProtoReflect() protoreflect.Message {
	mi := &file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAPIMethodOption.ProtoReflect.Descriptor instead.
func (*OpenAPIMethodOption) Descriptor() ([]byte, []int) {
	return file_internal_descriptor_openapiconfig_openapiconfig_proto_rawDescGZIP(), []int{1}
}

func (x *OpenAPIMethodOption) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *OpenAPIMethodOption) GetOption() *options.Operation {
	if x != nil {
		return x.Option
	}
	return nil
}

// OpenAPIMessageOption represents OpenAPI options on a message
type OpenAPIMessageOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string          `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Option  *options.Schema `protobuf:"bytes,2,opt,name=option,proto3" json:"option,omitempty"`
}

func (x *OpenAPIMessageOption) Reset() {
	*x = OpenAPIMessageOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAPIMessageOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAPIMessageOption) ProtoMessage() {}

func (x *OpenAPIMessageOption) ProtoReflect() protoreflect.Message {
	mi := &file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAPIMessageOption.ProtoReflect.Descriptor instead.
func (*OpenAPIMessageOption) Descriptor() ([]byte, []int) {
	return file_internal_descriptor_openapiconfig_openapiconfig_proto_rawDescGZIP(), []int{2}
}

func (x *OpenAPIMessageOption) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OpenAPIMessageOption) GetOption() *options.Schema {
	if x != nil {
		return x.Option
	}
	return nil
}

// OpenAPIServiceOption represents OpenAPI options on a service
type OpenAPIServiceOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string       `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"` // ex: Service
	Option  *options.Tag `protobuf:"bytes,2,opt,name=option,proto3" json:"option,omitempty"`
}

func (x *OpenAPIServiceOption) Reset() {
	*x = OpenAPIServiceOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAPIServiceOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAPIServiceOption) ProtoMessage() {}

func (x *OpenAPIServiceOption) ProtoReflect() protoreflect.Message {
	mi := &file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAPIServiceOption.ProtoReflect.Descriptor instead.
func (*OpenAPIServiceOption) Descriptor() ([]byte, []int) {
	return file_internal_descriptor_openapiconfig_openapiconfig_proto_rawDescGZIP(), []int{3}
}

func (x *OpenAPIServiceOption) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *OpenAPIServiceOption) GetOption() *options.Tag {
	if x != nil {
		return x.Option
	}
	return nil
}

// OpenAPIFieldOption represents OpenAPI options on a field
type OpenAPIFieldOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field  string              `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Option *options.JSONSchema `protobuf:"bytes,2,opt,name=option,proto3" json:"option,omitempty"`
}

func (x *OpenAPIFieldOption) Reset() {
	*x = OpenAPIFieldOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAPIFieldOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAPIFieldOption) ProtoMessage() {}

func (x *OpenAPIFieldOption) ProtoReflect() protoreflect.Message {
	mi := &file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAPIFieldOption.ProtoReflect.Descriptor instead.
func (*OpenAPIFieldOption) Descriptor() ([]byte, []int) {
	return file_internal_descriptor_openapiconfig_openapiconfig_proto_rawDescGZIP(), []int{4}
}

func (x *OpenAPIFieldOption) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *OpenAPIFieldOption) GetOption() *options.JSONSchema {
	if x != nil {
		return x.Option
	}
	return nil
}

// OpenAPIOptions represents OpenAPI protobuf options
type OpenAPIOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File    []*OpenAPIFileOption    `protobuf:"bytes,1,rep,name=file,proto3" json:"file,omitempty"`
	Method  []*OpenAPIMethodOption  `protobuf:"bytes,2,rep,name=method,proto3" json:"method,omitempty"`
	Message []*OpenAPIMessageOption `protobuf:"bytes,3,rep,name=message,proto3" json:"message,omitempty"`
	Service []*OpenAPIServiceOption `protobuf:"bytes,4,rep,name=service,proto3" json:"service,omitempty"`
	Field   []*OpenAPIFieldOption   `protobuf:"bytes,5,rep,name=field,proto3" json:"field,omitempty"`
}

func (x *OpenAPIOptions) Reset() {
	*x = OpenAPIOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAPIOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAPIOptions) ProtoMessage() {}

func (x *OpenAPIOptions) ProtoReflect() protoreflect.Message {
	mi := &file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAPIOptions.ProtoReflect.Descriptor instead.
func (*OpenAPIOptions) Descriptor() ([]byte, []int) {
	return file_internal_descriptor_openapiconfig_openapiconfig_proto_rawDescGZIP(), []int{5}
}

func (x *OpenAPIOptions) GetFile() []*OpenAPIFileOption {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *OpenAPIOptions) GetMethod() []*OpenAPIMethodOption {
	if x != nil {
		return x.Method
	}
	return nil
}

func (x *OpenAPIOptions) GetMessage() []*OpenAPIMessageOption {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *OpenAPIOptions) GetService() []*OpenAPIServiceOption {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *OpenAPIOptions) GetField() []*OpenAPIFieldOption {
	if x != nil {
		return x.Field
	}
	return nil
}

// OpenAPIConfig represents a set of OpenAPI options
type OpenAPIConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpenapiOptions *OpenAPIOptions `protobuf:"bytes,1,opt,name=openapi_options,json=openapiOptions,proto3" json:"openapi_options,omitempty"`
}

func (x *OpenAPIConfig) Reset() {
	*x = OpenAPIConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAPIConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAPIConfig) ProtoMessage() {}

func (x *OpenAPIConfig) ProtoReflect() protoreflect.Message {
	mi := &file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAPIConfig.ProtoReflect.Descriptor instead.
func (*OpenAPIConfig) Descriptor() ([]byte, []int) {
	return file_internal_descriptor_openapiconfig_openapiconfig_proto_rawDescGZIP(), []int{6}
}

func (x *OpenAPIConfig) GetOpenapiOptions() *OpenAPIOptions {
	if x != nil {
		return x.OpenapiOptions
	}
	return nil
}

var File_internal_descriptor_openapiconfig_openapiconfig_proto protoreflect.FileDescriptor

var file_internal_descriptor_openapiconfig_openapiconfig_proto_rawDesc = []byte{
	0x0a, 0x35, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x50, 0x49,
	0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x4a,
	0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x77, 0x61, 0x67, 0x67,
	0x65, 0x72, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7b, 0x0a, 0x13, 0x4f, 0x70,
	0x65, 0x6e, 0x41, 0x50, 0x49, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x4c, 0x0a, 0x06, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f,
	0x67, 0x65, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7b, 0x0a, 0x14, 0x4f, 0x70, 0x65, 0x6e, 0x41,
	0x50, 0x49, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f,
	0x67, 0x65, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x06, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x78, 0x0a, 0x14, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x50, 0x49, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e,
	0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x79,
	0x0a, 0x12, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x50, 0x49, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x4d, 0x0a, 0x06, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2e, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4a, 0x53, 0x4f, 0x4e, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xde, 0x03, 0x0a, 0x0e, 0x4f, 0x70,
	0x65, 0x6e, 0x41, 0x50, 0x49, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x55, 0x0a, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x6e,
	0x41, 0x50, 0x49, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x5b, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x50, 0x49, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x5e, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x44, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x50, 0x49, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x5e, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x44, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x58, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x42, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x50, 0x49, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x78, 0x0a, 0x0d, 0x4f, 0x70,
	0x65, 0x6e, 0x41, 0x50, 0x49, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x67, 0x0a, 0x0f, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x50, 0x49, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76,
	0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_descriptor_openapiconfig_openapiconfig_proto_rawDescOnce sync.Once
	file_internal_descriptor_openapiconfig_openapiconfig_proto_rawDescData = file_internal_descriptor_openapiconfig_openapiconfig_proto_rawDesc
)

func file_internal_descriptor_openapiconfig_openapiconfig_proto_rawDescGZIP() []byte {
	file_internal_descriptor_openapiconfig_openapiconfig_proto_rawDescOnce.Do(func() {
		file_internal_descriptor_openapiconfig_openapiconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_descriptor_openapiconfig_openapiconfig_proto_rawDescData)
	})
	return file_internal_descriptor_openapiconfig_openapiconfig_proto_rawDescData
}

var file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_internal_descriptor_openapiconfig_openapiconfig_proto_goTypes = []interface{}{
	(*OpenAPIFileOption)(nil),    // 0: grpc.gateway.internal.descriptor.openapiconfig.OpenAPIFileOption
	(*OpenAPIMethodOption)(nil),  // 1: grpc.gateway.internal.descriptor.openapiconfig.OpenAPIMethodOption
	(*OpenAPIMessageOption)(nil), // 2: grpc.gateway.internal.descriptor.openapiconfig.OpenAPIMessageOption
	(*OpenAPIServiceOption)(nil), // 3: grpc.gateway.internal.descriptor.openapiconfig.OpenAPIServiceOption
	(*OpenAPIFieldOption)(nil),   // 4: grpc.gateway.internal.descriptor.openapiconfig.OpenAPIFieldOption
	(*OpenAPIOptions)(nil),       // 5: grpc.gateway.internal.descriptor.openapiconfig.OpenAPIOptions
	(*OpenAPIConfig)(nil),        // 6: grpc.gateway.internal.descriptor.openapiconfig.OpenAPIConfig
	(*options.Swagger)(nil),      // 7: grpc.gateway.protoc_gen_openapiv2.options.Swagger
	(*options.Operation)(nil),    // 8: grpc.gateway.protoc_gen_openapiv2.options.Operation
	(*options.Schema)(nil),       // 9: grpc.gateway.protoc_gen_openapiv2.options.Schema
	(*options.Tag)(nil),          // 10: grpc.gateway.protoc_gen_openapiv2.options.Tag
	(*options.JSONSchema)(nil),   // 11: grpc.gateway.protoc_gen_openapiv2.options.JSONSchema
}
var file_internal_descriptor_openapiconfig_openapiconfig_proto_depIdxs = []int32{
	7,  // 0: grpc.gateway.internal.descriptor.openapiconfig.OpenAPIFileOption.option:type_name -> grpc.gateway.protoc_gen_openapiv2.options.Swagger
	8,  // 1: grpc.gateway.internal.descriptor.openapiconfig.OpenAPIMethodOption.option:type_name -> grpc.gateway.protoc_gen_openapiv2.options.Operation
	9,  // 2: grpc.gateway.internal.descriptor.openapiconfig.OpenAPIMessageOption.option:type_name -> grpc.gateway.protoc_gen_openapiv2.options.Schema
	10, // 3: grpc.gateway.internal.descriptor.openapiconfig.OpenAPIServiceOption.option:type_name -> grpc.gateway.protoc_gen_openapiv2.options.Tag
	11, // 4: grpc.gateway.internal.descriptor.openapiconfig.OpenAPIFieldOption.option:type_name -> grpc.gateway.protoc_gen_openapiv2.options.JSONSchema
	0,  // 5: grpc.gateway.internal.descriptor.openapiconfig.OpenAPIOptions.file:type_name -> grpc.gateway.internal.descriptor.openapiconfig.OpenAPIFileOption
	1,  // 6: grpc.gateway.internal.descriptor.openapiconfig.OpenAPIOptions.method:type_name -> grpc.gateway.internal.descriptor.openapiconfig.OpenAPIMethodOption
	2,  // 7: grpc.gateway.internal.descriptor.openapiconfig.OpenAPIOptions.message:type_name -> grpc.gateway.internal.descriptor.openapiconfig.OpenAPIMessageOption
	3,  // 8: grpc.gateway.internal.descriptor.openapiconfig.OpenAPIOptions.service:type_name -> grpc.gateway.internal.descriptor.openapiconfig.OpenAPIServiceOption
	4,  // 9: grpc.gateway.internal.descriptor.openapiconfig.OpenAPIOptions.field:type_name -> grpc.gateway.internal.descriptor.openapiconfig.OpenAPIFieldOption
	5,  // 10: grpc.gateway.internal.descriptor.openapiconfig.OpenAPIConfig.openapi_options:type_name -> grpc.gateway.internal.descriptor.openapiconfig.OpenAPIOptions
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_internal_descriptor_openapiconfig_openapiconfig_proto_init() }
func file_internal_descriptor_openapiconfig_openapiconfig_proto_init() {
	if File_internal_descriptor_openapiconfig_openapiconfig_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAPIFileOption); i {
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
		file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAPIMethodOption); i {
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
		file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAPIMessageOption); i {
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
		file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAPIServiceOption); i {
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
		file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAPIFieldOption); i {
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
		file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAPIOptions); i {
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
		file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAPIConfig); i {
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
			RawDescriptor: file_internal_descriptor_openapiconfig_openapiconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_descriptor_openapiconfig_openapiconfig_proto_goTypes,
		DependencyIndexes: file_internal_descriptor_openapiconfig_openapiconfig_proto_depIdxs,
		MessageInfos:      file_internal_descriptor_openapiconfig_openapiconfig_proto_msgTypes,
	}.Build()
	File_internal_descriptor_openapiconfig_openapiconfig_proto = out.File
	file_internal_descriptor_openapiconfig_openapiconfig_proto_rawDesc = nil
	file_internal_descriptor_openapiconfig_openapiconfig_proto_goTypes = nil
	file_internal_descriptor_openapiconfig_openapiconfig_proto_depIdxs = nil
}
