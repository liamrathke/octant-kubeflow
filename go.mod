module github.com/liamrathke/octant-kubeflow

go 1.16

require (
	github.com/vmware-tanzu/octant v0.22.1-0.20210729190618-fc5ac32211c7
	k8s.io/api v0.21.3
	k8s.io/apimachinery v0.21.3
)

replace (
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/docker/docker => github.com/moby/moby v17.12.0-ce-rc1.0.20200618181300-9dc6525e6118+incompatible
)
