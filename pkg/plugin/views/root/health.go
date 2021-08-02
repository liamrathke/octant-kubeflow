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

package root // import "github.com/liamrathke/octant-kubeflow/pkg/plugin/views/root"

import (
	"fmt"

	"github.com/vmware-tanzu/octant/pkg/view/component"

	"github.com/liamrathke/octant-kubeflow/pkg/kubeflow"
	"github.com/liamrathke/octant-kubeflow/pkg/plugin/utilities"
)

const (
	STATUS_OK      = "✅"
	STATUS_WARNING = "⛔"
)

const (
	COMPONENT  = "Kubeflow Component"
	CONTAINERS = "Containers Ready"
	PODS       = "Pods Running"
)

func BuildHealthView(cc utilities.ClientContext) (component.Component, error) {
	kubeflowComponents, err := kubeflow.GetHealth(cc)

	if err != nil {
		return nil, err
	}

	flexLayout := component.NewFlexLayout("Kubeflow Component Status")

	for _, kfc := range kubeflowComponents {
		var statusEmoji string
		if kfc.OK {
			statusEmoji = STATUS_OK
		} else {
			statusEmoji = STATUS_WARNING
		}

		title := fmt.Sprintf("%s %s", kfc.Name, statusEmoji)

		quadrant := component.NewQuadrant(title)
		quadrant.Set(component.QuadNW, "Running", fmt.Sprintf("%d", kfc.Containers.Up))
		quadrant.Set(component.QuadNE, "Containers", fmt.Sprintf("%d", kfc.Containers.Total))
		quadrant.Set(component.QuadSW, "Ready", fmt.Sprintf("%d", kfc.Pods.Up))
		quadrant.Set(component.QuadSE, "Pods", fmt.Sprintf("%d", kfc.Containers.Total))

		flexLayout.AddSections(component.FlexLayoutSection{
			{Width: component.WidthQuarter, View: quadrant},
		})
	}

	table := component.NewTableWithRows(
		"Failing Kubeflow Components", "No Kubeflow services found!",
		component.NewTableCols(COMPONENT, CONTAINERS, PODS),
		[]component.TableRow{})

	for _, kfc := range kubeflowComponents {
		tr := component.TableRow{
			COMPONENT:  component.NewText(kfc.Name),
			CONTAINERS: component.NewText(kfc.Containers.String()),
			PODS:       component.NewText(kfc.Pods.String()),
		}

		table.Add(tr)
	}

	flexLayout.AddSections(component.FlexLayoutSection{
		{Width: component.WidthFull, View: table},
	})

	return flexLayout, err
}
