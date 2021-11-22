package web

import "time"

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

func NewError(code ErrorCode, msg string) *Error {
	return &Error{
		Code:  code,
		Msg:   msg,
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

func NewErrorWithCauseAndMsg(code ErrorCode, cause error, msg string) *Error {
	return &Error{
		Code:  code,
		Msg:   msg,
		Cause: cause,
		Time:  time.Now(),
	}
}

func (a *Error) Error() string {
	if a.Msg != "" {
		return a.Msg
	}
	return a.Cause.Error()
}
