package runtime

import (
	"github.com/nicolaferraro/integration-operator/pkg/apis/camel/v1alpha1"
	"strings"
)

var (
	goSupportedComponents = map[string]bool{"http": true, "timer": true, "log": true}
)

type goIntegrationRuntime struct {

}

func (*goIntegrationRuntime) Name() v1alpha1.IntegrationRuntimeName {
	return v1alpha1.IntegrationRuntimeGo
}

func (*goIntegrationRuntime) Image() string {
	return "nferraro/camel-go-runtime:latest"
}

func (*goIntegrationRuntime) Supports(integration *v1alpha1.IntegrationSpec) bool {
	for _, flow := range integration.Flows {
		for _, step := range flow.Steps {
			if step.Type == "endpoint" {
				if parts := strings.Split(step.Uri, ":"); len(parts) > 1 {
					if s, ok := goSupportedComponents[parts[0]]; !ok || !s {
						return false
					}
				}
			}
		}
	}
	return true
}