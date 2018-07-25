package runtime

import "github.com/nicolaferraro/integration-operator/pkg/apis/camel/v1alpha1"

type classicIntegrationRuntime struct {

}

func (*classicIntegrationRuntime) Name() v1alpha1.IntegrationRuntimeName {
	return v1alpha1.IntegrationRuntimeClassic
}

func (*classicIntegrationRuntime) Image() string {
	return "nferraro/camel-classic-runtime-spring-boot:latest"
}

func (*classicIntegrationRuntime) Supports(integration *v1alpha1.IntegrationSpec) bool {
	// TODO integrate with Camel catalog
	return true
}