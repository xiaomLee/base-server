#!/usr/bin/env bash

BIN=/data/apps/gateway-ws/gateway-ws
STDLOG=/data/apps/gateway-ws/output.log

if [ $RUNMODE = "pre" ] ; then
        BIN=/data/apps/pre-gateway-ws/gateway-ws
        STDLOG=/data/apps/pre-gateway-ws/output.log
fi

chmod u+x $BIN

ID=$(/usr/sbin/pidof "$BIN")
if [ "$ID" ] ; then
        echo "kill -SIGINT $ID"
        kill -2 $ID
fi

while :
do
        ID=$(/usr/sbin/pidof "$BIN")
        if [ "$ID" ] ; then
                echo "gateway-ws still running...wait"
                sleep 0.1
        else
                echo "gateway-ws service was not started"
                echo "Starting service..."

                if [ $RUNMODE = "online" ] ; then
                        #nohup $BIN > /dev/null 2>&1 &
                        nohup $BIN > $STDLOG 2>&1 &
                else
                        nohup $BIN > $STDLOG 2>&1 &
                fi
                break
        fi
done
