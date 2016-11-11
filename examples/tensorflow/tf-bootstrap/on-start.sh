#!/bin/bash

set -e

PORT=2222
SERVICE=$1

PEERS="${SERVICE}="

while read -ra LINE; do
        PEERS=${PEERS}${LINE}:${PORT},
done

# Cut the last comma off
echo ${PEERS%?} > ${SERVICE}-peers