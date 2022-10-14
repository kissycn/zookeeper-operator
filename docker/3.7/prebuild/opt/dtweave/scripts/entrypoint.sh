#!/bin/bash

set -e

# Load logging library
. /opt/dtweave/scripts/utils/liblog.sh

export ZOO_BASE_DIR=/opt/modules/zookeeper
export ZOO_CONF_DIR=${ZOO_BASE_DIR}/conf
export ZOO_DATA_DIR=${ZOO_BASE_DIR}/data
export ZOO_DYNAMIC_CONF_FILE="${ZOO_CONF_DIR}/zoo.cfg.dynamic"
export MY_ID_FILE=${ZOO_DATA_DIR}/myid

# temp define
# TODO define env
export POD_NAME=zk-0
export ZK_NAME=zk
export HEADLESS_NAME=zk-headless
export NAMESPACE=default
export CLUSTER_DOMAIN=svc.cluster.local
export FOLLOWER_PORT=2888
export ELECTION_PORT=3888
export CLIENT_PORT=2181
export ZK_CONN_ADDR="zk1:2181,zk2:2182,zk3:2183"
export REPLICAS=4
export READY_REPLICAS=1

#HOSTNAME="$(hostname -s)"
HOSTNAME="web-4"
dynamicConf=""

function getMemberAddr() {
  echo "server.$1=${ZK_NAME}-$1.${HEADLESS_NAME}.${NAMESPACE}.${CLUSTER_DOMAIN}:${FOLLOWER_PORT}:${ELECTION_PORT}:participant;${CLIENT_PORT}"
}

function getDynamicConf() {
   for i in $(seq 0 $((${REPLICAS} - 1))); do
     dynamicConf="${dynamicConf}${dynamicConf:+"\n"}$(getMemberAddr $((i+1)))"
   done
}

function getMemberId() {
    if [[ -f $MY_ID_FILE ]]; then
      memberId="$(cat $MY_ID_FILE)"
    else
      if [[ $HOSTNAME =~ (.*)-([0-9]+)$ ]]; then
        ORD=${BASH_REMATCH[2]}
        memberId=$((ORD+1))
      else
        error "Failed to get index from hostname $HOST"
        exit 1
      fi
    fi
}

function setDynamicConf() {
    if [[ -f $ZOO_DYNAMIC_CONF_FILE ]]; then
      rm -rf $ZOO_DYNAMIC_CONF_FILE
    fi
    echo ${dynamicConf} > $ZOO_DYNAMIC_CONF_FILE
}

getMemberId
getDynamicConf
setDynamicConf

# re-joining after failure
if [[ -f $MY_ID_FILE ]]; then
  if [[ $READY_REPLICAS -gt 0 ]]; then
    echo "Re-joining zookeeper member by api"
    java -jar zk.jar add $ZK_CONN_ADDR $(getMemberAddr ${memberId})
    if [[ $? -ne 0 ]]; then
      echo "Re-joining failure"
    fi
  fi
else
  echo "Writing member id: $memberId to: MY_ID_FILE."
  echo ${memberId} > $MY_ID_FILE
  if [[ $READY_REPLICAS -gt 0 ]]; then
    echo "joining zookeeper member"
    java -jar zk.jar add $ZK_CONN_ADDR $(getMemberAddr ${memberId})
    if [[ $? -ne 0 ]]; then
      echo "Re-joining failure"
    fi
  fi
fi

exec "$@"