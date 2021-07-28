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

package router // import "github.com/liamrathke/octant-kubeflow/pkg/plugin/router"

import (
	"github.com/liamrathke/octant-kubeflow/pkg/kubeflow"
	"github.com/liamrathke/octant-kubeflow/pkg/plugin/utilities"
	"github.com/vmware-tanzu/octant/pkg/plugin/service"
	"github.com/vmware-tanzu/octant/pkg/view/component"
)

var r *service.Router

type Handler func(cc utilities.ClientContext) (component.Component, error)

func InitRoutes(router *service.Router) {
	r = router
	routeHelper("", rootHandler)
	routeHelper("/install", installHandler)
	routeHelper("/dashboard", dashboardHandler)
}

func routeHelper(routePath string, handler Handler) {
	r.HandleFunc(routePath, func(request service.Request) (component.ContentResponse, error) {
		return handleMiddleware(request, handler)
	})
}

func handleMiddleware(request service.Request, handler Handler) (component.ContentResponse, error) {
	cc := utilities.ClientContext{Client: request.DashboardClient(), Context: request.Context()}

	if !kubeflow.Validate(cc) {
		// redirectInstall(cc)
	}

	view, err := handler(cc)
	if err != nil {
		return component.EmptyContentResponse, err
	}
	response := component.NewContentResponse(nil)
	response.Add(view)
	return *response, nil
}

func redirectInstall(cc utilities.ClientContext) {
	// Redirect to install page
}
