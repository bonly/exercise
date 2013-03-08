#!/bin/bash
PID=$(ps ax|grep "online_watch" |grep -v "grep" | awk '{print $1}')
watch more /proc/$PID/status
