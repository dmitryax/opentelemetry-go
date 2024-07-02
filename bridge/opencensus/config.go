// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package opencensus // import "github.com/dmitryax/opentelemetry-go/bridge/opencensus"

import (
	"github.com/dmitryax/opentelemetry-go"
	"github.com/dmitryax/opentelemetry-go/trace"
)

const scopeName = "github.com/dmitryax/opentelemetry-go/bridge/opencensus"

// newTraceConfig returns a config configured with options.
func newTraceConfig(options []TraceOption) traceConfig {
	conf := traceConfig{tp: otel.GetTracerProvider()}
	for _, o := range options {
		conf = o.apply(conf)
	}
	return conf
}

type traceConfig struct {
	tp trace.TracerProvider
}

// TraceOption applies a configuration option value to an OpenCensus bridge
// Tracer.
type TraceOption interface {
	apply(traceConfig) traceConfig
}

// traceOptionFunc applies a set of options to a config.
type traceOptionFunc func(traceConfig) traceConfig

// apply returns a config with option(s) applied.
func (o traceOptionFunc) apply(conf traceConfig) traceConfig {
	return o(conf)
}

// WithTracerProvider specifies a tracer provider to use for creating a tracer.
func WithTracerProvider(tp trace.TracerProvider) TraceOption {
	return traceOptionFunc(func(conf traceConfig) traceConfig {
		conf.tp = tp
		return conf
	})
}

type metricConfig struct{}

// MetricOption applies a configuration option value to an OpenCensus bridge
// MetricProducer.
type MetricOption interface {
	apply(metricConfig) metricConfig
}
