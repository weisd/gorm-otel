package gormotel

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const (
	tracerName = "gorm-otel"
)

type options struct {
	tracer trace.Tracer

	// logResult means log SQL operation result into span log which causes span size grows up.
	// This is advised to only open in developing environment.
	logResult bool

	// Whether to log statement parameters or leave placeholders in the queries.
	logSqlParameters bool

	ignoreErrNotFound bool
}

func defaultOption() *options {
	return &options{
		tracer:            otel.GetTracerProvider().Tracer(tracerName),
		logResult:         false,
		logSqlParameters:  true,
		ignoreErrNotFound: true,
	}
}

type applyOption func(o *options)

// WithTracerProvider specifies a tracer provider to use for creating a tracer.
// If none is specified, the global provider is used.
func WithTracerProvider(provider trace.TracerProvider) applyOption {
	return func(o *options) {
		o.tracer = provider.Tracer(tracerName)
	}
}

// WithLogResult enable otelPlugin to log the result of each executed sql.
func WithLogResult(logResult bool) applyOption {
	return func(o *options) {
		o.logResult = logResult
	}
}

func WithSqlParameters(logSqlParameters bool) applyOption {
	return func(o *options) {
		o.logSqlParameters = logSqlParameters
	}
}

func WithIgnoreErrNotFound(ignoreErrNotFound bool) applyOption {
	return func(o *options) {
		o.ignoreErrNotFound = ignoreErrNotFound
	}
}
