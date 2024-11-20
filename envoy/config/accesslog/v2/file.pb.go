// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.28.3
// source: envoy/config/accesslog/v2/file.proto

package accesslogv2

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Custom configuration for an :ref:`AccessLog <envoy_api_msg_config.filter.accesslog.v2.AccessLog>`
// that writes log entries directly to a file. Configures the built-in *envoy.access_loggers.file*
// AccessLog.
type FileAccessLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A path to a local file to which to write the access log entries.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Types that are assignable to AccessLogFormat:
	//
	//	*FileAccessLog_Format
	//	*FileAccessLog_JsonFormat
	//	*FileAccessLog_TypedJsonFormat
	AccessLogFormat isFileAccessLog_AccessLogFormat `protobuf_oneof:"access_log_format"`
}

func (x *FileAccessLog) Reset() {
	*x = FileAccessLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_accesslog_v2_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileAccessLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileAccessLog) ProtoMessage() {}

func (x *FileAccessLog) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_accesslog_v2_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileAccessLog.ProtoReflect.Descriptor instead.
func (*FileAccessLog) Descriptor() ([]byte, []int) {
	return file_envoy_config_accesslog_v2_file_proto_rawDescGZIP(), []int{0}
}

func (x *FileAccessLog) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (m *FileAccessLog) GetAccessLogFormat() isFileAccessLog_AccessLogFormat {
	if m != nil {
		return m.AccessLogFormat
	}
	return nil
}

func (x *FileAccessLog) GetFormat() string {
	if x, ok := x.GetAccessLogFormat().(*FileAccessLog_Format); ok {
		return x.Format
	}
	return ""
}

func (x *FileAccessLog) GetJsonFormat() *structpb.Struct {
	if x, ok := x.GetAccessLogFormat().(*FileAccessLog_JsonFormat); ok {
		return x.JsonFormat
	}
	return nil
}

func (x *FileAccessLog) GetTypedJsonFormat() *structpb.Struct {
	if x, ok := x.GetAccessLogFormat().(*FileAccessLog_TypedJsonFormat); ok {
		return x.TypedJsonFormat
	}
	return nil
}

type isFileAccessLog_AccessLogFormat interface {
	isFileAccessLog_AccessLogFormat()
}

type FileAccessLog_Format struct {
	// Access log :ref:`format string<config_access_log_format_strings>`.
	// Envoy supports :ref:`custom access log formats <config_access_log_format>` as well as a
	// :ref:`default format <config_access_log_default_format>`.
	Format string `protobuf:"bytes,2,opt,name=format,proto3,oneof"`
}

type FileAccessLog_JsonFormat struct {
	// Access log :ref:`format dictionary<config_access_log_format_dictionaries>`. All values
	// are rendered as strings.
	JsonFormat *structpb.Struct `protobuf:"bytes,3,opt,name=json_format,json=jsonFormat,proto3,oneof"`
}

type FileAccessLog_TypedJsonFormat struct {
	// Access log :ref:`format dictionary<config_access_log_format_dictionaries>`. Values are
	// rendered as strings, numbers, or boolean values as appropriate. Nested JSON objects may
	// be produced by some command operators (e.g.FILTER_STATE or DYNAMIC_METADATA). See the
	// documentation for a specific command operator for details.
	TypedJsonFormat *structpb.Struct `protobuf:"bytes,4,opt,name=typed_json_format,json=typedJsonFormat,proto3,oneof"`
}

func (*FileAccessLog_Format) isFileAccessLog_AccessLogFormat() {}

func (*FileAccessLog_JsonFormat) isFileAccessLog_AccessLogFormat() {}

func (*FileAccessLog_TypedJsonFormat) isFileAccessLog_AccessLogFormat() {}

var File_envoy_config_accesslog_v2_file_proto protoreflect.FileDescriptor

var file_envoy_config_accesslog_v2_file_proto_rawDesc = []byte{
	0x0a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x76,
	0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x01, 0x0a, 0x0d, 0x46, 0x69, 0x6c, 0x65,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x12, 0x3a, 0x0a, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x00,
	0x52, 0x0a, 0x6a, 0x73, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x45, 0x0a, 0x11,
	0x74, 0x79, 0x70, 0x65, 0x64, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x48, 0x00, 0x52, 0x0f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x4a, 0x73, 0x6f, 0x6e, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x42, 0x13, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f,
	0x67, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x42, 0xbb, 0x01, 0xf2, 0x98, 0xfe, 0x8f, 0x05,
	0x29, 0x12, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02,
	0x10, 0x01, 0x0a, 0x27, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x32, 0x42, 0x09, 0x46, 0x69, 0x6c,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x32, 0x3b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x6c, 0x6f, 0x67, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_accesslog_v2_file_proto_rawDescOnce sync.Once
	file_envoy_config_accesslog_v2_file_proto_rawDescData = file_envoy_config_accesslog_v2_file_proto_rawDesc
)

func file_envoy_config_accesslog_v2_file_proto_rawDescGZIP() []byte {
	file_envoy_config_accesslog_v2_file_proto_rawDescOnce.Do(func() {
		file_envoy_config_accesslog_v2_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_accesslog_v2_file_proto_rawDescData)
	})
	return file_envoy_config_accesslog_v2_file_proto_rawDescData
}

var file_envoy_config_accesslog_v2_file_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_accesslog_v2_file_proto_goTypes = []interface{}{
	(*FileAccessLog)(nil),   // 0: envoy.config.accesslog.v2.FileAccessLog
	(*structpb.Struct)(nil), // 1: google.protobuf.Struct
}
var file_envoy_config_accesslog_v2_file_proto_depIdxs = []int32{
	1, // 0: envoy.config.accesslog.v2.FileAccessLog.json_format:type_name -> google.protobuf.Struct
	1, // 1: envoy.config.accesslog.v2.FileAccessLog.typed_json_format:type_name -> google.protobuf.Struct
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_config_accesslog_v2_file_proto_init() }
func file_envoy_config_accesslog_v2_file_proto_init() {
	if File_envoy_config_accesslog_v2_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_accesslog_v2_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileAccessLog); i {
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
	file_envoy_config_accesslog_v2_file_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FileAccessLog_Format)(nil),
		(*FileAccessLog_JsonFormat)(nil),
		(*FileAccessLog_TypedJsonFormat)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_accesslog_v2_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_accesslog_v2_file_proto_goTypes,
		DependencyIndexes: file_envoy_config_accesslog_v2_file_proto_depIdxs,
		MessageInfos:      file_envoy_config_accesslog_v2_file_proto_msgTypes,
	}.Build()
	File_envoy_config_accesslog_v2_file_proto = out.File
	file_envoy_config_accesslog_v2_file_proto_rawDesc = nil
	file_envoy_config_accesslog_v2_file_proto_goTypes = nil
	file_envoy_config_accesslog_v2_file_proto_depIdxs = nil
}
