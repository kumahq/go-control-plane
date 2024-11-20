// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.28.3
// source: envoy/extensions/load_balancing_policies/maglev/v3/maglev.proto

package maglevv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/load_balancing_policies/common/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This configuration allows the built-in Maglev LB policy to be configured via the LB policy
// extension point. See the :ref:`load balancing architecture overview
// <arch_overview_load_balancing_types>` and :ref:`Maglev<arch_overview_load_balancing_types_maglev>` for more information.
type Maglev struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The table size for Maglev hashing. Maglev aims for "minimal disruption" rather than an absolute guarantee.
	// Minimal disruption means that when the set of upstream hosts change, a connection will likely be sent to the same
	// upstream as it was before. Increasing the table size reduces the amount of disruption.
	// The table size must be prime number limited to 5000011. If it is not specified, the default is 65537.
	TableSize *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=table_size,json=tableSize,proto3" json:"table_size,omitempty"`
	// Common configuration for hashing-based load balancing policies.
	ConsistentHashingLbConfig *v3.ConsistentHashingLbConfig `protobuf:"bytes,2,opt,name=consistent_hashing_lb_config,json=consistentHashingLbConfig,proto3" json:"consistent_hashing_lb_config,omitempty"`
	// Enable locality weighted load balancing for maglev lb explicitly.
	LocalityWeightedLbConfig *v3.LocalityLbConfig_LocalityWeightedLbConfig `protobuf:"bytes,3,opt,name=locality_weighted_lb_config,json=localityWeightedLbConfig,proto3" json:"locality_weighted_lb_config,omitempty"`
}

func (x *Maglev) Reset() {
	*x = Maglev{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Maglev) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Maglev) ProtoMessage() {}

func (x *Maglev) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Maglev.ProtoReflect.Descriptor instead.
func (*Maglev) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_rawDescGZIP(), []int{0}
}

func (x *Maglev) GetTableSize() *wrapperspb.UInt64Value {
	if x != nil {
		return x.TableSize
	}
	return nil
}

func (x *Maglev) GetConsistentHashingLbConfig() *v3.ConsistentHashingLbConfig {
	if x != nil {
		return x.ConsistentHashingLbConfig
	}
	return nil
}

func (x *Maglev) GetLocalityWeightedLbConfig() *v3.LocalityLbConfig_LocalityWeightedLbConfig {
	if x != nil {
		return x.LocalityWeightedLbConfig
	}
	return nil
}

var File_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto protoreflect.FileDescriptor

var file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e,
	0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x67, 0x6c, 0x65,
	0x76, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x61, 0x67, 0x6c, 0x65, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69,
	0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x6d, 0x61, 0x67, 0x6c,
	0x65, 0x76, 0x2e, 0x76, 0x33, 0x1a, 0x3f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81,
	0x03, 0x0a, 0x06, 0x4d, 0x61, 0x67, 0x6c, 0x65, 0x76, 0x12, 0x47, 0x0a, 0x0a, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0a, 0xfa, 0x42, 0x07,
	0x32, 0x05, 0x18, 0xcb, 0x96, 0xb1, 0x02, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x8e, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x62, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67,
	0x4c, 0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x19, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x4c, 0x62, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x9c, 0x01, 0x0a, 0x1b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x5f, 0x6c, 0x62, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5d, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64,
	0x4c, 0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x18, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x4c, 0x62, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x42, 0xbd, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x40, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x6d, 0x61, 0x67, 0x6c, 0x65, 0x76, 0x2e, 0x76, 0x33, 0x42,
	0x0b, 0x4d, 0x61, 0x67, 0x6c, 0x65, 0x76, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x62,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f,
	0x6d, 0x61, 0x67, 0x6c, 0x65, 0x76, 0x2f, 0x76, 0x33, 0x3b, 0x6d, 0x61, 0x67, 0x6c, 0x65, 0x76,
	0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_rawDescOnce sync.Once
	file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_rawDescData = file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_rawDesc
)

func file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_rawDescGZIP() []byte {
	file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_rawDescData)
	})
	return file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_rawDescData
}

var file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_goTypes = []interface{}{
	(*Maglev)(nil),                                       // 0: envoy.extensions.load_balancing_policies.maglev.v3.Maglev
	(*wrapperspb.UInt64Value)(nil),                       // 1: google.protobuf.UInt64Value
	(*v3.ConsistentHashingLbConfig)(nil),                 // 2: envoy.extensions.load_balancing_policies.common.v3.ConsistentHashingLbConfig
	(*v3.LocalityLbConfig_LocalityWeightedLbConfig)(nil), // 3: envoy.extensions.load_balancing_policies.common.v3.LocalityLbConfig.LocalityWeightedLbConfig
}
var file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.load_balancing_policies.maglev.v3.Maglev.table_size:type_name -> google.protobuf.UInt64Value
	2, // 1: envoy.extensions.load_balancing_policies.maglev.v3.Maglev.consistent_hashing_lb_config:type_name -> envoy.extensions.load_balancing_policies.common.v3.ConsistentHashingLbConfig
	3, // 2: envoy.extensions.load_balancing_policies.maglev.v3.Maglev.locality_weighted_lb_config:type_name -> envoy.extensions.load_balancing_policies.common.v3.LocalityLbConfig.LocalityWeightedLbConfig
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_init() }
func file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_init() {
	if File_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Maglev); i {
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
			RawDescriptor: file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_msgTypes,
	}.Build()
	File_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto = out.File
	file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_rawDesc = nil
	file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_goTypes = nil
	file_envoy_extensions_load_balancing_policies_maglev_v3_maglev_proto_depIdxs = nil
}
