#!/bin/bash

LOG_FILE="/tmp/grpc_tensorflow_server.log"
touch "${LOG_FILE}"

SCRIPT_DIR=$( cd ${0%/*} && pwd -P )

# When running as a k8s StatefulSet, 
# let the task id be the ordinal assigned
# See http://kubernetes.io/docs/user-guide/petset/#ordinal-index
TASK_ID=${HOSTNAME##*-}

while true; do
	echo "."
	sleep 5;
done

# Wait with timeout for tf-worker peers to appear
# check /work-dir/worker-peers

echo Waiting at least one worker to be up...
while [ ! -f "/work-dir/worker_spec" ]; do 
	echo "."
	sleep 1; 
done

WORKER_SPEC="$(cat /work-dir/worker_spec)"

# Wait with timeout for tf-ps peers to appear
# /work-dir/ps-peers

echo Waiting for at least one parameter server to be up...
while [ ! -f "/work-dir/ps_spec" ]; do
	echo "."
	sleep 1;
done

PS_SPEC=$(cat /work-dir/ps_spec)

CLUSTER_SPEC=${PS_SPEC};${WORKER_SPEC}

# Start the server and log
python grpc_tensorflow_server.py --task_id=TASK_ID --cluster_spec=${CLUSTER_SPEC} $1 2>&1 | tee "${LOG_FILE}"