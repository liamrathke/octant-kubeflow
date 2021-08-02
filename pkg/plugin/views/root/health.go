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
	"github.com/vmware-tanzu/octant/pkg/view/component"

	"github.com/liamrathke/octant-kubeflow/pkg/kubeflow"
	"github.com/liamrathke/octant-kubeflow/pkg/plugin/utilities"
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

	flexLayout := component.NewFlexLayout("Kubeflow Health")

	services := buildServiceSection(cc, kubeflowComponents)
	flexLayout.AddSections(services)

	failing := buildFailingSection(cc, kubeflowComponents)
	flexLayout.AddSections(failing)

	return flexLayout, err
}

func buildServiceSection(cc utilities.ClientContext, kfcs []kubeflow.KubeflowComponent) component.FlexLayoutSection {
	services := make(component.FlexLayoutSection, len(kfcs))
	for index, kfc := range kfcs {
		title := component.NewText(kfc.Name)
		card := component.NewCard([]component.TitleComponent{title})

		cardLayout := component.NewFlexLayout("")
		cardLayout.AddSections(component.FlexLayoutSection{
			{Width: component.WidthHalf, View: donutFromStatus(kfc.Containers, "Containers", "Container")},
			{Width: component.WidthHalf, View: donutFromStatus(kfc.Pods, "Pods", "Pod")},
		})

		card.SetBody(cardLayout)

		services[index] = component.FlexLayoutItem{Width: component.WidthThird, View: card}
	}
	return services
}

func buildFailingSection(cc utilities.ClientContext, kfcs []kubeflow.KubeflowComponent) component.FlexLayoutSection {
	table := component.NewTableWithRows(
		"Failing Kubeflow Components", "No Kubeflow services found!",
		component.NewTableCols(COMPONENT, CONTAINERS, PODS),
		[]component.TableRow{})

	for _, kfc := range kfcs {
		tr := component.TableRow{
			COMPONENT:  component.NewText(kfc.Name),
			CONTAINERS: component.NewText(kfc.Containers.String()),
			PODS:       component.NewText(kfc.Pods.String()),
		}

		table.Add(tr)
	}

	return component.FlexLayoutSection{
		{Width: component.WidthFull, View: table},
	}
}

func donutFromStatus(status kubeflow.Status, plural, singular string) component.Component {
	donut := component.NewDonutChart()
	donut.SetLabels(plural, singular)
	donut.SetSize(component.DonutChartSizeMedium)

	segments := []component.DonutSegment{
		{Count: status.Up, Status: component.NodeStatusOK},
		{Count: status.Down, Status: component.NodeStatusWarning},
	}
	donut.SetSegments(segments)

	return donut
}
