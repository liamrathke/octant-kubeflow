# Octant Kubeflow

[![GitHub Actions status](https://github.com/liamrathke/octant-kubeflow/workflows/build/badge.svg)](https://github.com/liamrathke/octant-kubeflow/actions)

<img width="960" alt="Presentation1" src="https://user-images.githubusercontent.com/33555592/129067635-86bd36ac-1d20-4280-b473-f487213f1ac8.png">

An open-source Octant plugin that helps machine learning (ML) engineers debug and maintain their Kubeflow workloads, without relying on a Kubernetes background. 

## Features

![Screen Shot 2021-08-11 at 12 35 37 PM](https://user-images.githubusercontent.com/33555592/129068191-0c9de97c-a24e-4dbd-aced-7a464cab45be.png)

### Completed

- ✅ **Health Check**: Status page for all Kubeflow services, restart unready pods instantly
- ✅ **Central Dashboard**: Kubeflow dashboard embedded within Octant, manage ML tasks without switching windows

### Under Development

- ⚠️ **Installer**: Add Kubeflow to an existing Kubernetes cluster with a single click

## Uninstall

Run the following command to remove the plugin:

```
rm -f ~/.config/octant/plugins/octant-kubeflow
```

## Development

Requires Go 1.16+ and [fswatch](https://github.com/emcrisostomo/fswatch).

Run `make dev` at the root of this repo, which will do the following:

- Build the binary
- Start the Octant server at http://127.0.0.1:7777
- Watch for file changes, rebuild binary, restart Octant, repeat

Contributions are welcome!

## Credits

This plugin was adapted from [Octant Helm](https://github.com/bloodorangeio/octant-helm) by [Blood Orange](https://github.com/bloodorangeio).
