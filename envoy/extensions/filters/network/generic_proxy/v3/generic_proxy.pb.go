// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.29.2
// source: envoy/extensions/filters/network/generic_proxy/v3/generic_proxy.proto

package generic_proxyv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	v32 "github.com/envoyproxy/go-control-plane/envoy/config/accesslog/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v31 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
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

// [#next-free-field: 8]
type GenericProxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The human readable prefix to use when emitting statistics.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The codec which encodes and decodes the application protocol.
	// [#extension-category: envoy.generic_proxy.codecs]
	CodecConfig *v3.TypedExtensionConfig `protobuf:"bytes,2,opt,name=codec_config,json=codecConfig,proto3" json:"codec_config,omitempty"`
	// Types that are assignable to RouteSpecifier:
	//
	//	*GenericProxy_GenericRds
	//	*GenericProxy_RouteConfig
	RouteSpecifier isGenericProxy_RouteSpecifier `protobuf_oneof:"route_specifier"`
	// A list of individual Layer-7 filters that make up the filter chain for requests made to the
	// proxy. Order matters as the filters are processed sequentially as request events
	// happen.
	// [#extension-category: envoy.generic_proxy.filters]
	Filters []*v3.TypedExtensionConfig `protobuf:"bytes,5,rep,name=filters,proto3" json:"filters,omitempty"`
	// Tracing configuration for the generic proxy.
	Tracing *v31.HttpConnectionManager_Tracing `protobuf:"bytes,6,opt,name=tracing,proto3" json:"tracing,omitempty"`
	// Configuration for :ref:`access logs <arch_overview_access_logs>` emitted by generic proxy.
	AccessLog []*v32.AccessLog `protobuf:"bytes,7,rep,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
}

func (x *GenericProxy) Reset() {
	*x = GenericProxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericProxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericProxy) ProtoMessage() {}

func (x *GenericProxy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericProxy.ProtoReflect.Descriptor instead.
func (*GenericProxy) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *GenericProxy) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (x *GenericProxy) GetCodecConfig() *v3.TypedExtensionConfig {
	if x != nil {
		return x.CodecConfig
	}
	return nil
}

func (m *GenericProxy) GetRouteSpecifier() isGenericProxy_RouteSpecifier {
	if m != nil {
		return m.RouteSpecifier
	}
	return nil
}

func (x *GenericProxy) GetGenericRds() *GenericRds {
	if x, ok := x.GetRouteSpecifier().(*GenericProxy_GenericRds); ok {
		return x.GenericRds
	}
	return nil
}

func (x *GenericProxy) GetRouteConfig() *RouteConfiguration {
	if x, ok := x.GetRouteSpecifier().(*GenericProxy_RouteConfig); ok {
		return x.RouteConfig
	}
	return nil
}

func (x *GenericProxy) GetFilters() []*v3.TypedExtensionConfig {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *GenericProxy) GetTracing() *v31.HttpConnectionManager_Tracing {
	if x != nil {
		return x.Tracing
	}
	return nil
}

func (x *GenericProxy) GetAccessLog() []*v32.AccessLog {
	if x != nil {
		return x.AccessLog
	}
	return nil
}

type isGenericProxy_RouteSpecifier interface {
	isGenericProxy_RouteSpecifier()
}

type GenericProxy_GenericRds struct {
	// The generic proxies route table will be dynamically loaded via the meta RDS API.
	GenericRds *GenericRds `protobuf:"bytes,3,opt,name=generic_rds,json=genericRds,proto3,oneof"`
}

type GenericProxy_RouteConfig struct {
	// The route table for the generic proxy is static and is specified in this property.
	RouteConfig *RouteConfiguration `protobuf:"bytes,4,opt,name=route_config,json=routeConfig,proto3,oneof"`
}

func (*GenericProxy_GenericRds) isGenericProxy_RouteSpecifier() {}

func (*GenericProxy_RouteConfig) isGenericProxy_RouteSpecifier() {}

type GenericRds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration source specifier for RDS.
	ConfigSource *v3.ConfigSource `protobuf:"bytes,1,opt,name=config_source,json=configSource,proto3" json:"config_source,omitempty"`
	// The name of the route configuration. This name will be passed to the RDS API. This allows an
	// Envoy configuration with multiple generic proxies to use different route configurations.
	RouteConfigName string `protobuf:"bytes,2,opt,name=route_config_name,json=routeConfigName,proto3" json:"route_config_name,omitempty"`
}

func (x *GenericRds) Reset() {
	*x = GenericRds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericRds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericRds) ProtoMessage() {}

func (x *GenericRds) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericRds.ProtoReflect.Descriptor instead.
func (*GenericRds) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *GenericRds) GetConfigSource() *v3.ConfigSource {
	if x != nil {
		return x.ConfigSource
	}
	return nil
}

func (x *GenericRds) GetRouteConfigName() string {
	if x != nil {
		return x.RouteConfigName
	}
	return ""
}

var File_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_rawDesc = []byte{
	0x0a, 0x45, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x76, 0x33, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x31, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69,
	0x63, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x1a, 0x29, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c,
	0x6f, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x59, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76,
	0x33, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x04, 0x0a, 0x0c, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x28, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x12, 0x57, 0x0a, 0x0c, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x0b, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x60, 0x0a, 0x0b,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x64, 0x73,
	0x48, 0x00, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x64, 0x73, 0x12, 0x6a,
	0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0b, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x44, 0x0a, 0x07, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x33, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x12, 0x74, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x5a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e,
	0x48, 0x74, 0x74, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x74,
	0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x43, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x42, 0x16, 0x0a, 0x0f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x03,
	0xf8, 0x42, 0x01, 0x22, 0x94, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52,
	0x64, 0x73, 0x12, 0x51, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x11, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0xd0, 0x01, 0xba, 0x80, 0xc8,
	0xd1, 0x06, 0x02, 0x10, 0x02, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x0a, 0x3f, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x42, 0x11,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x68, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x69, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33, 0x3b, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x76, 0x33, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_rawDescData = file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_rawDesc
)

func file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_rawDescData)
	})
	return file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_rawDescData
}

var file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_goTypes = []interface{}{
	(*GenericProxy)(nil),                      // 0: envoy.extensions.filters.network.generic_proxy.v3.GenericProxy
	(*GenericRds)(nil),                        // 1: envoy.extensions.filters.network.generic_proxy.v3.GenericRds
	(*v3.TypedExtensionConfig)(nil),           // 2: envoy.config.core.v3.TypedExtensionConfig
	(*RouteConfiguration)(nil),                // 3: envoy.extensions.filters.network.generic_proxy.v3.RouteConfiguration
	(*v31.HttpConnectionManager_Tracing)(nil), // 4: envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager.Tracing
	(*v32.AccessLog)(nil),                     // 5: envoy.config.accesslog.v3.AccessLog
	(*v3.ConfigSource)(nil),                   // 6: envoy.config.core.v3.ConfigSource
}
var file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.filters.network.generic_proxy.v3.GenericProxy.codec_config:type_name -> envoy.config.core.v3.TypedExtensionConfig
	1, // 1: envoy.extensions.filters.network.generic_proxy.v3.GenericProxy.generic_rds:type_name -> envoy.extensions.filters.network.generic_proxy.v3.GenericRds
	3, // 2: envoy.extensions.filters.network.generic_proxy.v3.GenericProxy.route_config:type_name -> envoy.extensions.filters.network.generic_proxy.v3.RouteConfiguration
	2, // 3: envoy.extensions.filters.network.generic_proxy.v3.GenericProxy.filters:type_name -> envoy.config.core.v3.TypedExtensionConfig
	4, // 4: envoy.extensions.filters.network.generic_proxy.v3.GenericProxy.tracing:type_name -> envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager.Tracing
	5, // 5: envoy.extensions.filters.network.generic_proxy.v3.GenericProxy.access_log:type_name -> envoy.config.accesslog.v3.AccessLog
	6, // 6: envoy.extensions.filters.network.generic_proxy.v3.GenericRds.config_source:type_name -> envoy.config.core.v3.ConfigSource
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_init() }
func file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_init() {
	if File_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto != nil {
		return
	}
	file_envoy_extensions_filters_network_generic_proxy_v3_route_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericProxy); i {
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
		file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericRds); i {
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
	file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GenericProxy_GenericRds)(nil),
		(*GenericProxy_RouteConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto = out.File
	file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_rawDesc = nil
	file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_goTypes = nil
	file_envoy_extensions_filters_network_generic_proxy_v3_generic_proxy_proto_depIdxs = nil
}
