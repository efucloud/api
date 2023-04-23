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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ProjectArchitectureSpec defines the desired state of ProjectArchitecture
type ProjectArchitectureSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Import ant: Run "make" to regenerate code after modifying this file

	// Project Description
	// +optional
	Description string `json:"description,omitempty" protobuf:"bytes,1,opt,name=description"`
	// Project's services relation, key is service name
	// +optional
	Services map[string]ProjectService `json:"services,omitempty" protobuf:"bytes,2,rep,name=services"`
	// project's metrics prometheus information
	// +optional
	PrometheusInformation map[string]PrometheusInfo `json:"prometheusInformation,omitempty" protobuf:"bytes,3,rep,name=prometheusInformation"`
	// metrics scrape interval
	// +optional
	ScrapeInterval string `json:"scrapeInterval,omitempty" protobuf:"bytes,4,opt,name=scrapeInterval"`
	// metrics scrape length, default one hour
	// +kubebuilder:default:=3600
	ScrapeLength int64 `json:"scrapeLength,omitempty" protobuf:"varint,5,opt,name=scrapeLength"`
}

// Project service
type ProjectService struct {
	// Service Name
	// +kubebuilder:validation:Required
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	// Service type
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=frontend;backend;database;middleware
	Category string `json:"category,omitempty" protobuf:"bytes,2,opt,name=category"`
	// frontend icon display
	// +optional
	Icon string `json:"icon,omitempty" protobuf:"bytes,3,opt,name=icon"`
	// annotations
	// +optional
	Annotations map[string]string `json:"annotations,omitempty" protobuf:"bytes,4,rep,name=annotations"`
	// service description
	// +kubebuilder:validation:Optional
	// +optional
	Description string `json:"description,omitempty" protobuf:"bytes,5,opt,name=description"`
	// service item,map key is service version
	// +kubebuilder:validation:Required
	ProjectServiceItems map[string]ProjectServiceItem `json:"projectServiceItems,omitempty" protobuf:"bytes,6,rep,name=projectServiceItems"`
}

type ProjectServiceItem struct {
	// Service Name
	// +kubebuilder:validation:Required
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	// service description
	// +kubebuilder:validation:Optional
	// +optional
	Description string `json:"description,omitempty" protobuf:"bytes,2,opt,name=description"`
	// DeployPath
	// if server deployed in same namespace with ProjectArchitecture instance, deployPath is kind eg: [Deployment]
	// if service deployed in cluster other namespace, name has prefix with ClusterNamespace., eg: ClusterNamespace.[default].[Deployment]
	// if server deployed in other cluster's namespace, name has External with ExternalClusterNamespace, eg: ExternalClusterNamespace.[clustercode].[default].[Deployment]
	// if server is deployed in other system, name has prefix External. ,eg: External
	// +kubebuilder:validation:Required
	DeployPath string `json:"deployPath,omitempty" protobuf:"bytes,3,opt,name=deployPath"`

	// annotations
	// +optional
	Annotations map[string]string `json:"annotations,omitempty" protobuf:"bytes,4,rep,name=annotations"`
	// service relation
	// +kubebuilder:validation:Required
	ServiceRelations []ServiceRelation `json:"serviceRelations,omitempty" protobuf:"bytes,5,rep,name=serviceRelations"`
	// metrics
	// +optional
	Metrics []ServiceMetrics `json:"metrics,omitempty" protobuf:"bytes,6,rep,name=metrics"`
}

type ServiceRelation struct {
	// Project Service Name
	// +kubebuilder:validation:Required
	TargetRef string `json:"targetRef,omitempty" protobuf:"bytes,1,opt,name=targetRef"`
	// target service version
	// +kubebuilder:validation:Required
	TargetVersion string `json:"targetVersion,omitempty" protobuf:"bytes,2,opt,name=targetVersion"`

	// The IP protocol for this port. Supports "TCP", "UDP", and "SCTP".
	// Default is TCP.
	// +kubebuilder:validation:Enum=TCP;UDP;SCTP
	// +kubebuilder:default:=TCP
	// +optional
	Protocol string `json:"protocol,omitempty" protobuf:"bytes,3,opt,name=protocol"`
	// The port that will be exposed by this service.
	// +optional
	Port int32 `json:"port,omitempty" protobuf:"varint,4,opt,name=port"`
	// +kubebuilder:validation:Optional
	// +optional
	Description string `json:"description,omitempty" protobuf:"bytes,5,opt,name=description"`
	// annotations
	// +optional
	Annotations map[string]string `json:"annotations,omitempty" protobuf:"bytes,6,rep,name=annotations"`
}
type ServiceMetrics struct {
	// metrics name
	// +kubebuilder:validation:Required
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	// metrics description
	// +kubebuilder:validation:Required
	Description string `json:"description,omitempty" protobuf:"bytes,2,opt,name=description"`
	// Metrics PromQL
	// +kubebuilder:validation:Required
	PromQL string `json:"promQl,omitempty" protobuf:"bytes,3,opt,name=promQl"`
	//Query Type
	// +kubebuilder:validation:Enum=query;query_range
	// +kubebuilder:default:=query
	QueryType string `json:"queryType,omitempty" protobuf:"bytes,4,opt,name=queryType"`
	// Metrics Unit
	// +kubebuilder:validation:Required
	Unit string `json:"unit,omitempty" protobuf:"bytes,5,opt,name=unit"`
	// prometheus relation information
	// +kubebuilder:validation:Required
	PrometheusRef string `json:"prometheusRef,omitempty" protobuf:"bytes,6,opt,name=prometheusRef"`
}

type PrometheusInfo struct {
	// prometheus
	// +kubebuilder:validation:Required
	Address string `json:"address,omitempty" protobuf:"bytes,5,opt,name=address"`
	// +kubebuilder:default:=/var/run/secrets/kubernetes.io/serviceaccount/token
	// prometheus token path or value, defult is pod sa token path
	Token string `json:"token,omitempty" protobuf:"bytes,6,opt,name=token"`
	// Selects a key of a ConfigMap.
	// +optional
	ConfigMapKeyRef *corev1.ConfigMapKeySelector `json:"configMapKeyRef,omitempty" protobuf:"bytes,3,opt,name=configMapKeyRef"`
	// Selects a key of a secret in the pod's namespace
	// +optional
	SecretKeyRef *corev1.SecretKeySelector `json:"secretKeyRef,omitempty" protobuf:"bytes,4,opt,name=secretKeyRef"`
}

// ProjectArchitectureStatus defines the observed state of ProjectArchitecture
type ProjectArchitectureStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Hash string `json:"hash,omitempty" protobuf:"bytes,1,opt,name=hash"`
	// prometheus health check
	PrometheusStatus map[string]string `json:"prometheusStatus,omitempty" protobuf:"bytes,2,rep,name=prometheusStatus"`
	// scrape timestamp
	Timestamp int64 `json:"timestamp,omitempty" protobuf:"varint,3,opt,name=timestamp"`
	// project service, key is service name
	Service map[string]ServiceMetricsStatus `json:"service,omitempty" protobuf:"bytes,4,rep,name=service"`
}

// ServiceMetricsStatus defines project service diff versions
type ServiceMetricsStatus struct {
	// key is service version
	Items map[string]ServiceItemStatus `json:"items,omitempty" protobuf:"bytes,1,rep,name=items"`
}

// ServiceItemStatus one service version metrics
type ServiceItemStatus struct {
	// key is metrics name
	Metrics map[string]ServiceItemMetricsStatus `json:"metrics,omitempty" protobuf:"bytes,1,rep,name=metrics"`
}
type VectorDataLine struct {
	Metric map[string]string `json:"metric,omitempty" protobuf:"bytes,1,rep,name=metric"`
	Value  MetricValue       `json:"value,omitempty" protobuf:"bytes,2,opt,name=value"`
}

type VectorMatrixLine struct {
	Metric map[string]string `json:"metric,omitempty" protobuf:"bytes,1,rep,name=metric"`
	Values []MetricValue     `json:"values,omitempty" protobuf:"bytes,2,rep,name=values"`
}
type MetricValue struct {
	Timestamp int64  `json:"timestamp,omitempty" protobuf:"varint,1,opt,name=timestamp"`
	Value     string `json:"value,omitempty" protobuf:"bytes,2,opt,name=value"`
}
type ServiceItemMetricsStatus struct {
	// metrics name
	// +kubebuilder:validation:Required
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	// metrics description
	// +kubebuilder:validation:Required
	Description string `json:"description,omitempty" protobuf:"bytes,2,opt,name=description"`
	// metrics data
	VectorData   []VectorDataLine   `json:"vectorData,omitempty" protobuf:"bytes,3,rep,name=vectorData"`
	VectorMatrix []VectorMatrixLine `json:"vectorMatrix,omitempty" protobuf:"bytes,4,rep,name=vectorMatrix"`
	// prometheus response error
	Error string `json:"error,omitempty" protobuf:"bytes,5,opt,name=error"`
	//Query Type
	// +kubebuilder:validation:Enum=query;query_range
	// +kubebuilder:default:=query
	QueryType string `json:"queryType,omitempty" protobuf:"bytes,6,opt,name=queryType"`
	// Metrics Unit
	// +kubebuilder:validation:Required
	Unit string `json:"unit,omitempty" protobuf:"bytes,7,opt,name=unit"`
	// prometheus address
	Prometheus string `json:"prometheus,omitempty" protobuf:"bytes,8,opt,name=prometheus"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProjectArchitecture is the Schema for the projectarchitectures API
type ProjectArchitecture struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   ProjectArchitectureSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status ProjectArchitectureStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProjectArchitectureList contains a list of ProjectArchitecture
type ProjectArchitectureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []ProjectArchitecture `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="ClusterRef",type=string,JSONPath=`.spec.clusterRef`
// +kubebuilder:printcolumn:name="NamespaceNumber",type=integer,JSONPath=`.status.namespaceNumber`
// +kubebuilder:resource:scope=Cluster

// ClusterDashboard  cluster information
type ClusterDashboard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   ClusterDashboardSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status ClusterDashboardStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}
type ClusterDashboardSpec struct {
	// cluster related
	// +kubebuilder:validation:Required
	ClusterRef string `json:"clusterRef" yaml:"clusterRef" protobuf:"bytes,1,opt,name=clusterRef"`
	// agent will get metrics from prometheus, item is builtin promql
	MetricsSpec []DashboardMetrics `json:"metricsSpec" yaml:"metricsSpec" protobuf:"bytes,2,rep,name=metricsSpec"`
}
type ClusterDashboardStatus struct {
	// agent last report time
	// +optional
	LastReportTime *metav1.Time `json:"lastReportTime" yaml:"lastReportTime" protobuf:"bytes,1,opt,name=lastReportTime"`
	// cluster's nodes information, agent will report this information
	// key is node role
	// +optional
	Nodes []NodeInfo `json:"nodes" yaml:"nodes" protobuf:"bytes,2,rep,name=nodes"`
	// cluster's namespace number, agent will report this information
	// +optional
	NamespaceNumber *int32 `json:"namespaceNumber" yaml:"namespaceNumber" protobuf:"varint,3,opt,name=namespaceNumber"`
	// +optional
	Metrics []DashboardMetricsValue `json:"metrics" yaml:"metrics" protobuf:"bytes,4,rep,name=metrics"`
}
type NodeInfo struct {
	Role   string `json:"role" yaml:"role" protobuf:"bytes,1,opt,name=role"`
	Number *int32 `json:"number" yaml:"number" protobuf:"varint,2,opt,name=number"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterDashboardList contains a list of ClusterDashboard
type ClusterDashboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []ClusterDashboard `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="ClusterRef",type=string,JSONPath=`.spec.clusterRef`
// +kubebuilder:printcolumn:name="NamespaceNumber",type=integer,JSONPath=`.status.namespaceNumber`
// +kubebuilder:resource:scope=Cluster

// WorkspaceDashboard  Workspace information
type WorkspaceDashboard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   WorkspaceDashboardSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status WorkspaceDashboardStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}
type WorkspaceDashboardSpec struct {
	// agent will get metrics from prometheus, item is builtin promql
	MetricsSpec []DashboardMetrics `json:"metrics" yaml:"metrics" protobuf:"bytes,2,rep,name=metrics"`
}
type WorkspaceDashboardStatus struct {
	// agent last report time
	// +optional
	LastReportTime *metav1.Time `json:"lastReportTime" yaml:"lastReportTime" protobuf:"bytes,1,opt,name=lastReportTime"`
	// workspace namespace number, agent will report this information
	// +optional
	NamespaceNumber *int32 `json:"namespaceNumber" yaml:"namespaceNumber" protobuf:"varint,2,opt,name=namespaceNumber"`
	// +optional
	Metrics []DashboardMetricsValue `json:"metrics" yaml:"metrics" protobuf:"bytes,3,rep,name=metrics"`
	// meters
	Meters []Meter `json:"meters" yaml:"meters" protobuf:"bytes,4,rep,name=meters"`
}
type Meter struct {
	Name      string `json:"name" yaml:"name" protobuf:"bytes,1,opt,name=name"`
	Category  string `json:"category" yaml:"category" protobuf:"bytes,2,opt,name=category"`
	Namespace string `json:"namespace" yaml:"namespace" protobuf:"bytes,3,opt,name=namespace"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkspaceDashboardList contains a list of WorkspaceDashboard
type WorkspaceDashboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []WorkspaceDashboard `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type DashboardMetrics struct {
	Name      string `json:"name" yaml:"name" protobuf:"bytes,1,opt,name=name"`
	Promql    string `json:"promql" yaml:"promql" protobuf:"bytes,2,opt,name=promql"`
	ValueType string `json:"valueType" yaml:"valueType" protobuf:"bytes,3,opt,name=valueType"`
	ValueUnit string `json:"valueUnit" yaml:"valueUnit" protobuf:"bytes,4,opt,name=valueUnit"`
}

type DashboardMetricsValue struct {
	Name  string `json:"name" yaml:"name" protobuf:"bytes,1,opt,name=name"`
	Value string `json:"value" yaml:"value" protobuf:"bytes,2,opt,name=value"`
	Type  string `json:"type" yaml:"type" protobuf:"bytes,3,opt,name=type"`
	Unit  string `json:"unit" yaml:"unit" protobuf:"bytes,4,opt,name=unit"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="WorkloadKind",type=string,JSONPath=`.status.workloadKind`
// +kubebuilder:printcolumn:name="Replicas",type=integer,JSONPath=`.status.replicas`
// +kubebuilder:printcolumn:name="Timestamp",type="string",JSONPath=`.status.timestamp`
// +kubebuilder:resource:scope=Namespaced

// Workload workload
type Workload struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// +optional
	Status WorkloadStatus `json:"status" yaml:"status" protobuf:"bytes,2,opt,name=status"`
}
type WorkloadStatus struct {
	// pod top owner reference's resource kind, eg: Deployment,StatefulSet
	// +optional
	WorkloadKind string `json:"workloadKind" yaml:"workloadKind" protobuf:"bytes,19,opt,name=workloadKind"`
	// replica, DaemonSet's Replicas will equal number of nodes
	// +optional
	Replicas int32 `json:"replicas,omitempty" yaml:"replicas" protobuf:"varint,14,opt,name=replicas"`
	// Total number of non-terminated pods targeted
	// +optional
	UpdatedReplicas int32 `json:"updatedReplicas,omitempty" protobuf:"varint,15,opt,name=updatedReplicas"`
	// Total number of ready pods targeted
	// +optional
	ReadyReplicas int32 `json:"readyReplicas,omitempty" protobuf:"varint,16,opt,name=readyReplicas"`
	// Total number of unavailable pods targeted.
	// +optional
	UnavailableReplicas int32 `json:"unavailableReplicas,omitempty" protobuf:"varint,17,opt,name=unavailableReplicas"`
	// one pod's request and limit
	// +optional
	Resources corev1.ResourceRequirements `json:"resources,omitempty" yaml:"resources" protobuf:"bytes,18,opt,name=resources"`
	// The following fields define time interval from which metrics were
	// collected from the interval [Timestamp-Window, Timestamp].
	// +optional
	Timestamp metav1.Time `json:"timestamp,omitempty" yaml:"timestamp" protobuf:"bytes,8,opt,name=timestamp"`
	// +optional
	Window metav1.Duration `json:"window,omitempty" yaml:"window" protobuf:"bytes,9,opt,name=window"`
	// The memory usage is the memory working set.
	// +optional
	Usage corev1.ResourceList `json:"usage,omitempty" yaml:"usage" protobuf:"bytes,10,rep,name=usage,casttype=k8s.io/api/core/v1.ResourceList,castkey=k8s.io/api/core/v1.ResourceName"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkloadList contains a list of Workload
type WorkloadList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Workload `json:"items" protobuf:"bytes,2,rep,name=items"`
}
