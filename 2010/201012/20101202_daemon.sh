#!/bin/bash
loop_cnt=1
loop_lmt=105
proc_dir="/wlmz/gsrv"
proc_name="./gsrv ${proc_dir}/gsrv.conf"
sleep_sec=0.5

while [ $loop_cnt -le $loop_lmt ]
do
    proc_num=`ps -ef|grep "${proc_name}"|grep -v grep|awk '{print $2}'`
    if [ "${proc_num}" = "" ]
    then
        echo "Process is missing, respawn Process, `date +%F+%T+%N`"
        cd $proc_dir
        ulimit -c unlimited
        export LD_LIBRARY_PATH=/lib:/usr/lib:/usr/local/lib:/usr/local/mysql/lib
        `$proc_name`
    else
        echo "Process is running, `date +%F+%T+%N`"
    fi
    loop_cnt=$(($loop_cnt+1))
    sleep $sleep_sec
done
