# Install Kubeflow

It looks like Kubeflow hasn't been installed on this Kubernetes cluster. Fortunately, it's possible to install Kubeflow directly from this Octant plugin.

## How it works

In addition to packaged distributions, it's possible to install Kubeflow on existing clusters using the [Kubeflow Manifests](https://github.com/kubeflow/manifests#installation). The Octant Kubeflow plugin will download the latest manifests, validate that all prerequisites are installed, then attempt to install Kubeflow using the manifests.  

## ⚠️ Warning ⚠️

According to the [official Kubeflow documentation](https://www.kubeflow.org/docs/started/installing-kubeflow/#install-the-kubeflow-manifests-manually), 

> The Kubeflow community will not support environment-specific issues. If you need support, please consider using a [packaged distribution](https://www.kubeflow.org/docs/started/installing-kubeflow/#install-a-packaged-kubeflow-distribution) of Kubeflow. 