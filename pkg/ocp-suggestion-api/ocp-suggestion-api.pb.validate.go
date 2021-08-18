// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/ocp-suggestion-api/ocp-suggestion-api.proto

package ocp_suggestion_api

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
)

// Validate checks the field values on Suggestion with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Suggestion) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 0 {
		return SuggestionValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
	}

	if m.GetUserId() <= 0 {
		return SuggestionValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	if m.GetCourseId() <= 0 {
		return SuggestionValidationError{
			field:  "CourseId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// SuggestionValidationError is the validation error returned by
// Suggestion.Validate if the designated constraints aren't met.
type SuggestionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SuggestionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SuggestionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SuggestionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SuggestionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SuggestionValidationError) ErrorName() string { return "SuggestionValidationError" }

// Error satisfies the builtin error interface
func (e SuggestionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSuggestion.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SuggestionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SuggestionValidationError{}

// Validate checks the field values on CreateSuggestionV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateSuggestionV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUserId() <= 0 {
		return CreateSuggestionV1RequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	if m.GetCourseId() <= 0 {
		return CreateSuggestionV1RequestValidationError{
			field:  "CourseId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// CreateSuggestionV1RequestValidationError is the validation error returned by
// CreateSuggestionV1Request.Validate if the designated constraints aren't met.
type CreateSuggestionV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSuggestionV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSuggestionV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSuggestionV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSuggestionV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSuggestionV1RequestValidationError) ErrorName() string {
	return "CreateSuggestionV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSuggestionV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSuggestionV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSuggestionV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSuggestionV1RequestValidationError{}

// Validate checks the field values on CreateSuggestionV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateSuggestionV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for SuggestionId

	return nil
}

// CreateSuggestionV1ResponseValidationError is the validation error returned
// by CreateSuggestionV1Response.Validate if the designated constraints aren't met.
type CreateSuggestionV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSuggestionV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSuggestionV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSuggestionV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSuggestionV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSuggestionV1ResponseValidationError) ErrorName() string {
	return "CreateSuggestionV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSuggestionV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSuggestionV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSuggestionV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSuggestionV1ResponseValidationError{}

// Validate checks the field values on DescribeSuggestionV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeSuggestionV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetSuggestionId() <= 0 {
		return DescribeSuggestionV1RequestValidationError{
			field:  "SuggestionId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DescribeSuggestionV1RequestValidationError is the validation error returned
// by DescribeSuggestionV1Request.Validate if the designated constraints
// aren't met.
type DescribeSuggestionV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeSuggestionV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeSuggestionV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeSuggestionV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeSuggestionV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeSuggestionV1RequestValidationError) ErrorName() string {
	return "DescribeSuggestionV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeSuggestionV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeSuggestionV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeSuggestionV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeSuggestionV1RequestValidationError{}

// Validate checks the field values on DescribeSuggestionV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeSuggestionV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSuggestion()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeSuggestionV1ResponseValidationError{
				field:  "Suggestion",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeSuggestionV1ResponseValidationError is the validation error returned
// by DescribeSuggestionV1Response.Validate if the designated constraints
// aren't met.
type DescribeSuggestionV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeSuggestionV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeSuggestionV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeSuggestionV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeSuggestionV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeSuggestionV1ResponseValidationError) ErrorName() string {
	return "DescribeSuggestionV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeSuggestionV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeSuggestionV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeSuggestionV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeSuggestionV1ResponseValidationError{}

// Validate checks the field values on ListSuggestionV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListSuggestionV1Request) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListSuggestionV1RequestValidationError is the validation error returned by
// ListSuggestionV1Request.Validate if the designated constraints aren't met.
type ListSuggestionV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSuggestionV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSuggestionV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSuggestionV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSuggestionV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSuggestionV1RequestValidationError) ErrorName() string {
	return "ListSuggestionV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListSuggestionV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSuggestionV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSuggestionV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSuggestionV1RequestValidationError{}

// Validate checks the field values on ListSuggestionV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListSuggestionV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetSuggestions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListSuggestionV1ResponseValidationError{
					field:  fmt.Sprintf("Suggestions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListSuggestionV1ResponseValidationError is the validation error returned by
// ListSuggestionV1Response.Validate if the designated constraints aren't met.
type ListSuggestionV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSuggestionV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSuggestionV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSuggestionV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSuggestionV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSuggestionV1ResponseValidationError) ErrorName() string {
	return "ListSuggestionV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListSuggestionV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSuggestionV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSuggestionV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSuggestionV1ResponseValidationError{}

// Validate checks the field values on UpdateSuggestionV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateSuggestionV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSuggestion()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateSuggestionV1RequestValidationError{
				field:  "Suggestion",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateSuggestionV1RequestValidationError is the validation error returned by
// UpdateSuggestionV1Request.Validate if the designated constraints aren't met.
type UpdateSuggestionV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateSuggestionV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateSuggestionV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateSuggestionV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateSuggestionV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateSuggestionV1RequestValidationError) ErrorName() string {
	return "UpdateSuggestionV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateSuggestionV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateSuggestionV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateSuggestionV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateSuggestionV1RequestValidationError{}

// Validate checks the field values on UpdateSuggestionV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateSuggestionV1Response) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateSuggestionV1ResponseValidationError is the validation error returned
// by UpdateSuggestionV1Response.Validate if the designated constraints aren't met.
type UpdateSuggestionV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateSuggestionV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateSuggestionV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateSuggestionV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateSuggestionV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateSuggestionV1ResponseValidationError) ErrorName() string {
	return "UpdateSuggestionV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateSuggestionV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateSuggestionV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateSuggestionV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateSuggestionV1ResponseValidationError{}

// Validate checks the field values on RemoveSuggestionV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveSuggestionV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetSuggestionId() <= 0 {
		return RemoveSuggestionV1RequestValidationError{
			field:  "SuggestionId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// RemoveSuggestionV1RequestValidationError is the validation error returned by
// RemoveSuggestionV1Request.Validate if the designated constraints aren't met.
type RemoveSuggestionV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveSuggestionV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveSuggestionV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveSuggestionV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveSuggestionV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveSuggestionV1RequestValidationError) ErrorName() string {
	return "RemoveSuggestionV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveSuggestionV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveSuggestionV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveSuggestionV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveSuggestionV1RequestValidationError{}

// Validate checks the field values on RemoveSuggestionV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveSuggestionV1Response) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// RemoveSuggestionV1ResponseValidationError is the validation error returned
// by RemoveSuggestionV1Response.Validate if the designated constraints aren't met.
type RemoveSuggestionV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveSuggestionV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveSuggestionV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveSuggestionV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveSuggestionV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveSuggestionV1ResponseValidationError) ErrorName() string {
	return "RemoveSuggestionV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveSuggestionV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveSuggestionV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveSuggestionV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveSuggestionV1ResponseValidationError{}