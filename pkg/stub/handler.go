package stub

import (
	"context"

	"github.com/nicolaferraro/integration-operator/pkg/apis/camel/v1alpha1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	appsv1 "k8s.io/api/apps/v1"
	"github.com/nicolaferraro/integration-operator/pkg/util"
	"github.com/nicolaferraro/integration-operator/pkg/runtime/api"
	"github.com/nicolaferraro/integration-operator/pkg/runtime"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.Integration:

		r := runtime.GetIntegrationRuntimeFor(&o.Spec)
		cm := newConfigMap(o)
		deployment := newDeployment(o, r)

		err := sdk.Create(cm)
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("Failed to create config map : %v", err)
			return err
		} else if err != nil {
			err := sdk.Update(cm)
			if err != nil {
				logrus.Errorf("Failed to update config map : %v", err)
			}
		}

		err = sdk.Create(deployment)
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("Failed to create integration deployment : %v", err)
			return err
		} else if err != nil {
			err := sdk.Update(deployment)
			if err != nil {
				logrus.Errorf("Failed to update integration deployment : %v", err)
			}
		}
	}
	return nil
}

func newConfigMap(cr *v1alpha1.Integration) *v1.ConfigMap {
	integration, err := util.Serialize(cr.Spec)
	if err != nil {
		logrus.Error("Error while extracting integration", err)
	}

	return &v1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind: "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: cr.Name,
			Namespace: cr.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
					Group:   v1alpha1.SchemeGroupVersion.Group,
					Version: v1alpha1.SchemeGroupVersion.Version,
					Kind:    "Integration",
				}),
			},
			Labels: cr.Labels,
		},
		Data: map[string]string{
			"integration": integration,
		},
	}
}

func newDeployment(cr *v1alpha1.Integration, r api.IntegrationRuntime) *appsv1.Deployment {

	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name,
			Namespace: cr.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
					Group:   v1alpha1.SchemeGroupVersion.Group,
					Version: v1alpha1.SchemeGroupVersion.Version,
					Kind:    "Integration",
				}),
			},
			Labels: cr.Labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: cr.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"integration": cr.Name,
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"integration": cr.Name,
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    cr.Name,
							Image:   r.Image(),
							ImagePullPolicy: v1.PullIfNotPresent,
							VolumeMounts: []v1.VolumeMount {
								{
									Name: "integration",
									MountPath: "/etc/camel",
								},
							},
						},
					},
					Volumes: []v1.Volume {
						{
							Name: "integration",
							VolumeSource: v1.VolumeSource{
								ConfigMap: &v1.ConfigMapVolumeSource{
									LocalObjectReference: v1.LocalObjectReference{
										Name: cr.Name,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

