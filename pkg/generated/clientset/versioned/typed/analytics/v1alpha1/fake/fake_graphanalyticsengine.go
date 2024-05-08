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

	v1alpha1 "github.com/arangodb/kube-arangodb/pkg/apis/analytics/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGraphAnalyticsEngines implements GraphAnalyticsEngineInterface
type FakeGraphAnalyticsEngines struct {
	Fake *FakeAnalyticsV1alpha1
	ns   string
}

var graphanalyticsenginesResource = v1alpha1.SchemeGroupVersion.WithResource("graphanalyticsengines")

var graphanalyticsenginesKind = v1alpha1.SchemeGroupVersion.WithKind("GraphAnalyticsEngine")

// Get takes name of the graphAnalyticsEngine, and returns the corresponding graphAnalyticsEngine object, and an error if there is any.
func (c *FakeGraphAnalyticsEngines) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GraphAnalyticsEngine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(graphanalyticsenginesResource, c.ns, name), &v1alpha1.GraphAnalyticsEngine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GraphAnalyticsEngine), err
}

// List takes label and field selectors, and returns the list of GraphAnalyticsEngines that match those selectors.
func (c *FakeGraphAnalyticsEngines) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GraphAnalyticsEngineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(graphanalyticsenginesResource, graphanalyticsenginesKind, c.ns, opts), &v1alpha1.GraphAnalyticsEngineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GraphAnalyticsEngineList{ListMeta: obj.(*v1alpha1.GraphAnalyticsEngineList).ListMeta}
	for _, item := range obj.(*v1alpha1.GraphAnalyticsEngineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested graphAnalyticsEngines.
func (c *FakeGraphAnalyticsEngines) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(graphanalyticsenginesResource, c.ns, opts))

}

// Create takes the representation of a graphAnalyticsEngine and creates it.  Returns the server's representation of the graphAnalyticsEngine, and an error, if there is any.
func (c *FakeGraphAnalyticsEngines) Create(ctx context.Context, graphAnalyticsEngine *v1alpha1.GraphAnalyticsEngine, opts v1.CreateOptions) (result *v1alpha1.GraphAnalyticsEngine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(graphanalyticsenginesResource, c.ns, graphAnalyticsEngine), &v1alpha1.GraphAnalyticsEngine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GraphAnalyticsEngine), err
}

// Update takes the representation of a graphAnalyticsEngine and updates it. Returns the server's representation of the graphAnalyticsEngine, and an error, if there is any.
func (c *FakeGraphAnalyticsEngines) Update(ctx context.Context, graphAnalyticsEngine *v1alpha1.GraphAnalyticsEngine, opts v1.UpdateOptions) (result *v1alpha1.GraphAnalyticsEngine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(graphanalyticsenginesResource, c.ns, graphAnalyticsEngine), &v1alpha1.GraphAnalyticsEngine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GraphAnalyticsEngine), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGraphAnalyticsEngines) UpdateStatus(ctx context.Context, graphAnalyticsEngine *v1alpha1.GraphAnalyticsEngine, opts v1.UpdateOptions) (*v1alpha1.GraphAnalyticsEngine, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(graphanalyticsenginesResource, "status", c.ns, graphAnalyticsEngine), &v1alpha1.GraphAnalyticsEngine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GraphAnalyticsEngine), err
}

// Delete takes name of the graphAnalyticsEngine and deletes it. Returns an error if one occurs.
func (c *FakeGraphAnalyticsEngines) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(graphanalyticsenginesResource, c.ns, name, opts), &v1alpha1.GraphAnalyticsEngine{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGraphAnalyticsEngines) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(graphanalyticsenginesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.GraphAnalyticsEngineList{})
	return err
}

// Patch applies the patch and returns the patched graphAnalyticsEngine.
func (c *FakeGraphAnalyticsEngines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GraphAnalyticsEngine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(graphanalyticsenginesResource, c.ns, name, pt, data, subresources...), &v1alpha1.GraphAnalyticsEngine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GraphAnalyticsEngine), err
}
