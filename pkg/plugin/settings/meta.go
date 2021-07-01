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

const (
	name        = "kubeflow" // This should stay lowercase for routing purposes
	description = "Octant Kubeflow"
	rootNavIcon = "node-group" // See https://clarity.design/icons for all options
)

func GetName() string {
	return name
}

func GetDescription() string {
	return description
}
