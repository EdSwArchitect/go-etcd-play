
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

## Loading the JSON configuration file into ETCD server

> ./loadit.sh <filename> # Used when loading a non-pretty JSON into ETCD

> ./loadit-orig.sh <filename> # Used when loading the original JSON into ETCD


## ETCD Client 

Run after the loading of the configuration file. It will put a key/value into the ETCD. The key is foo.

It will also read the Common/Log entry from ETCD, returning the JSON.

*Tried to load JSON into etcd. The go client only worked without doing the JSON pretty print.*


Use this script to load the JSON config file: https://github.com/etcd-io/etcd/issues/8205#issuecomment-313230911

## ETCD CTL

From the command line, you can retrieve what's in ETCD with ETCD CTL

> etcdctl get Common/Log
>
> Common/Log
> {"disableLog4J":true,"maxHistorical":1000,"writers":{"console":{"type":"console","enabled":true,"levels":"ACCESS,ERROR,WARN,INFO,METRIC,DEBUG,DEVELOPMENT,SAMPLE","stdout":true,"stderr":false,"colorize":false,"bundling":true},"file":{"type":"file","enabled":false,"levels":"ACCESS,ERROR,WARN,INFO,DEBUG,DEVELOPMENT","folderPath":"./logs","filenamePrefix":"Logging","filenameSuffix":".log","delay":5000,"rollOverDays":5},"sampling":{"type":"file","enabled":false,"levels":"SAMPLE","folderPath":"./logs","filenamePrefix":"Sampling","filenameSuffix":".log","delay":5000,"rollOverDays":5},"metrics":{"type":"file","enabled":false,"levels":"METRIC","folderPath":"./logs","filenamePrefix":"Metrics","filenameSuffix":".log","delay":5000,"rollOverDays":5}}}
