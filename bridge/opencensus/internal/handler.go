// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/dmitryax/opentelemetry-go/bridge/opencensus/internal"

import "github.com/dmitryax/opentelemetry-go"

// Handle is the package level function to handle errors. It can be
// overwritten for testing.
var Handle = otel.Handle
