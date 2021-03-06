/*
Copyright 2022.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Example2Spec defines the desired state of Example2
type Example2Spec struct {
}

// Example2Status defines the observed state of Example2
type Example2Status struct {
	// +optional
	CustomStatus1 string `json:"customStatus1,omitempty"`
	// +optional
	CustomStatus2 *int32 `json:"customStatus2,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Example2 is the Schema for the example2s API
type Example2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Example2Spec   `json:"spec,omitempty"`
	Status Example2Status `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// Example2List contains a list of Example2
type Example2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Example2 `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Example2{}, &Example2List{})
}
