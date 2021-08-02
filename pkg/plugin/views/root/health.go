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
	"time"

	"github.com/vmware-tanzu/octant/pkg/view/component"

	"github.com/liamrathke/octant-kubeflow/pkg/kubeflow"
	"github.com/liamrathke/octant-kubeflow/pkg/plugin/utilities"
)

const (
	KUBEFLOW_HEALTH          = "Kubeflow Health"
	KUBEFLOW_SERVICES        = "Kubeflow Services"
	KUBEFLOW_UNREADY_PODS    = "Unready Pods"
	KUBEFLOW_NO_UNREADY_PODS = "No unready pods found!"
)

const (
	NAMESPACE = "Namespace"
	NAME      = "Pod Name"
	AGE       = "Age"
	ACTION    = "Action"
)

const (
	DONUT_SIZE      = component.DonutChartSizeMedium
	DONUT_THICKNESS = 25
)

func BuildHealthView(cc utilities.ClientContext) (component.Component, error) {
	kubeflowComponents, err := kubeflow.GetHealth(cc)

	if err != nil {
		return nil, err
	}

	flexLayout := component.NewFlexLayout(KUBEFLOW_HEALTH)

	services := buildServiceSection(cc, kubeflowComponents)
	flexLayout.AddSections(services)

	unready := buildUnreadySection(cc, kubeflowComponents)
	flexLayout.AddSections(unready)

	return flexLayout, err
}

func buildServiceSection(cc utilities.ClientContext, kfcs []kubeflow.KubeflowComponent) component.FlexLayoutSection {
	services := make(component.FlexLayoutSection, len(kfcs))
	for index, kfc := range kfcs {
		title := component.NewText(kfc.Name)
		card := component.NewCard([]component.TitleComponent{title})

		cardLayout := component.NewFlexLayout(KUBEFLOW_SERVICES)
		cardLayout.AddSections(component.FlexLayoutSection{
			{Width: component.WidthHalf, View: donutFromStatus(kfc.Containers, "Containers", "Container")},
			{Width: component.WidthHalf, View: donutFromStatus(kfc.Pods, "Pods", "Pod")},
		})

		card.SetBody(cardLayout)

		services[index] = component.FlexLayoutItem{Width: component.WidthThird, View: card}
	}
	return services
}

func buildUnreadySection(cc utilities.ClientContext, kfcs []kubeflow.KubeflowComponent) component.FlexLayoutSection {
	table := component.NewTableWithRows(
		KUBEFLOW_UNREADY_PODS, KUBEFLOW_NO_UNREADY_PODS,
		component.NewTableCols(NAMESPACE, NAME, AGE, ACTION),
		[]component.TableRow{})

	for _, kfc := range kfcs {
		for _, pod := range kfc.Unready {
			tr := component.TableRow{
				NAMESPACE: component.NewText(pod.Namespace),
				NAME:      component.NewText(pod.Name),
				AGE:       component.NewText(fmt.Sprintf("%d", time.Now().Sub(pod.Status.StartTime.Time))),
				ACTION:    component.NewText("Restart Pod"),
			}
			table.Add(tr)
		}
	}

	table.Sort(NAME)

	return component.FlexLayoutSection{
		{Width: component.WidthFull, View: table},
	}
}

func donutFromStatus(status kubeflow.Status, plural, singular string) component.Component {
	donut := component.NewDonutChart()
	donut.SetSize(DONUT_SIZE)
	donut.SetThickness(DONUT_THICKNESS)

	donut.SetLabels(plural, singular)

	segments := []component.DonutSegment{
		{Count: status.Up, Status: component.NodeStatusOK},
		{Count: status.Down, Status: component.NodeStatusWarning},
	}
	donut.SetSegments(segments)

	return donut
}
