//
// DISCLAIMER
//
// Copyright 2016-2023 ArangoDB GmbH, Cologne, Germany
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

package shared

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/arangodb/kube-arangodb/pkg/util/errors"
)

var (
	resourceNameRE = regexp.MustCompile(`^([0-9\-\.a-z])+$`)
)

// ValidateResourceName validates a kubernetes resource name.
// If not valid, an error is returned.
// See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/
func ValidateResourceName(name string) error {
	if len(name) > 253 {
		return errors.WithStack(errors.Newf("Name '%s' is too long", name))
	}
	if resourceNameRE.MatchString(name) {
		return nil
	}
	return errors.WithStack(errors.Newf("Name '%s' is not a valid resource name", name))
}

// ValidateOptionalResourceName validates a kubernetes resource name.
// If not empty and not valid, an error is returned.
func ValidateOptionalResourceName(name string) error {
	if name == "" {
		return nil
	}
	if err := ValidateResourceName(name); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// ValidateUID validates if it is valid Kubernetes UID
func ValidateUID(uid types.UID) error {
	_, err := uuid.Parse(string(uid))
	return err
}

// ValidatePullPolicy Validates core.PullPolicy
func ValidatePullPolicy(in core.PullPolicy) error {
	switch in {
	case core.PullAlways, core.PullNever, core.PullIfNotPresent:
		return nil
	}

	return errors.Newf("Unknown pull policy: '%s'", string(in))
}

type ValidateInterface interface {
	Validate() error
}

func Validate[T interface{}](in T) error {
	return validate(in)
}

func validate(in any) error {
	if v, ok := in.(ValidateInterface); ok {
		return v.Validate()
	}
	return nil
}

// ValidateOptional Validates object if is not nil
func ValidateOptional[T interface{}](in *T, validator func(T) error) error {
	if in != nil {
		return validator(*in)
	}

	return nil
}

// ValidateRequired Validates object and required not nil value
func ValidateRequired[T interface{}](in *T, validator func(T) error) error {
	if in != nil {
		return validator(*in)
	}

	return errors.Newf("should be not nil")
}

// ValidateList validates all elements on the list
func ValidateList[T interface{}](in []T, validator func(T) error) error {
	errors := make([]error, len(in))

	for id := range in {
		errors[id] = PrefixResourceError(fmt.Sprintf("[%d]", id), validator(in[id]))
	}

	return WithErrors(errors...)
}

// ValidateImage Validates if provided image is valid
func ValidateImage(image string) error {
	if image == "" {
		return errors.Newf("Image should be not empty")
	}

	return nil
}

// ValidateAnyNotNil Validates if any of the specified objects is not nil
func ValidateAnyNotNil[T any](msg string, obj ...*T) error {
	for _, o := range obj {
		if o != nil {
			return nil
		}
	}

	return errors.Newf(msg)
}

// ValidateServiceType checks that service type is supported
func ValidateServiceType(st core.ServiceType) error {
	switch st {
	case core.ServiceTypeClusterIP,
		core.ServiceTypeNodePort,
		core.ServiceTypeLoadBalancer,
		core.ServiceTypeExternalName:
		return nil
	}
	return errors.Newf("Unsupported service type %s", st)
}
