// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.28.3
// source: envoy/config/trace/v3/opentelemetry.proto

package tracev3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
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

// Configuration for the OpenTelemetry tracer.
//
//	[#extension: envoy.tracers.opentelemetry]
//
// [#next-free-field: 6]
type OpenTelemetryConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The upstream gRPC cluster that will receive OTLP traces.
	// Note that the tracer drops traces if the server does not read data fast enough.
	// This field can be left empty to disable reporting traces to the gRPC service.
	// Only one of “grpc_service“, “http_service“ may be used.
	GrpcService *v3.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	// The upstream HTTP cluster that will receive OTLP traces.
	// This field can be left empty to disable reporting traces to the HTTP service.
	// Only one of “grpc_service“, “http_service“ may be used.
	//
	// .. note::
	//
	//	Note: The ``request_headers_to_add`` property in the OTLP HTTP exporter service
	//	does not support the :ref:`format specifier <config_access_log_format>` as used for
	//	:ref:`HTTP access logging <config_access_log>`.
	//	The values configured are added as HTTP headers on the OTLP export request
	//	without any formatting applied.
	HttpService *v3.HttpService `protobuf:"bytes,3,opt,name=http_service,json=httpService,proto3" json:"http_service,omitempty"`
	// The name for the service. This will be populated in the ResourceSpan Resource attributes.
	// If it is not provided, it will default to "unknown_service:envoy".
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// An ordered list of resource detectors
	// [#extension-category: envoy.tracers.opentelemetry.resource_detectors]
	ResourceDetectors []*v3.TypedExtensionConfig `protobuf:"bytes,4,rep,name=resource_detectors,json=resourceDetectors,proto3" json:"resource_detectors,omitempty"`
	// Specifies the sampler to be used by the OpenTelemetry tracer.
	// The configured sampler implements the Sampler interface defined by the OpenTelemetry specification.
	// This field can be left empty. In this case, the default Envoy sampling decision is used.
	//
	// See: `OpenTelemetry sampler specification <https://opentelemetry.io/docs/specs/otel/trace/sdk/#sampler>`_
	// [#extension-category: envoy.tracers.opentelemetry.samplers]
	Sampler *v3.TypedExtensionConfig `protobuf:"bytes,5,opt,name=sampler,proto3" json:"sampler,omitempty"`
}

func (x *OpenTelemetryConfig) Reset() {
	*x = OpenTelemetryConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_trace_v3_opentelemetry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenTelemetryConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenTelemetryConfig) ProtoMessage() {}

func (x *OpenTelemetryConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_trace_v3_opentelemetry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenTelemetryConfig.ProtoReflect.Descriptor instead.
func (*OpenTelemetryConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_trace_v3_opentelemetry_proto_rawDescGZIP(), []int{0}
}

func (x *OpenTelemetryConfig) GetGrpcService() *v3.GrpcService {
	if x != nil {
		return x.GrpcService
	}
	return nil
}

func (x *OpenTelemetryConfig) GetHttpService() *v3.HttpService {
	if x != nil {
		return x.HttpService
	}
	return nil
}

func (x *OpenTelemetryConfig) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *OpenTelemetryConfig) GetResourceDetectors() []*v3.TypedExtensionConfig {
	if x != nil {
		return x.ResourceDetectors
	}
	return nil
}

func (x *OpenTelemetryConfig) GetSampler() *v3.TypedExtensionConfig {
	if x != nil {
		return x.Sampler
	}
	return nil
}

var File_envoy_config_trace_v3_opentelemetry_proto protoreflect.FileDescriptor

var file_envoy_config_trace_v3_opentelemetry_proto_rawDesc = []byte{
	0x0a, 0x29, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e,
	0x76, 0x33, 0x1a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x03, 0x0a, 0x13, 0x4f, 0x70,
	0x65, 0x6e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x5b, 0x0a, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x47,
	0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x15, 0xf2, 0x98, 0xfe, 0x8f,
	0x05, 0x0f, 0x12, 0x0d, 0x6f, 0x74, 0x6c, 0x70, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x52, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b,
	0x0a, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x15, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x0f, 0x12,
	0x0d, 0x6f, 0x74, 0x6c, 0x70, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x52, 0x0b,
	0x68, 0x74, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x59,
	0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x33, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x44, 0x0a, 0x07, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x33, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x42,
	0x89, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x23, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x42,
	0x12, 0x4f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x2f, 0x76, 0x33, 0x3b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_trace_v3_opentelemetry_proto_rawDescOnce sync.Once
	file_envoy_config_trace_v3_opentelemetry_proto_rawDescData = file_envoy_config_trace_v3_opentelemetry_proto_rawDesc
)

func file_envoy_config_trace_v3_opentelemetry_proto_rawDescGZIP() []byte {
	file_envoy_config_trace_v3_opentelemetry_proto_rawDescOnce.Do(func() {
		file_envoy_config_trace_v3_opentelemetry_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_trace_v3_opentelemetry_proto_rawDescData)
	})
	return file_envoy_config_trace_v3_opentelemetry_proto_rawDescData
}

var file_envoy_config_trace_v3_opentelemetry_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_trace_v3_opentelemetry_proto_goTypes = []interface{}{
	(*OpenTelemetryConfig)(nil),     // 0: envoy.config.trace.v3.OpenTelemetryConfig
	(*v3.GrpcService)(nil),          // 1: envoy.config.core.v3.GrpcService
	(*v3.HttpService)(nil),          // 2: envoy.config.core.v3.HttpService
	(*v3.TypedExtensionConfig)(nil), // 3: envoy.config.core.v3.TypedExtensionConfig
}
var file_envoy_config_trace_v3_opentelemetry_proto_depIdxs = []int32{
	1, // 0: envoy.config.trace.v3.OpenTelemetryConfig.grpc_service:type_name -> envoy.config.core.v3.GrpcService
	2, // 1: envoy.config.trace.v3.OpenTelemetryConfig.http_service:type_name -> envoy.config.core.v3.HttpService
	3, // 2: envoy.config.trace.v3.OpenTelemetryConfig.resource_detectors:type_name -> envoy.config.core.v3.TypedExtensionConfig
	3, // 3: envoy.config.trace.v3.OpenTelemetryConfig.sampler:type_name -> envoy.config.core.v3.TypedExtensionConfig
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_envoy_config_trace_v3_opentelemetry_proto_init() }
func file_envoy_config_trace_v3_opentelemetry_proto_init() {
	if File_envoy_config_trace_v3_opentelemetry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_trace_v3_opentelemetry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenTelemetryConfig); i {
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
			RawDescriptor: file_envoy_config_trace_v3_opentelemetry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_trace_v3_opentelemetry_proto_goTypes,
		DependencyIndexes: file_envoy_config_trace_v3_opentelemetry_proto_depIdxs,
		MessageInfos:      file_envoy_config_trace_v3_opentelemetry_proto_msgTypes,
	}.Build()
	File_envoy_config_trace_v3_opentelemetry_proto = out.File
	file_envoy_config_trace_v3_opentelemetry_proto_rawDesc = nil
	file_envoy_config_trace_v3_opentelemetry_proto_goTypes = nil
	file_envoy_config_trace_v3_opentelemetry_proto_depIdxs = nil
}
