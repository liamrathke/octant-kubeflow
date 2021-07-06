package actions

import (
	"fmt"

	"github.com/vmware-tanzu/octant/pkg/plugin/service"
)

const (
	UpdateHelmReleaseValues    = "octant-kubeflow.dev/update"
	UninstallHelmReleaseAction = "octant-kubeflow.dev/uninstall"
)

func ActionHandler(request *service.ActionRequest) error {
	actionName, err := request.Payload.String("action")
	if err != nil {
		return err
	}

	switch actionName {
	default:
		return fmt.Errorf("unable to find handler for plugin: %s", "octant-kubeflow")
	}
}
