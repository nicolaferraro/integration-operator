package main

import (
	"os"
	"io/ioutil"
	"github.com/nicolaferraro/integration-operator/pkg/util"
	"github.com/nicolaferraro/integration-operator/pkg/apis/camel/v1alpha1"
	"errors"
	"github.com/sirupsen/logrus"
	"path"
	"context"
	camelgo "github.com/lburgazzoli/camel-go/app"

	// load all camel component
	_ "github.com/lburgazzoli/camel-go/cmd"
)

const (
	CamelIntegrationFile = "/etc/camel/integration"
	CamelGoFlowsFile = "flow.yaml"
)

func main() {
	logrus.Info("Configuring flows for camel-go runtime")

	err := configure()
	if err != nil {
		panic(err)
	}

	app, err := camelgo.New("")
	if err != nil {
		panic(err)
	}

	logrus.Info("Starting camel-go application")
	app.Start()

	ctx := context.TODO()
	<- ctx.Done()
}

func configure() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	integrationYaml, err := ioutil.ReadFile(CamelIntegrationFile)
	if err != nil {
		return err
	}

	obj, err := util.LoadKubernetesResourceYaml(integrationYaml)
	if err != nil {
		return err
	}

	integration, ok := obj.(*v1alpha1.Integration)
	if !ok {
		return errors.New("no integration spec found in " + CamelIntegrationFile)
	}

	flows, err := util.MarshalFlows(integration)
	if err != nil {
		return err
	}

	flowsPath := path.Join(cwd, CamelGoFlowsFile)
	if err := ioutil.WriteFile(flowsPath, flows, 0755); err != nil {
		return err
	}

	return nil
}