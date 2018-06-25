package util

import (
	"github.com/nicolaferraro/integration-operator/pkg/apis/camel/v1alpha1"
	"gopkg.in/yaml.v2"
)

type serializationContainer struct {
	Routes *[]v1alpha1.RouteSpec `json:"routes"`
}

func Serialize(routes []v1alpha1.RouteSpec) (string, error) {

	container := serializationContainer{&routes}

	bin,err := yaml.Marshal(container)
	if err != nil {
		return "", err
	}
	return string(bin), nil
}