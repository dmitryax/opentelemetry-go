// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlptracegrpc_test

import (
	"context"

	"github.com/dmitryax/opentelemetry-go"
	"github.com/dmitryax/opentelemetry-go/exporters/otlp/otlptrace/otlptracegrpc"
	"github.com/dmitryax/opentelemetry-go/sdk/trace"
)

func Example() {
	ctx := context.Background()
	exp, err := otlptracegrpc.New(ctx)
	if err != nil {
		panic(err)
	}

	tracerProvider := trace.NewTracerProvider(trace.WithBatcher(exp))
	defer func() {
		if err := tracerProvider.Shutdown(ctx); err != nil {
			panic(err)
		}
	}()
	otel.SetTracerProvider(tracerProvider)

	// From here, the tracerProvider can be used by instrumentation to collect
	// telemetry.
}
