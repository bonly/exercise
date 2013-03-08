#!/bin/sh
#
# Created by bonly.
#
# 20080707.
#
# Backup mysql's full data.
#

DBNAME=$1
USERNAME=backup_user
PASSWD=123456
PREFIX=/home/david_yeung/backup_new/
DIRNAME=$PREFIX`date '+%Y%m%d'`
echo $TARNAME
# Add your own database name here.
case "$1" in
  t_girl);;
  *) exit;;
esac
if [ ! -d "$DIRNAME" ]
then
  mkdir "$DIRNAME"
fi
# Get all the tables' name.
NUM=`/usr/local/mysql/bin/mysql -u$USERNAME -p$PASSWD -S/tmp/mysql50.sock -s -vv -e "show tables" -D $DBNAME|wc -l`
HEADNUM=`expr ${NUM} - 3`
TAILNUM=`expr ${NUM} - 7`
ARR1=`/usr/local/mysql/bin/mysql -u$USERNAME -p$PASSWD -S/tmp/mysql50.sock -s -vv -e "show tables" -D $DBNAME| head -n"$HEADNUM" | tail -n "$TAILNUM"`
ARR2=($ARR1)

i=0
while [ "$i" -lt "${#ARR2[@]}" ]
do
 tmpFileName=${ARR2[$i]}
 # The real dump process.
 /usr/local/mysql/bin/mysqldump -u$USERNAME -p"$PASSWD" -S/tmp/mysql50.sock "$DBNAME" "$tmpFileName" > "$DIRNAME"/"$tmpFileName"
 let "i++"
done

