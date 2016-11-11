#!/bin/sh

# This volume is assumed to exist and is shared with parent of the init
# container.
DATA_VOLUME="/data"

# This volume is assumed to exist and is shared with the peer-finder
# init container. It contains on-start/change configuration scripts.
WORK_DIR="/work-dir"

for i in "$@"
do
case $i in
    -d=*|--data-dir=*)
    DATA_VOLUME="${i#*=}"
    shift
    ;;
    -u=*|--data-urls=*)
	DATA_URLS="${i#*=}"
	shift
	;;
    -w=*|--work-dir=*)
    WORK_DIR="${i#*=}"
    shift
    ;;
    *)
    # unknown option
    ;;
esac
done

echo Copying configuration scripts into ${WORK_DIR}
mkdir -p "${WORK_DIR}"
cp /on-start.sh "${WORK_DIR}"/
cp /peer-finder "${WORK_DIR}"/

echo Downloading data into ${DATA_VOLUME}
mkdir -p "${DATA_VOLUME}"

IFS=',' 
for i in "${DATA_URLS}"; do
	echo Downloading $i into ${DATA_VOLUME}...
	wget -P ${DATA_VOLUME} -q $i
done


