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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/arangodb/kube-arangodb/pkg/apis/ml/v1alpha1"
	scheme "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ArangoMLStoragesGetter has a method to return a ArangoMLStorageInterface.
// A group's client should implement this interface.
type ArangoMLStoragesGetter interface {
	ArangoMLStorages(namespace string) ArangoMLStorageInterface
}

// ArangoMLStorageInterface has methods to work with ArangoMLStorage resources.
type ArangoMLStorageInterface interface {
	Create(ctx context.Context, arangoMLStorage *v1alpha1.ArangoMLStorage, opts v1.CreateOptions) (*v1alpha1.ArangoMLStorage, error)
	Update(ctx context.Context, arangoMLStorage *v1alpha1.ArangoMLStorage, opts v1.UpdateOptions) (*v1alpha1.ArangoMLStorage, error)
	UpdateStatus(ctx context.Context, arangoMLStorage *v1alpha1.ArangoMLStorage, opts v1.UpdateOptions) (*v1alpha1.ArangoMLStorage, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ArangoMLStorage, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ArangoMLStorageList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ArangoMLStorage, err error)
	ArangoMLStorageExpansion
}

// arangoMLStorages implements ArangoMLStorageInterface
type arangoMLStorages struct {
	client rest.Interface
	ns     string
}

// newArangoMLStorages returns a ArangoMLStorages
func newArangoMLStorages(c *MlV1alpha1Client, namespace string) *arangoMLStorages {
	return &arangoMLStorages{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the arangoMLStorage, and returns the corresponding arangoMLStorage object, and an error if there is any.
func (c *arangoMLStorages) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ArangoMLStorage, err error) {
	result = &v1alpha1.ArangoMLStorage{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("arangomlstorages").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ArangoMLStorages that match those selectors.
func (c *arangoMLStorages) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ArangoMLStorageList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ArangoMLStorageList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("arangomlstorages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested arangoMLStorages.
func (c *arangoMLStorages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("arangomlstorages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a arangoMLStorage and creates it.  Returns the server's representation of the arangoMLStorage, and an error, if there is any.
func (c *arangoMLStorages) Create(ctx context.Context, arangoMLStorage *v1alpha1.ArangoMLStorage, opts v1.CreateOptions) (result *v1alpha1.ArangoMLStorage, err error) {
	result = &v1alpha1.ArangoMLStorage{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("arangomlstorages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(arangoMLStorage).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a arangoMLStorage and updates it. Returns the server's representation of the arangoMLStorage, and an error, if there is any.
func (c *arangoMLStorages) Update(ctx context.Context, arangoMLStorage *v1alpha1.ArangoMLStorage, opts v1.UpdateOptions) (result *v1alpha1.ArangoMLStorage, err error) {
	result = &v1alpha1.ArangoMLStorage{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("arangomlstorages").
		Name(arangoMLStorage.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(arangoMLStorage).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *arangoMLStorages) UpdateStatus(ctx context.Context, arangoMLStorage *v1alpha1.ArangoMLStorage, opts v1.UpdateOptions) (result *v1alpha1.ArangoMLStorage, err error) {
	result = &v1alpha1.ArangoMLStorage{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("arangomlstorages").
		Name(arangoMLStorage.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(arangoMLStorage).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the arangoMLStorage and deletes it. Returns an error if one occurs.
func (c *arangoMLStorages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("arangomlstorages").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *arangoMLStorages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("arangomlstorages").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched arangoMLStorage.
func (c *arangoMLStorages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ArangoMLStorage, err error) {
	result = &v1alpha1.ArangoMLStorage{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("arangomlstorages").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
