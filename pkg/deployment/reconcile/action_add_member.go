//
// DISCLAIMER
//
// Copyright 2016-2022 ArangoDB GmbH, Cologne, Germany
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

package reconcile

import (
	"context"

	"github.com/arangodb/kube-arangodb/pkg/deployment/topology"

	"github.com/arangodb/kube-arangodb/pkg/util/errors"

	api "github.com/arangodb/kube-arangodb/pkg/apis/deployment/v1"
	"github.com/arangodb/kube-arangodb/pkg/deployment/actions"
)

func init() {
	registerAction(api.ActionTypeAddMember, newAddMemberAction, addMemberTimeout)
}

// newAddMemberAction creates a new Action that implements the given
// planned AddMember action.
func newAddMemberAction(action api.Action, actionCtx ActionContext) Action {
	a := &actionAddMember{}

	a.actionImpl = newBaseActionImpl(action, actionCtx, &a.newMemberID)

	return a
}

var _ ActionPlanAppender = &actionAddMember{}

// actionAddMember implements an AddMemberAction.
type actionAddMember struct {
	// actionImpl implement timeout and member id functions
	actionImpl

	// actionEmptyCheckProgress implement check progress with empty implementation
	actionEmptyCheckProgress

	newMemberID string
}

// Start performs the start of the action.
// Returns true if the action is completely finished, false in case
// the start time needs to be recorded and a ready condition needs to be checked.
func (a *actionAddMember) Start(ctx context.Context) (bool, error) {
	newID, err := a.actionCtx.CreateMember(ctx, a.action.Group, a.action.MemberID, topology.WithTopologyMod)
	if err != nil {
		a.log.Err(err).Debug("Failed to create member")
		return false, errors.WithStack(err)
	}
	a.newMemberID = newID

	return true, nil
}

// ActionPlanAppender appends wait methods to the plan
func (a *actionAddMember) ActionPlanAppender(current api.Plan) (api.Plan, bool) {
	var app api.Plan

	if _, ok := a.action.Params[api.ActionTypeWaitForMemberUp.String()]; ok {
		app = append(app, actions.NewAction(api.ActionTypeWaitForMemberUp, a.action.Group, withPredefinedMember(a.newMemberID), "Wait for member in sync after creation"))
	}

	if _, ok := a.action.Params[api.ActionTypeWaitForMemberInSync.String()]; ok {
		app = append(app, actions.NewAction(api.ActionTypeWaitForMemberInSync, a.action.Group, withPredefinedMember(a.newMemberID), "Wait for member in sync after creation"))
	}

	if len(app) > 0 {
		return app.AfterFirst(func(a api.Action) bool {
			return a.Type == api.ActionTypeAddMember
		}, app...), true
	}

	return current, false
}
