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
	"github.com/vmware-tanzu/octant/pkg/action"
	"github.com/vmware-tanzu/octant/pkg/store"
	"github.com/vmware-tanzu/octant/pkg/view/component"
	v1 "k8s.io/api/core/v1"

	"github.com/liamrathke/octant-kubeflow/pkg/kubeflow"
	"github.com/liamrathke/octant-kubeflow/pkg/plugin/utilities"
)

const (
	KUBEFLOW_HEALTH          = "Kubeflow Health"
	KUBEFLOW_SERVICES        = "Kubeflow Services"
	KUBEFLOW_UNREADY_PODS    = "Unready Pods"
	KUBEFLOW_NO_UNREADY_PODS = "No unready pods found!"
	KUBEFLOW_RESTART_POD     = "Restart Pod"
)

const (
	NAMESPACE = "Namespace"
	NAME      = "Pod Name"
	AGE       = "Age"
	ACTION    = "Action"
)

const (
	DONUT_SIZE      = component.DonutChartSizeMedium
	DONUT_THICKNESS = 20
)

const (
	OCTANT_DELETE_OBJECT_ACTION = "action.octant.dev/deleteObject"
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
			tr := unreadyTableRow(pod)
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

func unreadyTableRow(pod v1.Pod) component.TableRow {
	namespace := component.NewText(pod.Namespace)
	name := component.NewText(pod.Name)
	age := component.NewTimestamp(pod.CreationTimestamp.Time)

	deleteAction := deletePodAction(pod.Namespace, pod.Name)
	action := component.NewButton(KUBEFLOW_RESTART_POD, deleteAction)

	return component.TableRow{
		NAMESPACE: namespace,
		NAME:      name,
		AGE:       age,
		ACTION:    action,
	}
}

func deletePodAction(namespace, name string) action.Payload {
	key := store.Key{
		APIVersion: "v1",
		Kind:       "Pod",
		Namespace:  namespace,
		Name:       name,
	}

	return action.CreatePayload(OCTANT_DELETE_OBJECT_ACTION, key.ToActionPayload())
}
