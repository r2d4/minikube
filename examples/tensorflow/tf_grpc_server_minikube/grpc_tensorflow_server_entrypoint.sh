LOG_FILE="/tmp/grpc_tensorflow_server.log"
touch "${LOG_FILE}"

SCRIPT_DIR=$( cd ${0%/*} && pwd -P )

# When running as a k8s StatefulSet, 
# let the task id be the ordinal assigned
# See http://kubernetes.io/docs/user-guide/petset/#ordinal-index
TASK_ID=${HOSTNAME##*-}

# Wait with timeout for tf-worker peers to appear
# check /work-dir/worker-peers

# Wait with timeout for tf-ps peers to appear
# /work-dir/ps-peers

# Start the server and log
python grpc_tensorflow_server.py --task_id=TASK_ID $1 2>&1 | tee "${LOG_FILE}"