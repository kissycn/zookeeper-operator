#!/bin/bash

set -e

# Load logging library
. /opt/dtweave/scripts/utils/liblog.sh

# Allow the container to be started with `--user`
if [[ "$1" = 'zkServer.sh' && "$(id -u)" = '0' ]]; then
    chown -R zookeeper "$DT_ROOT_DIR" "$ZOO_BASE_DIR" "$ZOO_CONF_DIR"
    exec gosu zookeeper "$0" "$@"
fi

#export ZOO_STATIC_CONF_FILE="${ZOO_CONF_DIR}/zoo.cfg"
#export ZOO_DYNAMIC_CONF_FILE="${ZOO_CONF_DIR}/zoo.cfg.dynamic"
export MY_ID_FILE=${ZOO_CONF_DIR}/myid

#HOSTNAME="$(hostname -s)"
HOSTNAME="web-1"

if [[ -f $MY_ID_FILE ]]; then
  export ZOO_SERVER_ID="$(cat $MY_ID_FILE)"
else
  if [[ $HOSTNAME =~ (.*)-([0-9]+)$ ]]; then
    ORD=${BASH_REMATCH[2]}
    export ZOO_SERVER_ID=$((ORD+1))
    info  "Writing my id: $ZOO_SERVER_ID to: MY_ID_FILE."
    echo $ZOO_SERVER_ID > $MY_ID_FILE
  else
    error "Failed to get index from hostname $HOST"
    exit 1
  fi
fi

exec "$@"