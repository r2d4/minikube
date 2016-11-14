#!/bin/bash

rm -r ./jobs
mkdir ./jobs
for i in `seq 0 $1`
do
  cat train.yaml.txt | sed "s/\$IDX/$i/" > ./jobs/train-$i.yaml
done

