package util

import (
	"github.com/nicolaferraro/integration-operator/pkg/apis/camel/v1alpha1"
	"github.com/ghodss/yaml"
)

type flowContainer struct {
	Flows	[]v1alpha1.FlowSpec `json:"flows,omitempty"`
}

func Serialize(integration *v1alpha1.Integration) (string, error) {
	bin,err := yaml.Marshal(*integration)
	if err != nil {
		return "", err
	}
	return string(bin), nil
}

func MarshalFlows(integration *v1alpha1.Integration) ([]byte, error) {
	container := flowContainer{
		Flows: integration.Spec.Flows,
	}
	bin,err := yaml.Marshal(container)
	if err != nil {
		return nil, err
	}
	return bin, nil
}