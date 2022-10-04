#!/usr/bin/env bash

#检测端口是否被占用
array=(9999 10000 10001 10002)
mark=0
for i in ${array[@]}
do
    port=$i
    #根据端口号查询对应的pid
    pid=`lsof -t -i:$port`
    if [  -n  "$pid"  ];  then
        echo "$port"
        mark=1
    fi
done

if [ $mark -eq 1 ]; then
    echo "端口被占用，请先停止进程！"
    exit
fi

if [ -f "./rpcserver" ];then
    echo "文件存在"
else
    go build rpcserver.go
fi

echo "RPC server started! ports: 9999 10000 10001 10002."

./rpcserver 9999 &
./rpcserver 10000 &
./rpcserver 10001 &
./rpcserver 10002 &



