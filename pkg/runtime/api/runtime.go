package api

import "github.com/nicolaferraro/integration-operator/pkg/apis/camel/v1alpha1"

type IntegrationRuntime interface {
	Name()	v1alpha1.IntegrationRuntimeName
	Image()	string
	Supports(integration *v1alpha1.IntegrationSpec) bool
}