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

type KubeflowComponent struct {
	Name       string
	Namespace  string
	Containers Status
	Pods       Status
	OK         bool
}

type Status struct {
	Up    int
	Down  int
	Total int
}

func (s *Status) String() string {
	return fmt.Sprintf("%d/%d", s.Up, s.Total)
}

var COMPONENTS = []KubeflowComponent{
	{Name: "Certificate Manager", Namespace: "cert-manager"},
	{Name: "Istio (System)", Namespace: "istio-system"},
	{Name: "Auth", Namespace: "auth"},
	{Name: "Knative (Eventing)", Namespace: "knative-eventing"},
	{Name: "Knative (Serving)", Namespace: "knative-serving"},
	{Name: "Kubeflow", Namespace: "kubeflow"},
}

func GetHealth(cc utilities.ClientContext) ([]KubeflowComponent, error) {
	var err error
	statuses := make([]KubeflowComponent, len(COMPONENTS))

	for c := range COMPONENTS {
		statuses[c], err = getHealthForComponent(cc, COMPONENTS[c])
	}

	return statuses, err
}

func getHealthForComponent(cc utilities.ClientContext, kfc KubeflowComponent) (KubeflowComponent, error) {
	pods, err := getPodsInNamespace(cc, kfc.Namespace)
	if err != nil {
		return KubeflowComponent{}, err
	}

	kfc.OK = true

	for _, pod := range pods {
		for _, status := range pod.Status.ContainerStatuses {
			kfc.Containers.Total++
			if status.Ready {
				kfc.Containers.Up++
			} else {
				kfc.Containers.Down++
				kfc.OK = false
			}
		}

		kfc.Pods.Total++
		if pod.Status.Phase == corev1.PodRunning {
			kfc.Pods.Up++
		} else {
			kfc.Pods.Down++
			kfc.OK = false
		}
	}

	return kfc, err
}
