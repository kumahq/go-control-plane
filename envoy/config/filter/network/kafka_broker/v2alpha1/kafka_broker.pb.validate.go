//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/network/kafka_broker/v2alpha1/kafka_broker.proto

package v2alpha1

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

// Validate checks the field values on KafkaBroker with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *KafkaBroker) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on KafkaBroker with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in KafkaBrokerMultiError, or
// nil if none found.
func (m *KafkaBroker) ValidateAll() error {
	return m.validate(true)
}

func (m *KafkaBroker) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetStatPrefix()) < 1 {
		err := KafkaBrokerValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for ForceResponseRewrite

	for idx, item := range m.GetApiKeysAllowed() {
		_, _ = idx, item

		if val := item; val < 0 || val > 32767 {
			err := KafkaBrokerValidationError{
				field:  fmt.Sprintf("ApiKeysAllowed[%v]", idx),
				reason: "value must be inside range [0, 32767]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	for idx, item := range m.GetApiKeysDenied() {
		_, _ = idx, item

		if val := item; val < 0 || val > 32767 {
			err := KafkaBrokerValidationError{
				field:  fmt.Sprintf("ApiKeysDenied[%v]", idx),
				reason: "value must be inside range [0, 32767]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	switch v := m.BrokerAddressRewriteSpec.(type) {
	case *KafkaBroker_IdBasedBrokerAddressRewriteSpec:
		if v == nil {
			err := KafkaBrokerValidationError{
				field:  "BrokerAddressRewriteSpec",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetIdBasedBrokerAddressRewriteSpec()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, KafkaBrokerValidationError{
						field:  "IdBasedBrokerAddressRewriteSpec",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, KafkaBrokerValidationError{
						field:  "IdBasedBrokerAddressRewriteSpec",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetIdBasedBrokerAddressRewriteSpec()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return KafkaBrokerValidationError{
					field:  "IdBasedBrokerAddressRewriteSpec",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return KafkaBrokerMultiError(errors)
	}

	return nil
}

// KafkaBrokerMultiError is an error wrapping multiple validation errors
// returned by KafkaBroker.ValidateAll() if the designated constraints aren't met.
type KafkaBrokerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m KafkaBrokerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m KafkaBrokerMultiError) AllErrors() []error { return m }

// KafkaBrokerValidationError is the validation error returned by
// KafkaBroker.Validate if the designated constraints aren't met.
type KafkaBrokerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KafkaBrokerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KafkaBrokerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KafkaBrokerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KafkaBrokerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KafkaBrokerValidationError) ErrorName() string { return "KafkaBrokerValidationError" }

// Error satisfies the builtin error interface
func (e KafkaBrokerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKafkaBroker.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KafkaBrokerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KafkaBrokerValidationError{}

// Validate checks the field values on IdBasedBrokerRewriteSpec with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IdBasedBrokerRewriteSpec) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IdBasedBrokerRewriteSpec with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IdBasedBrokerRewriteSpecMultiError, or nil if none found.
func (m *IdBasedBrokerRewriteSpec) ValidateAll() error {
	return m.validate(true)
}

func (m *IdBasedBrokerRewriteSpec) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, IdBasedBrokerRewriteSpecValidationError{
						field:  fmt.Sprintf("Rules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, IdBasedBrokerRewriteSpecValidationError{
						field:  fmt.Sprintf("Rules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IdBasedBrokerRewriteSpecValidationError{
					field:  fmt.Sprintf("Rules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return IdBasedBrokerRewriteSpecMultiError(errors)
	}

	return nil
}

// IdBasedBrokerRewriteSpecMultiError is an error wrapping multiple validation
// errors returned by IdBasedBrokerRewriteSpec.ValidateAll() if the designated
// constraints aren't met.
type IdBasedBrokerRewriteSpecMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdBasedBrokerRewriteSpecMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdBasedBrokerRewriteSpecMultiError) AllErrors() []error { return m }

// IdBasedBrokerRewriteSpecValidationError is the validation error returned by
// IdBasedBrokerRewriteSpec.Validate if the designated constraints aren't met.
type IdBasedBrokerRewriteSpecValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdBasedBrokerRewriteSpecValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdBasedBrokerRewriteSpecValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdBasedBrokerRewriteSpecValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdBasedBrokerRewriteSpecValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdBasedBrokerRewriteSpecValidationError) ErrorName() string {
	return "IdBasedBrokerRewriteSpecValidationError"
}

// Error satisfies the builtin error interface
func (e IdBasedBrokerRewriteSpecValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdBasedBrokerRewriteSpec.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdBasedBrokerRewriteSpecValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdBasedBrokerRewriteSpecValidationError{}

// Validate checks the field values on IdBasedBrokerRewriteRule with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IdBasedBrokerRewriteRule) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IdBasedBrokerRewriteRule with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IdBasedBrokerRewriteRuleMultiError, or nil if none found.
func (m *IdBasedBrokerRewriteRule) ValidateAll() error {
	return m.validate(true)
}

func (m *IdBasedBrokerRewriteRule) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() < 0 {
		err := IdBasedBrokerRewriteRuleValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetHost()) < 1 {
		err := IdBasedBrokerRewriteRuleValidationError{
			field:  "Host",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPort() > 65535 {
		err := IdBasedBrokerRewriteRuleValidationError{
			field:  "Port",
			reason: "value must be less than or equal to 65535",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return IdBasedBrokerRewriteRuleMultiError(errors)
	}

	return nil
}

// IdBasedBrokerRewriteRuleMultiError is an error wrapping multiple validation
// errors returned by IdBasedBrokerRewriteRule.ValidateAll() if the designated
// constraints aren't met.
type IdBasedBrokerRewriteRuleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdBasedBrokerRewriteRuleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdBasedBrokerRewriteRuleMultiError) AllErrors() []error { return m }

// IdBasedBrokerRewriteRuleValidationError is the validation error returned by
// IdBasedBrokerRewriteRule.Validate if the designated constraints aren't met.
type IdBasedBrokerRewriteRuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdBasedBrokerRewriteRuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdBasedBrokerRewriteRuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdBasedBrokerRewriteRuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdBasedBrokerRewriteRuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdBasedBrokerRewriteRuleValidationError) ErrorName() string {
	return "IdBasedBrokerRewriteRuleValidationError"
}

// Error satisfies the builtin error interface
func (e IdBasedBrokerRewriteRuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdBasedBrokerRewriteRule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdBasedBrokerRewriteRuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdBasedBrokerRewriteRuleValidationError{}
