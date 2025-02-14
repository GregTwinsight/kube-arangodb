//
// DISCLAIMER
//
// Copyright 2024 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

package v1beta1

import (
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/arangodb/kube-arangodb/pkg/apis/scheduler"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArangoSchedulerPodList is a list of Pods.
type ArangoSchedulerPodList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []ArangoSchedulerPod `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArangoSchedulerPod wraps core. ArangoSchedulerPod with profile details
type ArangoSchedulerPod struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArangoSchedulerPodSpec   `json:"spec"`
	Status ArangoSchedulerPodStatus `json:"status"`
}

type ArangoSchedulerPodSpec struct {
	// Profiles keeps list of the profiles
	Profiles []string `json:"profiles,omitempty"`

	core.PodSpec `json:",inline"`
}

type ArangoSchedulerPodStatus struct {
	ArangoSchedulerStatusMetadata `json:",inline"`

	core.PodStatus `json:",inline"`
}

// AsOwner creates an OwnerReference for the given  ArangoSchedulerPod
func (d *ArangoSchedulerPod) AsOwner() meta.OwnerReference {
	trueVar := true
	return meta.OwnerReference{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       scheduler.PodResourceKind,
		Name:       d.Name,
		UID:        d.UID,
		Controller: &trueVar,
	}
}

func (d *ArangoSchedulerPod) GetStatus() ArangoSchedulerPodStatus {
	return d.Status
}

func (d *ArangoSchedulerPod) SetStatus(status ArangoSchedulerPodStatus) {
	d.Status = status
}
