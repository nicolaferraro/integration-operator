package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type IntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Integration `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Integration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              IntegrationSpec   `json:"spec"`
	Status            IntegrationStatus `json:"status,omitempty"`
}

type IntegrationSpec struct {
	Replicas          *int32 `json:"replicas,omitempty"`
	Flows			  []FlowSpec `json:"flows,omitempty"`
}

type FlowSpec struct {
	Id				string `json:"id,omitempty"`
	Name			string `json:"name,omitempty"`
	Steps			[]StepSpec `json:"steps,omitempty"`
}

type StepSpec struct {
	Type			string  `json:"type,omitempty"`
	Uri				string	`json:"uri,omitempty"`
}

type IntegrationStatus struct {
	Phase	IntegrationPhase `json:"phase,omitempty"`
}

type IntegrationPhase string

const (
	IntegrationPhaseMissing	IntegrationPhase = ""
	IntegrationPhaseRunning	IntegrationPhase = "Running"
)

type StrategySpec struct {
	IntegrationRuntimeName `json:"runtime,omitempty"`
}

type IntegrationRuntimeName string

const (
	IntegrationRuntimeClassic IntegrationRuntimeName = "Classic"
)
