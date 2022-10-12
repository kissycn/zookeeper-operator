#!/bin/bash

set -ex

OK=$(echo ruok | nc localhost $1)

# Check to see if zookeeper service answers
if [[ "$OK" == "imok" ]]; then
  exit 0
else
  exit 1
fi