package util

import (
	"testing"
	"github.com/nicolaferraro/integration-operator/pkg/apis/camel/v1alpha1"
	"github.com/stretchr/testify/assert"
)

func TestSerialize(t *testing.T) {

	routes := []v1alpha1.RouteSpec{
		{
			Route: []string{
				"timer:tick",
				"log:info",
			},
		},
	}

	serialized, err := Serialize(routes)
	assert.Nil(t, err)
	assert.Contains(t, serialized, "routes")
	assert.Contains(t, serialized, "timer:tick")
	assert.Contains(t, serialized, "log:info")
}