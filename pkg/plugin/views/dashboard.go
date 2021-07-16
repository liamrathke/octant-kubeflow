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
	"errors"
	"fmt"

	"github.com/liamrathke/octant-kubeflow/pkg/kubeflow"
	"github.com/liamrathke/octant-kubeflow/pkg/state"

	"github.com/vmware-tanzu/octant/pkg/plugin/api"
	"github.com/vmware-tanzu/octant/pkg/plugin/service"
	"github.com/vmware-tanzu/octant/pkg/store"
	"github.com/vmware-tanzu/octant/pkg/view/component"
	"k8s.io/apimachinery/pkg/labels"

	"context"
)

const DASHBOARD_PORT = 8080

func BuildDashboardViewForRequest(request service.Request) (component.Component, error) {
	ctx := request.Context()
	client := request.DashboardClient()

	_, err := client.List(ctx, store.Key{
		APIVersion: "v1",
		Kind:       "Secret",
		Selector: &labels.Set{
			"owner": "kubeflow",
		},
	})

	if err != nil {
		return nil, err
	}

	dashboardPort, err := getDashboardPort()
	if err != nil {
		return nil, err
	} else if dashboardPort == 0 {
		dashboardPort, err = dashboardPortForward(client, ctx)
	}

	dashboardURL := fmt.Sprintf("http://localhost:%d", dashboardPort)
	dashboard := component.NewIFrame(dashboardURL, "")

	return dashboard, err
}

func getDashboardPort() (uint16, error) {
	state := state.GetState()
	if state.Dashboard.IsPortForwarded {
		return state.Dashboard.Port, nil
	}

	return 0, nil
}

func dashboardPortForward(client service.Dashboard, ctx context.Context) (uint16, error) {
	dashboardPodInfo, err := kubeflow.GetDashboardPod(client, ctx)
	if err != nil {
		return 0, err
	} else if !dashboardPodInfo.Found {
		return 0, errors.New("Dashboard pod could not be found")
	}

	request := api.PortForwardRequest{
		Namespace: dashboardPodInfo.Namespace,
		PodName:   dashboardPodInfo.PodName,
		Port:      DASHBOARD_PORT,
	}

	response, err := client.PortForward(ctx, request)
	if err != nil {
		return 0, err
	}

	port := response.Port

	state := state.GetState()
	state.Dashboard.IsPortForwarded = true
	state.Dashboard.Port = port

	return port, nil
}
