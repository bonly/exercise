#!/bin/bash
echo 'coming'
#read -p 'wait..' line
read line
echo "child says: Who's there?" $line
read -sp 'input:' line
echo "child says: Canoe who?" $line
read line
echo "child says: Groan..." $line
