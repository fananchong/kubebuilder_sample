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

// Example3Spec defines the desired state of Example3
type Example3Spec struct {
	// +kubebuilder:validation:Minimum=1
	InstanceNum *int64 `json:"instanceNum,omitempty"`
}

// Example3Status defines the observed state of Example3
type Example3Status struct {
	// +optional
	RealInstanceNum *int32 `json:"realInstanceNum,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Example3 is the Schema for the example3s API
type Example3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Example3Spec   `json:"spec,omitempty"`
	Status Example3Status `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// Example3List contains a list of Example3
type Example3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Example3 `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Example3{}, &Example3List{})
}
