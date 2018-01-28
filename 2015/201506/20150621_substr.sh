#!/bin/bash

filename="/home/bonly/abc/src/git"

echo "search src from front:" ${filename#*src}
echo "search src from back:" ${filename%src*}
