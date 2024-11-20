// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.28.3
// source: contrib/envoy/extensions/compression/qatzip/compressor/v3alpha/qatzip.proto

package v3alpha

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

type Qatzip_HardwareBufferSize int32

const (
	Qatzip_DEFAULT Qatzip_HardwareBufferSize = 0
	Qatzip_SZ_4K   Qatzip_HardwareBufferSize = 1
	Qatzip_SZ_8K   Qatzip_HardwareBufferSize = 2
	Qatzip_SZ_32K  Qatzip_HardwareBufferSize = 3
	Qatzip_SZ_64K  Qatzip_HardwareBufferSize = 4
	Qatzip_SZ_128K Qatzip_HardwareBufferSize = 5
	Qatzip_SZ_512K Qatzip_HardwareBufferSize = 6
)

// Enum value maps for Qatzip_HardwareBufferSize.
var (
	Qatzip_HardwareBufferSize_name = map[int32]string{
		0: "DEFAULT",
		1: "SZ_4K",
		2: "SZ_8K",
		3: "SZ_32K",
		4: "SZ_64K",
		5: "SZ_128K",
		6: "SZ_512K",
	}
	Qatzip_HardwareBufferSize_value = map[string]int32{
		"DEFAULT": 0,
		"SZ_4K":   1,
		"SZ_8K":   2,
		"SZ_32K":  3,
		"SZ_64K":  4,
		"SZ_128K": 5,
		"SZ_512K": 6,
	}
)

func (x Qatzip_HardwareBufferSize) Enum() *Qatzip_HardwareBufferSize {
	p := new(Qatzip_HardwareBufferSize)
	*p = x
	return p
}

func (x Qatzip_HardwareBufferSize) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Qatzip_HardwareBufferSize) Descriptor() protoreflect.EnumDescriptor {
	return file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_enumTypes[0].Descriptor()
}

func (Qatzip_HardwareBufferSize) Type() protoreflect.EnumType {
	return &file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_enumTypes[0]
}

func (x Qatzip_HardwareBufferSize) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Qatzip_HardwareBufferSize.Descriptor instead.
func (Qatzip_HardwareBufferSize) EnumDescriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_rawDescGZIP(), []int{0, 0}
}

// [#next-free-field: 6]
type Qatzip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value from 1 to 9 that controls the main compression speed-density lever.
	// The higher quality, the slower compression. The default value is 1.
	CompressionLevel *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=compression_level,json=compressionLevel,proto3" json:"compression_level,omitempty"`
	// A size of qat hardware buffer. This field will be set to "DEFAULT" if not specified.
	HardwareBufferSize Qatzip_HardwareBufferSize `protobuf:"varint,2,opt,name=hardware_buffer_size,json=hardwareBufferSize,proto3,enum=envoy.extensions.compression.qatzip.compressor.v3alpha.Qatzip_HardwareBufferSize" json:"hardware_buffer_size,omitempty"`
	// Threshold of compression service’s input size for software failover.
	// If the size of input request less than the threshold, qatzip will route the request to software
	// compressor. The default value is 1024. The maximum value is 512*1024.
	InputSizeThreshold *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=input_size_threshold,json=inputSizeThreshold,proto3" json:"input_size_threshold,omitempty"`
	// A size of stream buffer. The default value is 128 * 1024. The maximum value is 2*1024*1024 -
	// 5*1024
	StreamBufferSize *wrapperspb.UInt32Value `protobuf:"bytes,4,opt,name=stream_buffer_size,json=streamBufferSize,proto3" json:"stream_buffer_size,omitempty"`
	// Value for compressor's next output buffer. If not set, defaults to 4096.
	ChunkSize *wrapperspb.UInt32Value `protobuf:"bytes,5,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
}

func (x *Qatzip) Reset() {
	*x = Qatzip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Qatzip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Qatzip) ProtoMessage() {}

func (x *Qatzip) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Qatzip.ProtoReflect.Descriptor instead.
func (*Qatzip) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_rawDescGZIP(), []int{0}
}

func (x *Qatzip) GetCompressionLevel() *wrapperspb.UInt32Value {
	if x != nil {
		return x.CompressionLevel
	}
	return nil
}

func (x *Qatzip) GetHardwareBufferSize() Qatzip_HardwareBufferSize {
	if x != nil {
		return x.HardwareBufferSize
	}
	return Qatzip_DEFAULT
}

func (x *Qatzip) GetInputSizeThreshold() *wrapperspb.UInt32Value {
	if x != nil {
		return x.InputSizeThreshold
	}
	return nil
}

func (x *Qatzip) GetStreamBufferSize() *wrapperspb.UInt32Value {
	if x != nil {
		return x.StreamBufferSize
	}
	return nil
}

func (x *Qatzip) GetChunkSize() *wrapperspb.UInt32Value {
	if x != nil {
		return x.ChunkSize
	}
	return nil
}

var File_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto protoreflect.FileDescriptor

var file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x71, 0x61, 0x74, 0x7a, 0x69, 0x70, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2f, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x71, 0x61, 0x74, 0x7a, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x36, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x71, 0x61, 0x74, 0x7a,
	0x69, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x33,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x04,
	0x0a, 0x06, 0x51, 0x61, 0x74, 0x7a, 0x69, 0x70, 0x12, 0x54, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x18, 0x09, 0x28, 0x01, 0x52, 0x10, 0x63, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x8d,
	0x01, 0x0a, 0x14, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x51, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x71, 0x61, 0x74,
	0x7a, 0x69, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x76,
	0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x51, 0x61, 0x74, 0x7a, 0x69, 0x70, 0x2e, 0x48, 0x61,
	0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x12, 0x68, 0x61, 0x72, 0x64,
	0x77, 0x61, 0x72, 0x65, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x5c,
	0x0a, 0x14, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x74, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0c, 0xfa, 0x42, 0x09, 0x2a,
	0x07, 0x18, 0x80, 0x80, 0x20, 0x28, 0x80, 0x01, 0x52, 0x12, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x53,
	0x69, 0x7a, 0x65, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x58, 0x0a, 0x12,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0c, 0xfa, 0x42, 0x09, 0x2a, 0x07, 0x18, 0x80, 0xd8,
	0x7f, 0x28, 0x80, 0x08, 0x52, 0x10, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0c, 0xfa, 0x42, 0x09, 0x2a, 0x07, 0x18,
	0x80, 0x80, 0x04, 0x28, 0x80, 0x20, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a,
	0x65, 0x22, 0x69, 0x0a, 0x12, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x42, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55,
	0x4c, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x5a, 0x5f, 0x34, 0x4b, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x53, 0x5a, 0x5f, 0x38, 0x4b, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x5a,
	0x5f, 0x33, 0x32, 0x4b, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x5a, 0x5f, 0x36, 0x34, 0x4b,
	0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x5a, 0x5f, 0x31, 0x32, 0x38, 0x4b, 0x10, 0x05, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x5a, 0x5f, 0x35, 0x31, 0x32, 0x4b, 0x10, 0x06, 0x42, 0xc4, 0x01, 0xba,
	0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x44, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x71, 0x61, 0x74, 0x7a, 0x69, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0b, 0x51,
	0x61, 0x74, 0x7a, 0x69, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x65, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x71, 0x61, 0x74, 0x7a, 0x69, 0x70,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2f, 0x76, 0x33, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_rawDescOnce sync.Once
	file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_rawDescData = file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_rawDesc
)

func file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_rawDescGZIP() []byte {
	file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_rawDescOnce.Do(func() {
		file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_rawDescData = protoimpl.X.CompressGZIP(file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_rawDescData)
	})
	return file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_rawDescData
}

var file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_goTypes = []interface{}{
	(Qatzip_HardwareBufferSize)(0), // 0: envoy.extensions.compression.qatzip.compressor.v3alpha.Qatzip.HardwareBufferSize
	(*Qatzip)(nil),                 // 1: envoy.extensions.compression.qatzip.compressor.v3alpha.Qatzip
	(*wrapperspb.UInt32Value)(nil), // 2: google.protobuf.UInt32Value
}
var file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.compression.qatzip.compressor.v3alpha.Qatzip.compression_level:type_name -> google.protobuf.UInt32Value
	0, // 1: envoy.extensions.compression.qatzip.compressor.v3alpha.Qatzip.hardware_buffer_size:type_name -> envoy.extensions.compression.qatzip.compressor.v3alpha.Qatzip.HardwareBufferSize
	2, // 2: envoy.extensions.compression.qatzip.compressor.v3alpha.Qatzip.input_size_threshold:type_name -> google.protobuf.UInt32Value
	2, // 3: envoy.extensions.compression.qatzip.compressor.v3alpha.Qatzip.stream_buffer_size:type_name -> google.protobuf.UInt32Value
	2, // 4: envoy.extensions.compression.qatzip.compressor.v3alpha.Qatzip.chunk_size:type_name -> google.protobuf.UInt32Value
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_init() }
func file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_init() {
	if File_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Qatzip); i {
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
			RawDescriptor: file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_goTypes,
		DependencyIndexes: file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_depIdxs,
		EnumInfos:         file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_enumTypes,
		MessageInfos:      file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_msgTypes,
	}.Build()
	File_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto = out.File
	file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_rawDesc = nil
	file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_goTypes = nil
	file_contrib_envoy_extensions_compression_qatzip_compressor_v3alpha_qatzip_proto_depIdxs = nil
}
