#!/bin/bash
loop_cnt=1
loop_lmt=100
proc_name="nohup ./online_watch -a localhost:3306"
proc_name_old="./online_watch -a localhost:3306"
proc_dir="/home/bonly"
sleep_sec=0.5

echo "--------------- start `date +%F+%T+%N` ---------------"
while [ $loop_cnt -le $loop_lmt ]
do
    proc_num=`ps -ef|grep "${proc_name}"|grep -v grep|awk '{print $2}'`
    if [ "${proc_num}" = "" ]
    then
        proc_num=`ps -ef|grep "${proc_name_old}"|grep -v grep|awk '{print $2}'`
        if [ "${proc_num}" = "" ]
        then
            echo "Process is missing, respawn Process"
            cd $proc_dir
            ulimit -c unlimited
            export LD_LIBRARY_PATH=/lib:/usr/lib:/usr/local/lib:/usr/local/mysql/lib
            `$proc_name`
        else
            echo "Old process is running"
        fi
    else
        echo "Process is running"
    fi
    loop_cnt=$(($loop_cnt+1))
    sleep $sleep_sec
done

echo "---------------- end `date +%F+%T+%N` ----------------"
