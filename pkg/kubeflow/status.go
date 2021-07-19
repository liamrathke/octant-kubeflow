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
	"github.com/liamrathke/octant-kubeflow/pkg/plugin/utilities"
)

type KubeflowStatus struct {
	ServiceName string
	OK          bool
}

type ComponentStatus struct {
	Name        string
	Namespace   string
	TotalPods   int
	ReadyPods   int
	RunningPods int
}

var COMPONENTS = []ComponentStatus{
	{Name: "Certificate Manager", Namespace: "cert-manager"},
	{Name: "Istio (System)", Namespace: "istio-system"},
	{Name: "Auth", Namespace: "auth"},
	{Name: "Knative (Eventing)", Namespace: "knative-eventing"},
	{Name: "Knative (Serving)", Namespace: "knative-serving"},
	{Name: "Kubeflow", Namespace: "kubeflow"},
	{Name: "Kubeflow Example", Namespace: "kubeflow-user-example-com"},
}

func GetStatus() []KubeflowStatus {
	return []KubeflowStatus{
		{ServiceName: "Test1", OK: true},
		{ServiceName: "Test2", OK: true},
		{ServiceName: "Test3", OK: true},
		{ServiceName: "Test4", OK: true},
	}
}

func statusForComponent(cc utilities.ClientContext, component ComponentStatus) ComponentStatus {
	return ComponentStatus{}
}
