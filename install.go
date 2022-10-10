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

package api

import (
	"github.com/efucloud/api/cluster"
	"github.com/efucloud/api/monitor"
	"github.com/efucloud/api/rbac"
	"k8s.io/apimachinery/pkg/runtime"
)

var (
	schemeBuilder = runtime.NewSchemeBuilder(
		cluster.Install,
		rbac.Install,
		monitor.Install,
	)
	// Install is a function which adds every version of every openshift group to a scheme
	Install = schemeBuilder.AddToScheme
)
