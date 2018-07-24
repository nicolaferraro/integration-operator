package util

import (
	"testing"
	"github.com/nicolaferraro/integration-operator/pkg/apis/camel/v1alpha1"
	"github.com/stretchr/testify/assert"
)

func TestSerialize(t *testing.T) {
	replicas := int32(1)
	integration := v1alpha1.IntegrationSpec{
		Replicas: &replicas,
	}

	serialized, err := Serialize(integration)
	assert.Nil(t, err)
	assert.Contains(t, serialized, "replicas")
	assert.Contains(t, serialized, "1")
}