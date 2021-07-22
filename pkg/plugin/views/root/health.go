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
	TOTAL      = "Total Pods"
)

func BuildHealthTable(cc utilities.ClientContext) *component.Table {
	table := component.NewTableWithRows(
		"Kubeflow Component Health", "No Kubeflow services found!",
		component.NewTableCols(COMPONENT, CONTAINERS, PODS),
		[]component.TableRow{})

	for _, kfc := range kubeflow.GetHealth(cc) {
		tr := component.TableRow{
			COMPONENT:  component.NewText(kfc.Name),
			CONTAINERS: component.NewText(kfc.Containers.String()),
			PODS:       component.NewText(kfc.Pods.String()),
		}

		table.Add(tr)
	}

	table.Sort("Service")
	return table
}
