#!/bin/bash
for idx in {0..254}
do 
    ip addr del 192.168.$idx.2/24 dev enp1s0f0
done