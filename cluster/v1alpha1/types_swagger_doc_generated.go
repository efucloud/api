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
var map_Cluster = map[string]string{
	"": "Cluster is the schema for the clusters API",
}

func (Cluster) SwaggerDoc() map[string]string {
	return map_Cluster
}

var map_ClusterConfig = map[string]string{
	"": "ClusterConfig cluster",
}

func (ClusterConfig) SwaggerDoc() map[string]string {
	return map_ClusterConfig
}

var map_ClusterConfigList = map[string]string{
	"": "ClusterConfigList contains a list of Config",
}

func (ClusterConfigList) SwaggerDoc() map[string]string {
	return map_ClusterConfigList
}

var map_ClusterProfile = map[string]string{
	"": "ClusterProfile cluster",
}

func (ClusterProfile) SwaggerDoc() map[string]string {
	return map_ClusterProfile
}

var map_ClusterSpec = map[string]string{
	"code":        "cluster code",
	"provider":    "Provider of the cluster: Openshift, Kubernetes, or other cloud providers",
	"category":    "cluster category, such as: Strict???NonStrict???Dev???Test???Pro",
	"region":      "cluster region",
	"description": "\n cluster description",
	"master":      "cluster master url",
	"token":       "cluster's user or serviceaccount token",
	"certData":    "client certificate data",
	"keyData":     "client key",
	"caData":      "cluster Certificate Authority Data",
}

func (ClusterSpec) SwaggerDoc() map[string]string {
	return map_ClusterSpec
}

var map_ClusterStatus = map[string]string{
	"isManager":         "manager cluster, will auto judge",
	"version":           "kubernetes version",
	"encryptedToken":    "cluster sa token",
	"encryptedCertData": "encrypted client certificate data",
	"encryptedKeyData":  "encrypted client key",
	"encryptedCaData":   "encrypted cluster Certificate Authority Data",
}

func (ClusterStatus) SwaggerDoc() map[string]string {
	return map_ClusterStatus
}

var map_DeployTemplate = map[string]string{
	"": "DeployTemplate  specify workspace resource and namespace number,",
}

func (DeployTemplate) SwaggerDoc() map[string]string {
	return map_DeployTemplate
}

var map_DeployTemplateList = map[string]string{
	"": "DeployTemplateList contains a list of Config",
}

func (DeployTemplateList) SwaggerDoc() map[string]string {
	return map_DeployTemplateList
}

var map_Workspace = map[string]string{
	"": "Workspace is the Schema for the namespace groups API, use workspace name as namespace label value",
}

func (Workspace) SwaggerDoc() map[string]string {
	return map_Workspace
}

var map_WorkspaceList = map[string]string{
	"": "WorkspaceList contains a list of Workspace",
}

func (WorkspaceList) SwaggerDoc() map[string]string {
	return map_WorkspaceList
}

var map_WorkspaceResourceQuota = map[string]string{
	"": "WorkspaceResourceQuota  specify workspace resource and namespace number,",
}

func (WorkspaceResourceQuota) SwaggerDoc() map[string]string {
	return map_WorkspaceResourceQuota
}

var map_WorkspaceResourceQuotaList = map[string]string{
	"": "WorkspaceResourceQuotaList contains a list of Config",
}

func (WorkspaceResourceQuotaList) SwaggerDoc() map[string]string {
	return map_WorkspaceResourceQuotaList
}

var map_WorkspaceResourceQuotaSpec = map[string]string{
	"workspaceRef":    "WorkspaceRef is equal WorkspaceResourceQuota name",
	"clusterSelector": "namespace can be created in cluster",
}

func (WorkspaceResourceQuotaSpec) SwaggerDoc() map[string]string {
	return map_WorkspaceResourceQuotaSpec
}

var map_WorkspaceSpec = map[string]string{
	"code":        "workspace code",
	"description": "workspace description",
	"users":       "workspace users",
}

func (WorkspaceSpec) SwaggerDoc() map[string]string {
	return map_WorkspaceSpec
}

var map_WorkspaceStatus = map[string]string{
	"":                  "WorkspaceStatus defines the observed state of Workspace",
	"clusterNamespaces": "cluster's namespaces",
	"namespaces":        "all namespaces",
}

func (WorkspaceStatus) SwaggerDoc() map[string]string {
	return map_WorkspaceStatus
}

var map_WorkspaceUser = map[string]string{
	"userRef":        "ref kubeuser",
	"workspaceRoles": "ref workspace role",
}

func (WorkspaceUser) SwaggerDoc() map[string]string {
	return map_WorkspaceUser
}

// AUTO-GENERATED FUNCTIONS END HERE
