#!/bin/bash

while [[ $# -gt 1 ]]
do
key="$1"

case $key in
    --num_workers)
    NUM_WORKERS="$2"
    shift # past argument
    ;;
    --num_ps)
    NUM_PS="$2"
    shift # past argument
    ;;
    *)
            
    ;;
esac
shift 
done

