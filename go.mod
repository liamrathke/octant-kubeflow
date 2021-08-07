module github.com/liamrathke/octant-kubeflow

go 1.16

require (
	github.com/cheggaaa/pb v2.0.6+incompatible // indirect
	github.com/hashicorp/go-getter v1.5.6 // indirect
	github.com/vmware-tanzu/octant v0.22.1-0.20210729190618-fc5ac32211c7
	gopkg.in/VividCortex/ewma.v1 v1.1.1 // indirect
	gopkg.in/cheggaaa/pb.v2 v2.0.6 // indirect
	gopkg.in/fatih/color.v1 v1.7.0 // indirect
	gopkg.in/mattn/go-colorable.v0 v0.0.9 // indirect
	gopkg.in/mattn/go-isatty.v0 v0.0.4 // indirect
	gopkg.in/mattn/go-runewidth.v0 v0.0.3 // indirect
	k8s.io/api v0.21.3
	k8s.io/apimachinery v0.21.3
)

replace (
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/docker/docker => github.com/moby/moby v17.12.0-ce-rc1.0.20200618181300-9dc6525e6118+incompatible
)
