// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: service.proto

package pb

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

// define the regex for a UUID once up-front
var _service_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Post with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Post) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if utf8.RuneCountInString(m.GetTitle()) < 2 {
		return PostValidationError{
			field:  "Title",
			reason: "value length must be at least 2 runes",
		}
	}

	if utf8.RuneCountInString(m.GetDescription()) < 2 {
		return PostValidationError{
			field:  "Description",
			reason: "value length must be at least 2 runes",
		}
	}

	if v, ok := interface{}(m.GetCreatedOn()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PostValidationError{
				field:  "CreatedOn",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedOn()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PostValidationError{
				field:  "UpdatedOn",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PostValidationError is the validation error returned by Post.Validate if the
// designated constraints aren't met.
type PostValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PostValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PostValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PostValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PostValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PostValidationError) ErrorName() string { return "PostValidationError" }

// Error satisfies the builtin error interface
func (e PostValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPost.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PostValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PostValidationError{}

// Validate checks the field values on GetPostRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetPostRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return GetPostRequestValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// GetPostRequestValidationError is the validation error returned by
// GetPostRequest.Validate if the designated constraints aren't met.
type GetPostRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPostRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPostRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPostRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPostRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPostRequestValidationError) ErrorName() string { return "GetPostRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetPostRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPostRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPostRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPostRequestValidationError{}

// Validate checks the field values on UpdatePostRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UpdatePostRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return UpdatePostRequestValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetPost() == nil {
		return UpdatePostRequestValidationError{
			field:  "Post",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetPost()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdatePostRequestValidationError{
				field:  "Post",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdateMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdatePostRequestValidationError{
				field:  "UpdateMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdatePostRequestValidationError is the validation error returned by
// UpdatePostRequest.Validate if the designated constraints aren't met.
type UpdatePostRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePostRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePostRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePostRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePostRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePostRequestValidationError) ErrorName() string {
	return "UpdatePostRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePostRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePostRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePostRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePostRequestValidationError{}

// Validate checks the field values on DeletePostRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DeletePostRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return DeletePostRequestValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// DeletePostRequestValidationError is the validation error returned by
// DeletePostRequest.Validate if the designated constraints aren't met.
type DeletePostRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeletePostRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeletePostRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeletePostRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeletePostRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeletePostRequestValidationError) ErrorName() string {
	return "DeletePostRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeletePostRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeletePostRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeletePostRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeletePostRequestValidationError{}

// Validate checks the field values on ListPostResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListPostResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPosts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListPostResponseValidationError{
					field:  fmt.Sprintf("Posts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListPostResponseValidationError is the validation error returned by
// ListPostResponse.Validate if the designated constraints aren't met.
type ListPostResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPostResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPostResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPostResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPostResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPostResponseValidationError) ErrorName() string { return "ListPostResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListPostResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPostResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPostResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPostResponseValidationError{}
