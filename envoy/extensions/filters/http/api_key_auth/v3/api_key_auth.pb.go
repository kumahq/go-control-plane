// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.29.2
// source: envoy/extensions/filters/http/api_key_auth/v3/api_key_auth.proto

package api_key_authv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
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

// API Key HTTP authentication.
//
// For example, the following configuration configures the filter to authenticate the clients using
// the API key from the header “X-API-KEY“. And only the clients with the key “real-key“ are
// considered as authenticated.
//
// .. code-block:: yaml
//
//	credentials:
//	- key: real-key
//	  client: user
//	key_sources:
//	 - header: "X-API-KEY"
type ApiKeyAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The credentials that are used to authenticate the clients.
	Credentials []*Credential `protobuf:"bytes,1,rep,name=credentials,proto3" json:"credentials,omitempty"`
	// The key sources to fetch the key from the coming request.
	KeySources []*KeySource `protobuf:"bytes,2,rep,name=key_sources,json=keySources,proto3" json:"key_sources,omitempty"`
}

func (x *ApiKeyAuth) Reset() {
	*x = ApiKeyAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiKeyAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiKeyAuth) ProtoMessage() {}

func (x *ApiKeyAuth) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiKeyAuth.ProtoReflect.Descriptor instead.
func (*ApiKeyAuth) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDescGZIP(), []int{0}
}

func (x *ApiKeyAuth) GetCredentials() []*Credential {
	if x != nil {
		return x.Credentials
	}
	return nil
}

func (x *ApiKeyAuth) GetKeySources() []*KeySource {
	if x != nil {
		return x.KeySources
	}
	return nil
}

// API key auth configuration of per route or per virtual host or per route configuration.
type ApiKeyAuthPerRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The credentials that are used to authenticate the clients. If this field is non-empty, then the
	// credentials in the filter level configuration will be ignored and the credentials in this
	// configuration will be used.
	Credentials []*Credential `protobuf:"bytes,1,rep,name=credentials,proto3" json:"credentials,omitempty"`
	// The key sources to fetch the key from the coming request. If this field is non-empty, then the
	// key sources in the filter level configuration will be ignored and the key sources in this
	// configuration will be used.
	KeySources []*KeySource `protobuf:"bytes,2,rep,name=key_sources,json=keySources,proto3" json:"key_sources,omitempty"`
	// A list of clients that are allowed to access the route or vhost. The clients listed here
	// should be subset of the clients listed in the “credentials“ to provide authorization control
	// after the authentication is successful. If the list is empty, then all authenticated clients
	// are allowed. This provides very limited but simple authorization. If more complex authorization
	// is required, then use the :ref:`HTTP RBAC filter <config_http_filters_rbac>` instead.
	//
	// .. note::
	//
	//	Setting this field and ``credentials`` at the same configuration entry is not an error but
	//	also makes no much sense because they provide similar functionality. Please only use
	//	one of them at same configuration entry except for the case that you want to share the same
	//	credentials list across multiple routes but still use different allowed clients for each
	//	route.
	AllowedClients []string `protobuf:"bytes,3,rep,name=allowed_clients,json=allowedClients,proto3" json:"allowed_clients,omitempty"`
}

func (x *ApiKeyAuthPerRoute) Reset() {
	*x = ApiKeyAuthPerRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiKeyAuthPerRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiKeyAuthPerRoute) ProtoMessage() {}

func (x *ApiKeyAuthPerRoute) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiKeyAuthPerRoute.ProtoReflect.Descriptor instead.
func (*ApiKeyAuthPerRoute) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDescGZIP(), []int{1}
}

func (x *ApiKeyAuthPerRoute) GetCredentials() []*Credential {
	if x != nil {
		return x.Credentials
	}
	return nil
}

func (x *ApiKeyAuthPerRoute) GetKeySources() []*KeySource {
	if x != nil {
		return x.KeySources
	}
	return nil
}

func (x *ApiKeyAuthPerRoute) GetAllowedClients() []string {
	if x != nil {
		return x.AllowedClients
	}
	return nil
}

// Single credential entry that contains the API key and the related client id.
type Credential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The value of the unique API key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The unique id or identity that used to identify the client or consumer.
	Client string `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *Credential) Reset() {
	*x = Credential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credential) ProtoMessage() {}

func (x *Credential) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credential.ProtoReflect.Descriptor instead.
func (*Credential) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDescGZIP(), []int{2}
}

func (x *Credential) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Credential) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

type KeySource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The header name to fetch the key. If multiple header values are present, the first one will be
	// used. If the header value starts with 'Bearer ', this prefix will be stripped to get the
	// key value.
	//
	// If set, takes precedence over “query“ and “cookie“.
	Header string `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The query parameter name to fetch the key. If multiple query values are present, the first one
	// will be used.
	//
	// The field will be used if “header“ is not set. If set, takes precedence over “cookie“.
	Query string `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	// The cookie name to fetch the key.
	//
	// The field will be used if the “header“ and “query“ are not set.
	Cookie string `protobuf:"bytes,3,opt,name=cookie,proto3" json:"cookie,omitempty"`
}

func (x *KeySource) Reset() {
	*x = KeySource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeySource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeySource) ProtoMessage() {}

func (x *KeySource) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeySource.ProtoReflect.Descriptor instead.
func (*KeySource) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDescGZIP(), []int{3}
}

func (x *KeySource) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *KeySource) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *KeySource) GetCookie() string {
	if x != nil {
		return x.Cookie
	}
	return ""
}

var File_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDesc = []byte{
	0x0a, 0x40, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f,
	0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76,
	0x33, 0x1a, 0x1f, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x01, 0x0a,
	0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x63, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x39, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x33,
	0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x42, 0x06, 0xb8, 0xb7, 0x8b,
	0xa4, 0x02, 0x01, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x12, 0x59, 0x0a, 0x0b, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x33, 0x2e, 0x4b, 0x65, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x0a, 0x6b, 0x65, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0xfd, 0x01, 0x0a, 0x12,
	0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x12, 0x63, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x42, 0x06, 0xb8, 0xb7, 0x8b, 0xa4, 0x02, 0x01, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x59, 0x0a, 0x0b, 0x6b, 0x65, 0x79, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x70, 0x69,
	0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x33, 0x2e, 0x4b, 0x65, 0x79,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0a, 0x6b, 0x65, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x48, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x81, 0x01, 0x0a, 0x09, 0x4b, 0x65, 0x79, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x11, 0xfa, 0x42, 0x0e, 0x72, 0x0c, 0x18, 0x80, 0x08, 0xc8, 0x01, 0x00,
	0xd0, 0x01, 0x01, 0xc0, 0x01, 0x01, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1e,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x72, 0x03, 0x18, 0x80, 0x08, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x29,
	0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11,
	0xfa, 0x42, 0x0e, 0x72, 0x0c, 0x18, 0x80, 0x08, 0xc8, 0x01, 0x00, 0xd0, 0x01, 0x01, 0xc0, 0x01,
	0x01, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x42, 0xc5, 0x01, 0xba, 0x80, 0xc8, 0xd1,
	0x06, 0x02, 0x10, 0x02, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x0a, 0x3b, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65,
	0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x33, 0x42, 0x0f, 0x41, 0x70, 0x69, 0x4b, 0x65,
	0x79, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x63, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74,
	0x74, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x76, 0x33, 0x3b, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x76,
	0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDescData = file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDesc
)

func file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDescData
}

var file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_goTypes = []interface{}{
	(*ApiKeyAuth)(nil),         // 0: envoy.extensions.filters.http.api_key_auth.v3.ApiKeyAuth
	(*ApiKeyAuthPerRoute)(nil), // 1: envoy.extensions.filters.http.api_key_auth.v3.ApiKeyAuthPerRoute
	(*Credential)(nil),         // 2: envoy.extensions.filters.http.api_key_auth.v3.Credential
	(*KeySource)(nil),          // 3: envoy.extensions.filters.http.api_key_auth.v3.KeySource
}
var file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.filters.http.api_key_auth.v3.ApiKeyAuth.credentials:type_name -> envoy.extensions.filters.http.api_key_auth.v3.Credential
	3, // 1: envoy.extensions.filters.http.api_key_auth.v3.ApiKeyAuth.key_sources:type_name -> envoy.extensions.filters.http.api_key_auth.v3.KeySource
	2, // 2: envoy.extensions.filters.http.api_key_auth.v3.ApiKeyAuthPerRoute.credentials:type_name -> envoy.extensions.filters.http.api_key_auth.v3.Credential
	3, // 3: envoy.extensions.filters.http.api_key_auth.v3.ApiKeyAuthPerRoute.key_sources:type_name -> envoy.extensions.filters.http.api_key_auth.v3.KeySource
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_init() }
func file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_init() {
	if File_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiKeyAuth); i {
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
		file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiKeyAuthPerRoute); i {
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
		file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credential); i {
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
		file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeySource); i {
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
			RawDescriptor: file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto = out.File
	file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDesc = nil
	file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_goTypes = nil
	file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_depIdxs = nil
}
