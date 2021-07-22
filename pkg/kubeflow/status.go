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

package kubeflow // import "github.com/liamrathke/octant-kubeflow/pkg/kubeflow"

import (
	"fmt"

	"github.com/liamrathke/octant-kubeflow/pkg/plugin/utilities"
	corev1 "k8s.io/api/core/v1"
)

type ComponentStatus struct {
	Name        string
	Namespace   string
	Containers  Status
	Pods        Status
	TotalPods   int
	ReadyPods   int
	RunningPods int
}

type Status struct {
	OK    int
	Total int
}

func (s *Status) String() string {
	return fmt.Sprintf("%d/%d", s.OK, s.Total)
}

var COMPONENTS = []ComponentStatus{
	{Name: "Certificate Manager", Namespace: "cert-manager"},
	{Name: "Istio (System)", Namespace: "istio-system"},
	{Name: "Auth", Namespace: "auth"},
	{Name: "Knative (Eventing)", Namespace: "knative-eventing"},
	{Name: "Knative (Serving)", Namespace: "knative-serving"},
	{Name: "Kubeflow", Namespace: "kubeflow"},
}

func GetStatus(cc utilities.ClientContext) []ComponentStatus {
	statuses := make([]ComponentStatus, len(COMPONENTS))

	for c := range COMPONENTS {
		statuses[c], _ = getStatusForComponent(cc, COMPONENTS[c])
	}

	return statuses
}

func getStatusForComponent(cc utilities.ClientContext, component ComponentStatus) (ComponentStatus, error) {
	pods, err := getPodsInNamespace(cc, component.Namespace)
	if err != nil {
		return ComponentStatus{}, err
	}

	for _, pod := range pods {
		for _, status := range pod.Status.ContainerStatuses {
			component.Containers.Total++
			if status.Ready {
				component.Containers.OK++
			}
		}

		component.Pods.Total++
		if pod.Status.Phase == corev1.PodRunning {
			component.Pods.OK++
		}
	}

	return component, err
}
