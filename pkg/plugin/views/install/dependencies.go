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

package install // import "github.com/liamrathke/octant-kubeflow/pkg/plugin/views/install"

import (
	"fmt"

	"github.com/liamrathke/octant-kubeflow/pkg/plugin/utilities"
	"github.com/liamrathke/octant-kubeflow/pkg/state"
	"github.com/vmware-tanzu/octant/pkg/view/component"
)

const (
	OUTPUT_TEMPLATE = "```\n%s```"
)

func BuildDependencyCard(cc utilities.ClientContext) (component.Component, error) {
	title := component.NewText("Dependencies")
	dependencies := component.NewCard(component.Title(title))
	layout := component.NewFlexLayout("")

	dependencyState := state.GetState().Installer.Dependencies

	if dependencyState.Checked {
		layout.AddSections(component.FlexLayoutSection{
			{Width: component.WidthFull, View: component.NewText("Checking dependencies...")},
		})
	} else {
		output := dependencyState.Output
		outputMarkdown := component.NewMarkdownText(fmt.Sprintf(OUTPUT_TEMPLATE, output))
		layout.AddSections(component.FlexLayoutSection{
			{Width: component.WidthFull, View: outputMarkdown},
		})
	}

	return dependencies, nil
}
