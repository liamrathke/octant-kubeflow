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

package settings // import "github.com/liamrathke/octant-kubeflow/pkg/plugin/settings"

import (
	"strings"

	"github.com/liamrathke/octant-kubeflow/pkg/plugin/actions"

	"github.com/vmware-tanzu/octant/pkg/navigation"
	"github.com/vmware-tanzu/octant/pkg/plugin/service"

	"github.com/liamrathke/octant-kubeflow/pkg/plugin/router"
)

func GetOptions() []service.PluginOption {
	return []service.PluginOption{
		service.WithActionHandler(actions.ActionHandler),
		service.WithNavigation(
			func(request *service.NavigationRequest) (navigation.Navigation, error) {
				return navigation.Navigation{
					Title:    strings.Title(name),
					Path:     name,
					IconName: rootNavIcon,
					Children: []navigation.Navigation{
						{
							Title:    "Home",
							Path:     request.GeneratePath(""),
							IconName: "home",
						},
					},
				}, nil
			},
			router.InitRoutes,
		),
	}
}
