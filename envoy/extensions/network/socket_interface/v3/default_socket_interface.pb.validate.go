//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/network/socket_interface/v3/default_socket_interface.proto

package socket_interfacev3

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

// Validate checks the field values on DefaultSocketInterface with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DefaultSocketInterface) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DefaultSocketInterface with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DefaultSocketInterfaceMultiError, or nil if none found.
func (m *DefaultSocketInterface) ValidateAll() error {
	return m.validate(true)
}

func (m *DefaultSocketInterface) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetIoUringOptions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DefaultSocketInterfaceValidationError{
					field:  "IoUringOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DefaultSocketInterfaceValidationError{
					field:  "IoUringOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIoUringOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DefaultSocketInterfaceValidationError{
				field:  "IoUringOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DefaultSocketInterfaceMultiError(errors)
	}

	return nil
}

// DefaultSocketInterfaceMultiError is an error wrapping multiple validation
// errors returned by DefaultSocketInterface.ValidateAll() if the designated
// constraints aren't met.
type DefaultSocketInterfaceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DefaultSocketInterfaceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DefaultSocketInterfaceMultiError) AllErrors() []error { return m }

// DefaultSocketInterfaceValidationError is the validation error returned by
// DefaultSocketInterface.Validate if the designated constraints aren't met.
type DefaultSocketInterfaceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DefaultSocketInterfaceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DefaultSocketInterfaceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DefaultSocketInterfaceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DefaultSocketInterfaceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DefaultSocketInterfaceValidationError) ErrorName() string {
	return "DefaultSocketInterfaceValidationError"
}

// Error satisfies the builtin error interface
func (e DefaultSocketInterfaceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDefaultSocketInterface.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DefaultSocketInterfaceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DefaultSocketInterfaceValidationError{}

// Validate checks the field values on IoUringOptions with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IoUringOptions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IoUringOptions with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IoUringOptionsMultiError,
// or nil if none found.
func (m *IoUringOptions) ValidateAll() error {
	return m.validate(true)
}

func (m *IoUringOptions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetIoUringSize()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IoUringOptionsValidationError{
					field:  "IoUringSize",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IoUringOptionsValidationError{
					field:  "IoUringSize",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIoUringSize()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IoUringOptionsValidationError{
				field:  "IoUringSize",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for EnableSubmissionQueuePolling

	if all {
		switch v := interface{}(m.GetReadBufferSize()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IoUringOptionsValidationError{
					field:  "ReadBufferSize",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IoUringOptionsValidationError{
					field:  "ReadBufferSize",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReadBufferSize()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IoUringOptionsValidationError{
				field:  "ReadBufferSize",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetWriteTimeoutMs()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IoUringOptionsValidationError{
					field:  "WriteTimeoutMs",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IoUringOptionsValidationError{
					field:  "WriteTimeoutMs",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWriteTimeoutMs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IoUringOptionsValidationError{
				field:  "WriteTimeoutMs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return IoUringOptionsMultiError(errors)
	}

	return nil
}

// IoUringOptionsMultiError is an error wrapping multiple validation errors
// returned by IoUringOptions.ValidateAll() if the designated constraints
// aren't met.
type IoUringOptionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IoUringOptionsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IoUringOptionsMultiError) AllErrors() []error { return m }

// IoUringOptionsValidationError is the validation error returned by
// IoUringOptions.Validate if the designated constraints aren't met.
type IoUringOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IoUringOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IoUringOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IoUringOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IoUringOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IoUringOptionsValidationError) ErrorName() string { return "IoUringOptionsValidationError" }

// Error satisfies the builtin error interface
func (e IoUringOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIoUringOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IoUringOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IoUringOptionsValidationError{}
