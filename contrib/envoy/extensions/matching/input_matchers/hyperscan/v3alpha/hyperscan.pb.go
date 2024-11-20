// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.28.3
// source: contrib/envoy/extensions/matching/input_matchers/hyperscan/v3alpha/hyperscan.proto

package v3alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
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

// `Hyperscan <https://github.com/intel/hyperscan>`_ regex matcher. The matcher uses the Hyperscan
// engine which exploits x86 SIMD instructions to accelerate matching large numbers of regular
// expressions simultaneously across streams of data.
type Hyperscan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies a set of regex expressions that the input should match on.
	Regexes []*Hyperscan_Regex `protobuf:"bytes,1,rep,name=regexes,proto3" json:"regexes,omitempty"`
}

func (x *Hyperscan) Reset() {
	*x = Hyperscan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hyperscan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hyperscan) ProtoMessage() {}

func (x *Hyperscan) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hyperscan.ProtoReflect.Descriptor instead.
func (*Hyperscan) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_rawDescGZIP(), []int{0}
}

func (x *Hyperscan) GetRegexes() []*Hyperscan_Regex {
	if x != nil {
		return x.Regexes
	}
	return nil
}

// [#next-free-field: 11]
type Hyperscan_Regex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The regex expression.
	//
	// The expression must represent only the pattern to be matched, with no delimiters or flags.
	Regex string `protobuf:"bytes,1,opt,name=regex,proto3" json:"regex,omitempty"`
	// The ID of the regex expression.
	//
	// This option is designed to be used on the sub-expressions in logical combinations.
	Id uint32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// Matching will be performed case-insensitively.
	//
	// The expression may still use PCRE tokens (notably “(?i)“ and “(?-i)“) to switch
	// case-insensitive matching on and off.
	Caseless bool `protobuf:"varint,3,opt,name=caseless,proto3" json:"caseless,omitempty"`
	// Matching a “.“ will not exclude newlines.
	DotAll bool `protobuf:"varint,4,opt,name=dot_all,json=dotAll,proto3" json:"dot_all,omitempty"`
	// “^“ and “$“ anchors match any newlines in data.
	Multiline bool `protobuf:"varint,5,opt,name=multiline,proto3" json:"multiline,omitempty"`
	// Allow expressions which can match against an empty string.
	//
	// This option instructs the compiler to allow expressions that can match against empty buffers,
	// such as “.?“, “.*“, “(a|)“. Since Hyperscan can return every possible match for an expression,
	// such expressions generally execute very slowly.
	AllowEmpty bool `protobuf:"varint,6,opt,name=allow_empty,json=allowEmpty,proto3" json:"allow_empty,omitempty"`
	// Treat the pattern as a sequence of UTF-8 characters.
	Utf8 bool `protobuf:"varint,7,opt,name=utf8,proto3" json:"utf8,omitempty"`
	// Use Unicode properties for character classes.
	//
	// This option instructs Hyperscan to use Unicode properties, rather than the default ASCII
	// interpretations, for character mnemonics like “\w“ and “\s“ as well as the POSIX character
	// classes. It is only meaningful in conjunction with “utf8“.
	Ucp bool `protobuf:"varint,8,opt,name=ucp,proto3" json:"ucp,omitempty"`
	// Logical combination.
	//
	// This option instructs Hyperscan to parse this expression as logical combination syntax.
	// Logical constraints consist of operands, operators and parentheses. The operands are
	// expression indices, and operators can be “!“, “&“ or “|“.
	Combination bool `protobuf:"varint,9,opt,name=combination,proto3" json:"combination,omitempty"`
	// Don’t do any match reporting.
	//
	// This option instructs Hyperscan to ignore match reporting for this expression. It is
	// designed to be used on the sub-expressions in logical combinations.
	Quiet bool `protobuf:"varint,10,opt,name=quiet,proto3" json:"quiet,omitempty"`
}

func (x *Hyperscan_Regex) Reset() {
	*x = Hyperscan_Regex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hyperscan_Regex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hyperscan_Regex) ProtoMessage() {}

func (x *Hyperscan_Regex) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hyperscan_Regex.ProtoReflect.Descriptor instead.
func (*Hyperscan_Regex) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Hyperscan_Regex) GetRegex() string {
	if x != nil {
		return x.Regex
	}
	return ""
}

func (x *Hyperscan_Regex) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Hyperscan_Regex) GetCaseless() bool {
	if x != nil {
		return x.Caseless
	}
	return false
}

func (x *Hyperscan_Regex) GetDotAll() bool {
	if x != nil {
		return x.DotAll
	}
	return false
}

func (x *Hyperscan_Regex) GetMultiline() bool {
	if x != nil {
		return x.Multiline
	}
	return false
}

func (x *Hyperscan_Regex) GetAllowEmpty() bool {
	if x != nil {
		return x.AllowEmpty
	}
	return false
}

func (x *Hyperscan_Regex) GetUtf8() bool {
	if x != nil {
		return x.Utf8
	}
	return false
}

func (x *Hyperscan_Regex) GetUcp() bool {
	if x != nil {
		return x.Ucp
	}
	return false
}

func (x *Hyperscan_Regex) GetCombination() bool {
	if x != nil {
		return x.Combination
	}
	return false
}

func (x *Hyperscan_Regex) GetQuiet() bool {
	if x != nil {
		return x.Quiet
	}
	return false
}

var File_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto protoreflect.FileDescriptor

var file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_rawDesc = []byte{
	0x0a, 0x52, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x69, 0x6e, 0x67, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x73, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2f, 0x76, 0x33, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2e, 0x68,
	0x79, 0x70, 0x65, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x03, 0x0a, 0x09, 0x48, 0x79, 0x70,
	0x65, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x12, 0x6f, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x65, 0x78, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x69, 0x6e, 0x67, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x73, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x48, 0x79, 0x70, 0x65, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x52,
	0x65, 0x67, 0x65, 0x78, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x07,
	0x72, 0x65, 0x67, 0x65, 0x78, 0x65, 0x73, 0x1a, 0x88, 0x02, 0x0a, 0x05, 0x52, 0x65, 0x67, 0x65,
	0x78, 0x12, 0x1d, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x73, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x63, 0x61, 0x73, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x12, 0x17, 0x0a, 0x07,
	0x64, 0x6f, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64,
	0x6f, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x74, 0x66, 0x38, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x75, 0x74, 0x66, 0x38, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x63, 0x70, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x75, 0x63, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x71, 0x75, 0x69, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x71, 0x75, 0x69,
	0x65, 0x74, 0x42, 0xcf, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x48, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e,
	0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0e, 0x48, 0x79, 0x70, 0x65, 0x72, 0x73, 0x63,
	0x61, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x69, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x69, 0x6e, 0x67, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x73, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2f, 0x76, 0x33, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_rawDescOnce sync.Once
	file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_rawDescData = file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_rawDesc
)

func file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_rawDescGZIP() []byte {
	file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_rawDescOnce.Do(func() {
		file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_rawDescData = protoimpl.X.CompressGZIP(file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_rawDescData)
	})
	return file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_rawDescData
}

var file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_goTypes = []interface{}{
	(*Hyperscan)(nil),       // 0: envoy.extensions.matching.input_matchers.hyperscan.v3alpha.Hyperscan
	(*Hyperscan_Regex)(nil), // 1: envoy.extensions.matching.input_matchers.hyperscan.v3alpha.Hyperscan.Regex
}
var file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.matching.input_matchers.hyperscan.v3alpha.Hyperscan.regexes:type_name -> envoy.extensions.matching.input_matchers.hyperscan.v3alpha.Hyperscan.Regex
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_init()
}
func file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_init() {
	if File_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hyperscan); i {
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
		file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hyperscan_Regex); i {
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
			RawDescriptor: file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_goTypes,
		DependencyIndexes: file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_depIdxs,
		MessageInfos:      file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_msgTypes,
	}.Build()
	File_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto = out.File
	file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_rawDesc = nil
	file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_goTypes = nil
	file_contrib_envoy_extensions_matching_input_matchers_hyperscan_v3alpha_hyperscan_proto_depIdxs = nil
}
