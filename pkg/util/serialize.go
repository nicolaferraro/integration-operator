package util

import (
	"github.com/nicolaferraro/integration-operator/pkg/apis/camel/v1alpha1"
	"gopkg.in/yaml.v2"
)

func Serialize(integration v1alpha1.IntegrationSpec) (string, error) {
	bin,err := yaml.Marshal(integration)
	if err != nil {
		return "", err
	}
	return string(bin), nil
}