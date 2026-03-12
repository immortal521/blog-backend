// Package errs
package errs

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

// AppError represents an application-specific error with a code and a message.
// It wraps the underlying error and captures the stack trace for debugging purposes.
type AppError struct {
	Code  int
	Msg   string
	Err   error
	stack []uintptr
}

func New(code int, msg string, err error) *AppError {
	if ae, ok := err.(*AppError); ok {
		return ae
	}

	pcs := make([]uintptr, 32)
	n := runtime.Callers(2, pcs)
	pcs = pcs[:n]
	return &AppError{
		Code:  code,
		Msg:   msg,
		Err:   err,
		stack: pcs,
	}
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%d] %s: %+v", e.Code, e.Msg, e.Err)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Msg)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) StackString() string {
	if len(e.stack) == 0 {
		return ""
	}
	frames := runtime.CallersFrames(e.stack)
	var s strings.Builder
	for {
		frames, more := frames.Next()
		fmt.Fprintf(&s, "%s\n\t%s:%d\n", frames.Function, frames.File, frames.Line)
		if !more {
			break
		}
	}
	return s.String()
}

func (e *AppError) FormatStack() string {
	if e == nil {
		return ""
	}
	s := e.Error()
	stackStr := e.StackString()
	if stackStr != "" {
		s += "\n Stack trace:\n" + stackStr
	}
	if e.Err != nil {
		if nested, ok := e.Err.(*AppError); ok {
			s += "\nCaused by:\n" + nested.FormatStack()
		}
	}
	return s
}

func ToAppError(err error) *AppError {
	if err == nil {
		return nil
	}

	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr
	}

	return New(CodeInternalError, "Internal Server Error", err)
}
