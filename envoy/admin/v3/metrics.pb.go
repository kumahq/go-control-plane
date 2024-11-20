// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.28.3
// source: envoy/admin/v3/metrics.proto

package adminv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
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

type SimpleMetric_Type int32

const (
	SimpleMetric_COUNTER SimpleMetric_Type = 0
	SimpleMetric_GAUGE   SimpleMetric_Type = 1
)

// Enum value maps for SimpleMetric_Type.
var (
	SimpleMetric_Type_name = map[int32]string{
		0: "COUNTER",
		1: "GAUGE",
	}
	SimpleMetric_Type_value = map[string]int32{
		"COUNTER": 0,
		"GAUGE":   1,
	}
)

func (x SimpleMetric_Type) Enum() *SimpleMetric_Type {
	p := new(SimpleMetric_Type)
	*p = x
	return p
}

func (x SimpleMetric_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SimpleMetric_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_admin_v3_metrics_proto_enumTypes[0].Descriptor()
}

func (SimpleMetric_Type) Type() protoreflect.EnumType {
	return &file_envoy_admin_v3_metrics_proto_enumTypes[0]
}

func (x SimpleMetric_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SimpleMetric_Type.Descriptor instead.
func (SimpleMetric_Type) EnumDescriptor() ([]byte, []int) {
	return file_envoy_admin_v3_metrics_proto_rawDescGZIP(), []int{0, 0}
}

// Proto representation of an Envoy Counter or Gauge value.
type SimpleMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of the metric represented.
	Type SimpleMetric_Type `protobuf:"varint,1,opt,name=type,proto3,enum=envoy.admin.v3.SimpleMetric_Type" json:"type,omitempty"`
	// Current metric value.
	Value uint64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	// Name of the metric.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SimpleMetric) Reset() {
	*x = SimpleMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_admin_v3_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleMetric) ProtoMessage() {}

func (x *SimpleMetric) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_admin_v3_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleMetric.ProtoReflect.Descriptor instead.
func (*SimpleMetric) Descriptor() ([]byte, []int) {
	return file_envoy_admin_v3_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *SimpleMetric) GetType() SimpleMetric_Type {
	if x != nil {
		return x.Type
	}
	return SimpleMetric_COUNTER
}

func (x *SimpleMetric) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *SimpleMetric) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_envoy_admin_v3_metrics_proto protoreflect.FileDescriptor

var file_envoy_admin_v3_metrics_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x33,
	0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x33, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb8, 0x01, 0x0a, 0x0c, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x12, 0x35, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x33,
	0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x1e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f,
	0x55, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x41, 0x55, 0x47, 0x45,
	0x10, 0x01, 0x3a, 0x27, 0x9a, 0xc5, 0x88, 0x1e, 0x22, 0x0a, 0x20, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x42, 0x75, 0xba, 0x80, 0xc8,
	0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x1c, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x33, 0x42, 0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x33, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_admin_v3_metrics_proto_rawDescOnce sync.Once
	file_envoy_admin_v3_metrics_proto_rawDescData = file_envoy_admin_v3_metrics_proto_rawDesc
)

func file_envoy_admin_v3_metrics_proto_rawDescGZIP() []byte {
	file_envoy_admin_v3_metrics_proto_rawDescOnce.Do(func() {
		file_envoy_admin_v3_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_admin_v3_metrics_proto_rawDescData)
	})
	return file_envoy_admin_v3_metrics_proto_rawDescData
}

var file_envoy_admin_v3_metrics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_admin_v3_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_admin_v3_metrics_proto_goTypes = []interface{}{
	(SimpleMetric_Type)(0), // 0: envoy.admin.v3.SimpleMetric.Type
	(*SimpleMetric)(nil),   // 1: envoy.admin.v3.SimpleMetric
}
var file_envoy_admin_v3_metrics_proto_depIdxs = []int32{
	0, // 0: envoy.admin.v3.SimpleMetric.type:type_name -> envoy.admin.v3.SimpleMetric.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_admin_v3_metrics_proto_init() }
func file_envoy_admin_v3_metrics_proto_init() {
	if File_envoy_admin_v3_metrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_admin_v3_metrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleMetric); i {
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
			RawDescriptor: file_envoy_admin_v3_metrics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_admin_v3_metrics_proto_goTypes,
		DependencyIndexes: file_envoy_admin_v3_metrics_proto_depIdxs,
		EnumInfos:         file_envoy_admin_v3_metrics_proto_enumTypes,
		MessageInfos:      file_envoy_admin_v3_metrics_proto_msgTypes,
	}.Build()
	File_envoy_admin_v3_metrics_proto = out.File
	file_envoy_admin_v3_metrics_proto_rawDesc = nil
	file_envoy_admin_v3_metrics_proto_goTypes = nil
	file_envoy_admin_v3_metrics_proto_depIdxs = nil
}
