#!/bin/bash

set -e
echo Entering on-start...
if [[ "${HOSTNAME}" =~ "tf-worker" ]]; then
        JOB_NAME="worker"
elif [[ "${HOSTNAME}" =~ "tf-ps" ]]; then
        JOB_NAME="ps"
else
	echo "Unknown job"
	exit 1
fi

JOB_SPEC="${JOB_NAME}|"

while read -ra LINE; do
        echo Found peer ${LINE}
        JOB_SPEC=${JOB_SPEC}${LINE}:${PORT:-2222},
done

mkdir -p "${WORK_DIR:-/work-dir}"
# Cut the last comma off
echo ${JOB_SPEC%?} > "${WORK_DIR:-/work-dir}/${JOB_NAME}_spec"

