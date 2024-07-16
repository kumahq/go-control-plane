//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/quic/server_preferred_address/v3/datasource.proto

package server_preferred_addressv3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on DataSourceServerPreferredAddressConfig
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *DataSourceServerPreferredAddressConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// DataSourceServerPreferredAddressConfig with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// DataSourceServerPreferredAddressConfigMultiError, or nil if none found.
func (m *DataSourceServerPreferredAddressConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *DataSourceServerPreferredAddressConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetIpv4Config()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DataSourceServerPreferredAddressConfigValidationError{
					field:  "Ipv4Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DataSourceServerPreferredAddressConfigValidationError{
					field:  "Ipv4Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIpv4Config()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DataSourceServerPreferredAddressConfigValidationError{
				field:  "Ipv4Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetIpv6Config()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DataSourceServerPreferredAddressConfigValidationError{
					field:  "Ipv6Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DataSourceServerPreferredAddressConfigValidationError{
					field:  "Ipv6Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIpv6Config()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DataSourceServerPreferredAddressConfigValidationError{
				field:  "Ipv6Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DataSourceServerPreferredAddressConfigMultiError(errors)
	}

	return nil
}

// DataSourceServerPreferredAddressConfigMultiError is an error wrapping
// multiple validation errors returned by
// DataSourceServerPreferredAddressConfig.ValidateAll() if the designated
// constraints aren't met.
type DataSourceServerPreferredAddressConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DataSourceServerPreferredAddressConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DataSourceServerPreferredAddressConfigMultiError) AllErrors() []error { return m }

// DataSourceServerPreferredAddressConfigValidationError is the validation
// error returned by DataSourceServerPreferredAddressConfig.Validate if the
// designated constraints aren't met.
type DataSourceServerPreferredAddressConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataSourceServerPreferredAddressConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataSourceServerPreferredAddressConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataSourceServerPreferredAddressConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DataSourceServerPreferredAddressConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataSourceServerPreferredAddressConfigValidationError) ErrorName() string {
	return "DataSourceServerPreferredAddressConfigValidationError"
}

// Error satisfies the builtin error interface
func (e DataSourceServerPreferredAddressConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDataSourceServerPreferredAddressConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataSourceServerPreferredAddressConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DataSourceServerPreferredAddressConfigValidationError{}

// Validate checks the field values on
// DataSourceServerPreferredAddressConfig_AddressFamilyConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DataSourceServerPreferredAddressConfig_AddressFamilyConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// DataSourceServerPreferredAddressConfig_AddressFamilyConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DataSourceServerPreferredAddressConfig_AddressFamilyConfigMultiError, or
// nil if none found.
func (m *DataSourceServerPreferredAddressConfig_AddressFamilyConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *DataSourceServerPreferredAddressConfig_AddressFamilyConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetAddress() == nil {
		err := DataSourceServerPreferredAddressConfig_AddressFamilyConfigValidationError{
			field:  "Address",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DataSourceServerPreferredAddressConfig_AddressFamilyConfigValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DataSourceServerPreferredAddressConfig_AddressFamilyConfigValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DataSourceServerPreferredAddressConfig_AddressFamilyConfigValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPort()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DataSourceServerPreferredAddressConfig_AddressFamilyConfigValidationError{
					field:  "Port",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DataSourceServerPreferredAddressConfig_AddressFamilyConfigValidationError{
					field:  "Port",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPort()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DataSourceServerPreferredAddressConfig_AddressFamilyConfigValidationError{
				field:  "Port",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDnatAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DataSourceServerPreferredAddressConfig_AddressFamilyConfigValidationError{
					field:  "DnatAddress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DataSourceServerPreferredAddressConfig_AddressFamilyConfigValidationError{
					field:  "DnatAddress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDnatAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DataSourceServerPreferredAddressConfig_AddressFamilyConfigValidationError{
				field:  "DnatAddress",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DataSourceServerPreferredAddressConfig_AddressFamilyConfigMultiError(errors)
	}

	return nil
}

// DataSourceServerPreferredAddressConfig_AddressFamilyConfigMultiError is an
// error wrapping multiple validation errors returned by
// DataSourceServerPreferredAddressConfig_AddressFamilyConfig.ValidateAll() if
// the designated constraints aren't met.
type DataSourceServerPreferredAddressConfig_AddressFamilyConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DataSourceServerPreferredAddressConfig_AddressFamilyConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DataSourceServerPreferredAddressConfig_AddressFamilyConfigMultiError) AllErrors() []error {
	return m
}

// DataSourceServerPreferredAddressConfig_AddressFamilyConfigValidationError is
// the validation error returned by
// DataSourceServerPreferredAddressConfig_AddressFamilyConfig.Validate if the
// designated constraints aren't met.
type DataSourceServerPreferredAddressConfig_AddressFamilyConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataSourceServerPreferredAddressConfig_AddressFamilyConfigValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e DataSourceServerPreferredAddressConfig_AddressFamilyConfigValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e DataSourceServerPreferredAddressConfig_AddressFamilyConfigValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e DataSourceServerPreferredAddressConfig_AddressFamilyConfigValidationError) Key() bool {
	return e.key
}

// ErrorName returns error name.
func (e DataSourceServerPreferredAddressConfig_AddressFamilyConfigValidationError) ErrorName() string {
	return "DataSourceServerPreferredAddressConfig_AddressFamilyConfigValidationError"
}

// Error satisfies the builtin error interface
func (e DataSourceServerPreferredAddressConfig_AddressFamilyConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDataSourceServerPreferredAddressConfig_AddressFamilyConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataSourceServerPreferredAddressConfig_AddressFamilyConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DataSourceServerPreferredAddressConfig_AddressFamilyConfigValidationError{}