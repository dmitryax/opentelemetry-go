// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package propagation_test

import (
	"github.com/dmitryax/opentelemetry-go"
	"github.com/dmitryax/opentelemetry-go/propagation"
)

func ExampleTraceContext() {
	tc := propagation.TraceContext{}
	// Register the TraceContext propagator globally.
	otel.SetTextMapPropagator(tc)
}
