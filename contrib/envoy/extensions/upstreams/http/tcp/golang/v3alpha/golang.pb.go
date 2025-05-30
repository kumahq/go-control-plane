// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.29.3
// source: contrib/envoy/extensions/upstreams/http/tcp/golang/v3alpha/golang.proto

package v3alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// [#extension-category: envoy.upstreams]
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Globally unique ID for a dynamic library file.
	LibraryId string `protobuf:"bytes,1,opt,name=library_id,json=libraryId,proto3" json:"library_id,omitempty"`
	// Path to a dynamic library implementing the
	// :repo:`HttpTcpBridge API <contrib/golang/common/go/api.HttpTcpBridge>`
	// interface.
	LibraryPath string `protobuf:"bytes,2,opt,name=library_path,json=libraryPath,proto3" json:"library_path,omitempty"`
	// Globally unique name of the Go plugin.
	//
	// This name **must** be consistent with the name registered in “tcp::RegisterHttpTcpBridgeFactoryAndConfigParser“
	PluginName string `protobuf:"bytes,3,opt,name=plugin_name,json=pluginName,proto3" json:"plugin_name,omitempty"`
	// Configuration for the Go plugin.
	//
	// .. note::
	//
	//	This configuration is only parsed in the Golang plugin, and is therefore not validated
	//	by Envoy.
	//
	//	See the :repo:`HttpTcpBridge API <contrib/golang/common/go/api/filter.go>`
	//	for more information about how the plugin's configuration data can be accessed.
	PluginConfig *anypb.Any `protobuf:"bytes,4,opt,name=plugin_config,json=pluginConfig,proto3" json:"plugin_config,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetLibraryId() string {
	if x != nil {
		return x.LibraryId
	}
	return ""
}

func (x *Config) GetLibraryPath() string {
	if x != nil {
		return x.LibraryPath
	}
	return ""
}

func (x *Config) GetPluginName() string {
	if x != nil {
		return x.PluginName
	}
	return ""
}

func (x *Config) GetPluginConfig() *anypb.Any {
	if x != nil {
		return x.PluginConfig
	}
	return nil
}

var File_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto protoreflect.FileDescriptor

var file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_rawDesc = []byte{
	0x0a, 0x47, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x75, 0x70, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x74, 0x63, 0x70, 0x2f, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x75, 0x70, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x74, 0x63, 0x70, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc1, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x26, 0x0a, 0x0a,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x0c, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x0b, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x28, 0x0a, 0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0d, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0c, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0xc4, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02,
	0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x0a, 0x40, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x74, 0x63, 0x70, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0b, 0x47, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x61, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x75, 0x70, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x74, 0x63, 0x70, 0x2f, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_rawDescOnce sync.Once
	file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_rawDescData = file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_rawDesc
)

func file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_rawDescGZIP() []byte {
	file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_rawDescOnce.Do(func() {
		file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_rawDescData = protoimpl.X.CompressGZIP(file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_rawDescData)
	})
	return file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_rawDescData
}

var file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_goTypes = []interface{}{
	(*Config)(nil),    // 0: envoy.extensions.upstreams.http.tcp.golang.v3alpha.Config
	(*anypb.Any)(nil), // 1: google.protobuf.Any
}
var file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.upstreams.http.tcp.golang.v3alpha.Config.plugin_config:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_init() }
func file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_init() {
	if File_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_goTypes,
		DependencyIndexes: file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_depIdxs,
		MessageInfos:      file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_msgTypes,
	}.Build()
	File_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto = out.File
	file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_rawDesc = nil
	file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_goTypes = nil
	file_contrib_envoy_extensions_upstreams_http_tcp_golang_v3alpha_golang_proto_depIdxs = nil
}
