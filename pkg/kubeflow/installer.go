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
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"

	"github.com/liamrathke/octant-kubeflow/pkg/state"
)

const (
	DOWNLOAD_PATH = "/tmp/octant-kubeflow"
)

const (
	KUSTOMIZE_INSTALL_SCRIPT_URL      = "https://github.com/kubernetes-sigs/kustomize/blob/master/hack/install_kustomize.sh"
	KUSTOMIZE_INSTALL_SCRIPT_FILENAME = "install_kustomize.sh"
	KUSTOMIZE_VERSION                 = "3.2.0"
)

func Install() {
	var installer *state.Installer = &state.GetState().Installer
	installer.Stage = state.INSTALLING

	installer.Dependencies.Kustomize = isKustomizeInstalled()
	installer.Dependencies.Checked = true

	if !installer.Dependencies.Kustomize {
		out, err := installKustomize()
		if err != nil {
			installer.Dependencies.Errors = []string{err.Error()}
		} else {
			installer.Dependencies.Kustomize = true
			installer.Dependencies.Output = out
		}
	}

}

func isKustomizeInstalled() bool {
	out, err := exec.Command("kustomize", "help").Output()
	if err != nil {
		return false
	}

	return out != nil
}

func installKustomize() (string, error) {
	err := installKustomizeScript()
	if err != nil {
		return "", err
	}

	scriptPath := path.Join(DOWNLOAD_PATH, KUSTOMIZE_INSTALL_SCRIPT_FILENAME)
	out, err := exec.Command(scriptPath, KUSTOMIZE_VERSION).Output()

	return string(out), err
}

func installKustomizeScript() error {
	res, err := http.Get(KUSTOMIZE_INSTALL_SCRIPT_URL)
	script := res.Body
	defer script.Close()
	if err != nil {
		return err
	}

	filePath := path.Join(DOWNLOAD_PATH, KUSTOMIZE_INSTALL_SCRIPT_FILENAME)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, script)
	file.Close()

	return err
}
