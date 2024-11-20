// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.28.3
// source: envoy/extensions/compression/gzip/compressor/v3/gzip.proto

package compressorv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
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

// All the values of this enumeration translate directly to zlib's compression strategies.
// For more information about each strategy, please refer to zlib manual.
type Gzip_CompressionStrategy int32

const (
	Gzip_DEFAULT_STRATEGY Gzip_CompressionStrategy = 0
	Gzip_FILTERED         Gzip_CompressionStrategy = 1
	Gzip_HUFFMAN_ONLY     Gzip_CompressionStrategy = 2
	Gzip_RLE              Gzip_CompressionStrategy = 3
	Gzip_FIXED            Gzip_CompressionStrategy = 4
)

// Enum value maps for Gzip_CompressionStrategy.
var (
	Gzip_CompressionStrategy_name = map[int32]string{
		0: "DEFAULT_STRATEGY",
		1: "FILTERED",
		2: "HUFFMAN_ONLY",
		3: "RLE",
		4: "FIXED",
	}
	Gzip_CompressionStrategy_value = map[string]int32{
		"DEFAULT_STRATEGY": 0,
		"FILTERED":         1,
		"HUFFMAN_ONLY":     2,
		"RLE":              3,
		"FIXED":            4,
	}
)

func (x Gzip_CompressionStrategy) Enum() *Gzip_CompressionStrategy {
	p := new(Gzip_CompressionStrategy)
	*p = x
	return p
}

func (x Gzip_CompressionStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gzip_CompressionStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_enumTypes[0].Descriptor()
}

func (Gzip_CompressionStrategy) Type() protoreflect.EnumType {
	return &file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_enumTypes[0]
}

func (x Gzip_CompressionStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gzip_CompressionStrategy.Descriptor instead.
func (Gzip_CompressionStrategy) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_rawDescGZIP(), []int{0, 0}
}

type Gzip_CompressionLevel int32

const (
	Gzip_DEFAULT_COMPRESSION Gzip_CompressionLevel = 0
	Gzip_BEST_SPEED          Gzip_CompressionLevel = 1
	Gzip_COMPRESSION_LEVEL_1 Gzip_CompressionLevel = 1
	Gzip_COMPRESSION_LEVEL_2 Gzip_CompressionLevel = 2
	Gzip_COMPRESSION_LEVEL_3 Gzip_CompressionLevel = 3
	Gzip_COMPRESSION_LEVEL_4 Gzip_CompressionLevel = 4
	Gzip_COMPRESSION_LEVEL_5 Gzip_CompressionLevel = 5
	Gzip_COMPRESSION_LEVEL_6 Gzip_CompressionLevel = 6
	Gzip_COMPRESSION_LEVEL_7 Gzip_CompressionLevel = 7
	Gzip_COMPRESSION_LEVEL_8 Gzip_CompressionLevel = 8
	Gzip_COMPRESSION_LEVEL_9 Gzip_CompressionLevel = 9
	Gzip_BEST_COMPRESSION    Gzip_CompressionLevel = 9
)

// Enum value maps for Gzip_CompressionLevel.
var (
	Gzip_CompressionLevel_name = map[int32]string{
		0: "DEFAULT_COMPRESSION",
		1: "BEST_SPEED",
		// Duplicate value: 1: "COMPRESSION_LEVEL_1",
		2: "COMPRESSION_LEVEL_2",
		3: "COMPRESSION_LEVEL_3",
		4: "COMPRESSION_LEVEL_4",
		5: "COMPRESSION_LEVEL_5",
		6: "COMPRESSION_LEVEL_6",
		7: "COMPRESSION_LEVEL_7",
		8: "COMPRESSION_LEVEL_8",
		9: "COMPRESSION_LEVEL_9",
		// Duplicate value: 9: "BEST_COMPRESSION",
	}
	Gzip_CompressionLevel_value = map[string]int32{
		"DEFAULT_COMPRESSION": 0,
		"BEST_SPEED":          1,
		"COMPRESSION_LEVEL_1": 1,
		"COMPRESSION_LEVEL_2": 2,
		"COMPRESSION_LEVEL_3": 3,
		"COMPRESSION_LEVEL_4": 4,
		"COMPRESSION_LEVEL_5": 5,
		"COMPRESSION_LEVEL_6": 6,
		"COMPRESSION_LEVEL_7": 7,
		"COMPRESSION_LEVEL_8": 8,
		"COMPRESSION_LEVEL_9": 9,
		"BEST_COMPRESSION":    9,
	}
)

func (x Gzip_CompressionLevel) Enum() *Gzip_CompressionLevel {
	p := new(Gzip_CompressionLevel)
	*p = x
	return p
}

func (x Gzip_CompressionLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gzip_CompressionLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_enumTypes[1].Descriptor()
}

func (Gzip_CompressionLevel) Type() protoreflect.EnumType {
	return &file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_enumTypes[1]
}

func (x Gzip_CompressionLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gzip_CompressionLevel.Descriptor instead.
func (Gzip_CompressionLevel) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_rawDescGZIP(), []int{0, 1}
}

// [#next-free-field: 6]
type Gzip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value from 1 to 9 that controls the amount of internal memory used by zlib. Higher values
	// use more memory, but are faster and produce better compression results. The default value is 5.
	MemoryLevel *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=memory_level,json=memoryLevel,proto3" json:"memory_level,omitempty"`
	// A value used for selecting the zlib compression level. This setting will affect speed and
	// amount of compression applied to the content. "BEST_COMPRESSION" provides higher compression
	// at the cost of higher latency and is equal to "COMPRESSION_LEVEL_9". "BEST_SPEED" provides
	// lower compression with minimum impact on response time, the same as "COMPRESSION_LEVEL_1".
	// "DEFAULT_COMPRESSION" provides an optimal result between speed and compression. According
	// to zlib's manual this level gives the same result as "COMPRESSION_LEVEL_6".
	// This field will be set to "DEFAULT_COMPRESSION" if not specified.
	CompressionLevel Gzip_CompressionLevel `protobuf:"varint,2,opt,name=compression_level,json=compressionLevel,proto3,enum=envoy.extensions.compression.gzip.compressor.v3.Gzip_CompressionLevel" json:"compression_level,omitempty"`
	// A value used for selecting the zlib compression strategy which is directly related to the
	// characteristics of the content. Most of the time "DEFAULT_STRATEGY" will be the best choice,
	// which is also the default value for the parameter, though there are situations when
	// changing this parameter might produce better results. For example, run-length encoding (RLE)
	// is typically used when the content is known for having sequences which same data occurs many
	// consecutive times. For more information about each strategy, please refer to zlib manual.
	CompressionStrategy Gzip_CompressionStrategy `protobuf:"varint,3,opt,name=compression_strategy,json=compressionStrategy,proto3,enum=envoy.extensions.compression.gzip.compressor.v3.Gzip_CompressionStrategy" json:"compression_strategy,omitempty"`
	// Value from 9 to 15 that represents the base two logarithmic of the compressor's window size.
	// Larger window results in better compression at the expense of memory usage. The default is 12
	// which will produce a 4096 bytes window. For more details about this parameter, please refer to
	// zlib manual > deflateInit2.
	WindowBits *wrapperspb.UInt32Value `protobuf:"bytes,4,opt,name=window_bits,json=windowBits,proto3" json:"window_bits,omitempty"`
	// Value for Zlib's next output buffer. If not set, defaults to 4096.
	// See https://www.zlib.net/manual.html for more details. Also see
	// https://github.com/envoyproxy/envoy/issues/8448 for context on this filter's performance.
	ChunkSize *wrapperspb.UInt32Value `protobuf:"bytes,5,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
}

func (x *Gzip) Reset() {
	*x = Gzip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gzip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gzip) ProtoMessage() {}

func (x *Gzip) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gzip.ProtoReflect.Descriptor instead.
func (*Gzip) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_rawDescGZIP(), []int{0}
}

func (x *Gzip) GetMemoryLevel() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MemoryLevel
	}
	return nil
}

func (x *Gzip) GetCompressionLevel() Gzip_CompressionLevel {
	if x != nil {
		return x.CompressionLevel
	}
	return Gzip_DEFAULT_COMPRESSION
}

func (x *Gzip) GetCompressionStrategy() Gzip_CompressionStrategy {
	if x != nil {
		return x.CompressionStrategy
	}
	return Gzip_DEFAULT_STRATEGY
}

func (x *Gzip) GetWindowBits() *wrapperspb.UInt32Value {
	if x != nil {
		return x.WindowBits
	}
	return nil
}

func (x *Gzip) GetChunkSize() *wrapperspb.UInt32Value {
	if x != nil {
		return x.ChunkSize
	}
	return nil
}

var File_envoy_extensions_compression_gzip_compressor_v3_gzip_proto protoreflect.FileDescriptor

var file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67,
	0x7a, 0x69, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2f, 0x76,
	0x33, 0x2f, 0x67, 0x7a, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x7a, 0x69, 0x70, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x07, 0x0a, 0x04, 0x47, 0x7a, 0x69, 0x70, 0x12, 0x4a,
	0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x18, 0x09, 0x28, 0x01, 0x52, 0x0b, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x7d, 0x0a, 0x11, 0x63, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x46, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x7a, 0x69, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x7a, 0x69, 0x70, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x86, 0x01, 0x0a, 0x14, 0x63, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x49, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x7a, 0x69, 0x70, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x7a, 0x69, 0x70, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x13, 0x63,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x12, 0x48, 0x0a, 0x0b, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x62, 0x69, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x18, 0x0f, 0x28, 0x09,
	0x52, 0x0a, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x42, 0x69, 0x74, 0x73, 0x12, 0x49, 0x0a, 0x0a,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0c,
	0xfa, 0x42, 0x09, 0x2a, 0x07, 0x18, 0x80, 0x80, 0x04, 0x28, 0x80, 0x20, 0x52, 0x09, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x5f, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x14,
	0x0a, 0x10, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45,
	0x47, 0x59, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x48, 0x55, 0x46, 0x46, 0x4d, 0x41, 0x4e, 0x5f, 0x4f, 0x4e,
	0x4c, 0x59, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x09, 0x0a,
	0x05, 0x46, 0x49, 0x58, 0x45, 0x44, 0x10, 0x04, 0x22, 0xb6, 0x02, 0x0a, 0x10, 0x43, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x17, 0x0a,
	0x13, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x45, 0x53, 0x54, 0x5f, 0x53,
	0x50, 0x45, 0x45, 0x44, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x31, 0x10, 0x01, 0x12,
	0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4c,
	0x45, 0x56, 0x45, 0x4c, 0x5f, 0x32, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4d, 0x50,
	0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x33, 0x10,
	0x03, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x34, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f,
	0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x35, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x36, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13,
	0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x56, 0x45,
	0x4c, 0x5f, 0x37, 0x10, 0x07, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x38, 0x10, 0x08, 0x12, 0x17,
	0x0a, 0x13, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x45,
	0x56, 0x45, 0x4c, 0x5f, 0x39, 0x10, 0x09, 0x12, 0x14, 0x0a, 0x10, 0x42, 0x45, 0x53, 0x54, 0x5f,
	0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x09, 0x1a, 0x02, 0x10,
	0x01, 0x42, 0xb9, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x3d, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x7a, 0x69, 0x70, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x09, 0x47, 0x7a, 0x69,
	0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x63, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x7a,
	0x69, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2f, 0x76, 0x33,
	0x3b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x76, 0x33, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_rawDescOnce sync.Once
	file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_rawDescData = file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_rawDesc
)

func file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_rawDescGZIP() []byte {
	file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_rawDescData)
	})
	return file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_rawDescData
}

var file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_goTypes = []interface{}{
	(Gzip_CompressionStrategy)(0),  // 0: envoy.extensions.compression.gzip.compressor.v3.Gzip.CompressionStrategy
	(Gzip_CompressionLevel)(0),     // 1: envoy.extensions.compression.gzip.compressor.v3.Gzip.CompressionLevel
	(*Gzip)(nil),                   // 2: envoy.extensions.compression.gzip.compressor.v3.Gzip
	(*wrapperspb.UInt32Value)(nil), // 3: google.protobuf.UInt32Value
}
var file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_depIdxs = []int32{
	3, // 0: envoy.extensions.compression.gzip.compressor.v3.Gzip.memory_level:type_name -> google.protobuf.UInt32Value
	1, // 1: envoy.extensions.compression.gzip.compressor.v3.Gzip.compression_level:type_name -> envoy.extensions.compression.gzip.compressor.v3.Gzip.CompressionLevel
	0, // 2: envoy.extensions.compression.gzip.compressor.v3.Gzip.compression_strategy:type_name -> envoy.extensions.compression.gzip.compressor.v3.Gzip.CompressionStrategy
	3, // 3: envoy.extensions.compression.gzip.compressor.v3.Gzip.window_bits:type_name -> google.protobuf.UInt32Value
	3, // 4: envoy.extensions.compression.gzip.compressor.v3.Gzip.chunk_size:type_name -> google.protobuf.UInt32Value
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_init() }
func file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_init() {
	if File_envoy_extensions_compression_gzip_compressor_v3_gzip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gzip); i {
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
			RawDescriptor: file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_msgTypes,
	}.Build()
	File_envoy_extensions_compression_gzip_compressor_v3_gzip_proto = out.File
	file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_rawDesc = nil
	file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_goTypes = nil
	file_envoy_extensions_compression_gzip_compressor_v3_gzip_proto_depIdxs = nil
}
