#!/usr/bin/env bash
set -euo pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

REQUIRED_SYSTEM_COMMANDS=(
  "go"
  "octant"
  "fswatch"
)

function run_octant {
  local extra_args="$@"
  octant --plugin-path=${DIR}/../bin/ ${extra_args}
}

function run_octant_local {
  (cd $HOME/work/octant && ls && OCTANT_DISABLE_OPEN_BROWSER=true OCTANT_PLUGIN_PATH=${DIR}/../bin/ ./build/octant)
}

function log {
  echo "[dev.sh] $@"
}

function build_plugin {
  log "Building plugin..."
  ( cd ${DIR}/../ && go build -o bin/octant-kubeflow cmd/octant-kubeflow/main.go )
  log "Plugin built to bin/octant-kubeflow"
}

function start_octant {
  local extra_args="$@"
  log "Starting Octant..."
  run_octant_local &
  log "Octant available at http://127.0.0.1:7777/"
}

function stop_octant {
  log "Stopping Octant..."
  pkill -9 octant && kill -9 $(lsof -t -i:7777 && lsof -t -i:4200) || true
}

function watch_files {
  (
    cd ${DIR}/../ && eval "export $(go env | grep GOPATH)" && fswatch -0 -o \
      $(go list ./... | awk -v gp="${GOPATH}/src/" '{print gp $1 "/*.go"}' | xargs)
  )
}

function main {
  for c in ${REQUIRED_SYSTEM_COMMANDS[@]}; do
    if [[ ! -x "$(command -v ${c})" ]]; then
      log "System command missing: $c"
      exit 1
    fi
  done

  stop_octant
  build_plugin
  start_octant

  trap "log \"Goodbye\"" EXIT
  watch_files | while read -d "" event; do
    log "Source change detected, rebuilding plugin..."
    build_plugin
    log "Plugin rebuilt, restarting octant..."
    stop_octant
    start_octant --disable-open-browser
  done
}

main
