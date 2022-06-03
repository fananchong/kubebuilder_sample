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

// Example1Spec defines the desired state of Example1
type Example1Spec struct {
	//+kubebuilder:validation:Required
	//+kubebuilder:validation:Type=string
	Custom1 string `json:"custom1,omitempty"`

	//+kubebuilder:validation:Required
	//+kubebuilder:validation:Type=integer
	//+kubebuilder:validation:Minimum=0
	Custom2 int `json:"custom2,omitempty"`
}

// Example1Status defines the observed state of Example1
type Example1Status struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Example1 is the Schema for the example1s API
type Example1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Example1Spec   `json:"spec,omitempty"`
	Status Example1Status `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// Example1List contains a list of Example1
type Example1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Example1 `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Example1{}, &Example1List{})
}
