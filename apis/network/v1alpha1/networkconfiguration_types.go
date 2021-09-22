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

// NetworkConfigurationSpec defines the desired state of NetworkConfiguration
type NetworkConfigurationSpec struct {
	Description   string      `json:"description,omitempty"`
	Configuration interface{} `json:"configuration,omitempty"`
}

// NetworkConfigurationStatus defines the observed state of NetworkConfiguration
type NetworkConfigurationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// NetworkConfiguration is the Schema for the networkconfigurations API
type NetworkConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkConfigurationSpec   `json:"spec,omitempty"`
	Status NetworkConfigurationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NetworkConfigurationList contains a list of NetworkConfiguration
type NetworkConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkConfiguration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkConfiguration{}, &NetworkConfigurationList{})
}
