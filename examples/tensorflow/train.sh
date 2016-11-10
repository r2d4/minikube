#!/bin/bash

for i in `seq 0 $(($1 - 1))`
do
 IDX=$i 	
 cat << EOF 
apiVersion: batch/v1
kind: Job
metadata:
  name: train-mnist-data-$IDX
spec:
  template:
    metadata:
      name: train-mnist-data-$IDX
    spec:
      containers:
      - name: train-mnist-data
        image: tensorflow/tf_grpc_test_server
        imagePullPolicy: IfNotPresent
        command: 
          - python
          - /var/tf-k8s/python/mnist_replica.py
        args:
          - --job_name=worker
          - --task_index=$IDX
          - --num_gpus=0
          - --train_steps=500
          - --sync_replicas=False
          - --existing_servers=True
          - --ps_hosts=tf-ps-0.tf-ps.default.svc.cluster.local:2222
          - --worker_hosts=tf-worker-0.tf-worker.default.svc.cluster.local:2222,tf-worker-1.tf-worker.default.svc.cluster.local:2222
        volumeMounts:
        - name: tf-data
          mountPath: /tmp
      volumes:
      - name: tf-data
        hostPath:
          path: /shared
      restartPolicy: Never
EOF
  echo -e "---"
done



