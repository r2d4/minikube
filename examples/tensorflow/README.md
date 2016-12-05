### Distributed Tensorflow with Minikube

```bash
$ minikube start
````

#### Tensorflow GRPC Image
```
$ eval $(minikube docker-env)
$ cd tf_grpc_server_minikube
$ make
```

#### Create the Worker and Parameter Servers
```bash
$ python k8s_config.py --num_workers 2 | kubectl create -f -
```
#### Download the training data in a shared folder
```bash
$ kubectl create -f download.yaml
```
#### 

#### Train
```bash
$ ./train.sh 2 | kubectl create -f - 
```

