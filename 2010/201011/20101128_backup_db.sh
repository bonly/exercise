#!/bin/sh
#
# Created by bonly
#
# 20080707.
#
# Backup mysql's full data.
#
DBNAME=$1
BACKUPDIR=/home/david_yeung/backup_new
USERNAME=backup_user
PASSWD=123456
TARNAME="$BACKUPDIR"/backup"$1"`date '+%Y%m%d'`
# Add your own database name here.
case "$1" in
  my_site);;
  *) exit;; 
esac
# Get all the tables' name.
NUM=`/usr/local/mysql/bin/mysql -u$USERNAME -p$PASSWD -s -vv -e "show tables" -D $DBNAME|wc -l`
HEADNUM=`expr ${NUM} - 3`
TAILNUM=`expr ${NUM} - 7`
ARR1=`/usr/local/mysql/bin/mysql -u$USERNAME -p$PASSWD -s -vv -e "show tables" -D $DBNAME| head -n"$HEADNUM" | tail -n "$TAILNUM"`
ARR2=($ARR1)
i=0
while [ "$i" -lt "${#ARR2[@]}" ]
do
 tmpFileName=${ARR2[$i]}
 # The real dump process.
 /usr/local/mysql/bin/mysqldump -u$USERNAME -p"$PASSWD" "$DBNAME" "$tmpFileName" >> "$TARNAME" 
 let "i++"
done
