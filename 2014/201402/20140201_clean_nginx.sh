#! /bin/sh
#Auto Clean Nginx Cache Shell Scripts
#2013-06-12  wugk
#Define Path
CACHE_DIR=/data/www/proxy_cache_dir/
FILE="$*"

#To determine whether the input script，If not, then exit 判断脚本是否有输入，没有输入然后退出
if
  [  "$#" -eq "0" ];then
  echo "Please Insert clean Nginx cache File, Example: $0 index.html index.js"
  sleep 2 && exit
fi
  echo "The file : $FILE to be clean nginx Cache ,please waiting ....."

#Wrap processing for the input file, for grep lookup，对输入的文件进行换行处理，利于grep查找匹配相关内容
for i in `echo $FILE |sed 's//\n/g'`
do
   grep -ra  $i  ${CACHE_DIR}| awk -F':' '{print $1}'  > /tmp/cache_list.txt
    for j in `cat/tmp/cache_list.txt`
  do
    rm  -rf  $j
    echo "$i  $j  is  Deleted Success !"
  done
done
#The Scripts exec success and exit 0