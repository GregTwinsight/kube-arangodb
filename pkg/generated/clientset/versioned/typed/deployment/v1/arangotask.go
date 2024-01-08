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

package v1

import (
	"context"
	"time"

	v1 "github.com/arangodb/kube-arangodb/pkg/apis/deployment/v1"
	scheme "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ArangoTasksGetter has a method to return a ArangoTaskInterface.
// A group's client should implement this interface.
type ArangoTasksGetter interface {
	ArangoTasks(namespace string) ArangoTaskInterface
}

// ArangoTaskInterface has methods to work with ArangoTask resources.
type ArangoTaskInterface interface {
	Create(ctx context.Context, arangoTask *v1.ArangoTask, opts metav1.CreateOptions) (*v1.ArangoTask, error)
	Update(ctx context.Context, arangoTask *v1.ArangoTask, opts metav1.UpdateOptions) (*v1.ArangoTask, error)
	UpdateStatus(ctx context.Context, arangoTask *v1.ArangoTask, opts metav1.UpdateOptions) (*v1.ArangoTask, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ArangoTask, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ArangoTaskList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ArangoTask, err error)
	ArangoTaskExpansion
}

// arangoTasks implements ArangoTaskInterface
type arangoTasks struct {
	client rest.Interface
	ns     string
}

// newArangoTasks returns a ArangoTasks
func newArangoTasks(c *DatabaseV1Client, namespace string) *arangoTasks {
	return &arangoTasks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the arangoTask, and returns the corresponding arangoTask object, and an error if there is any.
func (c *arangoTasks) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ArangoTask, err error) {
	result = &v1.ArangoTask{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("arangotasks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ArangoTasks that match those selectors.
func (c *arangoTasks) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ArangoTaskList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ArangoTaskList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("arangotasks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested arangoTasks.
func (c *arangoTasks) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("arangotasks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a arangoTask and creates it.  Returns the server's representation of the arangoTask, and an error, if there is any.
func (c *arangoTasks) Create(ctx context.Context, arangoTask *v1.ArangoTask, opts metav1.CreateOptions) (result *v1.ArangoTask, err error) {
	result = &v1.ArangoTask{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("arangotasks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(arangoTask).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a arangoTask and updates it. Returns the server's representation of the arangoTask, and an error, if there is any.
func (c *arangoTasks) Update(ctx context.Context, arangoTask *v1.ArangoTask, opts metav1.UpdateOptions) (result *v1.ArangoTask, err error) {
	result = &v1.ArangoTask{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("arangotasks").
		Name(arangoTask.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(arangoTask).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *arangoTasks) UpdateStatus(ctx context.Context, arangoTask *v1.ArangoTask, opts metav1.UpdateOptions) (result *v1.ArangoTask, err error) {
	result = &v1.ArangoTask{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("arangotasks").
		Name(arangoTask.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(arangoTask).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the arangoTask and deletes it. Returns an error if one occurs.
func (c *arangoTasks) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("arangotasks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *arangoTasks) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("arangotasks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched arangoTask.
func (c *arangoTasks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ArangoTask, err error) {
	result = &v1.ArangoTask{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("arangotasks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
