#!/usr/bin/env bash

BIN=./base-server
STDLOG=output.log

RUNMODE=dev

if [["$1"]]; then
    RUNMODE=$1
fi

export RUNMODE=$RUNMODE

chmod u+x $BIN

ID=$(/usr/sbin/pidof "$BIN")
if [[ "$ID" ]] ; then
    echo "kill -SIGINT $ID"
    kill -2 $ID
fi

while :
do
    ID=$(/usr/sbin/pidof "$BIN")
    if [[ "$ID" ]] ; then
        echo "service still running...wait"
        sleep 0.1
    else
        echo "Starting service..."

        nohup $BIN > $STDLOG 2>&1 &
        echo "service started"
        break
    fi
done
