#!/bin/bash
filename = "$1"
while read line
do
  name=$line
  echo "txt= $name"
done < "$filename"
