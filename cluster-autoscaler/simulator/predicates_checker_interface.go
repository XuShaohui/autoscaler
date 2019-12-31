/*
Copyright 2016 The Kubernetes Authors.

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

package simulator

import (
	apiv1 "k8s.io/api/core/v1"
	scheduler_nodeinfo "k8s.io/kubernetes/pkg/scheduler/nodeinfo"
)

// PredicateChecker checks whether all required predicates pass for given Pod and Node.
type PredicateChecker interface {
	SnapshotClusterState() error
	FitsAny(pod *apiv1.Pod, nodeInfos map[string]*scheduler_nodeinfo.NodeInfo) (string, error)
	CheckPredicates(pod *apiv1.Pod, nodeInfo *scheduler_nodeinfo.NodeInfo) PredicateError
}