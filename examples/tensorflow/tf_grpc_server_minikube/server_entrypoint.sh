#!/bin/bash

python /grpc_tensorflow_server.py --task_id="${HOSTNAME##*-}" "$@"
