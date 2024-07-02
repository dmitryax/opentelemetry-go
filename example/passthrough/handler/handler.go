// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package handler // import "github.com/dmitryax/opentelemetry-go/example/passthrough/handler"

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/dmitryax/opentelemetry-go"
	"github.com/dmitryax/opentelemetry-go/propagation"
	"github.com/dmitryax/opentelemetry-go/trace"
)

// Handler is a minimal implementation of the handler and client from
// go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp for demonstration purposes.
// It handles an incoming http request, and makes an outgoing http request.
type Handler struct {
	propagators propagation.TextMapPropagator
	tracer      trace.Tracer
	next        func(r *http.Request)
}

// New returns a new Handler that will trace requests before handing them off
// to next.
func New(next func(r *http.Request)) *Handler {
	// Like most instrumentation packages, this handler defaults to using the
	// global progatators and tracer providers.
	return &Handler{
		propagators: otel.GetTextMapPropagator(),
		tracer:      otel.Tracer("examples/passthrough/handler"),
		next:        next,
	}
}

// HandleHTTPReq mimics what an instrumented http server does.
func (h *Handler) HandleHTTPReq(r *http.Request) {
	ctx := h.propagators.Extract(r.Context(), propagation.HeaderCarrier(r.Header))
	var span trace.Span
	log.Println("The \"handle passthrough request\" span should NOT be recorded, because it is recorded by a TracerProvider not backed by the SDK.")
	ctx, span = h.tracer.Start(ctx, "handle passthrough request")
	defer span.End()

	// Pretend to do work
	time.Sleep(time.Second)

	h.makeOutgoingRequest(ctx)
}

// makeOutgoingRequest mimics what an instrumented http client does.
func (h *Handler) makeOutgoingRequest(ctx context.Context) {
	// make a new http request
	r, err := http.NewRequest("", "", nil)
	if err != nil {
		panic(err)
	}

	log.Println("The \"make outgoing request from passthrough\" span should NOT be recorded, because it is recorded by a TracerProvider not backed by the SDK.")
	ctx, span := h.tracer.Start(ctx, "make outgoing request from passthrough")
	defer span.End()
	r = r.WithContext(ctx)
	h.propagators.Inject(ctx, propagation.HeaderCarrier(r.Header))
	h.next(r)
}
