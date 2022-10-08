#!/bin/bash

# Load logging library
. /opt/dtweave/scripts/utils/liblog.sh


# export paths
export DT_ROOT_DIR="/opt/dtweave"
export ZOO_BASE_DIR="${DT_ROOT_DIR}/zookeeper"
export ZOO_DATA_DIR="${DATA_DIR}"
export ZOO_DATA_LOG_DIR="${DATA_LOG_DIR}"
export ZOO_CONF_DIR="${ZOO_BASE_DIR}/conf"
#export ZOO_STATIC_CONF_FILE="${ZOO_CONF_DIR}/zoo.cfg"
#export ZOO_DYNAMIC_CONF_FILE="${ZOO_CONF_DIR}/zoo.cfg.dynamic"
export MY_ID_FILE=${ZOO_CONF_DIR}/myid

#HOSTNAME="$(hostname -s)"
HOSTNAME="web-1"

function setServerId() {
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
  return
}

setServerId