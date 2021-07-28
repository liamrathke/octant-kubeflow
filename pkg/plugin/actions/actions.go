package actions

import (
	"fmt"

	"github.com/liamrathke/octant-kubeflow/pkg/kubeflow"
	"github.com/vmware-tanzu/octant/pkg/plugin/service"
)

const (
	InstallKubeflow = "octant-kubeflow.dev/install"
)

func ActionHandler(request *service.ActionRequest) error {
	actionName, err := request.Payload.String("action")
	if err != nil {
		return err
	}

	switch actionName {
	case InstallKubeflow:
		kubeflow.Install()
		return nil
	default:
		return fmt.Errorf("unable to find handler for plugin: %s", "octant-kubeflow")
	}
}
