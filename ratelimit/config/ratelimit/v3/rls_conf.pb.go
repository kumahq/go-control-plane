// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: ratelimit/config/ratelimit/v3/rls_conf.proto

package ratelimitv3

import (
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

// Identifies the unit of of time for rate limit.
type RateLimitUnit int32

const (
	// The time unit is not known.
	RateLimitUnit_UNKNOWN RateLimitUnit = 0
	// The time unit representing a second.
	RateLimitUnit_SECOND RateLimitUnit = 1
	// The time unit representing a minute.
	RateLimitUnit_MINUTE RateLimitUnit = 2
	// The time unit representing an hour.
	RateLimitUnit_HOUR RateLimitUnit = 3
	// The time unit representing a day.
	RateLimitUnit_DAY RateLimitUnit = 4
	// The time unit representing a week.
	RateLimitUnit_WEEK RateLimitUnit = 7
	// The time unit representing a month.
	RateLimitUnit_MONTH RateLimitUnit = 5
	// The time unit representing a year.
	RateLimitUnit_YEAR RateLimitUnit = 6
)

// Enum value maps for RateLimitUnit.
var (
	RateLimitUnit_name = map[int32]string{
		0: "UNKNOWN",
		1: "SECOND",
		2: "MINUTE",
		3: "HOUR",
		4: "DAY",
		7: "WEEK",
		5: "MONTH",
		6: "YEAR",
	}
	RateLimitUnit_value = map[string]int32{
		"UNKNOWN": 0,
		"SECOND":  1,
		"MINUTE":  2,
		"HOUR":    3,
		"DAY":     4,
		"WEEK":    7,
		"MONTH":   5,
		"YEAR":    6,
	}
)

func (x RateLimitUnit) Enum() *RateLimitUnit {
	p := new(RateLimitUnit)
	*p = x
	return p
}

func (x RateLimitUnit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RateLimitUnit) Descriptor() protoreflect.EnumDescriptor {
	return file_ratelimit_config_ratelimit_v3_rls_conf_proto_enumTypes[0].Descriptor()
}

func (RateLimitUnit) Type() protoreflect.EnumType {
	return &file_ratelimit_config_ratelimit_v3_rls_conf_proto_enumTypes[0]
}

func (x RateLimitUnit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RateLimitUnit.Descriptor instead.
func (RateLimitUnit) EnumDescriptor() ([]byte, []int) {
	return file_ratelimit_config_ratelimit_v3_rls_conf_proto_rawDescGZIP(), []int{0}
}

// Rate limit configuration for a single domain.
type RateLimitConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the rate limit configuration. This should be unique for each configuration.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Domain name for the rate limit configuration.
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// List of rate limit configuration descriptors.
	Descriptors []*RateLimitDescriptor `protobuf:"bytes,3,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
}

func (x *RateLimitConfig) Reset() {
	*x = RateLimitConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ratelimit_config_ratelimit_v3_rls_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitConfig) ProtoMessage() {}

func (x *RateLimitConfig) ProtoReflect() protoreflect.Message {
	mi := &file_ratelimit_config_ratelimit_v3_rls_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitConfig.ProtoReflect.Descriptor instead.
func (*RateLimitConfig) Descriptor() ([]byte, []int) {
	return file_ratelimit_config_ratelimit_v3_rls_conf_proto_rawDescGZIP(), []int{0}
}

func (x *RateLimitConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RateLimitConfig) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *RateLimitConfig) GetDescriptors() []*RateLimitDescriptor {
	if x != nil {
		return x.Descriptors
	}
	return nil
}

// Rate limit configuration descriptor.
type RateLimitDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key of the descriptor.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Optional value of the descriptor.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// Rate limit policy of the descriptor.
	RateLimit *RateLimitPolicy `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	// List of sub rate limit descriptors.
	Descriptors []*RateLimitDescriptor `protobuf:"bytes,4,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	// Mark the descriptor as shadow. When the values is true, rate limit service allow requests to the backend.
	ShadowMode bool `protobuf:"varint,5,opt,name=shadow_mode,json=shadowMode,proto3" json:"shadow_mode,omitempty"`
	// Setting the `detailed_metric: true` for a descriptor will extend the metrics that are produced.
	DetailedMetric bool `protobuf:"varint,6,opt,name=detailed_metric,json=detailedMetric,proto3" json:"detailed_metric,omitempty"`
}

func (x *RateLimitDescriptor) Reset() {
	*x = RateLimitDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ratelimit_config_ratelimit_v3_rls_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitDescriptor) ProtoMessage() {}

func (x *RateLimitDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_ratelimit_config_ratelimit_v3_rls_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitDescriptor.ProtoReflect.Descriptor instead.
func (*RateLimitDescriptor) Descriptor() ([]byte, []int) {
	return file_ratelimit_config_ratelimit_v3_rls_conf_proto_rawDescGZIP(), []int{1}
}

func (x *RateLimitDescriptor) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *RateLimitDescriptor) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *RateLimitDescriptor) GetRateLimit() *RateLimitPolicy {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

func (x *RateLimitDescriptor) GetDescriptors() []*RateLimitDescriptor {
	if x != nil {
		return x.Descriptors
	}
	return nil
}

func (x *RateLimitDescriptor) GetShadowMode() bool {
	if x != nil {
		return x.ShadowMode
	}
	return false
}

func (x *RateLimitDescriptor) GetDetailedMetric() bool {
	if x != nil {
		return x.DetailedMetric
	}
	return false
}

// Rate-limit policy.
type RateLimitPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unit of time for the rate limit.
	Unit RateLimitUnit `protobuf:"varint,1,opt,name=unit,proto3,enum=ratelimit.config.ratelimit.v3.RateLimitUnit" json:"unit,omitempty"`
	// Number of requests allowed in the policy within `unit` time.
	RequestsPerUnit uint32 `protobuf:"varint,2,opt,name=requests_per_unit,json=requestsPerUnit,proto3" json:"requests_per_unit,omitempty"`
	// Mark the rate limit policy as unlimited. All requests are allowed to the backend.
	Unlimited bool `protobuf:"varint,3,opt,name=unlimited,proto3" json:"unlimited,omitempty"`
	// Optional name for the rate limit policy. Name the policy, if it should be replaced (dropped evaluation) by
	// another policy.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// List of rate limit policies, this rate limit policy will replace (drop evaluation)
	// For more information: https://github.com/envoyproxy/ratelimit/tree/0b2f4d5fb04bf55e1873e2c5e2bb28da67c0643f#replaces
	// Example: https://github.com/envoyproxy/ratelimit/tree/0b2f4d5fb04bf55e1873e2c5e2bb28da67c0643f#example-7
	Replaces []*RateLimitReplace `protobuf:"bytes,5,rep,name=replaces,proto3" json:"replaces,omitempty"`
}

func (x *RateLimitPolicy) Reset() {
	*x = RateLimitPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ratelimit_config_ratelimit_v3_rls_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitPolicy) ProtoMessage() {}

func (x *RateLimitPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_ratelimit_config_ratelimit_v3_rls_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitPolicy.ProtoReflect.Descriptor instead.
func (*RateLimitPolicy) Descriptor() ([]byte, []int) {
	return file_ratelimit_config_ratelimit_v3_rls_conf_proto_rawDescGZIP(), []int{2}
}

func (x *RateLimitPolicy) GetUnit() RateLimitUnit {
	if x != nil {
		return x.Unit
	}
	return RateLimitUnit_UNKNOWN
}

func (x *RateLimitPolicy) GetRequestsPerUnit() uint32 {
	if x != nil {
		return x.RequestsPerUnit
	}
	return 0
}

func (x *RateLimitPolicy) GetUnlimited() bool {
	if x != nil {
		return x.Unlimited
	}
	return false
}

func (x *RateLimitPolicy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RateLimitPolicy) GetReplaces() []*RateLimitReplace {
	if x != nil {
		return x.Replaces
	}
	return nil
}

// Replace specifies the rate limit policy that should be replaced (dropped evaluation).
// For more information: https://github.com/envoyproxy/ratelimit/tree/0b2f4d5fb04bf55e1873e2c5e2bb28da67c0643f#replaces
type RateLimitReplace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the rate limit policy, that is being replaced (dropped evaluation).
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RateLimitReplace) Reset() {
	*x = RateLimitReplace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ratelimit_config_ratelimit_v3_rls_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitReplace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitReplace) ProtoMessage() {}

func (x *RateLimitReplace) ProtoReflect() protoreflect.Message {
	mi := &file_ratelimit_config_ratelimit_v3_rls_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitReplace.ProtoReflect.Descriptor instead.
func (*RateLimitReplace) Descriptor() ([]byte, []int) {
	return file_ratelimit_config_ratelimit_v3_rls_conf_proto_rawDescGZIP(), []int{3}
}

func (x *RateLimitReplace) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_ratelimit_config_ratelimit_v3_rls_conf_proto protoreflect.FileDescriptor

var file_ratelimit_config_ratelimit_v3_rls_conf_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x76, 0x33, 0x2f,
	0x72, 0x6c, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d,
	0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33, 0x22, 0x93, 0x01,
	0x0a, 0x0f, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x54, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x32, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e,
	0x76, 0x33, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x73, 0x22, 0xac, 0x02, 0x0a, 0x13, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x54, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x61, 0x64,
	0x6f, 0x77, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73,
	0x68, 0x61, 0x64, 0x6f, 0x77, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0e, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x22, 0xfe, 0x01, 0x0a, 0x0f, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x40, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x55, 0x6e,
	0x69, 0x74, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x50, 0x65, 0x72,
	0x55, 0x6e, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x6e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x75, 0x6e, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x61, 0x74, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x73, 0x22, 0x26, 0x0a, 0x10, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x2a, 0x66, 0x0a, 0x0d, 0x52,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x43,
	0x4f, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x10,
	0x02, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x55, 0x52, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x44,
	0x41, 0x59, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x45, 0x45, 0x4b, 0x10, 0x07, 0x12, 0x09,
	0x0a, 0x05, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x59, 0x45, 0x41,
	0x52, 0x10, 0x06, 0x42, 0x91, 0x01, 0x0a, 0x2b, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x2e, 0x76, 0x33, 0x42, 0x0e, 0x52, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x72, 0x61,
	0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72,
	0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x76, 0x33, 0x3b, 0x72, 0x61, 0x74, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ratelimit_config_ratelimit_v3_rls_conf_proto_rawDescOnce sync.Once
	file_ratelimit_config_ratelimit_v3_rls_conf_proto_rawDescData = file_ratelimit_config_ratelimit_v3_rls_conf_proto_rawDesc
)

func file_ratelimit_config_ratelimit_v3_rls_conf_proto_rawDescGZIP() []byte {
	file_ratelimit_config_ratelimit_v3_rls_conf_proto_rawDescOnce.Do(func() {
		file_ratelimit_config_ratelimit_v3_rls_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_ratelimit_config_ratelimit_v3_rls_conf_proto_rawDescData)
	})
	return file_ratelimit_config_ratelimit_v3_rls_conf_proto_rawDescData
}

var file_ratelimit_config_ratelimit_v3_rls_conf_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ratelimit_config_ratelimit_v3_rls_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ratelimit_config_ratelimit_v3_rls_conf_proto_goTypes = []interface{}{
	(RateLimitUnit)(0),          // 0: ratelimit.config.ratelimit.v3.RateLimitUnit
	(*RateLimitConfig)(nil),     // 1: ratelimit.config.ratelimit.v3.RateLimitConfig
	(*RateLimitDescriptor)(nil), // 2: ratelimit.config.ratelimit.v3.RateLimitDescriptor
	(*RateLimitPolicy)(nil),     // 3: ratelimit.config.ratelimit.v3.RateLimitPolicy
	(*RateLimitReplace)(nil),    // 4: ratelimit.config.ratelimit.v3.RateLimitReplace
}
var file_ratelimit_config_ratelimit_v3_rls_conf_proto_depIdxs = []int32{
	2, // 0: ratelimit.config.ratelimit.v3.RateLimitConfig.descriptors:type_name -> ratelimit.config.ratelimit.v3.RateLimitDescriptor
	3, // 1: ratelimit.config.ratelimit.v3.RateLimitDescriptor.rate_limit:type_name -> ratelimit.config.ratelimit.v3.RateLimitPolicy
	2, // 2: ratelimit.config.ratelimit.v3.RateLimitDescriptor.descriptors:type_name -> ratelimit.config.ratelimit.v3.RateLimitDescriptor
	0, // 3: ratelimit.config.ratelimit.v3.RateLimitPolicy.unit:type_name -> ratelimit.config.ratelimit.v3.RateLimitUnit
	4, // 4: ratelimit.config.ratelimit.v3.RateLimitPolicy.replaces:type_name -> ratelimit.config.ratelimit.v3.RateLimitReplace
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ratelimit_config_ratelimit_v3_rls_conf_proto_init() }
func file_ratelimit_config_ratelimit_v3_rls_conf_proto_init() {
	if File_ratelimit_config_ratelimit_v3_rls_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ratelimit_config_ratelimit_v3_rls_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitConfig); i {
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
		file_ratelimit_config_ratelimit_v3_rls_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitDescriptor); i {
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
		file_ratelimit_config_ratelimit_v3_rls_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitPolicy); i {
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
		file_ratelimit_config_ratelimit_v3_rls_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitReplace); i {
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
			RawDescriptor: file_ratelimit_config_ratelimit_v3_rls_conf_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ratelimit_config_ratelimit_v3_rls_conf_proto_goTypes,
		DependencyIndexes: file_ratelimit_config_ratelimit_v3_rls_conf_proto_depIdxs,
		EnumInfos:         file_ratelimit_config_ratelimit_v3_rls_conf_proto_enumTypes,
		MessageInfos:      file_ratelimit_config_ratelimit_v3_rls_conf_proto_msgTypes,
	}.Build()
	File_ratelimit_config_ratelimit_v3_rls_conf_proto = out.File
	file_ratelimit_config_ratelimit_v3_rls_conf_proto_rawDesc = nil
	file_ratelimit_config_ratelimit_v3_rls_conf_proto_goTypes = nil
	file_ratelimit_config_ratelimit_v3_rls_conf_proto_depIdxs = nil
}
