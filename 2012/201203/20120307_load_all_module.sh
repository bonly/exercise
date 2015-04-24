#!/bin/bash
awk '{system("modprobe " $NF)}' modules.alias 
