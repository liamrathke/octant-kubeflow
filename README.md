# Octant Kubeflow

[![GitHub Actions status](https://github.com/liamrathke/octant-kubeflow/workflows/build/badge.svg)](https://github.com/liamrathke/octant-kubeflow/actions)

## Uninstall

Run the following command to remove the plugin:

```
rm -f ~/.config/octant/plugins/octant-kubeflow
```

## Development

Requires Go 1.13+ and [fswatch](https://github.com/emcrisostomo/fswatch).

Run `make dev` at the root of this repo, which will do the following:

- Build the binary
- Start the Octant server at http://127.0.0.1:7777
- Watch for file changes, rebuild binary, restart Octant, repeat

Contributions are welcome!

## Credits

This plugin was adapted from [Octant Helm](https://github.com/bloodorangeio/octant-helm) by [Blood Orange](https://github.com/bloodorangeio).
