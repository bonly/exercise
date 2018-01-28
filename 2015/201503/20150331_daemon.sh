#!/bin/bash
proc_name="v2ray"
sleep_sec=5
proc_dir=root

while [ true ]
do
    proc_num=`ps -ef|grep "${proc_name}"|grep -v grep|awk '{print $2}'`
    if [ "${proc_num}" = "" ]
    then
        echo "Process is missing, respawn Process, `date +%F+%T+%N`"
        cd $proc_dir
        nohup nice -n 15 /root/`$proc_name` -config /root/config.json >/dev/null & 
    else
        echo "Process is running, `date +%F+%T+%N`"
    fi
    sleep $sleep_sec
done
