#!/bin/bash

set -e

# Load logging library
. /opt/dtweave/scripts/utils/liblog.sh

export ZOO_BASE_DIR=/opt/modules/zookeeper-3.7.1/zk
export ZOO_CONF_DIR=${ZOO_BASE_DIR}/conf
export ZOO_DATA_DIR=${ZOO_BASE_DIR}/data
export ZOO_DYNAMIC_CONF_FILE="${ZOO_CONF_DIR}/zoo.cfg.dynamic"
export MY_ID_FILE=${ZOO_DATA_DIR}/myid
export HOST=web-0
export DOMAIN=zookeeker-headless
export FOLLOWER_PORT=2888
export ELECTION_PORT=3888
export CLIENT_PORT=2181

#HOSTNAME="$(hostname -s)"
HOSTNAME="web-1"
ConnStr=""
Replicas=3
#ReadyReplicas=3

function zkConfig() {
  echo "server.$1=$HOST.$DOMAIN:$FOLLOWER_PORT:$ELECTION_PORT;$CLIENT_PORT"
}

function getConnStr() {
   ConnStr=""
   for i in $(seq 0 $((${Replicas} - 1))); do
     ConnStr="${ConnStr}${ConnStr:+"\n"}" zkConfig $((i+1))
   done
}

getConnStr
echo ${ConnStr}


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


#exec "$@"