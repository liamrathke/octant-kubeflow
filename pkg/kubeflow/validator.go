/*
Copyright 2021 Liam Rathke/VMware

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kubeflow // import "github.com/liamrathke/octant-kubeflow/pkg/kubeflow"

import (
	"time"

	"github.com/liamrathke/octant-kubeflow/pkg/plugin/utilities"
	"github.com/liamrathke/octant-kubeflow/pkg/state"
)

func Validate(cc utilities.ClientContext) bool {
	state := state.GetState()
	if !state.Validator.Timestamp.IsZero() {
		state.Validator.Installed = validateComponents(cc)
		state.Validator.Timestamp = time.Now()
	}
	return state.Validator.Installed
}

func validateComponents(cc utilities.ClientContext) bool {
	_, err := GetHealth(cc)
	return err == nil
}
