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

package integrations

import (
	"context"

	"github.com/spf13/cobra"

	pbImplSchedulerV1 "github.com/arangodb/kube-arangodb/integrations/scheduler/v1"
	"github.com/arangodb/kube-arangodb/pkg/util/constants"
	"github.com/arangodb/kube-arangodb/pkg/util/errors"
	"github.com/arangodb/kube-arangodb/pkg/util/kclient"
	"github.com/arangodb/kube-arangodb/pkg/util/svc"
)

func init() {
	register(func() Integration {
		return &schedulerV1{}
	})
}

type schedulerV1 struct {
	Configuration pbImplSchedulerV1.Configuration
}

func (b *schedulerV1) Name() string {
	return "scheduler.v1"
}

func (b *schedulerV1) Description() string {
	return "SchedulerV1 Integration"
}

func (b *schedulerV1) Register(cmd *cobra.Command, arg ArgGen) error {
	f := cmd.Flags()

	f.StringVar(&b.Configuration.Namespace, arg("namespace"), constants.NamespaceWithDefault("default"), "Kubernetes Namespace")
	f.BoolVar(&b.Configuration.VerifyAccess, arg("verify-access"), true, "Verify the CRD Access")

	return nil
}

func (b *schedulerV1) Handler(ctx context.Context) (svc.Handler, error) {
	client, ok := kclient.GetDefaultFactory().Client()
	if !ok {
		return nil, errors.Errorf("Unable to create Kubernetes Client")
	}

	return pbImplSchedulerV1.New(ctx, client, b.Configuration)
}
