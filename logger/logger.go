// Package logger
package logger

import (
	"time"
)

type Logger interface {
	Info(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
	Warn(msg string, fields ...Field)

	Fatal(msg string, fields ...Field)
	Panic(msg string, fields ...Field)

	Infof(format string, args ...any)
	Errorf(format string, args ...any)
	Debugf(format string, args ...any)
	Warnf(format string, args ...any)
	Fatalf(format string, args ...any)
	Panicf(format string, args ...any)

	WithFields(fields ...Field) Logger

	Sync() error
}

type fieldType int

const (
	fieldAny fieldType = iota
	fieldError
	fieldString
	fieldInt
	fieldInt64
	fieldBool
	fieldFloat64
	fieldDuration
	fieldTime
)

type Field struct {
	Key  string
	Type fieldType
	Any  any
	Err  error
}

func Any(key string, value any) Field {
	return Field{Key: key, Type: fieldAny, Any: value}
}

func Error(err error) Field {
	return Field{Key: "error", Type: fieldError, Err: err}
}

func String(key, value string) Field {
	return Field{Key: key, Type: fieldString, Any: value}
}

func Int(key string, value int) Field {
	return Field{Key: key, Type: fieldInt, Any: value}
}

func Int64(key string, value int64) Field {
	return Field{Key: key, Type: fieldInt64, Any: value}
}

func Bool(key string, value bool) Field {
	return Field{Key: key, Type: fieldBool, Any: value}
}

func Float64(key string, value float64) Field {
	return Field{Key: key, Type: fieldFloat64, Any: value}
}

func Duration(key string, value time.Duration) Field {
	return Field{Key: key, Type: fieldDuration, Any: value}
}

func Time(key string, value time.Time) Field {
	return Field{Key: key, Type: fieldTime, Any: value}
}
