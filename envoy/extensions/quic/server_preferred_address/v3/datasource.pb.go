// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.29.2
// source: envoy/extensions/quic/server_preferred_address/v3/datasource.proto

package server_preferred_addressv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
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

// Configuration for DataSourceServerPreferredAddressConfig.
type DataSourceServerPreferredAddressConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The IPv4 address to advertise to clients for Server Preferred Address.
	Ipv4Config *DataSourceServerPreferredAddressConfig_AddressFamilyConfig `protobuf:"bytes,1,opt,name=ipv4_config,json=ipv4Config,proto3" json:"ipv4_config,omitempty"`
	// The IPv6 address to advertise to clients for Server Preferred Address.
	Ipv6Config *DataSourceServerPreferredAddressConfig_AddressFamilyConfig `protobuf:"bytes,2,opt,name=ipv6_config,json=ipv6Config,proto3" json:"ipv6_config,omitempty"`
}

func (x *DataSourceServerPreferredAddressConfig) Reset() {
	*x = DataSourceServerPreferredAddressConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSourceServerPreferredAddressConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceServerPreferredAddressConfig) ProtoMessage() {}

func (x *DataSourceServerPreferredAddressConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceServerPreferredAddressConfig.ProtoReflect.Descriptor instead.
func (*DataSourceServerPreferredAddressConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_rawDescGZIP(), []int{0}
}

func (x *DataSourceServerPreferredAddressConfig) GetIpv4Config() *DataSourceServerPreferredAddressConfig_AddressFamilyConfig {
	if x != nil {
		return x.Ipv4Config
	}
	return nil
}

func (x *DataSourceServerPreferredAddressConfig) GetIpv6Config() *DataSourceServerPreferredAddressConfig_AddressFamilyConfig {
	if x != nil {
		return x.Ipv6Config
	}
	return nil
}

// Addresses for server preferred address for a single address family (IPv4 or IPv6).
type DataSourceServerPreferredAddressConfig_AddressFamilyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The server preferred address sent to clients. The data must contain an IP address string.
	Address *v3.DataSource `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// The server preferred address port sent to clients. The data must contain a integer port value.
	//
	// If this is not specified, the listener's port is used.
	//
	// Note: Envoy currently must receive all packets for a QUIC connection on the same port, so unless
	// :ref:`dnat_address <envoy_v3_api_field_extensions.quic.server_preferred_address.v3.DataSourceServerPreferredAddressConfig.AddressFamilyConfig.dnat_address>`
	// is configured, this must be left unset.
	Port *v3.DataSource `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	// If there is a DNAT between the client and Envoy, the address that Envoy will observe
	// server preferred address packets being sent to. If this is not specified, it is assumed
	// there is no DNAT and the server preferred address packets will be sent to the address advertised
	// to clients for server preferred address.
	DnatAddress *v3.DataSource `protobuf:"bytes,3,opt,name=dnat_address,json=dnatAddress,proto3" json:"dnat_address,omitempty"`
}

func (x *DataSourceServerPreferredAddressConfig_AddressFamilyConfig) Reset() {
	*x = DataSourceServerPreferredAddressConfig_AddressFamilyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSourceServerPreferredAddressConfig_AddressFamilyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceServerPreferredAddressConfig_AddressFamilyConfig) ProtoMessage() {}

func (x *DataSourceServerPreferredAddressConfig_AddressFamilyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceServerPreferredAddressConfig_AddressFamilyConfig.ProtoReflect.Descriptor instead.
func (*DataSourceServerPreferredAddressConfig_AddressFamilyConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DataSourceServerPreferredAddressConfig_AddressFamilyConfig) GetAddress() *v3.DataSource {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *DataSourceServerPreferredAddressConfig_AddressFamilyConfig) GetPort() *v3.DataSource {
	if x != nil {
		return x.Port
	}
	return nil
}

func (x *DataSourceServerPreferredAddressConfig_AddressFamilyConfig) GetDnatAddress() *v3.DataSource {
	if x != nil {
		return x.DnatAddress
	}
	return nil
}

var File_envoy_extensions_quic_server_preferred_address_v3_datasource_proto protoreflect.FileDescriptor

var file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_rawDesc = []byte{
	0x0a, 0x42, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x2f, 0x76, 0x33, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x31, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa3, 0x04, 0x0a, 0x26, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x8e, 0x01, 0x0a, 0x0b,
	0x69, 0x70, 0x76, 0x34, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x6d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x0a, 0x69, 0x70, 0x76, 0x34, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x8e, 0x01, 0x0a,
	0x0b, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x6d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x0a, 0x69, 0x70, 0x76, 0x36, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0xd6, 0x01,
	0x0a, 0x13, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x44, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x34, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x43, 0x0a, 0x0c, 0x64, 0x6e, 0x61, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0b, 0x64, 0x6e, 0x61, 0x74, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0xd1, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10,
	0x02, 0x0a, 0x3f, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x76, 0x33, 0x42, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x73, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x71,
	0x75, 0x69, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x33, 0x3b,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_rawDescOnce sync.Once
	file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_rawDescData = file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_rawDesc
)

func file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_rawDescGZIP() []byte {
	file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_rawDescData)
	})
	return file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_rawDescData
}

var file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_goTypes = []interface{}{
	(*DataSourceServerPreferredAddressConfig)(nil),                     // 0: envoy.extensions.quic.server_preferred_address.v3.DataSourceServerPreferredAddressConfig
	(*DataSourceServerPreferredAddressConfig_AddressFamilyConfig)(nil), // 1: envoy.extensions.quic.server_preferred_address.v3.DataSourceServerPreferredAddressConfig.AddressFamilyConfig
	(*v3.DataSource)(nil), // 2: envoy.config.core.v3.DataSource
}
var file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.quic.server_preferred_address.v3.DataSourceServerPreferredAddressConfig.ipv4_config:type_name -> envoy.extensions.quic.server_preferred_address.v3.DataSourceServerPreferredAddressConfig.AddressFamilyConfig
	1, // 1: envoy.extensions.quic.server_preferred_address.v3.DataSourceServerPreferredAddressConfig.ipv6_config:type_name -> envoy.extensions.quic.server_preferred_address.v3.DataSourceServerPreferredAddressConfig.AddressFamilyConfig
	2, // 2: envoy.extensions.quic.server_preferred_address.v3.DataSourceServerPreferredAddressConfig.AddressFamilyConfig.address:type_name -> envoy.config.core.v3.DataSource
	2, // 3: envoy.extensions.quic.server_preferred_address.v3.DataSourceServerPreferredAddressConfig.AddressFamilyConfig.port:type_name -> envoy.config.core.v3.DataSource
	2, // 4: envoy.extensions.quic.server_preferred_address.v3.DataSourceServerPreferredAddressConfig.AddressFamilyConfig.dnat_address:type_name -> envoy.config.core.v3.DataSource
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_init() }
func file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_init() {
	if File_envoy_extensions_quic_server_preferred_address_v3_datasource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSourceServerPreferredAddressConfig); i {
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
		file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSourceServerPreferredAddressConfig_AddressFamilyConfig); i {
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
			RawDescriptor: file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_msgTypes,
	}.Build()
	File_envoy_extensions_quic_server_preferred_address_v3_datasource_proto = out.File
	file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_rawDesc = nil
	file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_goTypes = nil
	file_envoy_extensions_quic_server_preferred_address_v3_datasource_proto_depIdxs = nil
}
