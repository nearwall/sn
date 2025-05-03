package logger

import (
	"context"
	"errors"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type logLevel string

const (
	debugLevel logLevel = "debug"
	infoLevel  logLevel = "info"
	warnLevel  logLevel = "warn"
	errorLevel logLevel = "error"
	fatalLevel logLevel = "fatal"
)

func (logger *ServiceLogger) otel(ctx context.Context, level logLevel, msg string, fields ...interface{}) {
	// Try to extract the current span from the context.
	span := trace.SpanFromContext(ctx)
	if span == nil || !span.SpanContext().IsValid() {
		return // no active span found, nothing to do
	}

	if level == errorLevel || level == fatalLevel {
		span.RecordError(errors.New(msg))
		span.SetStatus(codes.Error, msg)
	}

	// Convert zap fields into OTel attributes.
	attrs := fieldsToAttributes(string(level), fields...)

	// Add an event to the current span.
	span.AddEvent(msg, trace.WithAttributes(attrs...))
}

func fieldsToAttributes(level string, args ...interface{}) []attribute.KeyValue {
	var attrs []attribute.KeyValue

	if len(args)%2 != 0 {
		args = args[:len(args)-1]
	}

	attrs = append(attrs, attribute.String("level", level))

	for i := 0; i < len(args); i += 2 {
		attrs = append(attrs, attribute.String(fmt.Sprintf("%v", args[i]), fmt.Sprintf("%v", args[i+1])))
	}

	return attrs
}
