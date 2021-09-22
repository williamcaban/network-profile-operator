/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NetworkProfileClassSpec defines the desired state of NetworkProfileClass
type NetworkProfileClassSpec struct {
	DscpName  string `json:"dscpName,omitempty"`
	MaxLoss   int    `json:"maxLoss,omitempty"`
	MaxRTT    int    `json:"maxRTT,omitempty"`
	MaxJitter int    `json:"maxJitter,omitempty"`
}

// NetworkProfileClassStatus defines the observed state of NetworkProfileClass
type NetworkProfileClassStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// NetworkProfileClass is the Schema for the networkprofileclasses API
type NetworkProfileClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkProfileClassSpec   `json:"spec,omitempty"`
	Status NetworkProfileClassStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NetworkProfileClassList contains a list of NetworkProfileClass
type NetworkProfileClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkProfileClass `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkProfileClass{}, &NetworkProfileClassList{})
}
