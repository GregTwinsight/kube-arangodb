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

package v1beta1

import (
	"context"
	time "time"

	schedulerv1beta1 "github.com/arangodb/kube-arangodb/pkg/apis/scheduler/v1beta1"
	versioned "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/arangodb/kube-arangodb/pkg/generated/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/arangodb/kube-arangodb/pkg/generated/listers/scheduler/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ArangoSchedulerDeploymentInformer provides access to a shared informer and lister for
// ArangoSchedulerDeployments.
type ArangoSchedulerDeploymentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ArangoSchedulerDeploymentLister
}

type arangoSchedulerDeploymentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewArangoSchedulerDeploymentInformer constructs a new informer for ArangoSchedulerDeployment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewArangoSchedulerDeploymentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredArangoSchedulerDeploymentInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredArangoSchedulerDeploymentInformer constructs a new informer for ArangoSchedulerDeployment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredArangoSchedulerDeploymentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SchedulerV1beta1().ArangoSchedulerDeployments(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SchedulerV1beta1().ArangoSchedulerDeployments(namespace).Watch(context.TODO(), options)
			},
		},
		&schedulerv1beta1.ArangoSchedulerDeployment{},
		resyncPeriod,
		indexers,
	)
}

func (f *arangoSchedulerDeploymentInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredArangoSchedulerDeploymentInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *arangoSchedulerDeploymentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&schedulerv1beta1.ArangoSchedulerDeployment{}, f.defaultInformer)
}

func (f *arangoSchedulerDeploymentInformer) Lister() v1beta1.ArangoSchedulerDeploymentLister {
	return v1beta1.NewArangoSchedulerDeploymentLister(f.Informer().GetIndexer())
}
