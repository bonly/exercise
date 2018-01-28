#!/bin/bash
#
# Created by bonly.
#
# 20170714.
#
# Backup mysql's full data.
#
NOW=`date '+%c'`
#存储过程及函数
#TARNAME=/home/bonly/db_backup/backup_"$1"`date '+%Y%m%d'`_proc.sql
TARNAME=/home/bonly/db_backup/backup_proc.sql
/usr/bin/mysqldump -utecha -ptechappen -n -t -d -R techappen > "$TARNAME"
#数据
#TARNAME=/home/bonly/db_backup/backup_"$1"`date '+%Y%m%d'`_data.sql
TARNAME=/home/bonly/db_backup/backup_data.sql
/usr/bin/mysqldump -utecha -ptechappen -n -t techappen > "$TARNAME"
#表结构
#TARNAME=/home/bonly/db_backup/backup_"$1"`date '+%Y%m%d'`_tbl.sql
TARNAME=/home/bonly/db_backup/backup_tbl.sql
/usr/bin/mysqldump -utecha -ptechappen -d techappen > "$TARNAME"

cd /home/bonly/db_backup
git add .
git commit -am  "备份时间 $NOW"

#echo $NOW

#sync_db
#!/bin/bash
#cd /home/bonly/db_backup
#/bin/git pull origin master

#crontab -e
#0 * * * * /home/bonly/db_backup/backup.sh

#crontabl -e
#0 * * * * /bin/bash /home/bonly/sync_db