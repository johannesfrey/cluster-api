/*
Copyright 2020 The Kubernetes Authors.

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

package remote

import (
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/cluster-api/api/core/v1beta2/index"
)

// Index is a helper to model the info passed to cache.IndexField.
//
// Deprecated: This will be removed in Cluster API v1.10, use clustercache.CacheOptionsIndex instead.
type Index struct {
	Object       client.Object
	Field        string
	ExtractValue client.IndexerFunc
}

// NodeProviderIDIndex is used to index Nodes by ProviderID.
//
// Deprecated: This will be removed in Cluster API v1.10, use clustercache.NodeProviderIDIndex instead.
var NodeProviderIDIndex = Index{
	Object:       &corev1.Node{},
	Field:        index.NodeProviderIDField,
	ExtractValue: index.NodeByProviderID,
}
