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
// +kubebuilder:printcolumn:name="AuthSameWithLuffy",type=string,JSONPath=`.spec.authenticatedSameWithLuffy`
// +kubebuilder:printcolumn:name="Version",type=string,JSONPath=`.status.version.version`
// +kubebuilder:printcolumn:name="Category",type=string,JSONPath=`.spec.category`
// +kubebuilder:printcolumn:name="LastCheck",type="string",JSONPath=`.status.lastCheck`
// +kubebuilder:printcolumn:name="Provider",type=string,JSONPath=`.spec.provider`
// +kubebuilder:printcolumn:name="Region",type=string,JSONPath=`.spec.region`
// +kubebuilder:resource:scope=Cluster

// Cluster cluster
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
	// authenticated by oidc
	AuthenticatedSameWithLuffy bool `json:"authenticatedSameWithLuffy" yaml:"authenticatedSameWithLuffy" protobuf:"varint,4,opt,name=authenticatedSameWithLuffy"`
	// cluster region
	// +optional
	Region string `json:"region" yaml:"region" protobuf:"bytes,5,opt,name=region"`
	//  cluster description
	// +optional
	Description string `json:"description" yaml:"description" protobuf:"bytes,6,opt,name=description"`
	// cluster master url
	// +kubebuilder:validation:Required
	Master string `json:"master" yaml:"master" protobuf:"bytes,7,opt,name=master"`
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
	// user can create namespace
	// +kubebuilder:default:=false
	// +optional
	UserCanCreateNamespace bool `json:"userCanCreateNamespace" yaml:"userCanCreateNamespace" protobuf:"varint,12,opt,name=userCanCreateNamespace"`
	// agent token
	// +optional
	AgentToken string `json:"agentToken" yaml:"agentToken" protobuf:"bytes,13,opt,name=agentToken"`
	// if not empty, luffy will auto deploy efucloud agent application
	// +optional
	EfuCloudAppImagePrefix string `json:"efuCloudAppImagePrefix" yaml:"efuCloudAppImagePrefix" protobuf:"bytes,14,opt,name=efuCloudAppImagePrefix"`
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
	// encrypted agent token
	// +optional
	EncryptedAgentToken []byte `json:"encryptedAgentToken" yaml:"encryptedAgentToken" protobuf:"bytes,18,opt,name=encryptedAgentToken"`
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

// Workspace only exist on manager cluster
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              WorkspaceSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            WorkspaceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}
type WorkspaceSpec struct {
	// workspace code
	// +kubebuilder:validation:Required
	Code string `json:"code" yaml:"code" protobuf:"bytes,1,opt,name=code"`
	// eauth group's account will auto add in workspace
	// +optional
	EAuthGroups []string `json:"eAuthGroups" yaml:"eAuthGroups" protobuf:"bytes,4,rep,name=eAuthGroups"`
	// eauth group's account with prefix will auto add in workspace
	// +optional
	EAuthGroupPrefix string `json:"eAuthGroupPrefix" yaml:"eAuthGroupPrefix" protobuf:"bytes,5,opt,name=eAuthGroupPrefix"`
	// workspace description
	// +kubebuilder:validation:Required
	Description string `json:"description" yaml:"description" protobuf:"bytes,2,opt,name=description"`
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
// +kubebuilder:printcolumn:name="Workspace",type=string,JSONPath=`.spec.workspaceRef`
// +kubebuilder:printcolumn:name="Code",type=string,JSONPath=`.spec.code`
// +kubebuilder:resource:scope=Cluster

// ClusterWorkspace will save on member cluster
type ClusterWorkspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              ClusterWorkspaceSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            ClusterWorkspaceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}
type ClusterWorkspaceSpec struct {
	// workspace name
	// +kubebuilder:validation:Required
	WorkspaceRef string `json:"workspaceRef" yaml:"workspaceRef" protobuf:"bytes,3,opt,name=workspaceRef"`
	// workspace code
	// +optional
	Code string `json:"code" yaml:"code" protobuf:"bytes,1,opt,name=code"`
	// workspace description
	// +optional
	Description string `json:"description" yaml:"description" protobuf:"bytes,2,opt,name=description"`
}

// ClusterWorkspaceStatus defines the observed state of Workspace
type ClusterWorkspaceStatus struct {
	// all namespaces
	// +optional
	Namespaces []string `json:"namespaces" yaml:"namespaces" protobuf:"bytes,2,rep,name=namespaces"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterWorkspaceList contains a list of Workspace
type ClusterWorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []ClusterWorkspace `json:"items" protobuf:"bytes,2,rep,name=items"`
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

// DeployTemplateList contains a list of Config
type DeployTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []DeployTemplate `json:"items" protobuf:"bytes,2,rep,name=items"`
}
