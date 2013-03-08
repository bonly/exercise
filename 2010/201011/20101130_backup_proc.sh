#!/bin/sh
# Created by bonly 20080122.
#
# Backup site's routine.
TARNAME=/home/bonly/backup_new/spBackup"$1"`date '+%Y%m%d'`
/usr/local/mysql/bin/mysqldump -ubackup_user -p123456 -n -t -d -R my_site > "$TARNAME"
