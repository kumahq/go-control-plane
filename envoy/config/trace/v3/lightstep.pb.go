// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.28.3
// source: envoy/config/trace/v3/lightstep.proto

package tracev3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Available propagation modes
type LightstepConfig_PropagationMode int32

const (
	// Propagate trace context in the single header x-ot-span-context.
	LightstepConfig_ENVOY LightstepConfig_PropagationMode = 0
	// Propagate trace context using LightStep's native format.
	LightstepConfig_LIGHTSTEP LightstepConfig_PropagationMode = 1
	// Propagate trace context using the b3 format.
	LightstepConfig_B3 LightstepConfig_PropagationMode = 2
	// Propagation trace context using the w3 trace-context standard.
	LightstepConfig_TRACE_CONTEXT LightstepConfig_PropagationMode = 3
)

// Enum value maps for LightstepConfig_PropagationMode.
var (
	LightstepConfig_PropagationMode_name = map[int32]string{
		0: "ENVOY",
		1: "LIGHTSTEP",
		2: "B3",
		3: "TRACE_CONTEXT",
	}
	LightstepConfig_PropagationMode_value = map[string]int32{
		"ENVOY":         0,
		"LIGHTSTEP":     1,
		"B3":            2,
		"TRACE_CONTEXT": 3,
	}
)

func (x LightstepConfig_PropagationMode) Enum() *LightstepConfig_PropagationMode {
	p := new(LightstepConfig_PropagationMode)
	*p = x
	return p
}

func (x LightstepConfig_PropagationMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LightstepConfig_PropagationMode) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_config_trace_v3_lightstep_proto_enumTypes[0].Descriptor()
}

func (LightstepConfig_PropagationMode) Type() protoreflect.EnumType {
	return &file_envoy_config_trace_v3_lightstep_proto_enumTypes[0]
}

func (x LightstepConfig_PropagationMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LightstepConfig_PropagationMode.Descriptor instead.
func (LightstepConfig_PropagationMode) EnumDescriptor() ([]byte, []int) {
	return file_envoy_config_trace_v3_lightstep_proto_rawDescGZIP(), []int{0, 0}
}

// Configuration for the LightStep tracer.
// [#extension: envoy.tracers.lightstep]
// [#not-implemented-hide:]
type LightstepConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The cluster manager cluster that hosts the LightStep collectors.
	CollectorCluster string `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	// File containing the access token to the `LightStep
	// <https://lightstep.com/>`_ API.
	//
	// Deprecated: Marked as deprecated in envoy/config/trace/v3/lightstep.proto.
	AccessTokenFile string `protobuf:"bytes,2,opt,name=access_token_file,json=accessTokenFile,proto3" json:"access_token_file,omitempty"`
	// Access token to the `LightStep <https://lightstep.com/>`_ API.
	AccessToken *v3.DataSource `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// Propagation modes to use by LightStep's tracer.
	PropagationModes []LightstepConfig_PropagationMode `protobuf:"varint,3,rep,packed,name=propagation_modes,json=propagationModes,proto3,enum=envoy.config.trace.v3.LightstepConfig_PropagationMode" json:"propagation_modes,omitempty"`
}

func (x *LightstepConfig) Reset() {
	*x = LightstepConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_trace_v3_lightstep_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LightstepConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LightstepConfig) ProtoMessage() {}

func (x *LightstepConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_trace_v3_lightstep_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LightstepConfig.ProtoReflect.Descriptor instead.
func (*LightstepConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_trace_v3_lightstep_proto_rawDescGZIP(), []int{0}
}

func (x *LightstepConfig) GetCollectorCluster() string {
	if x != nil {
		return x.CollectorCluster
	}
	return ""
}

// Deprecated: Marked as deprecated in envoy/config/trace/v3/lightstep.proto.
func (x *LightstepConfig) GetAccessTokenFile() string {
	if x != nil {
		return x.AccessTokenFile
	}
	return ""
}

func (x *LightstepConfig) GetAccessToken() *v3.DataSource {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

func (x *LightstepConfig) GetPropagationModes() []LightstepConfig_PropagationMode {
	if x != nil {
		return x.PropagationModes
	}
	return nil
}

var File_envoy_config_trace_v3_lightstep_proto protoreflect.FileDescriptor

var file_envoy_config_trace_v3_lightstep_proto_rawDesc = []byte{
	0x0a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x74, 0x65,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x1f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xaf, 0x03, 0x0a, 0x0f, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x74, 0x65, 0x70, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x34, 0x0a, 0x11, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x10, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x11, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0x92, 0xc7, 0x86, 0xd8, 0x04, 0x03, 0x33, 0x2e, 0x30, 0x18,
	0x01, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x72, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x70, 0x61,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x36, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74,
	0x73, 0x74, 0x65, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x61,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0x92,
	0x01, 0x07, 0x22, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x70, 0x61,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x46, 0x0a, 0x0f, 0x50,
	0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x09,
	0x0a, 0x05, 0x45, 0x4e, 0x56, 0x4f, 0x59, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x49, 0x47,
	0x48, 0x54, 0x53, 0x54, 0x45, 0x50, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x42, 0x33, 0x10, 0x02,
	0x12, 0x11, 0x0a, 0x0d, 0x54, 0x52, 0x41, 0x43, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x58,
	0x54, 0x10, 0x03, 0x3a, 0x2c, 0x9a, 0xc5, 0x88, 0x1e, 0x27, 0x0a, 0x25, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x74, 0x65, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x42, 0xb7, 0x01, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x2c, 0x12, 0x2a, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x72, 0x73, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x74, 0x65, 0x70, 0x2e, 0x76,
	0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x23,
	0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x2e, 0x76, 0x33, 0x42, 0x0e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x74, 0x65, 0x70, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x2f, 0x76, 0x33, 0x3b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_trace_v3_lightstep_proto_rawDescOnce sync.Once
	file_envoy_config_trace_v3_lightstep_proto_rawDescData = file_envoy_config_trace_v3_lightstep_proto_rawDesc
)

func file_envoy_config_trace_v3_lightstep_proto_rawDescGZIP() []byte {
	file_envoy_config_trace_v3_lightstep_proto_rawDescOnce.Do(func() {
		file_envoy_config_trace_v3_lightstep_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_trace_v3_lightstep_proto_rawDescData)
	})
	return file_envoy_config_trace_v3_lightstep_proto_rawDescData
}

var file_envoy_config_trace_v3_lightstep_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_config_trace_v3_lightstep_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_trace_v3_lightstep_proto_goTypes = []interface{}{
	(LightstepConfig_PropagationMode)(0), // 0: envoy.config.trace.v3.LightstepConfig.PropagationMode
	(*LightstepConfig)(nil),              // 1: envoy.config.trace.v3.LightstepConfig
	(*v3.DataSource)(nil),                // 2: envoy.config.core.v3.DataSource
}
var file_envoy_config_trace_v3_lightstep_proto_depIdxs = []int32{
	2, // 0: envoy.config.trace.v3.LightstepConfig.access_token:type_name -> envoy.config.core.v3.DataSource
	0, // 1: envoy.config.trace.v3.LightstepConfig.propagation_modes:type_name -> envoy.config.trace.v3.LightstepConfig.PropagationMode
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_config_trace_v3_lightstep_proto_init() }
func file_envoy_config_trace_v3_lightstep_proto_init() {
	if File_envoy_config_trace_v3_lightstep_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_trace_v3_lightstep_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LightstepConfig); i {
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
			RawDescriptor: file_envoy_config_trace_v3_lightstep_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_trace_v3_lightstep_proto_goTypes,
		DependencyIndexes: file_envoy_config_trace_v3_lightstep_proto_depIdxs,
		EnumInfos:         file_envoy_config_trace_v3_lightstep_proto_enumTypes,
		MessageInfos:      file_envoy_config_trace_v3_lightstep_proto_msgTypes,
	}.Build()
	File_envoy_config_trace_v3_lightstep_proto = out.File
	file_envoy_config_trace_v3_lightstep_proto_rawDesc = nil
	file_envoy_config_trace_v3_lightstep_proto_goTypes = nil
	file_envoy_config_trace_v3_lightstep_proto_depIdxs = nil
}
