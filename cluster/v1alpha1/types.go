/*
Copyright 2022 The efucloud.com Authors.

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
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	WorkspaceNamespaceSelectorKey = "com.efucloud.workspace.namespace"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Code",type=string,JSONPath=`.spec.code`
// +kubebuilder:printcolumn:name="Healthy",type=string,JSONPath=`.status.healthy`
// +kubebuilder:printcolumn:name="Manager",type=string,JSONPath=`.status.isManager`
// +kubebuilder:printcolumn:name="Platform",type=string,JSONPath=`.status.version.platform`
// +kubebuilder:printcolumn:name="Version",type=string,JSONPath=`.status.version.version`
// +kubebuilder:printcolumn:name="Category",type=string,JSONPath=`.spec.category`
// +kubebuilder:printcolumn:name="Namespaces",type=string,JSONPath=`.status.namespaceNumber`
// +kubebuilder:printcolumn:name="LastCheck",type="string",JSONPath=`.status.lastCheck`
// +kubebuilder:printcolumn:name="Provider",type=string,JSONPath=`.spec.provider`
// +kubebuilder:printcolumn:name="Region",type=string,JSONPath=`.spec.region`
// +kubebuilder:resource:scope=Cluster

// Cluster is the schema for the clusters API
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   ClusterSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status ClusterStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}
type ClusterSpec struct {
	// cluster code
	// +kubebuilder:validation:Required
	Code string `json:"code" yaml:"code" protobuf:"bytes,1,opt,name=code"`
	// Provider of the cluster: Openshift, Kubernetes, or other cloud providers
	// +kubebuilder:validation:Required
	Provider string `json:"provider" yaml:"provider" protobuf:"bytes,2,opt,name=provider"`
	// cluster category, such as: Strict、NonStrict、Dev、Test、Pro
	// +kubebuilder:validation:Required
	Category string `json:"category" yaml:"category" protobuf:"bytes,3,opt,name=category"`
	// cluster region
	// +optional
	Region string `json:"region" yaml:"region" protobuf:"bytes,4,opt,name=region"`
	//  cluster description
	// +optional
	Description string `json:"description" yaml:"description" protobuf:"bytes,5,opt,name=description"`
	// cluster master url
	// +kubebuilder:validation:Required
	Master string `json:"master" yaml:"master" protobuf:"bytes,6,opt,name=master"`
	// cluster's user or serviceaccount token
	// +optional
	Token string `json:"token" yaml:"token" protobuf:"bytes,8,opt,name=token"`
	// client certificate data
	// +optional
	CertData string `json:"certData" yaml:"certData" protobuf:"bytes,9,opt,name=certData"`
	// client key
	// +optional
	KeyData string `json:"keyData" yaml:"keyData" protobuf:"bytes,10,opt,name=keyData"`
	// cluster Certificate Authority Data
	// +kubebuilder:validation:Required
	CaData string `json:"caData" yaml:"caData" protobuf:"bytes,11,opt,name=caData"`
}
type ClusterStatus struct {
	// manager cluster, will auto judge
	// +optional
	IsManager bool `json:"isManager" yaml:"isManager" protobuf:"varint,10,opt,name=isManager"`
	// +optional
	Healthy bool `json:"healthy" yaml:"healthy" protobuf:"varint,11,opt,name=healthy"`
	// +optional
	LastCheck *metav1.Time `json:"lastCheck" yaml:"lastCheck" protobuf:"bytes,12,opt,name=lastCheck"`
	// kubernetes version
	// +optional
	Version Version `json:"version,omitempty" yaml:"version" protobuf:"bytes,13,opt,name=version"`
	// cluster sa token
	// +optional
	EncryptedToken []byte `json:"encryptedToken" yaml:"encryptedToken" protobuf:"bytes,14,opt,name=encryptedToken"`
	// +optional
	Hash string `json:"hash" yaml:"hash" protobuf:"bytes,8,opt,name=hash"`
	// +optional
	Namespaces []string `json:"namespaces" yaml:"namespaces" protobuf:"bytes,9,rep,name=namespaces"`
	// encrypted client certificate data
	// +optional
	EncryptedCertData []byte `json:"encryptedCertData" yaml:"encryptedCertData" protobuf:"bytes,15,opt,name=encryptedCertData"`
	// encrypted client key
	// +optional
	EncryptedKeyData []byte `json:"encryptedKeyData" yaml:"encryptedKeyData" protobuf:"bytes,16,opt,name=encryptedKeyData"`
	// encrypted cluster Certificate Authority Data
	// +optional
	EncryptedCaData []byte `json:"encryptedCaData" yaml:"encryptedCaData" protobuf:"bytes,17,opt,name=encryptedCaData"`
}

type Version struct {
	// +optional
	Major string `json:"major" protobuf:"bytes,1,opt,name=major"`
	// +optional
	Minor string `json:"minor" protobuf:"bytes,2,opt,name=minor"`
	// +optional
	GitVersion string `json:"gitVersion" protobuf:"bytes,3,opt,name=gitVersion"`
	// +optional
	GitCommit string `json:"gitCommit" protobuf:"bytes,4,opt,name=gitCommit"`
	// +optional
	GitTreeState string `json:"gitTreeState" protobuf:"bytes,5,opt,name=gitTreeState"`
	// +optional
	BuildDate string `json:"buildDate" protobuf:"bytes,6,opt,name=buildDate"`
	// +optional
	GoVersion string `json:"goVersion" protobuf:"bytes,7,opt,name=goVersion"`
	// +optional
	Compiler string `json:"compiler" protobuf:"bytes,8,opt,name=compiler"`
	// +optional
	Platform string `json:"platform" protobuf:"bytes,9,opt,name=platform"`
	// +optional
	Version string `json:"version" protobuf:"bytes,10,opt,name=version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Cluster `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Cluster",type=string,JSONPath=`.spec.clusterRef`
// +kubebuilder:resource:scope=Cluster

// ClusterProfile cluster
type ClusterProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              ClusterProfileSpec   `json:"spec" yaml:"spec" protobuf:"bytes,2,opt,name=spec"`
	Status            ClusterProfileStatus `json:"status" yaml:"status" protobuf:"bytes,3,opt,name=status"`
}
type ClusterProfileSpec struct {
	// +kubebuilder:validation:Required
	ClusterRef string `json:"clusterRef" yaml:"clusterRef" protobuf:"bytes,1,opt,name=clusterRef"`
	// +kubebuilder:validation:Required
	PrometheusConfigMapRef string `json:"prometheusConfigMapRef" yaml:"prometheusConfigMapRef" protobuf:"bytes,2,opt,name=prometheusConfigMapRef"`
	// +kubebuilder:validation:Required
	GrafanaConfigMapRef string `json:"grafanaConfigMapRef" yaml:"grafanaConfigMapRef" protobuf:"bytes,3,opt,name=grafanaConfigMapRef"`
	// +kubebuilder:validation:Required
	HarborConfigMapRef string `json:"harborConfigMapRef" yaml:"harborConfigMapRef" protobuf:"bytes,4,opt,name=harborConfigMapRef"`
}

type ClusterProfileStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []ClusterProfile `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Code",type=string,JSONPath=`.spec.code`
// +kubebuilder:printcolumn:name="Cascade Delete",type=string,JSONPath=`.spec.cascadeDelete`
// +kubebuilder:resource:scope=Cluster

// Workspace is the Schema for the namespace groups API, use workspace name as namespace label value
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   WorkspaceSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status WorkspaceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}
type WorkspaceSpec struct {
	// workspace code
	// +kubebuilder:validation:Required
	Code string `json:"code" yaml:"code" protobuf:"bytes,1,opt,name=code"`
	// workspace description
	// +kubebuilder:validation:Required
	Description string `json:"description" yaml:"description" protobuf:"bytes,2,opt,name=description"`
	// workspace users
	// +optional
	Users []WorkspaceUser `json:"users" yaml:"users" protobuf:"bytes,3,rep,name=users"`
}
type WorkspaceUser struct {
	// ref kubeuser
	// +kubebuilder:validation:Required
	UserRef string `json:"userRef" yaml:"userRef" protobuf:"bytes,1,opt,name=userRef"`
	// ref workspace role
	// +optional
	WorkspaceRoles []string `json:"workspaceRoles" yaml:"workspaceRoles" protobuf:"bytes,2,rep,name=workspaceRoles"`
}

// WorkspaceStatus defines the observed state of Workspace
type WorkspaceStatus struct {
	// cluster's namespaces
	// +optional
	ClusterNamespaces []ClusterNamespace `json:"clusterNamespaces" yaml:"clusterNamespaces" protobuf:"bytes,1,rep,name=clusterNamespaces"`
	// all namespaces
	// +optional
	Namespaces []string `json:"namespaces" yaml:"namespaces" protobuf:"bytes,2,rep,name=namespaces"`
}
type ClusterNamespace struct {
	// +kubebuilder:validation:Required
	Cluster string `json:"cluster" yaml:"cluster" protobuf:"bytes,1,opt,name=cluster"`
	// +kubebuilder:validation:Required

	Namespaces []string `json:"namespaces" yaml:"namespaces" protobuf:"bytes,2,rep,name=namespaces"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkspaceList contains a list of Workspace
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Workspace `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Cluster",type=string,JSONPath=`.metadata.name`
// +kubebuilder:printcolumn:name="Description",type=string,JSONPath=`.spec.description`
// +kubebuilder:resource:scope=Cluster

// ClusterConfig cluster
type ClusterConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              ClusterConfigSpec   `json:"spec" yaml:"spec" protobuf:"bytes,2,opt,name=spec"`
	Status            ClusterConfigStatus `json:"status" yaml:"status" protobuf:"bytes,3,opt,name=status"`
}
type ClusterConfigSpec struct {
	Description string `json:"description" yaml:"description" protobuf:"bytes,2,opt,name=description"`
	Data        string `json:"data" yaml:"data" protobuf:"bytes,3,opt,name=data"`
}

type ClusterConfigStatus struct {
	Available   bool   `json:"available" yaml:"available" protobuf:"varint,1,opt,name=available"`
	EncryptData []byte `json:"encryptData" yaml:"encryptData" protobuf:"bytes,2,opt,name=encryptData"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterConfigList contains a list of Config
type ClusterConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []ClusterConfig `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:printcolumn:name="Workspace",type=string,JSONPath=`.spec.workspaceRef`
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// WorkspaceResourceQuota  specify workspace resource and namespace number,
type WorkspaceResourceQuota struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              WorkspaceResourceQuotaSpec   `json:"spec" yaml:"spec" protobuf:"bytes,2,opt,name=spec"`
	Status            WorkspaceResourceQuotaStatus `json:"status" yaml:"status" protobuf:"bytes,3,opt,name=status"`
}
type WorkspaceResourceQuotaSpec struct {
	// WorkspaceRef is equal WorkspaceResourceQuota name
	// +kubebuilder:validation:Required
	WorkspaceRef string `json:"workspaceRef" yaml:"workspaceRef" protobuf:"bytes,1,opt,name=workspaceRef"`
	// +optional
	Hard            v1.ResourceList `json:"hard" yaml:"hard" protobuf:"bytes,2,rep,name=hard,casttype=k8s.io/api/core/v1.ResourceList,castkey=k8s.io/api/core/v1.ResourceName"`
	NamespaceNumber int64           `json:"namespaceNumber" yaml:"namespaceNumber" protobuf:"varint,3,opt,name=namespaceNumber"`
	// namespace can be created in cluster
	// +kubebuilder:validation:Required
	ClusterSelector *metav1.LabelSelector `json:"clusterSelector" yaml:"clusterSelector" protobuf:"bytes,4,opt,name=clusterSelector"`
}

type WorkspaceResourceQuotaStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

//WorkspaceResourceQuotaList contains a list of Config
type WorkspaceResourceQuotaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []WorkspaceResourceQuota `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type Parameter struct {
	Name        string `json:"name" yaml:"name" protobuf:"bytes,1,opt,name=name"`
	DisplayName string `json:"displayName" yaml:"displayName" protobuf:"bytes,2,opt,name=displayName"`
	Description string `json:"description" yaml:"description" protobuf:"bytes,3,opt,name=description"`
	Value       string `json:"value" yaml:"value" protobuf:"bytes,4,opt,name=value"`
	// +optional
	Generate string `json:"generate" yaml:"generate" protobuf:"bytes,5,opt,name=generate"`
	// +optional
	From string `json:"from" yaml:"from" protobuf:"bytes,6,opt,name=from"`
	// +optional
	Required bool `json:"required" yaml:"required" protobuf:"varint,7,opt,name=required"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:printcolumn:name="Workspace",type=string,JSONPath=`.spec.workspaceRef`
// +kubebuilder:resource:scope=Cluster

// DeployTemplate  specify workspace resource and namespace number,
type DeployTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Description string                 `json:"description" yaml:"description" protobuf:"bytes,2,opt,name=description"`
	Objects     []runtime.RawExtension `json:"objects" yaml:"objects" protobuf:"bytes,3,rep,name=objects"`
	// +optional
	Parameters []Parameter `json:"parameters" yaml:"parameters" protobuf:"bytes,4,rep,name=parameters"`
	// +optional
	ObjectLabels map[string]string `json:"objectLabels" yaml:"objectLabels" protobuf:"bytes,5,rep,name=objectLabels"`
	// +kubebuilder:default:=5
	History int32 `json:"history" yaml:"history" protobuf:"varint,6,opt,name=history"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

//DeployTemplateList contains a list of Config
type DeployTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []DeployTemplate `json:"items" protobuf:"bytes,2,rep,name=items"`
}
