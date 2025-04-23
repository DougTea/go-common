package web

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type ErrorCode string

const CommonError ErrorCode = "500000"

func (e ErrorCode) String() string {
	return string(e)
}

type Error struct {
	Code  ErrorCode
	Msg   string
	Cause error
	Time  time.Time
}

func NewError(code ErrorCode, msg string, args ...interface{}) *Error {
	return &Error{
		Code:  code,
		Msg:   fmt.Sprintf(msg, args...),
		Cause: nil,
		Time:  time.Now(),
	}
}

func NewErrorWithCause(code ErrorCode, cause error) *Error {
	return &Error{
		Code:  code,
		Msg:   "",
		Cause: cause,
		Time:  time.Now(),
	}
}

func NewErrorWithCauseAndMsg(code ErrorCode, cause error, msg string, args ...interface{}) *Error {
	return &Error{
		Code:  code,
		Msg:   fmt.Sprintf(msg, args...),
		Cause: cause,
		Time:  time.Now(),
	}
}

func (e *Error) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	return e.Cause.Error()
}

func (e *Error) Unwrap() error {
	return e.Cause
}

func (e *Error) ErrorWithCause() string {
	msg := new(strings.Builder)
	msg.WriteString(e.Msg)
	for err := e.Cause; err != nil; err = errors.Unwrap(err) {
		msg.WriteString("caused by: ")
		msg.WriteString(err.Error())
	}
	return msg.String()
}
