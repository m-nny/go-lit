package lit

import (
	"errors"
	"fmt"
)

const (
	ECONFLICT       = "conflict"
	EINTERNAL       = "internal"
	EINVALID        = "invalid"
	ENOTFOUND       = "not_found"
	ENOTIMPLEMENTED = "not_implemented"
	EUNAUHORIZED    = "unauhorized"
)

type Error struct {
	// Machine-readable error code.
	Code string

	// User-readable error code.
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("wtf error: code=%s message=%s", e.Code, e.Message)
}

func ErrorCode(err error) string {
	var e *Error
	if err == nil {
		return ""
	} else if errors.As(err, &e) {
		return e.Code
	}
	return EINTERNAL
}

func ErrorMessage(err error) string {
	var e *Error
	if err == nil {
		return ""
	} else if errors.As(err, &e) {
		return e.Message
	}
	return "Internal error."
}

func ErrorF(code string, format string, args ...interface{}) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}
