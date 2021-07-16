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
	"context"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/vmware-tanzu/octant/pkg/plugin/service"
	"github.com/vmware-tanzu/octant/pkg/store"
)

type PodInfo struct {
	Found   bool
	PodName string
}

type PodSpec struct {
	Namespace       string
	PodNameContains string
}

var DASHBOARD_POD_SPEC = PodSpec{
	Namespace:       "istio-system",
	PodNameContains: "istio-ingressgateway",
}

func GetDashboardPod(client service.Dashboard) (PodInfo, error) {
	return GetPodInfo(client, nil, DASHBOARD_POD_SPEC)
}

func GetPodInfo(client service.Dashboard, ctx context.Context, podSpec PodSpec) (PodInfo, error) {
	unstructuredPods, err := client.List(ctx, store.Key{
		APIVersion: "v1",
		Kind:       "Pod",
		Namespace:  podSpec.Namespace,
	})
	if err != nil {
		return PodInfo{}, err
	}

	var podList corev1.PodList
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredPods.UnstructuredContent(), &podList)
	if err != nil {
		return PodInfo{}, err
	}

	for _, pod := range podList.Items {
		isValidPodName := strings.Contains(pod.Name, podSpec.PodNameContains)
		isRunningPod := pod.Status.Phase == corev1.PodRunning
		if isValidPodName && isRunningPod {
			return PodInfo{Found: true, PodName: pod.Name}, nil
		}
	}

	return PodInfo{}, nil
}
