//
// DISCLAIMER
//
// Copyright 2023 ArangoDB GmbH, Cologne, Germany
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

package v1alpha1

import (
	"github.com/arangodb/kube-arangodb/pkg/apis/shared"
	"github.com/arangodb/kube-arangodb/pkg/util/errors"
)

type ArangoMLStorageSpecBackend struct {
	// S3 backend implements storage as a proxy to the provided S3 API endpoint
	S3 *ArangoMLStorageSpecBackendS3 `json:"s3,omitempty"`
}

func (s *ArangoMLStorageSpecBackend) GetS3() *ArangoMLStorageSpecBackendS3 {
	if s == nil || s.S3 == nil {
		return nil
	}
	return s.S3
}

func (s *ArangoMLStorageSpecBackend) Validate() error {
	if s == nil {
		return errors.Newf("Backend is not specified")
	}

	if s.S3 == nil {
		return errors.Newf("At least one backend needs to be defined")
	}

	return shared.WithErrors(shared.PrefixResourceError("s3", s.S3.Validate()))
}
