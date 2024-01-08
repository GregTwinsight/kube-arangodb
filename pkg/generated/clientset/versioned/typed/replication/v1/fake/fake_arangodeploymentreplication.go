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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/arangodb/kube-arangodb/pkg/apis/replication/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeArangoDeploymentReplications implements ArangoDeploymentReplicationInterface
type FakeArangoDeploymentReplications struct {
	Fake *FakeReplicationV1
	ns   string
}

var arangodeploymentreplicationsResource = v1.SchemeGroupVersion.WithResource("arangodeploymentreplications")

var arangodeploymentreplicationsKind = v1.SchemeGroupVersion.WithKind("ArangoDeploymentReplication")

// Get takes name of the arangoDeploymentReplication, and returns the corresponding arangoDeploymentReplication object, and an error if there is any.
func (c *FakeArangoDeploymentReplications) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ArangoDeploymentReplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(arangodeploymentreplicationsResource, c.ns, name), &v1.ArangoDeploymentReplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ArangoDeploymentReplication), err
}

// List takes label and field selectors, and returns the list of ArangoDeploymentReplications that match those selectors.
func (c *FakeArangoDeploymentReplications) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ArangoDeploymentReplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(arangodeploymentreplicationsResource, arangodeploymentreplicationsKind, c.ns, opts), &v1.ArangoDeploymentReplicationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ArangoDeploymentReplicationList{ListMeta: obj.(*v1.ArangoDeploymentReplicationList).ListMeta}
	for _, item := range obj.(*v1.ArangoDeploymentReplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested arangoDeploymentReplications.
func (c *FakeArangoDeploymentReplications) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(arangodeploymentreplicationsResource, c.ns, opts))

}

// Create takes the representation of a arangoDeploymentReplication and creates it.  Returns the server's representation of the arangoDeploymentReplication, and an error, if there is any.
func (c *FakeArangoDeploymentReplications) Create(ctx context.Context, arangoDeploymentReplication *v1.ArangoDeploymentReplication, opts metav1.CreateOptions) (result *v1.ArangoDeploymentReplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(arangodeploymentreplicationsResource, c.ns, arangoDeploymentReplication), &v1.ArangoDeploymentReplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ArangoDeploymentReplication), err
}

// Update takes the representation of a arangoDeploymentReplication and updates it. Returns the server's representation of the arangoDeploymentReplication, and an error, if there is any.
func (c *FakeArangoDeploymentReplications) Update(ctx context.Context, arangoDeploymentReplication *v1.ArangoDeploymentReplication, opts metav1.UpdateOptions) (result *v1.ArangoDeploymentReplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(arangodeploymentreplicationsResource, c.ns, arangoDeploymentReplication), &v1.ArangoDeploymentReplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ArangoDeploymentReplication), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeArangoDeploymentReplications) UpdateStatus(ctx context.Context, arangoDeploymentReplication *v1.ArangoDeploymentReplication, opts metav1.UpdateOptions) (*v1.ArangoDeploymentReplication, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(arangodeploymentreplicationsResource, "status", c.ns, arangoDeploymentReplication), &v1.ArangoDeploymentReplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ArangoDeploymentReplication), err
}

// Delete takes name of the arangoDeploymentReplication and deletes it. Returns an error if one occurs.
func (c *FakeArangoDeploymentReplications) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(arangodeploymentreplicationsResource, c.ns, name, opts), &v1.ArangoDeploymentReplication{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeArangoDeploymentReplications) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(arangodeploymentreplicationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ArangoDeploymentReplicationList{})
	return err
}

// Patch applies the patch and returns the patched arangoDeploymentReplication.
func (c *FakeArangoDeploymentReplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ArangoDeploymentReplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(arangodeploymentreplicationsResource, c.ns, name, pt, data, subresources...), &v1.ArangoDeploymentReplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ArangoDeploymentReplication), err
}
