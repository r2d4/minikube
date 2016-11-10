LOG_FILE="/tmp/grpc_tensorflow_server.log"
SCRIPT_DIR=$( cd ${0%/*} && pwd -P )

# When running as a k8s StatefulSet, 
# let the task id be the ordinal assigned
# See http://kubernetes.io/docs/user-guide/petset/#ordinal-index
TASK_ID=${HOSTNAME##*-}

# We find the peers of the StatefulSet
# using a helper binary that ouputs 
CLUSTER_SPEC=$(/peer-finder)

touch "${LOG_FILE}"

python grpc_tensorflow_server.py --task_id=$@ 2>&1 | tee "${LOG_FILE}"