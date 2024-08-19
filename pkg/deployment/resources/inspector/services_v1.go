//
// DISCLAIMER
//
// Copyright 2016-2024 ArangoDB GmbH, Cologne, Germany
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

package inspector

import (
	"context"

	core "k8s.io/api/core/v1"
	apiErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/arangodb/kube-arangodb/pkg/util/errors"
	"github.com/arangodb/kube-arangodb/pkg/util/k8sutil/inspector/constants"
	ins "github.com/arangodb/kube-arangodb/pkg/util/k8sutil/inspector/service/v1"
)

func (p *servicesInspector) V1() ins.Inspector {
	return p.v1
}

type servicesInspectorV1 struct {
	serviceInspector *servicesInspector

	services map[string]*core.Service
	err      error
}

func (p *servicesInspectorV1) validate() error {
	if p == nil {
		return errors.Errorf("ServicesV1Inspector is nil")
	}

	if p.serviceInspector == nil {
		return errors.Errorf("Parent is nil")
	}

	if p.services == nil {
		return errors.Errorf("Services or err should be not nil")
	}

	if p.err != nil {
		return errors.Errorf("Services or err cannot be not nil together")
	}

	return nil
}

func (p *servicesInspectorV1) GetSimple(name string) (*core.Service, bool) {
	service, ok := p.services[name]
	if !ok {
		return nil, false
	}

	return service, true
}

func (p *servicesInspectorV1) Iterate(action ins.Action, filters ...ins.Filter) error {
	for _, service := range p.services {
		if err := p.iterateService(service, action, filters...); err != nil {
			return err
		}
	}

	return nil
}

func (p *servicesInspectorV1) iterateService(service *core.Service, action ins.Action, filters ...ins.Filter) error {
	for _, f := range filters {
		if f == nil {
			continue
		}

		if !f(service) {
			return nil
		}
	}

	return action(service)
}

func (p *servicesInspectorV1) Read() ins.ReadInterface {
	return p
}

func (p *servicesInspectorV1) Get(ctx context.Context, name string, opts meta.GetOptions) (*core.Service, error) {
	if s, ok := p.GetSimple(name); !ok {
		return nil, apiErrors.NewNotFound(constants.ServiceGR(), name)
	} else {
		return s, nil
	}
}
