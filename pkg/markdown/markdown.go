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

package markdown // import "github.com/liamrathke/octant-kubeflow/pkg/plugin/markdown"

import (
	"embed"
	"path"

	"github.com/vmware-tanzu/octant/pkg/view/component"
)

//go:embed templates/*
var fs embed.FS

func FileToComponent(file string) (component.Component, error) {
	path := path.Join("templates", file)
	buf, err := fs.ReadFile(path)
	if err != nil {
		return nil, err
	}

	text := string(buf)
	return component.NewMarkdownText(text), err
}
