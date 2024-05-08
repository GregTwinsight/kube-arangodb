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

package v1alpha1

import (
	"context"
	time "time"

	analyticsv1alpha1 "github.com/arangodb/kube-arangodb/pkg/apis/analytics/v1alpha1"
	versioned "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/arangodb/kube-arangodb/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/arangodb/kube-arangodb/pkg/generated/listers/analytics/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GraphAnalyticsEngineInformer provides access to a shared informer and lister for
// GraphAnalyticsEngines.
type GraphAnalyticsEngineInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.GraphAnalyticsEngineLister
}

type graphAnalyticsEngineInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewGraphAnalyticsEngineInformer constructs a new informer for GraphAnalyticsEngine type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGraphAnalyticsEngineInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGraphAnalyticsEngineInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredGraphAnalyticsEngineInformer constructs a new informer for GraphAnalyticsEngine type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGraphAnalyticsEngineInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AnalyticsV1alpha1().GraphAnalyticsEngines(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AnalyticsV1alpha1().GraphAnalyticsEngines(namespace).Watch(context.TODO(), options)
			},
		},
		&analyticsv1alpha1.GraphAnalyticsEngine{},
		resyncPeriod,
		indexers,
	)
}

func (f *graphAnalyticsEngineInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGraphAnalyticsEngineInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *graphAnalyticsEngineInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&analyticsv1alpha1.GraphAnalyticsEngine{}, f.defaultInformer)
}

func (f *graphAnalyticsEngineInformer) Lister() v1alpha1.GraphAnalyticsEngineLister {
	return v1alpha1.NewGraphAnalyticsEngineLister(f.Informer().GetIndexer())
}
