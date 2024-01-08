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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/arangodb/kube-arangodb/pkg/apis/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ArangoJobLister helps list ArangoJobs.
// All objects returned here must be treated as read-only.
type ArangoJobLister interface {
	// List lists all ArangoJobs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ArangoJob, err error)
	// ArangoJobs returns an object that can list and get ArangoJobs.
	ArangoJobs(namespace string) ArangoJobNamespaceLister
	ArangoJobListerExpansion
}

// arangoJobLister implements the ArangoJobLister interface.
type arangoJobLister struct {
	indexer cache.Indexer
}

// NewArangoJobLister returns a new ArangoJobLister.
func NewArangoJobLister(indexer cache.Indexer) ArangoJobLister {
	return &arangoJobLister{indexer: indexer}
}

// List lists all ArangoJobs in the indexer.
func (s *arangoJobLister) List(selector labels.Selector) (ret []*v1.ArangoJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ArangoJob))
	})
	return ret, err
}

// ArangoJobs returns an object that can list and get ArangoJobs.
func (s *arangoJobLister) ArangoJobs(namespace string) ArangoJobNamespaceLister {
	return arangoJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ArangoJobNamespaceLister helps list and get ArangoJobs.
// All objects returned here must be treated as read-only.
type ArangoJobNamespaceLister interface {
	// List lists all ArangoJobs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ArangoJob, err error)
	// Get retrieves the ArangoJob from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ArangoJob, error)
	ArangoJobNamespaceListerExpansion
}

// arangoJobNamespaceLister implements the ArangoJobNamespaceLister
// interface.
type arangoJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ArangoJobs in the indexer for a given namespace.
func (s arangoJobNamespaceLister) List(selector labels.Selector) (ret []*v1.ArangoJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ArangoJob))
	})
	return ret, err
}

// Get retrieves the ArangoJob from the indexer for a given namespace and name.
func (s arangoJobNamespaceLister) Get(name string) (*v1.ArangoJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("arangojob"), name)
	}
	return obj.(*v1.ArangoJob), nil
}
