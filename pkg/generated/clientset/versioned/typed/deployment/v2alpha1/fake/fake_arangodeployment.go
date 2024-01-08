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

	v2alpha1 "github.com/arangodb/kube-arangodb/pkg/apis/deployment/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeArangoDeployments implements ArangoDeploymentInterface
type FakeArangoDeployments struct {
	Fake *FakeDatabaseV2alpha1
	ns   string
}

var arangodeploymentsResource = v2alpha1.SchemeGroupVersion.WithResource("arangodeployments")

var arangodeploymentsKind = v2alpha1.SchemeGroupVersion.WithKind("ArangoDeployment")

// Get takes name of the arangoDeployment, and returns the corresponding arangoDeployment object, and an error if there is any.
func (c *FakeArangoDeployments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.ArangoDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(arangodeploymentsResource, c.ns, name), &v2alpha1.ArangoDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ArangoDeployment), err
}

// List takes label and field selectors, and returns the list of ArangoDeployments that match those selectors.
func (c *FakeArangoDeployments) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.ArangoDeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(arangodeploymentsResource, arangodeploymentsKind, c.ns, opts), &v2alpha1.ArangoDeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2alpha1.ArangoDeploymentList{ListMeta: obj.(*v2alpha1.ArangoDeploymentList).ListMeta}
	for _, item := range obj.(*v2alpha1.ArangoDeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested arangoDeployments.
func (c *FakeArangoDeployments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(arangodeploymentsResource, c.ns, opts))

}

// Create takes the representation of a arangoDeployment and creates it.  Returns the server's representation of the arangoDeployment, and an error, if there is any.
func (c *FakeArangoDeployments) Create(ctx context.Context, arangoDeployment *v2alpha1.ArangoDeployment, opts v1.CreateOptions) (result *v2alpha1.ArangoDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(arangodeploymentsResource, c.ns, arangoDeployment), &v2alpha1.ArangoDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ArangoDeployment), err
}

// Update takes the representation of a arangoDeployment and updates it. Returns the server's representation of the arangoDeployment, and an error, if there is any.
func (c *FakeArangoDeployments) Update(ctx context.Context, arangoDeployment *v2alpha1.ArangoDeployment, opts v1.UpdateOptions) (result *v2alpha1.ArangoDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(arangodeploymentsResource, c.ns, arangoDeployment), &v2alpha1.ArangoDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ArangoDeployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeArangoDeployments) UpdateStatus(ctx context.Context, arangoDeployment *v2alpha1.ArangoDeployment, opts v1.UpdateOptions) (*v2alpha1.ArangoDeployment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(arangodeploymentsResource, "status", c.ns, arangoDeployment), &v2alpha1.ArangoDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ArangoDeployment), err
}

// Delete takes name of the arangoDeployment and deletes it. Returns an error if one occurs.
func (c *FakeArangoDeployments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(arangodeploymentsResource, c.ns, name, opts), &v2alpha1.ArangoDeployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeArangoDeployments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(arangodeploymentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v2alpha1.ArangoDeploymentList{})
	return err
}

// Patch applies the patch and returns the patched arangoDeployment.
func (c *FakeArangoDeployments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.ArangoDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(arangodeploymentsResource, c.ns, name, pt, data, subresources...), &v2alpha1.ArangoDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ArangoDeployment), err
}
