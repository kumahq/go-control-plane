//go:build vtprotobuf
// +build vtprotobuf

// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/extensions/filters/http/oauth2/v3/oauth.proto

package oauth2v3

import (
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	durationpb "github.com/planetscale/vtprotobuf/types/known/durationpb"
	wrapperspb "github.com/planetscale/vtprotobuf/types/known/wrapperspb"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *OAuth2Credentials_CookieNames) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OAuth2Credentials_CookieNames) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OAuth2Credentials_CookieNames) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.OauthNonce) > 0 {
		i -= len(m.OauthNonce)
		copy(dAtA[i:], m.OauthNonce)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.OauthNonce)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.RefreshToken) > 0 {
		i -= len(m.RefreshToken)
		copy(dAtA[i:], m.RefreshToken)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.RefreshToken)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.IdToken) > 0 {
		i -= len(m.IdToken)
		copy(dAtA[i:], m.IdToken)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.IdToken)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.OauthExpires) > 0 {
		i -= len(m.OauthExpires)
		copy(dAtA[i:], m.OauthExpires)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.OauthExpires)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OauthHmac) > 0 {
		i -= len(m.OauthHmac)
		copy(dAtA[i:], m.OauthHmac)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.OauthHmac)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BearerToken) > 0 {
		i -= len(m.BearerToken)
		copy(dAtA[i:], m.BearerToken)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.BearerToken)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OAuth2Credentials) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OAuth2Credentials) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OAuth2Credentials) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.CookieDomain) > 0 {
		i -= len(m.CookieDomain)
		copy(dAtA[i:], m.CookieDomain)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.CookieDomain)))
		i--
		dAtA[i] = 0x2a
	}
	if m.CookieNames != nil {
		size, err := m.CookieNames.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if msg, ok := m.TokenFormation.(*OAuth2Credentials_HmacSecret); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if m.TokenSecret != nil {
		if vtmsg, ok := interface{}(m.TokenSecret).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.TokenSecret)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OAuth2Credentials_HmacSecret) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OAuth2Credentials_HmacSecret) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HmacSecret != nil {
		if vtmsg, ok := interface{}(m.HmacSecret).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.HmacSecret)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x1a
	} else {
		i = protohelpers.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *OAuth2Config) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OAuth2Config) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OAuth2Config) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.DisableRefreshTokenSetCookie {
		i--
		if m.DisableRefreshTokenSetCookie {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa0
	}
	if m.DisableAccessTokenSetCookie {
		i--
		if m.DisableAccessTokenSetCookie {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x98
	}
	if m.RetryPolicy != nil {
		if vtmsg, ok := interface{}(m.RetryPolicy).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.RetryPolicy)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x92
	}
	if m.DisableIdTokenSetCookie {
		i--
		if m.DisableIdTokenSetCookie {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if m.PreserveAuthorizationHeader {
		i--
		if m.PreserveAuthorizationHeader {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.DefaultRefreshTokenExpiresIn != nil {
		size, err := (*durationpb.Duration)(m.DefaultRefreshTokenExpiresIn).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x7a
	}
	if len(m.DenyRedirectMatcher) > 0 {
		for iNdEx := len(m.DenyRedirectMatcher) - 1; iNdEx >= 0; iNdEx-- {
			if vtmsg, ok := interface{}(m.DenyRedirectMatcher[iNdEx]).(interface {
				MarshalToSizedBufferVTStrict([]byte) (int, error)
			}); ok {
				size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			} else {
				encoded, err := proto.Marshal(m.DenyRedirectMatcher[iNdEx])
				if err != nil {
					return 0, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
			}
			i--
			dAtA[i] = 0x72
		}
	}
	if m.DefaultExpiresIn != nil {
		size, err := (*durationpb.Duration)(m.DefaultExpiresIn).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x6a
	}
	if m.UseRefreshToken != nil {
		size, err := (*wrapperspb.BoolValue)(m.UseRefreshToken).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x62
	}
	if m.AuthType != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.AuthType))
		i--
		dAtA[i] = 0x58
	}
	if len(m.Resources) > 0 {
		for iNdEx := len(m.Resources) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Resources[iNdEx])
			copy(dAtA[i:], m.Resources[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Resources[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.AuthScopes) > 0 {
		for iNdEx := len(m.AuthScopes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AuthScopes[iNdEx])
			copy(dAtA[i:], m.AuthScopes[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.AuthScopes[iNdEx])))
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.PassThroughMatcher) > 0 {
		for iNdEx := len(m.PassThroughMatcher) - 1; iNdEx >= 0; iNdEx-- {
			if vtmsg, ok := interface{}(m.PassThroughMatcher[iNdEx]).(interface {
				MarshalToSizedBufferVTStrict([]byte) (int, error)
			}); ok {
				size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			} else {
				encoded, err := proto.Marshal(m.PassThroughMatcher[iNdEx])
				if err != nil {
					return 0, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if m.ForwardBearerToken {
		i--
		if m.ForwardBearerToken {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.SignoutPath != nil {
		if vtmsg, ok := interface{}(m.SignoutPath).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.SignoutPath)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.RedirectPathMatcher != nil {
		if vtmsg, ok := interface{}(m.RedirectPathMatcher).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.RedirectPathMatcher)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.RedirectUri) > 0 {
		i -= len(m.RedirectUri)
		copy(dAtA[i:], m.RedirectUri)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.RedirectUri)))
		i--
		dAtA[i] = 0x22
	}
	if m.Credentials != nil {
		size, err := m.Credentials.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AuthorizationEndpoint) > 0 {
		i -= len(m.AuthorizationEndpoint)
		copy(dAtA[i:], m.AuthorizationEndpoint)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.AuthorizationEndpoint)))
		i--
		dAtA[i] = 0x12
	}
	if m.TokenEndpoint != nil {
		if vtmsg, ok := interface{}(m.TokenEndpoint).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.TokenEndpoint)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OAuth2) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OAuth2) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OAuth2) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.Config != nil {
		size, err := m.Config.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OAuth2Credentials_CookieNames) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BearerToken)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.OauthHmac)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.OauthExpires)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.IdToken)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.RefreshToken)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.OauthNonce)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *OAuth2Credentials) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.TokenSecret != nil {
		if size, ok := interface{}(m.TokenSecret).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.TokenSecret)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if vtmsg, ok := m.TokenFormation.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	if m.CookieNames != nil {
		l = m.CookieNames.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.CookieDomain)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *OAuth2Credentials_HmacSecret) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HmacSecret != nil {
		if size, ok := interface{}(m.HmacSecret).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.HmacSecret)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 2
	}
	return n
}
func (m *OAuth2Config) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TokenEndpoint != nil {
		if size, ok := interface{}(m.TokenEndpoint).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.TokenEndpoint)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.AuthorizationEndpoint)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Credentials != nil {
		l = m.Credentials.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.RedirectUri)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.RedirectPathMatcher != nil {
		if size, ok := interface{}(m.RedirectPathMatcher).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.RedirectPathMatcher)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.SignoutPath != nil {
		if size, ok := interface{}(m.SignoutPath).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.SignoutPath)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.ForwardBearerToken {
		n += 2
	}
	if len(m.PassThroughMatcher) > 0 {
		for _, e := range m.PassThroughMatcher {
			if size, ok := interface{}(e).(interface {
				SizeVT() int
			}); ok {
				l = size.SizeVT()
			} else {
				l = proto.Size(e)
			}
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if len(m.AuthScopes) > 0 {
		for _, s := range m.AuthScopes {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if len(m.Resources) > 0 {
		for _, s := range m.Resources {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if m.AuthType != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.AuthType))
	}
	if m.UseRefreshToken != nil {
		l = (*wrapperspb.BoolValue)(m.UseRefreshToken).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.DefaultExpiresIn != nil {
		l = (*durationpb.Duration)(m.DefaultExpiresIn).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.DenyRedirectMatcher) > 0 {
		for _, e := range m.DenyRedirectMatcher {
			if size, ok := interface{}(e).(interface {
				SizeVT() int
			}); ok {
				l = size.SizeVT()
			} else {
				l = proto.Size(e)
			}
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if m.DefaultRefreshTokenExpiresIn != nil {
		l = (*durationpb.Duration)(m.DefaultRefreshTokenExpiresIn).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.PreserveAuthorizationHeader {
		n += 3
	}
	if m.DisableIdTokenSetCookie {
		n += 3
	}
	if m.RetryPolicy != nil {
		if size, ok := interface{}(m.RetryPolicy).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.RetryPolicy)
		}
		n += 2 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.DisableAccessTokenSetCookie {
		n += 3
	}
	if m.DisableRefreshTokenSetCookie {
		n += 3
	}
	n += len(m.unknownFields)
	return n
}

func (m *OAuth2) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Config != nil {
		l = m.Config.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}
