// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tracetransform

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dmitryax/opentelemetry-go/attribute"
	"github.com/dmitryax/opentelemetry-go/sdk/resource"
)

func TestNilResource(t *testing.T) {
	assert.Empty(t, Resource(nil))
}

func TestEmptyResource(t *testing.T) {
	assert.Empty(t, Resource(&resource.Resource{}))
}

/*
* This does not include any testing on the ordering of Resource Attributes.
* They are stored as a map internally to the Resource and their order is not
* guaranteed.
 */

func TestResourceAttributes(t *testing.T) {
	attrs := []attribute.KeyValue{attribute.Int("one", 1), attribute.Int("two", 2)}

	got := Resource(resource.NewSchemaless(attrs...)).GetAttributes()
	if !assert.Len(t, attrs, 2) {
		return
	}
	assert.ElementsMatch(t, KeyValues(attrs), got)
}
