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

	"github.com/vmware-tanzu/octant/pkg/plugin/api"
	"github.com/vmware-tanzu/octant/pkg/plugin/service"
	"github.com/vmware-tanzu/octant/pkg/store"
	"github.com/vmware-tanzu/octant/pkg/view/component"
	"k8s.io/apimachinery/pkg/labels"

	"context"
)

func BuildDashboardViewForRequest(request service.Request) (component.Component, error) {
	context := request.Context()
	client := request.DashboardClient()

	_, err := client.List(context, store.Key{
		APIVersion: "v1",
		Kind:       "Secret",
		Selector: &labels.Set{
			"owner": "kubeflow",
		},
	})

	if err != nil {
		return nil, err
	}

	// change namespace to istio-system
	// get list of port forwards
	// if istio-ingressgateway is port forwarded, use that port
	// otherwise, forward the port and use it

	dashboardPort, err := getDashboardPort()
	if err != nil {
		return nil, err
	} else if dashboardPort < 0 {
		dashboardPort, err = dashboardPortForward(client, context)
	}

	dashboardURL := fmt.Sprintf("http://localhost:%d", dashboardPort)
	dashboard := component.NewIFrame(dashboardURL, "")

	return dashboard, err
}

func getDashboardPort() (int, error) {
	return -1, nil
}

func dashboardPortForward(client service.Dashboard, context context.Context) (int, error) {
	request := api.PortForwardRequest{
		Namespace: "istio-system",
		PodName:   "istio-ingressgateway-56d9b7fdb-krwjh",
		Port:      8080,
	}

	response, err := client.PortForward(context, request)
	if err != nil {
		return -1, err
	}

	return int(response.Port), nil
}
