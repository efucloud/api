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

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_KubeUserAPIKeyList = map[string]string{
	"": "KubeUserAPIKeyList contains a list of KubeUserAPIKey",
}

func (KubeUserAPIKeyList) SwaggerDoc() map[string]string {
	return map_KubeUserAPIKeyList
}

var map_KubeUserAPIKeySpec = map[string]string{
	"userRef": "relate user",
	"expired": "expired time, if nil will not expire",
}

func (KubeUserAPIKeySpec) SwaggerDoc() map[string]string {
	return map_KubeUserAPIKeySpec
}

var map_KubeUserList = map[string]string{
	"": "KubeUserList contains a list of KubeUser",
}

func (KubeUserList) SwaggerDoc() map[string]string {
	return map_KubeUserList
}

var map_KubeUserSpec = map[string]string{
	"username":    "username",
	"email":       "user email",
	"language":    "default language",
	"password":    "login password",
	"mobilePhone": "user's mobile phone",
}

func (KubeUserSpec) SwaggerDoc() map[string]string {
	return map_KubeUserSpec
}

var map_KubeUserStatus = map[string]string{
	"lastLoginTime":   "last login time",
	"lastRemoteIp":    "last login remote ip",
	"enable":          "if available is false, user will not login system",
	"passwordEncrypt": "etcd only save password with encrypt",
}

func (KubeUserStatus) SwaggerDoc() map[string]string {
	return map_KubeUserStatus
}

var map_UserKubeConfig = map[string]string{
	"": "UserKubeConfig is the Schema for the usermanages API",
}

func (UserKubeConfig) SwaggerDoc() map[string]string {
	return map_UserKubeConfig
}

var map_UserKubeConfigList = map[string]string{
	"": "UserKubeConfigList contains a list of UserKubeConfig",
}

func (UserKubeConfigList) SwaggerDoc() map[string]string {
	return map_UserKubeConfigList
}

var map_UserKubeConfigSpec = map[string]string{
	"":                      "UserKubeConfigSpec defines the desired state of UserKubeConfig",
	"kubeUserRef":           "ref kubeuser",
	"clusterRef":            "cluster resource name",
	"expiredTime":           "expire time",
	"clientCertificateData": "user ClientCertificateData if content is raw data will auto base64 encode is csr.Status.Certificate",
	"clientKeyData":         "user ClientKeyData if content is raw data will auto base64 encode csr private key",
}

func (UserKubeConfigSpec) SwaggerDoc() map[string]string {
	return map_UserKubeConfigSpec
}

var map_UserKubeConfigStatus = map[string]string{
	"":                               "UserKubeConfigStatus defines the observed state of UserKubeConfig",
	"enable":                         "if false app will not use this kubeconfig although available is true",
	"available":                      "if true, app can use kubeconfig connect with cluster",
	"type":                           "Only one condition of a given type is allowed.",
	"csrRef":                         "ref  cluster's CertificateSigningRequest",
	"encryptedClientCertificateData": "user ClientCertificateData if content is raw data will auto base64 encode",
	"encryptedClientKeyData":         "user ClientKeyData if content is raw data will auto base64 encode",
}

func (UserKubeConfigStatus) SwaggerDoc() map[string]string {
	return map_UserKubeConfigStatus
}

var map_WorkspaceRole = map[string]string{
	"": "WorkspaceRole workspace role",
}

func (WorkspaceRole) SwaggerDoc() map[string]string {
	return map_WorkspaceRole
}

var map_WorkspaceRoleList = map[string]string{
	"": "WorkspaceRoleList contains a list of Workspace",
}

func (WorkspaceRoleList) SwaggerDoc() map[string]string {
	return map_WorkspaceRoleList
}

var map_WorkspaceRoleSpec = map[string]string{
	"clusterRoleRefs": "ref cluster roles, it must have label: efucloud.com/custom`",
	"description":     "Description  about workspace role",
	"roleRefs":        "only ref pod's namespace role,  it must have label: efucloud.com/custom",
}

func (WorkspaceRoleSpec) SwaggerDoc() map[string]string {
	return map_WorkspaceRoleSpec
}

// AUTO-GENERATED FUNCTIONS END HERE
