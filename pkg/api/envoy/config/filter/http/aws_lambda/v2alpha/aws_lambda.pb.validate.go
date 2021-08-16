// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/http/aws_lambda/v2alpha/aws_lambda.proto

package envoy_config_filter_http_aws_lambda_v2alpha

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Config) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetArn()) < 1 {
		return ConfigValidationError{
			field:  "Arn",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for PayloadPassthrough

	if _, ok := Config_InvocationMode_name[int32(m.GetInvocationMode())]; !ok {
		return ConfigValidationError{
			field:  "InvocationMode",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// ConfigValidationError is the validation error returned by Config.Validate if
// the designated constraints aren't met.
type ConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigValidationError) ErrorName() string { return "ConfigValidationError" }

// Error satisfies the builtin error interface
func (e ConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigValidationError{}

// Validate checks the field values on PerRouteConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *PerRouteConfig) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetInvokeConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PerRouteConfigValidationError{
				field:  "InvokeConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PerRouteConfigValidationError is the validation error returned by
// PerRouteConfig.Validate if the designated constraints aren't met.
type PerRouteConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PerRouteConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PerRouteConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PerRouteConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PerRouteConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PerRouteConfigValidationError) ErrorName() string { return "PerRouteConfigValidationError" }

// Error satisfies the builtin error interface
func (e PerRouteConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPerRouteConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PerRouteConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PerRouteConfigValidationError{}
