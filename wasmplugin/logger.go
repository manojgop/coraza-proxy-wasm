// Copyright The OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package wasmplugin

import (
	"io"

	"github.com/corazawaf/coraza/v3/loggers"
	"github.com/rs/zerolog"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
)

type debugLogger struct {
	level  loggers.LogLevel
	logger *zerolog.Logger
}

var _ loggers.DebugLogger = (*debugLogger)(nil)

func (l *debugLogger) Info(message string, args ...interface{}) {
	if l.level >= loggers.LogLevelInfo {
		proxywasm.LogInfof(message, args...)
	}
}

func (l *debugLogger) Warn(message string, args ...interface{}) {
	if l.level >= loggers.LogLevelWarn {
		proxywasm.LogWarnf(message, args...)
	}
}

func (l *debugLogger) Error(message string, args ...interface{}) {
	if l.level >= loggers.LogLevelError {
		proxywasm.LogErrorf(message, args...)
	}
}

func (l *debugLogger) Debug(message string, args ...interface{}) {
	if l.level >= loggers.LogLevelDebug {
		proxywasm.LogDebugf(message, args...)
	}
}

func (l *debugLogger) Trace(message string, args ...interface{}) {
	if l.level >= loggers.LogLevelTrace {
		proxywasm.LogTracef(message, args...)
	}
}

func (l *debugLogger) SetLevel(level loggers.LogLevel) {
	l.level = level
	zerolog.SetGlobalLevel(getZeroLogLevel(level))
}

func (l *debugLogger) SetOutput(w io.WriteCloser) {
	proxywasm.LogWarn("ignoring SecDebugLog directive, debug logs are always routed to proxy logs")
}

func (l *debugLogger) GetLogger() *zerolog.Logger {
	return l.logger
}

func getZeroLogLevel(level loggers.LogLevel) zerolog.Level {
	switch level {
	case loggers.LogLevelNoLog:
		return zerolog.NoLevel
	case loggers.LogLevelError:
		return zerolog.ErrorLevel
	case loggers.LogLevelWarn:
		return zerolog.WarnLevel
	case loggers.LogLevelInfo:
		return zerolog.InfoLevel
	case loggers.LogLevelDebug:
		return zerolog.DebugLevel
	case loggers.LogLevelTrace:
		return zerolog.TraceLevel
	default:
		return zerolog.InfoLevel
	}
}
