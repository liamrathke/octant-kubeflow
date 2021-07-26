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

package views // import "github.com/liamrathke/octant-kubeflow/pkg/plugin/views"

import (
	"github.com/liamrathke/octant-kubeflow/pkg/plugin/utilities"
	"github.com/vmware-tanzu/octant/pkg/view/component"
)

func BuildInstallViewForCC(cc utilities.ClientContext) (component.Component, error) {
	// cc := utilities.ClientContext{Client: request.DashboardClient(), Context: request.Context()}

	prompt := component.NewText("Install Kubeflow!")

	return prompt, nil
}
