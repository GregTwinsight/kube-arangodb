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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha

import (
	"context"
	time "time"

	storagev1alpha "github.com/arangodb/kube-arangodb/pkg/apis/storage/v1alpha"
	versioned "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/arangodb/kube-arangodb/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha "github.com/arangodb/kube-arangodb/pkg/generated/listers/storage/v1alpha"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ArangoLocalStorageInformer provides access to a shared informer and lister for
// ArangoLocalStorages.
type ArangoLocalStorageInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha.ArangoLocalStorageLister
}

type arangoLocalStorageInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewArangoLocalStorageInformer constructs a new informer for ArangoLocalStorage type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewArangoLocalStorageInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredArangoLocalStorageInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredArangoLocalStorageInformer constructs a new informer for ArangoLocalStorage type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredArangoLocalStorageInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1alpha().ArangoLocalStorages().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1alpha().ArangoLocalStorages().Watch(context.TODO(), options)
			},
		},
		&storagev1alpha.ArangoLocalStorage{},
		resyncPeriod,
		indexers,
	)
}

func (f *arangoLocalStorageInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredArangoLocalStorageInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *arangoLocalStorageInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&storagev1alpha.ArangoLocalStorage{}, f.defaultInformer)
}

func (f *arangoLocalStorageInformer) Lister() v1alpha.ArangoLocalStorageLister {
	return v1alpha.NewArangoLocalStorageLister(f.Informer().GetIndexer())
}
