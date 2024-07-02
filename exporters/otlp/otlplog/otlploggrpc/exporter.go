// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlploggrpc // import "github.com/dmitryax/opentelemetry-go/exporters/otlp/otlplog/otlploggrpc"

import (
	"context"

	"github.com/dmitryax/opentelemetry-go/sdk/log"
)

// Exporter is a OpenTelemetry log Exporter. It transports log data encoded as
// OTLP protobufs using gRPC.
type Exporter struct {
	// TODO: implement.
}

// Compile-time check Exporter implements [log.Exporter].
var _ log.Exporter = (*Exporter)(nil)

// New returns a new [Exporter].
func New(_ context.Context, options ...Option) (*Exporter, error) {
	cfg := newConfig(options)
	c, err := newClient(cfg)
	if err != nil {
		return nil, err
	}
	return newExporter(c, cfg)
}

func newExporter(*client, config) (*Exporter, error) {
	// TODO: implement
	return &Exporter{}, nil
}

// Export transforms and transmits log records to an OTLP receiver.
func (e *Exporter) Export(ctx context.Context, records []log.Record) error {
	// TODO: implement.
	return nil
}

// Shutdown shuts down the Exporter. Calls to Export or ForceFlush will perform
// no operation after this is called.
func (e *Exporter) Shutdown(ctx context.Context) error {
	// TODO: implement.
	return nil
}

// ForceFlush does nothing. The Exporter holds no state.
func (e *Exporter) ForceFlush(ctx context.Context) error {
	// TODO: implement.
	return nil
}
