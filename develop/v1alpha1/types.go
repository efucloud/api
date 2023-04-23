/*
Copyright The efucloud Authors.

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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CoderTemplateSpec defines the desired state of CoderTemplate
type CoderTemplateSpec struct {
	//template name
	Name string `json:"name" yaml:"name" protobuf:"bytes,1,opt,name=name"` //
	//template description
	Description string `json:"description" yaml:"description" protobuf:"bytes,2,opt,name=description"` //
	//template' image author
	Author string `json:"author" yaml:"author" protobuf:"bytes,3,opt,name=author"` //
	// template language
	// +kubebuilder:validation:Enum=golang;java;python;ruby;rust
	Language string `json:"language" yaml:"language" protobuf:"bytes,4,opt,name=language"` //
	// template containers
	Containers []v1.Container `json:"containers" yaml:"containers" protobuf:"bytes,5,rep,name=containers"` //
}

// CoderTemplateStatus defines the observed state of CoderTemplate
type CoderTemplateStatus struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CoderTemplate is the Schema for the CoderTemplates API
type CoderTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   CoderTemplateSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status CoderTemplateStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CoderTemplateList contains a list of CoderTemplate
type CoderTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []CoderTemplate `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// CoderVirtSpec defines the desired state of CoderVirt
type CoderVirtSpec struct {
	Running *bool `json:"running,omitempty" yaml:"running,omitempty" ¬optional:"true" protobuf:"varint,1,opt,name=running"`
	// +kubebuilder:validation:Enum=golang;java;python;ruby;rust
	Language    string `json:"language" yaml:"language" protobuf:"bytes,2,opt,name=language"`          //
	TemplateRef string `json:"templateRef" yaml:"templateRef" protobuf:"bytes,3,opt,name=templateRef"` //
}

// CoderVirtStatus defines the observed state of CoderVirt
type CoderVirtStatus struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CoderVirt is the Schema for the CoderVirts API
type CoderVirt struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   CoderVirtSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status CoderVirtStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CoderVirtList contains a list of CoderVirt
type CoderVirtList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []CoderVirt `json:"items" protobuf:"bytes,2,rep,name=items"`
}
