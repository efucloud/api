/*
Copyright The efucloud.com Authors.

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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Workspace",type="string",JSONPath=`.status.workspace`
// +kubebuilder:resource:scope=Namespaced

//Meter meter
type Meter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              MeterSpec `json:"spec" yaml:"spec" protobuf:"bytes,2,opt,name=spec"`
	// +optional
	Status MeterStatus `json:"status" yaml:"status" protobuf:"bytes,3,opt,name=status"`
}
type MeterSpec struct {
	// repository category: api/mock/web/app/jmeter/bdd
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum:=mock;api;web;app;jmeter
	Category string `json:"category" yaml:"category" protobuf:"bytes,2,opt,name=category"`
	// meter description
	// +kubebuilder:validation:Required
	Description string `json:"description" yaml:"description" protobuf:"bytes,3,opt,name=description"`
}
type MeterStatus struct {
	// workspace ref
	// +optional
	Workspace string `json:"workspace" json:"workspace" protobuf:"bytes,1,opt,name=workspace"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MeterList contains a list of Meter
type MeterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Meter `json:"items" protobuf:"bytes,2,rep,name=items"`
}
