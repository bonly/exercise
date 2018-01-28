#!/bin/bash
DBNAME=$1
USERNAME=techa
PASSWD=techappen
PREFIX=/home/bonly/db_backup/
DIRNAME=$PREFIX`date '+%Y%m%d'`
echo $TARNAME
# Add your own database name here.
case "$1" in
  techappen);;
  *) exit;;
esac
if [ ! -d "$DIRNAME" ]
then
  mkdir "$DIRNAME"
fi
# Get all the tables' name.
NUM=`/usr/bin/mysql -u$USERNAME -p$PASSWD -S/var/run/mysqld/mysqld.sock -s -vv -e "show tables" -D $DBNAME|wc -l`
HEADNUM=`expr ${NUM} - 3`
TAILNUM=`expr ${NUM} - 7`
ARR1=`/usr/bin/mysql -u$USERNAME -p$PASSWD -S/var/run/mysqld/mysqld.sock -s -vv -e "show tables" -D $DBNAME| head -n"$HEADNUM" | tail -n "$TAILNUM"`
ARR2=${ARR1}

i=0
while [ "$i" -lt "${#ARR2[@]}" ]
do
 tmpFileName=${ARR2[$i]}
 # The real dump process.
 /usr/bin/mysqldump -u$USERNAME -p"$PASSWD" -S/var/run/mysqld/mysqld.sock "$DBNAME" "$tmpFileName" > "$DIRNAME"/"$tmpFileName"
 let "i++"
done

