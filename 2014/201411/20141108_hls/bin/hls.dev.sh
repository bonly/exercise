#!/bin/bash
nohup ./hls -log_dir=/tmp/ -logtostderr=true -v=99  -box_srv=0.0.0.0:5020 -callback=http://devpay.xbed.com.cn:9030/hls/callback -oms_srv=0.0.0.0:5010 >/dev/null &
