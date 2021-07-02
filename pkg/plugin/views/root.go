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

	"github.com/vmware-tanzu/octant/pkg/plugin/service"
	"github.com/vmware-tanzu/octant/pkg/store"
	"github.com/vmware-tanzu/octant/pkg/view/component"
	"k8s.io/apimachinery/pkg/labels"
)

func BuildRootViewForRequest(request service.Request) (component.Component, error) {
	ctx := request.Context()
	client := request.DashboardClient()

	_, err := client.List(ctx, store.Key{
		APIVersion: "v1",
		Kind:       "Secret",
		Selector: &labels.Set{
			"owner": "helm",
		},
	})

	if err != nil {
		return nil, err
	}

	header := component.NewMarkdownText(fmt.Sprintf("## Kubeflow"))

	flexLayout := component.NewFlexLayout("Home")
	flexLayout.AddSections(component.FlexLayoutSection{
		{Width: component.WidthFull, View: header},
	})

	return flexLayout, nil
}
