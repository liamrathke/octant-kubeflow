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
	"github.com/vmware-tanzu/octant/pkg/plugin/service"
	"github.com/vmware-tanzu/octant/pkg/view/component"
)

var r *service.Router

func InitRoutes(router *service.Router) {
	r = router
	router.HandleFunc("/", func(request service.Request) (component.ContentResponse, error) {
		return component.ContentResponse{}, nil
	})
	routeHelper("", rootHandler)
	routeHelper("/dashboard", dashboardHandler)
}

func routeHelper(routePath string, handleFunc service.HandleFunc) {
	// var wrapMiddleware service.HandleFunc
	// wrapMiddleware = func(request service.Request) {
	// 	return handleMiddleware(request, handleFunc)
	// }
	r.HandleFunc(routePath, func(request service.Request) {
		return handleMiddleware(request, handleFunc)
	})
}

func handleMiddleware(request service.Request, handleFunc service.HandleFunc) (component.ContentResponse, error) {
	// Middleware code goes here
	// If Kubeflow not installed, redirect to information page (later installation page)
	view, err := handleFunc(request)
	if err != nil {
		return component.EmptyContentResponse, err
	}
	return view, nil
}
