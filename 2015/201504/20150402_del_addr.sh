#!/bin/bash
min1=1
min2=0

echo begin > /tmp/ping.log
for idx in {0..254}
do
    echo 192.168.$idx.2
    ip addr add 192.168.$idx.2/24 dev enp1s0f0
    #for idy in {2..254} 
    #do
        #echo 192.168.$idx.$idy
        #ip addr add 192.168.$idx.$idy/24 dev enp1s0f0
        #ip addr del 192.168.$idx.$idy/24 dev enp1s0f0
    #done
done

for idx in {0..254}
do 
    (ping 192.168.$idx.1 -c 1 >> /tmp/ping.log) &
done

#for idx in {0..254}
#do 
#    ip addr del 192.168.$idx.2/24 dev enp1s0f0
#done