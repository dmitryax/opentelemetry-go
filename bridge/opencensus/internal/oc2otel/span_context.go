// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package oc2otel // import "github.com/dmitryax/opentelemetry-go/bridge/opencensus/internal/oc2otel"

import (
	octrace "go.opencensus.io/trace"

	"github.com/dmitryax/opentelemetry-go/trace"
)

func SpanContext(sc octrace.SpanContext) trace.SpanContext {
	var traceFlags trace.TraceFlags
	if sc.IsSampled() {
		traceFlags = trace.FlagsSampled
	}
	return trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    trace.TraceID(sc.TraceID),
		SpanID:     trace.SpanID(sc.SpanID),
		TraceFlags: traceFlags,
	})
}
