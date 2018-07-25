package runtime

import (
	"github.com/nicolaferraro/integration-operator/pkg/runtime/api"
	"github.com/nicolaferraro/integration-operator/pkg/apis/camel/v1alpha1"
)

var (
	// All integration runtimes in strict priority order
	runtimes = []api.IntegrationRuntime {
		&classicIntegrationRuntime{},
	}
)

func GetIntegrationRuntimes() []api.IntegrationRuntime {
	return runtimes
}

func GetIntegrationRuntime(name v1alpha1.IntegrationRuntimeName) api.IntegrationRuntime {
	for _, r := range runtimes {
		if r.Name() == name {
			return r
		}
	}
	return nil
}

func GetIntegrationRuntimeFor(integration *v1alpha1.IntegrationSpec) api.IntegrationRuntime {
	for _, r := range runtimes {
		if r.Supports(integration) {
			return r
		}
	}
	return nil
}
