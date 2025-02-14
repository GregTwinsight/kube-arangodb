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

package metric_descriptions

import (
	"github.com/arangodb/kube-arangodb/pkg/util/metrics"
)

var (
	arangodbOperatorEnginePanicsRecovered = metrics.NewDescription("arangodb_operator_engine_panics_recovered", "Number of Panics recovered inside Operator reconciliation loop", []string{`section`}, nil)
)

func init() {
	registerDescription(arangodbOperatorEnginePanicsRecovered)
}

func NewArangodbOperatorEnginePanicsRecoveredCounterFactory() metrics.FactoryCounter[ArangodbOperatorEnginePanicsRecoveredInput] {
	return metrics.NewFactoryCounter[ArangodbOperatorEnginePanicsRecoveredInput]()
}

func NewArangodbOperatorEnginePanicsRecoveredInput(section string) ArangodbOperatorEnginePanicsRecoveredInput {
	return ArangodbOperatorEnginePanicsRecoveredInput{
		Section: section,
	}
}

type ArangodbOperatorEnginePanicsRecoveredInput struct {
	Section string `json:"section"`
}

func (i ArangodbOperatorEnginePanicsRecoveredInput) Counter(value float64) metrics.Metric {
	return ArangodbOperatorEnginePanicsRecoveredCounter(value, i.Section)
}

func (i ArangodbOperatorEnginePanicsRecoveredInput) Desc() metrics.Description {
	return ArangodbOperatorEnginePanicsRecovered()
}

func ArangodbOperatorEnginePanicsRecovered() metrics.Description {
	return arangodbOperatorEnginePanicsRecovered
}

func ArangodbOperatorEnginePanicsRecoveredCounter(value float64, section string) metrics.Metric {
	return ArangodbOperatorEnginePanicsRecovered().Counter(value, section)
}
