// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.29.2
// source: envoy/extensions/matching/input_matchers/metadata/v3/metadata.proto

package metadatav3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
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

// Metadata matcher for metadata from http matching input data.
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Metadata is matched if the value retrieved by metadata matching input is matched to this value.
	Value *v3.ValueMatcher `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// If true, the match result will be inverted.
	Invert bool `protobuf:"varint,4,opt,name=invert,proto3" json:"invert,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetValue() *v3.ValueMatcher {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Metadata) GetInvert() bool {
	if x != nil {
		return x.Invert
	}
	return false
}

var File_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto protoreflect.FileDescriptor

var file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_rawDesc = []byte{
	0x0a, 0x43, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x34, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67,
	0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x1a, 0x21, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f,
	0x76, 0x33, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x43, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x42,
	0xc5, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x42, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x69, 0x6e, 0x67, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x42, 0x0d,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x66, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x33, 0x3b, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_rawDescOnce sync.Once
	file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_rawDescData = file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_rawDesc
)

func file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_rawDescGZIP() []byte {
	file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_rawDescData)
	})
	return file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_rawDescData
}

var file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_goTypes = []interface{}{
	(*Metadata)(nil),        // 0: envoy.extensions.matching.input_matchers.metadata.v3.Metadata
	(*v3.ValueMatcher)(nil), // 1: envoy.type.matcher.v3.ValueMatcher
}
var file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.matching.input_matchers.metadata.v3.Metadata.value:type_name -> envoy.type.matcher.v3.ValueMatcher
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_init() }
func file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_init() {
	if File_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
			RawDescriptor: file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_msgTypes,
	}.Build()
	File_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto = out.File
	file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_rawDesc = nil
	file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_goTypes = nil
	file_envoy_extensions_matching_input_matchers_metadata_v3_metadata_proto_depIdxs = nil
}
