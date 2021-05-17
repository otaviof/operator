// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ShipwrightBuildSpec defines the characteristics of Shipwright-Build instances.
type ShipwrightBuildSpec struct {
	// Namespace the target namespace oh which Shipwright-Build will be deployed.
	Namespace string `json:"namespace,omitempty"`
}

// ShipwrightBuildStatus defines the observed state of Shipwright-Build
type ShipwrightBuildStatus struct{}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:subresource:status

// ShipwrightBuild is the Schema for this operator API. It represents a instance of Shipwright Build
// Kubernetes Operator.
type ShipwrightBuild struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ShipwrightBuildSpec   `json:"spec,omitempty"`
	Status ShipwrightBuildStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ShipwrightBuildList contains a list of ShipwrightBuild
type ShipwrightBuildList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ShipwrightBuild `json:"items"`
}

// init registers the current Schema on the Scheme Builder during initialization.
func init() {
	SchemeBuilder.Register(&ShipwrightBuild{}, &ShipwrightBuildList{})
}
