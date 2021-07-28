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
	"fmt"

	"github.com/liamrathke/octant-kubeflow/pkg/markdown"
	"github.com/liamrathke/octant-kubeflow/pkg/plugin/actions"
	"github.com/liamrathke/octant-kubeflow/pkg/plugin/utilities"
	"github.com/liamrathke/octant-kubeflow/pkg/state"
	"github.com/vmware-tanzu/octant/pkg/action"
	"github.com/vmware-tanzu/octant/pkg/view/component"
)

func BuildInstallViewForCC(cc utilities.ClientContext) (component.Component, error) {
	// cc := utilities.ClientContext{Client: request.DashboardClient(), Context: request.Context()}

	switch state.GetState().Installer.Stage {
	case state.NOT_INSTALLED:
		return buildNotInstalledView(cc)
	case state.INSTALLING:
		return buildInstallingView(cc)
	default:
		return nil, fmt.Errorf("unable to find installer view based on state")
	}
}

func buildNotInstalledView(cc utilities.ClientContext) (component.Component, error) {
	prompt, err := markdown.FileToComponent("install/not_installed.md")

	payload := action.Payload{
		"action": actions.InstallKubeflow,
	}
	button := component.NewButton("I understand, install Kubeflow", payload)

	flexLayout := component.NewFlexLayout("Install Kubeflow")
	flexLayout.AddSections(component.FlexLayoutSection{
		{Width: component.WidthFull, View: prompt},
		{Width: component.WidthFull, View: button},
	})

	return flexLayout, err
}

func buildInstallingView(cc utilities.ClientContext) (component.Component, error) {
	text := component.NewText("Installing Kubeflow!")
	return text, nil
}
