// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.29.2
// source: envoy/extensions/outlier_detection_monitors/common/v3/error_types.proto

package commonv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
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

// [#protodoc-title: Outlier detection error buckets]
// Error bucket for HTTP codes.
// [#not-implemented-hide:]
type HttpErrors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Range *v3.Int32Range `protobuf:"bytes,1,opt,name=range,proto3" json:"range,omitempty"`
}

func (x *HttpErrors) Reset() {
	*x = HttpErrors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpErrors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpErrors) ProtoMessage() {}

func (x *HttpErrors) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpErrors.ProtoReflect.Descriptor instead.
func (*HttpErrors) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_rawDescGZIP(), []int{0}
}

func (x *HttpErrors) GetRange() *v3.Int32Range {
	if x != nil {
		return x.Range
	}
	return nil
}

// Error bucket for locally originated errors.
// [#not-implemented-hide:]
type LocalOriginErrors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LocalOriginErrors) Reset() {
	*x = LocalOriginErrors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalOriginErrors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalOriginErrors) ProtoMessage() {}

func (x *LocalOriginErrors) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalOriginErrors.ProtoReflect.Descriptor instead.
func (*LocalOriginErrors) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_rawDescGZIP(), []int{1}
}

// Error bucket for database errors.
// Sub-parameters may be added later, like malformed response, error on write, etc.
// [#not-implemented-hide:]
type DatabaseErrors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DatabaseErrors) Reset() {
	*x = DatabaseErrors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseErrors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseErrors) ProtoMessage() {}

func (x *DatabaseErrors) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseErrors.ProtoReflect.Descriptor instead.
func (*DatabaseErrors) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_rawDescGZIP(), []int{2}
}

// Union of possible error buckets.
// [#not-implemented-hide:]
type ErrorBuckets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of buckets "catching" HTTP codes.
	HttpErrors []*HttpErrors `protobuf:"bytes,1,rep,name=http_errors,json=httpErrors,proto3" json:"http_errors,omitempty"`
	// List of buckets "catching" locally originated errors.
	LocalOriginErrors []*LocalOriginErrors `protobuf:"bytes,2,rep,name=local_origin_errors,json=localOriginErrors,proto3" json:"local_origin_errors,omitempty"`
	// List of buckets "catching" database errors.
	DatabaseErrors []*DatabaseErrors `protobuf:"bytes,3,rep,name=database_errors,json=databaseErrors,proto3" json:"database_errors,omitempty"`
}

func (x *ErrorBuckets) Reset() {
	*x = ErrorBuckets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorBuckets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorBuckets) ProtoMessage() {}

func (x *ErrorBuckets) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorBuckets.ProtoReflect.Descriptor instead.
func (*ErrorBuckets) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_rawDescGZIP(), []int{3}
}

func (x *ErrorBuckets) GetHttpErrors() []*HttpErrors {
	if x != nil {
		return x.HttpErrors
	}
	return nil
}

func (x *ErrorBuckets) GetLocalOriginErrors() []*LocalOriginErrors {
	if x != nil {
		return x.LocalOriginErrors
	}
	return nil
}

func (x *ErrorBuckets) GetDatabaseErrors() []*DatabaseErrors {
	if x != nil {
		return x.DatabaseErrors
	}
	return nil
}

var File_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto protoreflect.FileDescriptor

var file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_rawDesc = []byte{
	0x0a, 0x47, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x35, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x6c,
	0x69, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33,
	0x1a, 0x19, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x33, 0x2f,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x0a, 0x48, 0x74,
	0x74, 0x70, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x10,
	0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x22, 0xdc, 0x02, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x12, 0x62, 0x0a, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x6c, 0x69, 0x65,
	0x72, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x48,
	0x74, 0x74, 0x70, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x52, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x78, 0x0a, 0x13, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x48, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x64, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x52, 0x11, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12,
	0x6e, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x6c,
	0x69, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x52,
	0x0e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42,
	0xc7, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x43, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x6c, 0x69,
	0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x42,
	0x0f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x65, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x6c,
	0x69, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x33,
	0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_rawDescOnce sync.Once
	file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_rawDescData = file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_rawDesc
)

func file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_rawDescGZIP() []byte {
	file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_rawDescData)
	})
	return file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_rawDescData
}

var file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_goTypes = []interface{}{
	(*HttpErrors)(nil),        // 0: envoy.extensions.outlier_detection_monitors.common.v3.HttpErrors
	(*LocalOriginErrors)(nil), // 1: envoy.extensions.outlier_detection_monitors.common.v3.LocalOriginErrors
	(*DatabaseErrors)(nil),    // 2: envoy.extensions.outlier_detection_monitors.common.v3.DatabaseErrors
	(*ErrorBuckets)(nil),      // 3: envoy.extensions.outlier_detection_monitors.common.v3.ErrorBuckets
	(*v3.Int32Range)(nil),     // 4: envoy.type.v3.Int32Range
}
var file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_depIdxs = []int32{
	4, // 0: envoy.extensions.outlier_detection_monitors.common.v3.HttpErrors.range:type_name -> envoy.type.v3.Int32Range
	0, // 1: envoy.extensions.outlier_detection_monitors.common.v3.ErrorBuckets.http_errors:type_name -> envoy.extensions.outlier_detection_monitors.common.v3.HttpErrors
	1, // 2: envoy.extensions.outlier_detection_monitors.common.v3.ErrorBuckets.local_origin_errors:type_name -> envoy.extensions.outlier_detection_monitors.common.v3.LocalOriginErrors
	2, // 3: envoy.extensions.outlier_detection_monitors.common.v3.ErrorBuckets.database_errors:type_name -> envoy.extensions.outlier_detection_monitors.common.v3.DatabaseErrors
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_init() }
func file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_init() {
	if File_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpErrors); i {
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
		file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalOriginErrors); i {
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
		file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseErrors); i {
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
		file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorBuckets); i {
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
			RawDescriptor: file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_msgTypes,
	}.Build()
	File_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto = out.File
	file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_rawDesc = nil
	file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_goTypes = nil
	file_envoy_extensions_outlier_detection_monitors_common_v3_error_types_proto_depIdxs = nil
}
