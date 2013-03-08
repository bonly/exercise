#!/bin/bash
kill $(ps ax|grep "./GServer ./yx_gsrv.conf" |grep -v "grep" | awk '{print $1}')
