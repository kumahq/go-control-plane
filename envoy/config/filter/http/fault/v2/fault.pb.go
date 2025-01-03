// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.29.2
// source: envoy/config/filter/http/fault/v2/fault.proto

package faultv2

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/fault/v2"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
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

type FaultAbort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ErrorType:
	//
	//	*FaultAbort_HttpStatus
	//	*FaultAbort_HeaderAbort_
	ErrorType isFaultAbort_ErrorType `protobuf_oneof:"error_type"`
	// The percentage of requests/operations/connections that will be aborted with the error code
	// provided.
	Percentage *_type.FractionalPercent `protobuf:"bytes,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
}

func (x *FaultAbort) Reset() {
	*x = FaultAbort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_http_fault_v2_fault_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaultAbort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaultAbort) ProtoMessage() {}

func (x *FaultAbort) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_fault_v2_fault_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaultAbort.ProtoReflect.Descriptor instead.
func (*FaultAbort) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_fault_v2_fault_proto_rawDescGZIP(), []int{0}
}

func (m *FaultAbort) GetErrorType() isFaultAbort_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return nil
}

func (x *FaultAbort) GetHttpStatus() uint32 {
	if x, ok := x.GetErrorType().(*FaultAbort_HttpStatus); ok {
		return x.HttpStatus
	}
	return 0
}

func (x *FaultAbort) GetHeaderAbort() *FaultAbort_HeaderAbort {
	if x, ok := x.GetErrorType().(*FaultAbort_HeaderAbort_); ok {
		return x.HeaderAbort
	}
	return nil
}

func (x *FaultAbort) GetPercentage() *_type.FractionalPercent {
	if x != nil {
		return x.Percentage
	}
	return nil
}

type isFaultAbort_ErrorType interface {
	isFaultAbort_ErrorType()
}

type FaultAbort_HttpStatus struct {
	// HTTP status code to use to abort the HTTP request.
	HttpStatus uint32 `protobuf:"varint,2,opt,name=http_status,json=httpStatus,proto3,oneof"`
}

type FaultAbort_HeaderAbort_ struct {
	// Fault aborts are controlled via an HTTP header (if applicable).
	HeaderAbort *FaultAbort_HeaderAbort `protobuf:"bytes,4,opt,name=header_abort,json=headerAbort,proto3,oneof"`
}

func (*FaultAbort_HttpStatus) isFaultAbort_ErrorType() {}

func (*FaultAbort_HeaderAbort_) isFaultAbort_ErrorType() {}

// [#next-free-field: 14]
type HTTPFault struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If specified, the filter will inject delays based on the values in the
	// object.
	Delay *v2.FaultDelay `protobuf:"bytes,1,opt,name=delay,proto3" json:"delay,omitempty"`
	// If specified, the filter will abort requests based on the values in
	// the object. At least *abort* or *delay* must be specified.
	Abort *FaultAbort `protobuf:"bytes,2,opt,name=abort,proto3" json:"abort,omitempty"`
	// Specifies the name of the (destination) upstream cluster that the
	// filter should match on. Fault injection will be restricted to requests
	// bound to the specific upstream cluster.
	UpstreamCluster string `protobuf:"bytes,3,opt,name=upstream_cluster,json=upstreamCluster,proto3" json:"upstream_cluster,omitempty"`
	// Specifies a set of headers that the filter should match on. The fault
	// injection filter can be applied selectively to requests that match a set of
	// headers specified in the fault filter config. The chances of actual fault
	// injection further depend on the value of the :ref:`percentage
	// <envoy_api_field_config.filter.http.fault.v2.FaultAbort.percentage>` field.
	// The filter will check the request's headers against all the specified
	// headers in the filter config. A match will happen if all the headers in the
	// config are present in the request with the same values (or based on
	// presence if the *value* field is not in the config).
	Headers []*route.HeaderMatcher `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	// Faults are injected for the specified list of downstream hosts. If this
	// setting is not set, faults are injected for all downstream nodes.
	// Downstream node name is taken from :ref:`the HTTP
	// x-envoy-downstream-service-node
	// <config_http_conn_man_headers_downstream-service-node>` header and compared
	// against downstream_nodes list.
	DownstreamNodes []string `protobuf:"bytes,5,rep,name=downstream_nodes,json=downstreamNodes,proto3" json:"downstream_nodes,omitempty"`
	// The maximum number of faults that can be active at a single time via the configured fault
	// filter. Note that because this setting can be overridden at the route level, it's possible
	// for the number of active faults to be greater than this value (if injected via a different
	// route). If not specified, defaults to unlimited. This setting can be overridden via
	// `runtime <config_http_filters_fault_injection_runtime>` and any faults that are not injected
	// due to overflow will be indicated via the `faults_overflow
	// <config_http_filters_fault_injection_stats>` stat.
	//
	// .. attention::
	//
	//	Like other :ref:`circuit breakers <arch_overview_circuit_break>` in Envoy, this is a fuzzy
	//	limit. It's possible for the number of active faults to rise slightly above the configured
	//	amount due to the implementation details.
	MaxActiveFaults *wrapperspb.UInt32Value `protobuf:"bytes,6,opt,name=max_active_faults,json=maxActiveFaults,proto3" json:"max_active_faults,omitempty"`
	// The response rate limit to be applied to the response body of the stream. When configured,
	// the percentage can be overridden by the :ref:`fault.http.rate_limit.response_percent
	// <config_http_filters_fault_injection_runtime>` runtime key.
	//
	// .. attention::
	//
	//	This is a per-stream limit versus a connection level limit. This means that concurrent streams
	//	will each get an independent limit.
	ResponseRateLimit *v2.FaultRateLimit `protobuf:"bytes,7,opt,name=response_rate_limit,json=responseRateLimit,proto3" json:"response_rate_limit,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.delay.fixed_delay_percent
	DelayPercentRuntime string `protobuf:"bytes,8,opt,name=delay_percent_runtime,json=delayPercentRuntime,proto3" json:"delay_percent_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.abort.abort_percent
	AbortPercentRuntime string `protobuf:"bytes,9,opt,name=abort_percent_runtime,json=abortPercentRuntime,proto3" json:"abort_percent_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.delay.fixed_duration_ms
	DelayDurationRuntime string `protobuf:"bytes,10,opt,name=delay_duration_runtime,json=delayDurationRuntime,proto3" json:"delay_duration_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.abort.http_status
	AbortHttpStatusRuntime string `protobuf:"bytes,11,opt,name=abort_http_status_runtime,json=abortHttpStatusRuntime,proto3" json:"abort_http_status_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.max_active_faults
	MaxActiveFaultsRuntime string `protobuf:"bytes,12,opt,name=max_active_faults_runtime,json=maxActiveFaultsRuntime,proto3" json:"max_active_faults_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.rate_limit.response_percent
	ResponseRateLimitPercentRuntime string `protobuf:"bytes,13,opt,name=response_rate_limit_percent_runtime,json=responseRateLimitPercentRuntime,proto3" json:"response_rate_limit_percent_runtime,omitempty"`
}

func (x *HTTPFault) Reset() {
	*x = HTTPFault{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_http_fault_v2_fault_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPFault) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPFault) ProtoMessage() {}

func (x *HTTPFault) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_fault_v2_fault_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPFault.ProtoReflect.Descriptor instead.
func (*HTTPFault) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_fault_v2_fault_proto_rawDescGZIP(), []int{1}
}

func (x *HTTPFault) GetDelay() *v2.FaultDelay {
	if x != nil {
		return x.Delay
	}
	return nil
}

func (x *HTTPFault) GetAbort() *FaultAbort {
	if x != nil {
		return x.Abort
	}
	return nil
}

func (x *HTTPFault) GetUpstreamCluster() string {
	if x != nil {
		return x.UpstreamCluster
	}
	return ""
}

func (x *HTTPFault) GetHeaders() []*route.HeaderMatcher {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *HTTPFault) GetDownstreamNodes() []string {
	if x != nil {
		return x.DownstreamNodes
	}
	return nil
}

func (x *HTTPFault) GetMaxActiveFaults() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxActiveFaults
	}
	return nil
}

func (x *HTTPFault) GetResponseRateLimit() *v2.FaultRateLimit {
	if x != nil {
		return x.ResponseRateLimit
	}
	return nil
}

func (x *HTTPFault) GetDelayPercentRuntime() string {
	if x != nil {
		return x.DelayPercentRuntime
	}
	return ""
}

func (x *HTTPFault) GetAbortPercentRuntime() string {
	if x != nil {
		return x.AbortPercentRuntime
	}
	return ""
}

func (x *HTTPFault) GetDelayDurationRuntime() string {
	if x != nil {
		return x.DelayDurationRuntime
	}
	return ""
}

func (x *HTTPFault) GetAbortHttpStatusRuntime() string {
	if x != nil {
		return x.AbortHttpStatusRuntime
	}
	return ""
}

func (x *HTTPFault) GetMaxActiveFaultsRuntime() string {
	if x != nil {
		return x.MaxActiveFaultsRuntime
	}
	return ""
}

func (x *HTTPFault) GetResponseRateLimitPercentRuntime() string {
	if x != nil {
		return x.ResponseRateLimitPercentRuntime
	}
	return ""
}

// Fault aborts are controlled via an HTTP header (if applicable). See the
// :ref:`HTTP fault filter <config_http_filters_fault_injection_http_header>` documentation for
// more information.
type FaultAbort_HeaderAbort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FaultAbort_HeaderAbort) Reset() {
	*x = FaultAbort_HeaderAbort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_http_fault_v2_fault_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaultAbort_HeaderAbort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaultAbort_HeaderAbort) ProtoMessage() {}

func (x *FaultAbort_HeaderAbort) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_fault_v2_fault_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaultAbort_HeaderAbort.ProtoReflect.Descriptor instead.
func (*FaultAbort_HeaderAbort) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_fault_v2_fault_proto_rawDescGZIP(), []int{0, 0}
}

var File_envoy_config_filter_http_fault_v2_fault_proto protoreflect.FileDescriptor

var file_envoy_config_filter_http_fault_v2_fault_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x2f, 0x76, 0x32, 0x2f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e,
	0x76, 0x32, 0x1a, 0x29, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32,
	0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x02, 0x0a, 0x0a, 0x46, 0x61,
	0x75, 0x6c, 0x74, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x68, 0x74, 0x74, 0x70,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0b, 0xfa,
	0x42, 0x08, 0x2a, 0x06, 0x10, 0xd8, 0x04, 0x28, 0xc8, 0x01, 0x48, 0x00, 0x52, 0x0a, 0x68, 0x74,
	0x74, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x5e, 0x0a, 0x0c, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x5f, 0x61, 0x62, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e,
	0x76, 0x32, 0x2e, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x12, 0x3d, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x70, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x42, 0x11, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x22,
	0xad, 0x06, 0x0a, 0x09, 0x48, 0x54, 0x54, 0x50, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x3e, 0x0a,
	0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x61, 0x75, 0x6c,
	0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x43, 0x0a,
	0x05, 0x61, 0x62, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x32,
	0x2e, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x61, 0x62, 0x6f,
	0x72, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x75, 0x70,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x3b, 0x0a,
	0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x6f,
	0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x11, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f,
	0x6d, 0x61, 0x78, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x12,
	0x5c, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x61, 0x75, 0x6c,
	0x74, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x11, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x32, 0x0a,
	0x15, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x64, 0x65,
	0x6c, 0x61, 0x79, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x32, 0x0a, 0x15, 0x61, 0x62, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x5f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x61, 0x62, 0x6f, 0x72, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x19, 0x61,
	0x62, 0x6f, 0x72, 0x74, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16,
	0x61, 0x62, 0x6f, 0x72, 0x74, 0x48, 0x74, 0x74, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x19, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x5f, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6d, 0x61, 0x78, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x4c, 0x0a, 0x23, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x72, 0x61,
	0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x5f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x42,
	0xc7, 0x01, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x28, 0x12, 0x26, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x33,
	0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x0a, 0x2f, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x32, 0x42, 0x0a, 0x46, 0x61, 0x75, 0x6c, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67,
	0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x76,
	0x32, 0x3b, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_envoy_config_filter_http_fault_v2_fault_proto_rawDescOnce sync.Once
	file_envoy_config_filter_http_fault_v2_fault_proto_rawDescData = file_envoy_config_filter_http_fault_v2_fault_proto_rawDesc
)

func file_envoy_config_filter_http_fault_v2_fault_proto_rawDescGZIP() []byte {
	file_envoy_config_filter_http_fault_v2_fault_proto_rawDescOnce.Do(func() {
		file_envoy_config_filter_http_fault_v2_fault_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_filter_http_fault_v2_fault_proto_rawDescData)
	})
	return file_envoy_config_filter_http_fault_v2_fault_proto_rawDescData
}

var file_envoy_config_filter_http_fault_v2_fault_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_config_filter_http_fault_v2_fault_proto_goTypes = []interface{}{
	(*FaultAbort)(nil),              // 0: envoy.config.filter.http.fault.v2.FaultAbort
	(*HTTPFault)(nil),               // 1: envoy.config.filter.http.fault.v2.HTTPFault
	(*FaultAbort_HeaderAbort)(nil),  // 2: envoy.config.filter.http.fault.v2.FaultAbort.HeaderAbort
	(*_type.FractionalPercent)(nil), // 3: envoy.type.FractionalPercent
	(*v2.FaultDelay)(nil),           // 4: envoy.config.filter.fault.v2.FaultDelay
	(*route.HeaderMatcher)(nil),     // 5: envoy.api.v2.route.HeaderMatcher
	(*wrapperspb.UInt32Value)(nil),  // 6: google.protobuf.UInt32Value
	(*v2.FaultRateLimit)(nil),       // 7: envoy.config.filter.fault.v2.FaultRateLimit
}
var file_envoy_config_filter_http_fault_v2_fault_proto_depIdxs = []int32{
	2, // 0: envoy.config.filter.http.fault.v2.FaultAbort.header_abort:type_name -> envoy.config.filter.http.fault.v2.FaultAbort.HeaderAbort
	3, // 1: envoy.config.filter.http.fault.v2.FaultAbort.percentage:type_name -> envoy.type.FractionalPercent
	4, // 2: envoy.config.filter.http.fault.v2.HTTPFault.delay:type_name -> envoy.config.filter.fault.v2.FaultDelay
	0, // 3: envoy.config.filter.http.fault.v2.HTTPFault.abort:type_name -> envoy.config.filter.http.fault.v2.FaultAbort
	5, // 4: envoy.config.filter.http.fault.v2.HTTPFault.headers:type_name -> envoy.api.v2.route.HeaderMatcher
	6, // 5: envoy.config.filter.http.fault.v2.HTTPFault.max_active_faults:type_name -> google.protobuf.UInt32Value
	7, // 6: envoy.config.filter.http.fault.v2.HTTPFault.response_rate_limit:type_name -> envoy.config.filter.fault.v2.FaultRateLimit
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_envoy_config_filter_http_fault_v2_fault_proto_init() }
func file_envoy_config_filter_http_fault_v2_fault_proto_init() {
	if File_envoy_config_filter_http_fault_v2_fault_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_filter_http_fault_v2_fault_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaultAbort); i {
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
		file_envoy_config_filter_http_fault_v2_fault_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPFault); i {
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
		file_envoy_config_filter_http_fault_v2_fault_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaultAbort_HeaderAbort); i {
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
	file_envoy_config_filter_http_fault_v2_fault_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FaultAbort_HttpStatus)(nil),
		(*FaultAbort_HeaderAbort_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_filter_http_fault_v2_fault_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_filter_http_fault_v2_fault_proto_goTypes,
		DependencyIndexes: file_envoy_config_filter_http_fault_v2_fault_proto_depIdxs,
		MessageInfos:      file_envoy_config_filter_http_fault_v2_fault_proto_msgTypes,
	}.Build()
	File_envoy_config_filter_http_fault_v2_fault_proto = out.File
	file_envoy_config_filter_http_fault_v2_fault_proto_rawDesc = nil
	file_envoy_config_filter_http_fault_v2_fault_proto_goTypes = nil
	file_envoy_config_filter_http_fault_v2_fault_proto_depIdxs = nil
}
