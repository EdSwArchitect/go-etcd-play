
# Go ETCD Docker Server and ETCD Client

## ETCD Docker Server

The following command is used to run the Docker server.

DATA_DIR is set to a directory location to save data placed into the server. NODE1 is the IP address or the host name of the server.

> REGISTRY=gcr.io/etcd-development/etcd
>  
> docker run \
>   -p 2379:2379 \
>   -p 2380:2380 \
>   --volume=${DATA_DIR}:/etcd-data \
>   --name etcd ${REGISTRY}:latest \
>   /usr/local/bin/etcd \
>   --data-dir=/etcd-data --name node1 \
>   --initial-advertise-peer-urls http://${NODE1}:2380 --listen-peer-urls http://0.0.0.0:2380 \
>   --advertise-client-urls http://${NODE1}:2379 --listen-client-urls http://0.0.0.0:2379 \
>   --initial-cluster node1=http://${NODE1}:2380
>


Tried to load JSON into etcd. The go client only worked without doing the JSON pretty print.

Used another application to load it.

Use this script to load the JSON config file: https://github.com/etcd-io/etcd/issues/8205#issuecomment-313230911

I modified the 

Before running
> export ETCDCTL_API=3

Command to run etcd in a container:

> REGISTRY=gcr.io/etcd-development/etcd
>  
> docker run \
>   -p 2379:2379 \
>   -p 2380:2380 \
>   --volume=${DATA_DIR}:/etcd-data \
>   --name etcd ${REGISTRY}:latest \
>   /usr/local/bin/etcd \
>   --data-dir=/etcd-data --name node1 \
>   --initial-advertise-peer-urls http://${NODE1}:2380 --listen-peer-urls http://0.0.0.0:2380 \
>   --advertise-client-urls http://${NODE1}:2379 --listen-client-urls http://0.0.0.0:2379 \
>   --initial-cluster node1=http://${NODE1}:2380
>
