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
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:printcolumn:name="Description",type=string,JSONPath=`.description`
// +genclient:nonNamespaced
// +kubebuilder:resource:scope=Cluster

// WorkspaceRole workspace role
type WorkspaceRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              WorkspaceRoleSpec `json:"spec" yaml:"spec" protobuf:"bytes,2,opt,name=spec"`
}
type WorkspaceRoleSpec struct {
	// +optional
	Rules []rbacv1.PolicyRule `json:"rules" yaml:"rules" protobuf:"bytes,1,rep,name=rules"`
	//ref cluster roles, it must have label: efucloud.com/custom`
	// +optional
	ClusterRoleRefs []string `json:"clusterRoleRefs" yaml:"clusterRoleRefs" protobuf:"bytes,2,rep,name=clusterRoleRefs"`
	// Description  about workspace role
	// +kubebuilder:validation:Required
	Description string `json:"description" yaml:"description" protobuf:"bytes,3,opt,name=description"`
	// only ref pod's namespace role,  it must have label: efucloud.com/custom
	// +optional
	RoleRefs []string `json:"roleRefs" yaml:"roleRefs" protobuf:"bytes,4,rep,name=roleRefs"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkspaceRoleList contains a list of Workspace
type WorkspaceRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []WorkspaceRole `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Username",type=string,JSONPath=`.spec.username`
// +kubebuilder:printcolumn:name="Email",type=string,JSONPath=`.spec.email`
// +kubebuilder:printcolumn:name="Enable",type=boolean,JSONPath=`.status.enable`
// +kubebuilder:printcolumn:name="LastLoginTime",type="string",JSONPath=`.status.lastLoginTime`
// +kubebuilder:printcolumn:name="LastRemoteIP",type=string,JSONPath=`.status.lastRemoteIP`
// +kubebuilder:resource:scope=Cluster

type KubeUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              KubeUserSpec   `json:"spec,omitempty" yaml:"spec,omitempty,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            KubeUserStatus `json:"status,omitempty" yaml:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}
type KubeUserSpec struct {
	// username
	// +kubebuilder:validation:Required
	Username string `json:"username" yaml:"username" protobuf:"bytes,1,opt,name=username"`
	// user email
	// +optional
	Email string `json:"email" yaml:"email" protobuf:"bytes,2,opt,name=email"`
	// default language
	// +kubebuilder:validation:Enum=zh;en
	// +optional
	Language string `json:"language" yaml:"language" protobuf:"bytes,3,opt,name=language"`
	// user's mobile phone
	// +optional
	Phone string `json:"phone" yaml:"phone" protobuf:"bytes,4,opt,name=phone"`
	// 是否有效
	Enable uint     `json:"enable" yaml:"enable" protobuf:"varint,5,opt,name=enable"`
	Groups []string `json:"groups" yaml:"groups" protobuf:"bytes,6,rep,name=groups"`
}
type KubeUserStatus struct {
	// last login time
	// +optional
	LastLoginTime *metav1.Time `json:"lastLoginTime" yaml:"lastLoginTime" protobuf:"bytes,1,opt,name=lastLoginTime"`
	// last login remote ip
	// +optional
	LastRemoteIP string `json:"lastRemoteIp" yaml:"lastRemoteIp" protobuf:"bytes,2,opt,name=lastRemoteIp"`
	// if available is false, user will not login system
	// +kubebuilder:default:=true
	Enable bool `json:"enable" yaml:"enable" protobuf:"varint,3,opt,name=enable"`
	// etcd only save password with encrypt
	// +optional
	PasswordEncrypt string `json:"passwordEncrypt" yaml:"passwordEncrypt" protobuf:"bytes,4,opt,name=passwordEncrypt"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeUserList contains a list of KubeUser
type KubeUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []KubeUser `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +genclient:nonNamespaced
// +kubebuilder:printcolumn:name="KubeUserRef",type=string,JSONPath=`.spec.kubeUserRef`
// +kubebuilder:printcolumn:name="ClusterRef",type=string,JSONPath=`.spec.clusterRef`
// +kubebuilder:printcolumn:name="Enable",type=boolean,JSONPath=`.status.enable`
// +kubebuilder:printcolumn:name="Available",type=boolean,JSONPath=`.status.available`
// +kubebuilder:printcolumn:name="ExpiredTime",type=string,JSONPath=`.spec.expiredTime`
// +kubebuilder:printcolumn:name="LastCheck",type="string",JSONPath=`.status.lastCheck`
// +kubebuilder:resource:scope=Cluster

// UserKubeConfig is the Schema for the usermanages API
type UserKubeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   UserKubeConfigSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status UserKubeConfigStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// RequestConditionType is the type of CertificateSigningRequestCondition
type RequestConditionType string

// Well-known condition types for certificate requests.
const (
	// CertificateApproved Approved indicates the request was approved and should be issued by the signer.
	CertificateApproved RequestConditionType = "Approved"
	// CertificateDenied Denied indicates the request was denied and should not be issued by the signer.
	CertificateDenied RequestConditionType = "Denied"
	// CertificateFailed Failed indicates the signer failed to issue the certificate.
	CertificateFailed RequestConditionType = "Failed"
	// CertificateWaited waiting to issue the certificate.
	CertificateWaited RequestConditionType = "Waiting"
	// CertificateCreated created in cluster
	CertificateCreated RequestConditionType = "Created"
	//CertificateReceived kubeconfig get csr
	CertificateReceived RequestConditionType = "Received"
)

// UserKubeConfigSpec defines the desired state of UserKubeConfig
type UserKubeConfigSpec struct {
	// ref kubeuser
	// +kubebuilder:validation:Required
	KubeUserRef string `json:"kubeUserRef" yaml:"kubeUserRef" protobuf:"bytes,1,opt,name=kubeUserRef"`
	// cluster resource name
	// +kubebuilder:validation:Required
	ClusterRef string `json:"clusterRef" yaml:"clusterRef" protobuf:"bytes,2,opt,name=clusterRef"`
	// expire time
	// +optional
	ExpiredTime *metav1.Time `json:"expiredTime" yaml:"expiredTime" protobuf:"bytes,3,opt,name=expiredTime"`
	// user ClientCertificateData if content is raw data will auto base64 encode
	// is csr.Status.Certificate
	ClientCertificateData string `json:"clientCertificateData" yaml:"clientCertificateData" protobuf:"bytes,5,opt,name=clientCertificateData"`
	// user ClientKeyData if content is raw data will auto base64 encode
	// csr private key
	ClientKeyData string `json:"clientKeyData" yaml:"clientKeyData" protobuf:"bytes,6,opt,name=clientKeyData"`
}

// UserKubeConfigStatus defines the observed state of UserKubeConfig
type UserKubeConfigStatus struct {
	// if false app will not use this kubeconfig although available is true
	// +optional
	// +kubebuilder:default:=true
	Enable bool `json:"enable" yaml:"enable" protobuf:"varint,1,opt,name=enable"`
	// if true, app can use kubeconfig connect with cluster
	// +optional
	// +kubebuilder:default:=true
	Available bool `json:"available" yaml:"available" protobuf:"varint,2,opt,name=available"`
	// Only one condition of a given type is allowed.
	Type RequestConditionType `json:"type" yaml:"type" protobuf:"bytes,3,opt,name=type,casttype=RequestConditionType"`
	//ref  cluster's CertificateSigningRequest
	// +optional
	CSRRef string `json:"csrRef" yaml:"csrRef" protobuf:"bytes,5,opt,name=csrRef"`
	// +optional
	Hash string `json:"hash" yaml:"hash" protobuf:"bytes,6,opt,name=hash"`
	// user ClientCertificateData if content is raw data will auto base64 encode
	// +optional
	EncryptedClientCertificateData []byte `json:"encryptedClientCertificateData" yaml:"encryptedClientCertificateData" protobuf:"bytes,7,opt,name=encryptedClientCertificateData"`
	// user ClientKeyData if content is raw data will auto base64 encode
	// +optional
	EncryptedClientKeyData []byte `json:"encryptedClientKeyData" yaml:"encryptedClientKeyData" protobuf:"bytes,8,opt,name=encryptedClientKeyData"`
	// +optional
	LastCheck *metav1.Time `json:"lastCheck" yaml:"lastCheck" protobuf:"bytes,9,opt,name=lastCheck"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserKubeConfigList contains a list of UserKubeConfig
type UserKubeConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []UserKubeConfig `json:"items" protobuf:"bytes,2,rep,name=items"`
}
