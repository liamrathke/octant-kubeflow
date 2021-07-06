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
)

func BuildStatusTable() *component.Table {
	table := component.NewTableWithRows(
		"Status", "No Kubeflow services found!",
		component.NewTableCols("Service", "Status"),
		[]component.TableRow{})

	for _, s := range kubeflow.GetStatus() {
		tr := component.TableRow{
			"Service": component.NewText(s.ServiceName),
		}

		if s.OK {
			tr["Status"] = component.NewIcon("success-standard")
		} else {
			tr["Status"] = component.NewIcon("error-standard")
		}
		table.Add(tr)
	}

	table.Sort("Service")
	return table
}
