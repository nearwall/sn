package logger

import (
	"context"
	"runtime"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	Logger interface {
		Errorf(ctx context.Context, msg string, args ...any)
		Warnf(ctx context.Context, msg string, args ...any)
		Infof(ctx context.Context, msg string, args ...any)
		Debugf(ctx context.Context, msg string, args ...any)
		Fatalf(ctx context.Context, msg string, args ...any)

		Error(ctx context.Context, msg string, args ...any)
		Warn(ctx context.Context, msg string, args ...any)
		Info(ctx context.Context, msg string, args ...any)
		Debug(ctx context.Context, msg string, args ...any)
		Fatal(ctx context.Context, msg string, args ...any)
	}

	ServiceLogger struct {
		*zap.Logger
		serviceName string
	}
)

func (logger *ServiceLogger) Errorf(
	ctx context.Context,
	msg string,
	args ...any,
) {
	logger.prepare(ctx).Sugar().Errorf(msg, args...)
	logger.otel(ctx, errorLevel, msg, args...)
}

func (logger *ServiceLogger) Warnf(ctx context.Context, msg string, args ...any) {
	logger.prepare(ctx).Sugar().Warnf(msg, args...)
	logger.otel(ctx, warnLevel, msg, args...)
}

func (logger *ServiceLogger) Infof(ctx context.Context, msg string, args ...any) {
	logger.prepare(ctx).Sugar().Infof(msg, args...)
	logger.otel(ctx, infoLevel, msg, args...)
}

func (logger *ServiceLogger) Debugf(ctx context.Context, msg string, args ...any) {
	logger.prepare(ctx).Sugar().Debugf(msg, args...)
	logger.otel(ctx, debugLevel, msg, args...)
}

func (logger *ServiceLogger) Fatalf(ctx context.Context, msg string, args ...any) {
	logger.prepare(ctx).Sugar().Fatalf(msg, args...)
	logger.otel(ctx, fatalLevel, msg, args...)
}

func (logger *ServiceLogger) Error(
	ctx context.Context,
	msg string,
	args ...any,
) {
	logger.prepare(ctx).Sugar().Errorw(msg, args...)
	logger.otel(ctx, errorLevel, msg, args...)
}

func (logger *ServiceLogger) Warn(ctx context.Context, msg string, args ...any) {
	logger.prepare(ctx).Sugar().Warnw(msg, args...)
	logger.otel(ctx, warnLevel, msg, args...)
}

func (logger *ServiceLogger) Info(ctx context.Context, msg string, args ...any) {
	logger.prepare(ctx).Sugar().Infow(msg, args...)
	logger.otel(ctx, infoLevel, msg, args...)
}

func (logger *ServiceLogger) Debug(ctx context.Context, msg string, args ...any) {
	logger.prepare(ctx).Sugar().Debugw(msg, args...)
	logger.otel(ctx, debugLevel, msg, args...)
}

func (logger *ServiceLogger) Fatal(ctx context.Context, msg string, args ...any) {
	logger.prepare(ctx).Sugar().Fatalw(msg, args...)
	logger.otel(ctx, fatalLevel, msg, args...)
}

func (logger *ServiceLogger) prepare(
	ctx context.Context,
	additionalFields ...zap.Field,
) *zap.Logger {
	reqID, ok := ctx.Value(RequestIDLabel).(string)
	if !ok || reqID == "" {
		reqID = "n/a"
	}

	traceID, ok := ctx.Value(TraceIDLabel).(string)
	if !ok || traceID == "" {
		traceID = "n/a"
	}

	_, caller, line, ok := runtime.Caller(2)
	if !ok {
		caller = "n/a"
		line = 0
	}

	additionalFields = append(additionalFields,
		zap.String("req_id", reqID),
		zap.String("service_name", logger.serviceName),
		zap.String("caller", caller+":"+strconv.Itoa(line)),
		zap.String("trace_id", traceID),
	)

	return logger.With(additionalFields...)
}

var log *ServiceLogger

func Log() Logger {
	if log == nil || log.Logger == nil {
		SetDebugLogger("service_name_not_set", true)
	}
	return log
}

func SetProductionLogger(serviceName string) {
	log = newProductionLogger(serviceName)
}

func SetDebugLogger(serviceName string, disableStackTrace bool) {
	log = newDebugLogger(serviceName, disableStackTrace)
}

func SetPepeLogger(debugMode, disableStackTrace bool) {
	log = newPepeLogger(debugMode, disableStackTrace)
}

func newProductionLogger(serviceName string) *ServiceLogger {
	cfg := zap.NewProductionConfig()

	cfg.DisableStacktrace = true
	cfg.DisableCaller = true
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	return &ServiceLogger{zap.Must(cfg.Build()), serviceName}
}

func newDebugLogger(serviceName string, disableStackTrace bool) *ServiceLogger {
	cfg := zap.NewProductionConfig()

	cfg.DisableStacktrace = disableStackTrace
	cfg.DisableCaller = true
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	cfg.Level.SetLevel(zap.DebugLevel)

	return &ServiceLogger{zap.Must(cfg.Build()), serviceName}
}

func newPepeLogger(debugMode, disableStackTrace bool) *ServiceLogger {
	cfg := zap.NewDevelopmentConfig()

	cfg.DisableStacktrace = disableStackTrace
	cfg.DisableCaller = true
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.ConsoleSeparator = "\t"
	if debugMode {
		cfg.Level.SetLevel(zap.DebugLevel)
	}

	return &ServiceLogger{zap.Must(cfg.Build()), "cli-tool"}
}
