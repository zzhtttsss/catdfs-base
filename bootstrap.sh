#!/bin/bash
echo "Hello World !"
# delete etcd key-value

# bootstrap master
echo "bootstrap master!"
docker build -t master-docker /Users/zzh/GolandProjects/tinydfs/tinydfs-master
docker run -itd --name master-test-1 --net=es-network --ip 172.18.0.11 master-docker

# bootstrap cs
echo "bootstrap cs!"


# bootstrap client
echo "bootstrap client!"
