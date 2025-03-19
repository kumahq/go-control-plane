// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.29.3
// source: envoy/extensions/http/ext_proc/response_processors/save_processing_response/v3/save_processing_response.proto

package save_processing_responsev3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
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

// Extension to save the :ref:`response
// <envoy_v3_api_msg_service.ext_proc.v3.ProcessingResponse>` from the external processor as
// filter state with name
// "envoy.http.ext_proc.response_processors.save_processing_response[.:ref:`filter_state_name_suffix
// <envoy_v3_api_field_extensions.http.ext_proc.response_processors.save_processing_response.v3.SaveProcessingResponse.filter_state_name>`].
// This extension supports saving of request and response headers and trailers,
// and immediate response.
//
// .. note::
//
//	Response processors are currently in alpha.
//
// [#next-free-field: 7]
type SaveProcessingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The default filter state name is
	// "envoy.http.ext_proc.response_processors.save_processing_response".
	// If defined, “filter_state_name_suffix“ is appended to this.
	// For example, setting “filter_state_name_suffix“ to "xyz" will set the
	// filter state name to "envoy.http.ext_proc.response_processors.save_processing_response.xyz"
	FilterStateNameSuffix string `protobuf:"bytes,1,opt,name=filter_state_name_suffix,json=filterStateNameSuffix,proto3" json:"filter_state_name_suffix,omitempty"`
	// Save the response to filter state when :ref:`request_headers
	// <envoy_v3_api_field_service.ext_proc.v3.ProcessingResponse.request_headers>` is set.
	SaveRequestHeaders *SaveProcessingResponse_SaveOptions `protobuf:"bytes,2,opt,name=save_request_headers,json=saveRequestHeaders,proto3" json:"save_request_headers,omitempty"`
	// Save the response to filter state when :ref:`response_headers
	// <envoy_v3_api_field_service.ext_proc.v3.ProcessingResponse.response_headers>` is set.
	SaveResponseHeaders *SaveProcessingResponse_SaveOptions `protobuf:"bytes,3,opt,name=save_response_headers,json=saveResponseHeaders,proto3" json:"save_response_headers,omitempty"`
	// Save the response to filter state when :ref:`request_trailers
	// <envoy_v3_api_field_service.ext_proc.v3.ProcessingResponse.request_trailers>` is set.
	SaveRequestTrailers *SaveProcessingResponse_SaveOptions `protobuf:"bytes,4,opt,name=save_request_trailers,json=saveRequestTrailers,proto3" json:"save_request_trailers,omitempty"`
	// Save the response to filter state when :ref:`response_trailers
	// <envoy_v3_api_field_service.ext_proc.v3.ProcessingResponse.response_trailers>` is set.
	SaveResponseTrailers *SaveProcessingResponse_SaveOptions `protobuf:"bytes,5,opt,name=save_response_trailers,json=saveResponseTrailers,proto3" json:"save_response_trailers,omitempty"`
	// Save the response to filter state when :ref:`immediate_response
	// <envoy_v3_api_field_service.ext_proc.v3.ProcessingResponse.immediate_response>` is set.
	SaveImmediateResponse *SaveProcessingResponse_SaveOptions `protobuf:"bytes,6,opt,name=save_immediate_response,json=saveImmediateResponse,proto3" json:"save_immediate_response,omitempty"`
}

func (x *SaveProcessingResponse) Reset() {
	*x = SaveProcessingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveProcessingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveProcessingResponse) ProtoMessage() {}

func (x *SaveProcessingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveProcessingResponse.ProtoReflect.Descriptor instead.
func (*SaveProcessingResponse) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_rawDescGZIP(), []int{0}
}

func (x *SaveProcessingResponse) GetFilterStateNameSuffix() string {
	if x != nil {
		return x.FilterStateNameSuffix
	}
	return ""
}

func (x *SaveProcessingResponse) GetSaveRequestHeaders() *SaveProcessingResponse_SaveOptions {
	if x != nil {
		return x.SaveRequestHeaders
	}
	return nil
}

func (x *SaveProcessingResponse) GetSaveResponseHeaders() *SaveProcessingResponse_SaveOptions {
	if x != nil {
		return x.SaveResponseHeaders
	}
	return nil
}

func (x *SaveProcessingResponse) GetSaveRequestTrailers() *SaveProcessingResponse_SaveOptions {
	if x != nil {
		return x.SaveRequestTrailers
	}
	return nil
}

func (x *SaveProcessingResponse) GetSaveResponseTrailers() *SaveProcessingResponse_SaveOptions {
	if x != nil {
		return x.SaveResponseTrailers
	}
	return nil
}

func (x *SaveProcessingResponse) GetSaveImmediateResponse() *SaveProcessingResponse_SaveOptions {
	if x != nil {
		return x.SaveImmediateResponse
	}
	return nil
}

type SaveProcessingResponse_SaveOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether or not to save the response for the response type.
	SaveResponse bool `protobuf:"varint,1,opt,name=save_response,json=saveResponse,proto3" json:"save_response,omitempty"`
	// When true, saves the response if there was an error when processing
	// the response from the external processor.
	SaveOnError bool `protobuf:"varint,2,opt,name=save_on_error,json=saveOnError,proto3" json:"save_on_error,omitempty"`
}

func (x *SaveProcessingResponse_SaveOptions) Reset() {
	*x = SaveProcessingResponse_SaveOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveProcessingResponse_SaveOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveProcessingResponse_SaveOptions) ProtoMessage() {}

func (x *SaveProcessingResponse_SaveOptions) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveProcessingResponse_SaveOptions.ProtoReflect.Descriptor instead.
func (*SaveProcessingResponse_SaveOptions) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SaveProcessingResponse_SaveOptions) GetSaveResponse() bool {
	if x != nil {
		return x.SaveResponse
	}
	return false
}

func (x *SaveProcessingResponse_SaveOptions) GetSaveOnError() bool {
	if x != nil {
		return x.SaveOnError
	}
	return false
}

var File_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto protoreflect.FileDescriptor

var file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_rawDesc = []byte{
	0x0a, 0x6d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63,
	0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0x73, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x76, 0x33,
	0x2f, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x4e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x2e,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x73, 0x2e, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x33, 0x1a,
	0x1f, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfa, 0x07, 0x0a, 0x16, 0x53, 0x61, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x18, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f,
	0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x75, 0x66,
	0x66, 0x69, 0x78, 0x12, 0xa4, 0x01, 0x0a, 0x14, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x72, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x63, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x2e, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x76, 0x33, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x12, 0x73, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0xa6, 0x01, 0x0a, 0x15, 0x73,
	0x61, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x72, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x2e,
	0x73, 0x61, 0x76, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x61, 0x76, 0x65,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x13,
	0x73, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x12, 0xa6, 0x01, 0x0a, 0x15, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x72, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x65, 0x78, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x63, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x2e, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x61, 0x76, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x13, 0x73, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x12, 0xa8, 0x01, 0x0a,
	0x16, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74,
	0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x72, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f,
	0x72, 0x73, 0x2e, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53,
	0x61, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x14, 0x73, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54,
	0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x12, 0xaa, 0x01, 0x0a, 0x17, 0x73, 0x61, 0x76, 0x65,
	0x5f, 0x69, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x72, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x2e, 0x73,
	0x61, 0x76, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x15, 0x73,
	0x61, 0x76, 0x65, 0x49, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x56, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x61, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x61, 0x76, 0x65,
	0x5f, 0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x73, 0x61, 0x76, 0x65, 0x4f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0xa0, 0x02, 0xba,
	0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x0a,
	0x5c, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x73, 0x2e, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x1b, 0x53,
	0x61, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x90, 0x01, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x65, 0x78, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2f, 0x76, 0x33, 0x3b, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x76, 0x33, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_rawDescOnce sync.Once
	file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_rawDescData = file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_rawDesc
)

func file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_rawDescGZIP() []byte {
	file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_rawDescData)
	})
	return file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_rawDescData
}

var file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_goTypes = []interface{}{
	(*SaveProcessingResponse)(nil),             // 0: envoy.extensions.http.ext_proc.response_processors.save_processing_response.v3.SaveProcessingResponse
	(*SaveProcessingResponse_SaveOptions)(nil), // 1: envoy.extensions.http.ext_proc.response_processors.save_processing_response.v3.SaveProcessingResponse.SaveOptions
}
var file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.http.ext_proc.response_processors.save_processing_response.v3.SaveProcessingResponse.save_request_headers:type_name -> envoy.extensions.http.ext_proc.response_processors.save_processing_response.v3.SaveProcessingResponse.SaveOptions
	1, // 1: envoy.extensions.http.ext_proc.response_processors.save_processing_response.v3.SaveProcessingResponse.save_response_headers:type_name -> envoy.extensions.http.ext_proc.response_processors.save_processing_response.v3.SaveProcessingResponse.SaveOptions
	1, // 2: envoy.extensions.http.ext_proc.response_processors.save_processing_response.v3.SaveProcessingResponse.save_request_trailers:type_name -> envoy.extensions.http.ext_proc.response_processors.save_processing_response.v3.SaveProcessingResponse.SaveOptions
	1, // 3: envoy.extensions.http.ext_proc.response_processors.save_processing_response.v3.SaveProcessingResponse.save_response_trailers:type_name -> envoy.extensions.http.ext_proc.response_processors.save_processing_response.v3.SaveProcessingResponse.SaveOptions
	1, // 4: envoy.extensions.http.ext_proc.response_processors.save_processing_response.v3.SaveProcessingResponse.save_immediate_response:type_name -> envoy.extensions.http.ext_proc.response_processors.save_processing_response.v3.SaveProcessingResponse.SaveOptions
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_init()
}
func file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_init() {
	if File_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveProcessingResponse); i {
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
		file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveProcessingResponse_SaveOptions); i {
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
			RawDescriptor: file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_msgTypes,
	}.Build()
	File_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto = out.File
	file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_rawDesc = nil
	file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_goTypes = nil
	file_envoy_extensions_http_ext_proc_response_processors_save_processing_response_v3_save_processing_response_proto_depIdxs = nil
}
